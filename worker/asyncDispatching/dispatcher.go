package asyncDispatching

import (
	"PP/worker/Math"
	"fmt"
	"reflect"
	"sync"
)

type UnaryFloatFunction func(*Math.FloatSequence) float64

type BinaryFloatFunction func(float64, float64) float64

type Dispatcher struct {
	Target *Node
	cache  *Cache
}

type Node struct {
	InitialSequence *Math.FloatSequence
	Left            *Node
	Right           *Node
	Parent          *Node
	UnaryAction     UnaryFloatFunction
	BinaryAction    BinaryFloatFunction
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

func NewDispatcher(target *Node) *Dispatcher {
	disp := Dispatcher{
		Target: target,
		cache:  NewCache(),
	}

	return &disp
}

// ShowTree traverse  to show
func ShowTree(node *Node) {
	if node.Left == nil && node.Right == nil {
		fmt.Println(node)
		return
	}

	if node.Left != nil {
		ShowTree(node.Left)
	}
	if node.Right != nil {
		ShowTree(node.Right)
	}
	fmt.Println(node)
}

func (disp *Dispatcher) Drop() {
	var traverseTree func(node *Node)

	traverseTree = func(node *Node) {
		if node.Left == nil && node.Right == nil {
			node.InitialSequence = nil
		}
		if node.Left != nil {
			node.InitialSequence = nil
			traverseTree(node.Left)
		}
		if node.Right != nil {
			node.InitialSequence = nil
			traverseTree(node.Right)
		}
	}

	traverseTree(disp.Target)
	disp.cache.Drop()
}

func Traverse(disp *Dispatcher) float64 {
	defer disp.Drop()

	var wg sync.WaitGroup
	channel := make(chan float64, 1)

	wg.Add(1)
	go traverseUtil(disp.Target, channel, &wg, disp)
	wg.Wait()
	close(channel)

	return <-channel
}

func traverseUtil(node *Node, out chan float64, wg *sync.WaitGroup, disp *Dispatcher) {
	defer wg.Done()

	var nwg sync.WaitGroup
	channel := make(chan float64, 2)

	if node.Left != nil {
		nwg.Add(1)
		go traverseUtil(node.Left, channel, &nwg, disp)
	}
	if node.Right != nil {
		nwg.Add(1)
		go traverseUtil(node.Right, channel, &nwg, disp)
	}
	nwg.Wait()
	close(channel)

	nodeData := make([]float64, 0) //read data from channel to array(nodeData)
	for val := range channel {     // from left and right nodes
		nodeData = append(nodeData, val)
	}

	result := 0.0

	if node.BinaryAction != nil && node.UnaryAction == nil { // can add this in separate function
		hash := BinaryIntFuncHash(node.BinaryAction, nodeData[0], nodeData[1])

		item, err := disp.cache.GetItem(hash)
		if err == nil { //if it has in cache just return
			result = item.(float64)
		} else {
			resCh := make(chan *Result, 1)
			go DoJob(&Job{
				Data:         nodeData,
				BinaryAction: node.BinaryAction,
				Result:       resCh,
			})

			select {
			case res := <-resCh: //wait for answer
				disp.cache.SetItem(hash, res.Data)
				result = res.Data.(float64)
				close(resCh)
			}
		}
	} else {
		hash := UnaryIntSeqFuncHash(node.UnaryAction, node.InitialSequence)

		item, err := disp.cache.GetItem(hash)
		if err == nil {
			result = item.(float64)
		} else {
			resCh := make(chan *Result, 1)
			go DoJob(&Job{
				Data:        node.InitialSequence,
				UnaryAction: node.UnaryAction,
				Result:      resCh,
			})

			select {
			case res := <-resCh: //wait for answer
				disp.cache.SetItem(hash, res.Data)
				result = res.Data.(float64)
				close(resCh)
			}
		}
	}

	out <- result
}

func DoJob(job *Job) {
	if job.UnaryAction != nil && job.BinaryAction == nil {
		result := &Result{
			Data: job.UnaryAction(job.Data.(*Math.FloatSequence)),
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
