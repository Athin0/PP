package asyncDispatching

import (
	"fmt"
	"sync"
	"testing"

	"PP/worker/genericMath"
)

var TestingDispersion = "0.8997599663"

var TestingF = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetVariance()
	return res
}

func prepareTestWorkersPool(numberOfWorkers int) (*WorkersPool, []chan *Result) {
	channels := make([]chan *Result, 6)
	for i := 0; i < 6; i++ {
		channels[i] = make(chan *Result, 1)
	}

	wp := NewWorkersPool(numberOfWorkers)

	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[0]})
	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[1]})
	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[2]})
	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[3]})
	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[4]})
	wp.AddJob(&Job{Data: TestingSequence, UnaryAction: TestingF, Result: channels[5]})

	return wp, channels
}

func TestWorkersPool_TestFlow(t *testing.T) {
	wp, channels := prepareTestWorkersPool(4)

	wg := &sync.WaitGroup{}
	wg.Add(6)

	for i, ch := range channels {
		go func(ch chan *Result, i int) {
			defer wg.Done()
			select {
			case res := <-ch:
				if fmt.Sprintf("%.10f", res.Data.(float64)) != TestingDispersion {
					t.Errorf("Expected %s, got %s", TestingDispersion, fmt.Sprintf("%.10f", res.Data.(float64)))
				}
				close(ch)
			}
		}(ch, i)
	}

	wg.Wait()
	wp.KillWorkersPool()
}

func BenchmarkWorkersPoolFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wp, channels := prepareTestWorkersPool(4)

		wg := &sync.WaitGroup{}
		wg.Add(6)

		for i, ch := range channels {
			func(ch chan *Result, i int) {
				defer wg.Done()
				select {
				case <-ch:
					close(ch)
				}
			}(ch, i)
		}

		wg.Wait()
		wp.KillWorkersPool()
	}
}
