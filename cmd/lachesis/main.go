package main

import (
	"context"
	"fmt"
	"github.com/Fantom-foundation/go-lachesis/utils/delayer"
	"os"
	godebug "runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-lachesis/cmd/lachesis/metrics"
	"github.com/Fantom-foundation/go-lachesis/cmd/lachesis/tracing"
	"github.com/Fantom-foundation/go-lachesis/debug"
	"github.com/Fantom-foundation/go-lachesis/flags"
	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/integration"
	"github.com/Fantom-foundation/go-lachesis/utils/errlock"
	_ "github.com/Fantom-foundation/go-lachesis/version"
)

const (
	// clientIdentifier to advertise over the network.
	clientIdentifier = "go-lachesis"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
	// The app that holds all commands and flags.
	app = flags.NewApp(gitCommit, gitDate, "the go-lachesis command line interface")

	testFlags    []cli.Flag
	nodeFlags    []cli.Flag
	rpcFlags     []cli.Flag
	metricsFlags []cli.Flag
)

// init the CLI app.
func init() {
	overrideFlags()
	overrideParams()

	// Flags for testing purpose.
	testFlags = []cli.Flag{
		FakeNetFlag,
		LegacyTestnetFlag,
	}

	// Flags that configure the node.
	nodeFlags = []cli.Flag{
		utils.IdentityFlag,
		utils.UnlockedAccountFlag,
		utils.PasswordFileFlag,
		utils.BootnodesFlag,
		utils.LegacyBootnodesV4Flag,
		utils.LegacyBootnodesV5Flag,
		DataDirFlag,
		utils.KeyStoreDirFlag,
		utils.ExternalSignerFlag,
		utils.NoUSBFlag,
		utils.SmartCardDaemonPathFlag,
		utils.TxPoolLocalsFlag,
		utils.TxPoolNoLocalsFlag,
		utils.TxPoolJournalFlag,
		utils.TxPoolRejournalFlag,
		utils.TxPoolPriceLimitFlag,
		utils.TxPoolPriceBumpFlag,
		utils.TxPoolAccountSlotsFlag,
		utils.TxPoolGlobalSlotsFlag,
		utils.TxPoolAccountQueueFlag,
		utils.TxPoolGlobalQueueFlag,
		utils.TxPoolLifetimeFlag,
		utils.ExitWhenSyncedFlag,
		utils.CacheFlag,
		utils.CacheDatabaseFlag,
		utils.CacheTrieFlag,
		utils.CacheGCFlag,
		utils.CacheNoPrefetchFlag,
		utils.ListenPortFlag,
		utils.MaxPeersFlag,
		utils.MaxPendingPeersFlag,
		utils.NATFlag,
		utils.NoDiscoverFlag,
		utils.DiscoveryV5Flag,
		utils.NetrestrictFlag,
		utils.NodeKeyFileFlag,
		utils.NodeKeyHexFlag,
		utils.VMEnableDebugFlag,
		utils.NetworkIdFlag,
		utils.EthStatsURLFlag,
		utils.NoCompactionFlag,
		utils.GpoBlocksFlag,
		utils.LegacyGpoBlocksFlag,
		utils.GpoPercentileFlag,
		utils.LegacyGpoPercentileFlag,
		utils.GpoMaxGasPriceFlag,
		utils.EWASMInterpreterFlag,
		utils.EVMInterpreterFlag,
		configFileFlag,
		operaConfigFileFlag,
		validatorFlag,
	}

	rpcFlags = []cli.Flag{
		utils.HTTPEnabledFlag,
		utils.HTTPListenAddrFlag,
		utils.HTTPPortFlag,
		utils.HTTPCORSDomainFlag,
		utils.HTTPVirtualHostsFlag,
		utils.LegacyRPCEnabledFlag,
		utils.LegacyRPCListenAddrFlag,
		utils.LegacyRPCPortFlag,
		utils.LegacyRPCCORSDomainFlag,
		utils.LegacyRPCVirtualHostsFlag,
		utils.GraphQLEnabledFlag,
		utils.GraphQLCORSDomainFlag,
		utils.GraphQLVirtualHostsFlag,
		utils.HTTPApiFlag,
		utils.LegacyRPCApiFlag,
		utils.WSEnabledFlag,
		utils.WSListenAddrFlag,
		utils.LegacyWSListenAddrFlag,
		utils.WSPortFlag,
		utils.LegacyWSPortFlag,
		utils.WSApiFlag,
		utils.LegacyWSApiFlag,
		utils.WSAllowedOriginsFlag,
		utils.LegacyWSAllowedOriginsFlag,
		utils.IPCDisabledFlag,
		utils.IPCPathFlag,
		utils.InsecureUnlockAllowedFlag,
		utils.RPCGlobalGasCap,
		utils.RPCGlobalTxFeeCap,
	}

	metricsFlags = []cli.Flag{
		utils.MetricsEnabledFlag,
		utils.MetricsEnabledExpensiveFlag,
		utils.MetricsEnableInfluxDBFlag,
		utils.MetricsInfluxDBEndpointFlag,
		utils.MetricsInfluxDBDatabaseFlag,
		utils.MetricsInfluxDBUsernameFlag,
		utils.MetricsInfluxDBPasswordFlag,
		utils.MetricsInfluxDBTagsFlag,
		metrics.PrometheusEndpointFlag,
		tracing.EnableFlag,
	}

	// App.

	app.Action = lachesisOperaMigrationMain
	app.Version = params.VersionWithCommit(gitCommit, gitDate)
	app.HideVersion = true // we have a command to print the version
	app.Commands = []cli.Command{
		// See accountcmd.go:
		accountCommand,
		walletCommand,
		// See consolecmd.go:
		consoleCommand,
		attachCommand,
		javascriptCommand,
		// See config.go:
		dumpConfigCommand,
		// See misccmd.go:
		versionCommand,
		licenseCommand,
		// See chaincmd.go
		importCommand,
		exportCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags, testFlags...)
	app.Flags = append(app.Flags, nodeFlags...)
	app.Flags = append(app.Flags, rpcFlags...)
	app.Flags = append(app.Flags, consoleFlags...)
	app.Flags = append(app.Flags, debug.Flags...)
	app.Flags = append(app.Flags, metricsFlags...)

	app.Before = func(ctx *cli.Context) error {
		if err := debug.Setup(ctx); err != nil {
			return err
		}

		// Start metrics export if enabled
		utils.SetupMetrics(ctx)
		metrics.SetupPrometheus(ctx)

		return nil
	}

	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		prompt.Stdin.Close() // Resets terminal mode.
		return nil
	}

	// Tune Go's GC to be more aggressive
	godebug.SetGCPercent(65)
}

