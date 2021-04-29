package eventcheck

import (
	"errors"

	"github.com/Fantom-foundation/go-lachesis/eventcheck/epochcheck"
)

var (
	ErrAlreadyConnectedEvent = errors.New("event is connected already")
	ErrMigration             = errors.New("network migration")
)

func IsBan(err error) bool {
	if err == epochcheck.ErrNotRelevant ||
		err == ErrAlreadyConnectedEvent ||
		err == ErrMigration {
		return false
	}
	return err != nil
}
