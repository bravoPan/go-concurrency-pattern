package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutine)
	for i := 0; i < maxGoroutine; i++ {
		go func() {
			// 1st: with range - will not iterate until channel is close
			//for w := range p.work {
			//	w.Task()
			//}
			// 2nd: with for select continuously to check if channel is close, if close then stop
			for {
				w, ok := <-p.work
				if !ok {
					break
				}
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
}