func main() {
	overrideFlags()
	overrideParams()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func lachesisMain(ctx *cli.Context, cfg *config) error {
	// TODO: tracing flags
	tracingStop, err := tracing.Start(ctx)
	if err != nil {
		return err
	}
	defer tracingStop()

	node, _ := makeNode(ctx, makeAllConfigs(ctx))
	defer node.Close()
	startNode(ctx, node)
	node.Wait()
	return nil
}

func makeNode(ctx *cli.Context, cfg *config) (*node.Node, *gossip.Service) {
	stack := makeConfigNode(ctx, &cfg.Node)

	// check errlock file
	// TODO: do the same with with stack.OpenDatabaseWithFreezer()
	errlock.SetDefaultDatadir(cfg.Node.DataDir)
	errlock.Check()

	engine, _, cdb, gdb := integration.MakeEngine(cfg.Node.DataDir, &cfg.Lachesis)
	metrics.SetDataDir(cfg.Node.DataDir)

	// configure emitter
	setValidator(ctx, &cfg.Lachesis.Emitter)

	// Create and register a gossip network service.

	svc, err := gossip.NewService(stack, &cfg.Lachesis, gdb, engine)
	if err != nil {
		utils.Fatalf("Failed to create the service: %v", err)
	}

	delayer.New(svc.IsMigration, 10*time.Second, func() {
		myValidatorID, myValidatorAddr := svc.Emitter().GetValidator()
		operaMigrationCtx.Store(&OperaMigration{
			validatorID:   myValidatorID,
			validatorAddr: myValidatorAddr,
			gdb:           gdb,
			cdb:           cdb,
		})
		err := stack.Close()
		if err != nil {
			utils.Fatalf("Service stopping error during migration: %v", err)
		}
	}).Start()

	stack.RegisterAPIs(svc.APIs())
	stack.RegisterProtocols(svc.Protocols())
	stack.RegisterLifecycle(svc)

	return stack, svc
}

func makeConfigNode(ctx *cli.Context, cfg *node.Config) *node.Node {
	stack, err := node.New(cfg)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}

	addFakeAccount(ctx, stack)

	return stack
}

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces.
func startNode(ctx *cli.Context, stack *node.Node) {
	debug.Memsize.Add("node", stack)

	// Start up the node itself
	utils.StartNode(stack)

	// Unlock any account specifically requested
	unlockAccounts(ctx, stack)

	// Register wallet event handlers to open and auto-derive wallets
	events := make(chan accounts.WalletEvent, 16)
	stack.AccountManager().Subscribe(events)

	// Create a client to interact with local lachesis node.
	rpcClient, err := stack.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to self: %v", err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	/*
		// Set contract backend for ethereum service if local node
		// is serving LES requests.
		if ctx.GlobalInt(utils.LightLegacyServFlag.Name) > 0 || ctx.GlobalInt(utils.LightServeFlag.Name) > 0 {
			var ethService *eth.Ethereum
			if err := stack.Service(&ethService); err != nil {
				utils.Fatalf("Failed to retrieve ethereum service: %v", err)
			}
			ethService.SetContractBackend(ethClient)
		}
		// Set contract backend for les service if local node is
		// running as a light client.
		if ctx.GlobalString(utils.SyncModeFlag.Name) == "light" {
			var lesService *les.LightEthereum
			if err := stack.Service(&lesService); err != nil {
				utils.Fatalf("Failed to retrieve light ethereum service: %v", err)
			}
			lesService.SetContractBackend(ethClient)
		}
	*/
	go func() {
		// Open any wallets already attached
		for _, wallet := range stack.AccountManager().Wallets() {
			if err := wallet.Open(""); err != nil {
				log.Warn("Failed to open wallet", "url", wallet.URL(), "err", err)
			}
		}
		// Listen for wallet event till termination
		for event := range events {
			switch event.Kind {
			case accounts.WalletArrived:
				if err := event.Wallet.Open(""); err != nil {
					log.Warn("New wallet appeared, failed to open", "url", event.Wallet.URL(), "err", err)
				}
			case accounts.WalletOpened:
				status, _ := event.Wallet.Status()
				log.Info("New wallet appeared", "url", event.Wallet.URL(), "status", status)

				var derivationPaths []accounts.DerivationPath
				if event.Wallet.URL().Scheme == "ledger" {
					derivationPaths = append(derivationPaths, accounts.LegacyLedgerBaseDerivationPath)
				}
				derivationPaths = append(derivationPaths, accounts.DefaultBaseDerivationPath)

				event.Wallet.SelfDerive(derivationPaths, ethClient)

			case accounts.WalletDropped:
				log.Info("Old wallet dropped", "url", event.Wallet.URL())
				event.Wallet.Close()
			}
		}
	}()

	// Spawn a standalone goroutine for status synchronization monitoring,
	// close the node when synchronization is complete if user required.
	if ctx.GlobalBool(utils.ExitWhenSyncedFlag.Name) {
		go func() {
			for first := true; ; first = false {
				// Call ftm_syncing until it returns false
				time.Sleep(5 * time.Second)

				var syncing bool
				err := rpcClient.CallContext(context.TODO(), &syncing, "ftm_syncing")
				if err != nil {
					continue
				}
				if !syncing {
					if !first {
						time.Sleep(time.Minute)
					}
					log.Info("Synchronisation completed. Exiting due to exitwhensynced flag.")
					err = stack.Close()
					if err != nil {
						continue
					}
					return
				}
			}
		}()
	}
}

// unlockAccounts unlocks any account specifically requested.
func unlockAccounts(ctx *cli.Context, stack *node.Node) {
	var unlocks []string
	inputs := strings.Split(ctx.GlobalString(utils.UnlockedAccountFlag.Name), ",")
	for _, input := range inputs {
		if trimmed := strings.TrimSpace(input); trimmed != "" {
			unlocks = append(unlocks, trimmed)
		}
	}
	// Short circuit if there is no account to unlock.
	if len(unlocks) == 0 {
		return
	}
	// If insecure account unlocking is not allowed if node's APIs are exposed to external.
	// Print warning log to user and skip unlocking.
	if !stack.Config().InsecureUnlockAllowed && stack.Config().ExtRPCEnabled() {
		utils.Fatalf("Account unlock with HTTP access is forbidden!")
	}
	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	passwords := utils.MakePasswordList(ctx)
	for i, account := range unlocks {
		unlockAccount(ks, account, i, passwords)
	}
}
