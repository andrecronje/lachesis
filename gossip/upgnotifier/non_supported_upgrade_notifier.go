package upgnotifier

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/Fantom-foundation/go-lachesis/version"
)

type Reader interface {
	GetNetworkVersion() *big.Int
	GetNonSupportedUpgrade() *big.Int
}

type Logger struct {
	logger.Instance
	reader Reader
	done   chan struct{}
	wg     sync.WaitGroup
}

func New(reader Reader) *Logger {
	return &Logger{
		reader:   reader,
		done:     make(chan struct{}),
		Instance: logger.MakeInstance(),
	}
}

func (l *Logger) Start() {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if l.reader.GetNetworkVersion().Cmp(version.AsBigInt()) > 0 {
					l.Log.Warn(fmt.Sprintf("Network upgrade %s was activated. Current node version is %s. "+
						"Please upgrade your node and re-sync the chain data.", version.BigToString(l.reader.GetNetworkVersion()), version.AsString()))
				} else if l.reader.GetNonSupportedUpgrade().Sign() > 0 {
					l.Log.Warn(fmt.Sprintf("Node's state is dirty because node was upgraded after the network upgrade %s was activated. "+
						"Please re-sync the chain data to continue.", version.BigToString(l.reader.GetNonSupportedUpgrade())))
				}
			case <-l.done:
				return
			}
		}
	}()
}

func (l *Logger) Stop() {
	close(l.done)
	l.wg.Wait()
}
