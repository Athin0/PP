package asyncDispatching

import (
	"reflect"

	"PP/worker/genericMath"
)

type WorkersPool struct {
	workersCount int
	jobs         chan *Job
}

type Job struct {
	Data         interface{}
	UnaryAction  UnaryFloatFunction
	BinaryAction BinaryFloatFunction
	Result       chan *Result
}

type Result struct {
	Data interface{}
}

func NewWorkersPool(numberOfWorkers int) *WorkersPool {
	wp := WorkersPool{
		workersCount: numberOfWorkers,
		jobs:         make(chan *Job, numberOfWorkers),
	}

	for i := 0; i < numberOfWorkers; i++ { //todo make different version
		go func(in chan *Job) {
			for job := range in {
				if job.UnaryAction != nil && job.BinaryAction == nil {
					result := &Result{
						Data: job.UnaryAction(job.Data.(*genericMath.FloatSequence)),
					}
					job.Result <- result
				} else {
					result := &Result{
						Data: job.BinaryAction(
							reflect.ValueOf(job.Data).Index(0).Float(),
							reflect.ValueOf(job.Data).Index(1).Float(),
						),
					}
					job.Result <- result
				}
			}
		}(wp.jobs)
	}

	return &wp
}

func (wp *WorkersPool) AddJob(job *Job) { //todo добавить err если нельзя вкинуть работу?
	wp.jobs <- job
}

func (wp *WorkersPool) KillWorkersPool() {
	close(wp.jobs)
	wp.workersCount = 0
	wp.jobs = nil
}
