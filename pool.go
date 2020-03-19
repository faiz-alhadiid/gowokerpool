package goworkerpool

import (
	"sync"
)

// Pool ...
type Pool struct {
	Workers    []WorkerAdaptor
	WorkerChan chan WorkerAdaptor
	wg         sync.WaitGroup
}

// NewPool ...
func NewPool(workers ...WorkerAdaptor) *Pool {
	return &Pool{
		Workers:    workers,
		WorkerChan: make(chan WorkerAdaptor, len(workers)),
	}
}

// Init ...
func (p *Pool) Init() error {
	for _, worker := range p.Workers {
		if err := worker.Init(); err != nil {
			return err
		}
		p.WorkerChan <- worker
	}
	return nil
}

// Run ...
func (p *Pool) Run(processes <-chan interface{}) error {
	outChannel := make(chan Result)
	for {
		go func() {
			p.wg.Add(1)
			worker := <-p.WorkerChan
			worker.Execute(processes, outChannel)
			p.WorkerChan <- worker
			p.wg.Done()
		}()
	}
}
