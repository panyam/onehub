package dbsync

import (
	"log"
	"sync"
	"time"
)

type cmd struct {
	Name string
	Data interface{}
}

type LogEventCallbackFunc func([]PGMSG, error) (numProcessed int, stop bool)

type LogQueue struct {
	// How many items will be peeked at a time
	Callback LogEventCallbackFunc

	// A channel letting us know how many items are to be committed
	commitReqChan chan int

	pgdb *PGDB

	isRunning                 bool
	cmdChan                   chan cmd
	wg                        sync.WaitGroup
	TimerDelayBackoffFactor   float32
	DelayBetweenPeekRetries   time.Duration
	MaxDelayBetweenEmptyPeeks time.Duration
}

func NewLogQueue(pgdb *PGDB, callback LogEventCallbackFunc) *LogQueue {
	out := &LogQueue{
		Callback:                  callback,
		pgdb:                      pgdb,
		TimerDelayBackoffFactor:   1.5,
		DelayBetweenPeekRetries:   100 * time.Millisecond,
		MaxDelayBetweenEmptyPeeks: 5 * time.Second,
	}
	return out
}

func (l *LogQueue) IsRunning() bool {
	return l.isRunning
}

func (l *LogQueue) Stop() {
	l.cmdChan <- cmd{Name: "stop"}
	l.wg.Wait()
}

/**
 * Resume getting events.
 */
func (l *LogQueue) Start() {
	l.isRunning = true
	log.Println("Log Processor Started")
	l.wg.Add(1)
	l.commitReqChan = make(chan int)
	timerDelay := 0 * time.Second
	readTimer := time.NewTimer(0)
	defer func() {
		l.isRunning = false
		close(l.commitReqChan)
		if !readTimer.Stop() {
			<-readTimer.C
		}
		l.commitReqChan = nil
		l.wg.Done()
		log.Println("Log Processor Stopped")
	}()
	for {
		select {
		case cmd := <-l.cmdChan:
			if cmd.Name == "stop" {
				// clean up and stop
				return
			} else if cmd.Name == "resetreadtimer" {
				readTimer.Reset(0)
			} else {
				log.Println("Invalid command: ", cmd)
			}
		case <-readTimer.C:
			// TODO - make 1024 upto the caller
			msgs, err := l.pgdb.GetMessages(10000, false, nil)
			if err != nil {
				l.Callback(nil, err)
				return
			}
			if len(msgs) == 0 {
				// Nothing found - try again after a delay
				log.Println("Timer delay: ", timerDelay)
				if timerDelay == 0 {
					timerDelay += l.DelayBetweenPeekRetries
				} else {
					timerDelay = (3 * timerDelay) / 2
				}
				if timerDelay > l.MaxDelayBetweenEmptyPeeks {
					timerDelay = l.MaxDelayBetweenEmptyPeeks
				}
			} else {
				timerDelay = 0

				// Here we should process the messages so we can tell the DB
				numProcessed, stop := l.Callback(msgs, err)
				err = l.pgdb.Forward(numProcessed)
				if err != nil {
					panic(err)
				}
				if stop {
					return
				}
			}
			readTimer = time.NewTimer(timerDelay)
		}
	}
}
