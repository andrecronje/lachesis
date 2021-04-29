package delayer

import (
	"github.com/Fantom-foundation/go-lachesis/logger"
	"sync"
	"time"
)

type Delayer struct {
	logger.Instance
	done      chan struct{}
	wg        sync.WaitGroup
	condition func() bool
	action    func()
	delay     time.Duration
}

func New(condition func() bool, delay time.Duration, action func()) *Delayer {
	return &Delayer{
		Instance:  logger.MakeInstance(),
		done:      make(chan struct{}),
		condition: condition,
		delay:     delay,
		action:    action,
	}
}

func (w *Delayer) Start() {
	operaMigrationStarted := time.Time{}
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		ticker := time.NewTicker(w.delay / 20)
		for {
			select {
			case <-ticker.C:
				if operaMigrationStarted.IsZero() && w.condition() {
					operaMigrationStarted = time.Now()
				}
				if !operaMigrationStarted.IsZero() && time.Since(operaMigrationStarted) >= w.delay {
					w.action()
					return
				}
			case <-w.done:
				return
			}
		}
	}()
}

func (w *Delayer) Stop() {
	close(w.done)
	w.wg.Wait()
}
