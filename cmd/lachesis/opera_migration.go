package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/go-opera/cmd/opera/launcher"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/valkeystore/encryption"
	oidx "github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/kvdb/leveldb"
	"github.com/Fantom-foundation/go-lachesis/poset"
	"github.com/Fantom-foundation/go-lachesis/utils/iocopy"
)

type OperaMigration struct {
	validatorID   idx.StakerID
	validatorAddr common.Address
	cdb           *poset.Store
	gdb           *gossip.Store
}

var operaMigrationCtx atomic.Value

func excludeArg(args []string, exclude string, excludeValue bool) []string {
	filteredArgs := make([]string, 0, len(os.Args))
	skip := 0
	for _, arg := range args {
		if strings.HasPrefix(arg, exclude+"=") {
			skip++
		} else if arg == exclude {
			skip++
			if excludeValue {
				skip++
			}
		}
		if skip == 0 {
			filteredArgs = append(filteredArgs, arg)
		} else {
			skip--
		}
	}
	return filteredArgs
}

func addFrontArgs(args []string, add []string) []string {
	return append(append([]string{args[0]}, add...), args[1:]...)
}

func operaMigration(m *OperaMigration, cfg *config, operaDatadir, operaGenesisPath string, migrationDB kvdb.Writer) {
	time.Sleep(5 * time.Second) // sleep just in case not all the lachesis services are terminated

	start := time.Now()

	// migrate state
	const dirPerm = 0700
	_ = os.MkdirAll(operaDatadir, dirPerm)
	tmpPath := path.Join(cfg.Node.DataDir, "tmp", "export")
	_ = os.RemoveAll(tmpPath)
	log.Info("Exporting Opera-compatible state", "tmp", tmpPath, "path", operaGenesisPath)
	genStore, err := ExportEncodedState(operaGenesisPath, tmpPath, m.gdb, m.cdb, &cfg.Lachesis.Net)
	if genStore != nil {
		defer genStore.Close()
		defer os.RemoveAll(tmpPath)
	}
	if err != nil {
		utils.Fatalf("Failed to export state: %v\n", err)
	}
	log.Info("Exported Opera-compatible state", "elapsed", common.PrettyDuration(time.Since(start)))

	// migrate keys
	_, _, srckeydir, _ := cfg.Node.AccountConfig()
	dstkeydir := path.Join(operaDatadir, "keystore")
	// account keys
	_ = iocopy.Dir(srckeydir, dstkeydir)
	// nodekey
	_ = os.MkdirAll(path.Join(operaDatadir, "go-opera"), dirPerm)
	_ = iocopy.File(path.Join(path.Join(cfg.Node.DataDir, "go-lachesis"), "nodekey"), path.Join(path.Join(operaDatadir, "go-opera"), "nodekey"))
	// validator key (if any)
	var pubkey validatorpk.PubKey
	if m.validatorID != 0 {
		validators := genStore.GetMetadata().Validators.Map()
		me, ok := validators[oidx.ValidatorID(m.validatorID)]
		if ok {
			pubkey = me.PubKey
			acckeypath, err := launcher.FindAccountKeypath(m.validatorAddr, srckeydir)
			if err != nil {
				log.Warn("Failed to migrate validator key", "err", err)
			}
			if len(acckeypath) != 0 {
				valkeypath := path.Join(path.Join(dstkeydir, "validator"), common.Bytes2Hex(pubkey.Bytes()))
				err = encryption.MigrateAccountToValidatorKey(acckeypath, valkeypath, pubkey)
				if err != nil {
					log.Warn("Failed to migrate validator key", "err", err)
				}
			}
		}
	}

	// After a successful migration, save migration context to DB
	if m.validatorID != 0 && !pubkey.Empty() {
		_ = migrationDB.Put([]byte("validatorID"), m.validatorID.Bytes())
		_ = migrationDB.Put([]byte("validatorAddr"), m.validatorAddr.Bytes())
		_ = migrationDB.Put([]byte("validatorPubkey"), []byte(pubkey.String()))
	}
	_ = migrationDB.Put([]byte("migrated"), []byte{1})
}

func resolvePath(path, datadir string) string {
	if filepath.IsAbs(path) {
		return path
	}
	if datadir == "" || path == "" {
		return ""
	}
	return filepath.Join(datadir, path)
}

