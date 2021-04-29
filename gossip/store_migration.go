package gossip

import (
	"github.com/ethereum/go-ethereum/ethdb"

	"github.com/Fantom-foundation/go-lachesis/utils/migration"
)

func isEmptyDB(db ethdb.Iteratee) bool {
	it := db.NewIterator(nil, nil)
	defer it.Release()
	return !it.Next()
}

func (s *Store) Migrate() error {
	versions := migration.NewKvdbIDStore(s.table.Version)
	if isEmptyDB(s.mainDb) && isEmptyDB(s.async.mainDb) {
		// short circuit if empty DB
		versions.SetID(s.migrations().ID())
		return nil
	}
	return s.migrations().Exec(versions)
}

func (s *Store) migrations() *migration.Migration {
	return migration.
		Begin("lachesis-gossip-store").
		Next("remove async data from sync DBs",
			func() error {
				s.rmPrefix(s.table.PackInfos, "serverPool")
				s.rmPrefix(s.mainDb, "Z")
				return nil
			}).
		Next("remove legacy genesis field",
			s.app.MigrateEraseGenesisField).
		Next("multi-delegations",
			s.app.MigrateMultiDelegations).
		Next("adjustable offline pruning time",
			s.app.MigrateAdjustableOfflinePeriod).
		Next("adjustable minimum gas price",
			s.app.MigrateAdjustableMinGasPrice)
}
