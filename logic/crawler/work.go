package crawler

import "sync"

type Worker interface {
	Work()
}

type Pool struct {
	worker chan Worker
	wg sync.WaitGroup
}

func NewPool(maxGoroutines int) *Pool {
	p := Pool{
		worker: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.worker {
				w.Work()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.worker <- w
}

func (p *Pool) ShutDown() {
	close(p.worker)
	p.wg.Wait()
}