func lachesisOperaMigrationMain(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	if ctx.GlobalIsSet(configFileFlag.Name) && !ctx.GlobalIsSet(operaConfigFileFlag.Name) {
		return fmt.Errorf("legacy Lachesis config file is specified and Opera config is not specified. Use --%s flag to specify Opera-compatible config path", operaConfigFileFlag.Name)
	}
	if !ctx.GlobalIsSet(configFileFlag.Name) && ctx.GlobalIsSet(operaConfigFileFlag.Name) {
		return fmt.Errorf("Opera config file is specified and legacy Lachesis config is not specified. Use --%s to specify Lachesis-compatible config path", configFileFlag.Name)
	}
	cfg := makeAllConfigs(ctx)
	operaDatadir := path.Join(cfg.Node.DataDir, "opera")
	operaArgs := excludeArg(os.Args, "--"+validatorFlag.Name, true)
	operaArgs = excludeArg(operaArgs, "--"+configFileFlag.Name, true)
	operaArgs = excludeArg(operaArgs, "--"+DataDirFlag.Name, true)
	operaArgs = excludeArg(operaArgs, "--"+utils.MetricsEnabledFlag.Name, false)
	operaArgs = excludeArg(operaArgs, "--"+utils.LegacyTestnetFlag.Name, false)
	operaArgs = addFrontArgs(operaArgs, []string{"--" + DataDirFlag.Name, operaDatadir})

	migrationDB, err := leveldb.New(path.Join(cfg.Node.DataDir, "migration"), 1, 1, "", nil, nil)
	if err != nil {
		return fmt.Errorf("Failed to Opera migration DB: %v", err)
	}
	defer migrationDB.Close()

	migrated, _ := migrationDB.Get([]byte("migrated"))
	if migrated == nil {
		testOperaArgs := excludeArg(operaArgs, "--"+utils.UnlockedAccountFlag.Name, true)
		testOperaArgs = excludeArg(testOperaArgs, "--"+utils.PasswordFileFlag.Name, true)
		testOperaArgs = append(testOperaArgs, "checkconfig")
		// test that Opera is launchable
		metricsFlag := metrics.Enabled
		metrics.Enabled = false
		err := launcher.Launch(testOperaArgs)
		if err != nil {
			return fmt.Errorf("Opera is not launchable: %v", err)
		}
		metrics.Enabled = metricsFlag
		// launch Lachesis
		err = lachesisMain(ctx, cfg)
		if err != nil {
			return err
		}
	}
	operaGenesisPath := path.Join(operaDatadir, "genesis.g")
	operaArgs = addFrontArgs(operaArgs, []string{"--genesis", operaGenesisPath})
	operaArgs = excludeArg(operaArgs, "--"+FakeNetFlag.Name, true)

	if v := operaMigrationCtx.Load(); v != nil {
		// perform migration
		m := v.(*OperaMigration)
		operaMigration(m, cfg, operaDatadir, operaGenesisPath, migrationDB)
	}

	// launch opera
	migrated, _ = migrationDB.Get([]byte("migrated"))
	if migrated != nil {
		// set validator
		validatorIDBytes, _ := migrationDB.Get([]byte("validatorID"))
		validatorPubkeyBytes, _ := migrationDB.Get([]byte("validatorPubkey"))
		zeroAddr := common.Address{}
		if getValidatorAddr(ctx, &cfg.Lachesis.Emitter) != zeroAddr && validatorIDBytes != nil && validatorPubkeyBytes != nil {
			operaArgs = addFrontArgs(operaArgs, []string{"--validator.id", fmt.Sprintf("%d", idx.BytesToStakerID(validatorIDBytes))})
			operaArgs = addFrontArgs(operaArgs, []string{"--validator.pubkey", string(validatorPubkeyBytes)})
			if ctx.GlobalIsSet(utils.PasswordFileFlag.Name) {
				operaArgs = addFrontArgs(operaArgs, []string{"--validator.password", ctx.GlobalString(utils.PasswordFileFlag.Name)})
			}
		}
		// set IPC
		operaArgs = excludeArg(operaArgs, "--"+utils.IPCPathFlag.Name, true)
		operaArgs = addFrontArgs(operaArgs, []string{"--" + utils.IPCPathFlag.Name, resolvePath(cfg.Node.IPCPath, cfg.Node.DataDir)})
		metrics.Enabled = false
		err = launcher.Launch(operaArgs)
		if err != nil {
			return fmt.Errorf("Opera error: %v", err)
		}
	}

	return nil
}
