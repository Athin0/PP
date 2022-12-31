package asyncDispatching

import (
	"PP/worker/genericMath"
	"sync"
)

type UnaryFloatFunction func(*genericMath.FloatSequence) float64

type BinaryFloatFunction func(float64, float64) float64

type Dispatcher struct {
	Target *Node
	cache  *Cache
	wp     *WorkersPool
}

type Node struct {
	InitialSequence *genericMath.FloatSequence
	Left            *Node
	Right           *Node
	Parent          *Node
	UnaryAction     UnaryFloatFunction
	BinaryAction    BinaryFloatFunction
}

func NewDispatcher(target *Node, sequence *genericMath.FloatSequence, numberOfWorkers int) *Dispatcher {
	disp := Dispatcher{
		Target: target,
		cache:  NewCache(),
		wp:     NewWorkersPool(numberOfWorkers),
	}

	var traverseTree func(node *Node, sequence *genericMath.FloatSequence)

	traverseTree = func(node *Node, sequence *genericMath.FloatSequence) {
		if node.Left == nil && node.Right == nil {
			node.InitialSequence = sequence
		}

		if node.Left != nil {
			traverseTree(node.Left, sequence)
		}
		if node.Right != nil {
			traverseTree(node.Right, sequence)
		}
	}

	traverseTree(target, sequence)

	return &disp
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
	disp.wp.KillWorkersPool()
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

	if node.Left != nil { //todo  нужно ли канал?
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

	if node.BinaryAction != nil && node.UnaryAction == nil { //todo проверку в функцию??
		hash := BinaryIntFuncHash(node.BinaryAction, nodeData[0], nodeData[1])

		item, err := disp.cache.GetItem(hash)
		if err == nil { //if it has in cache just return
			result = item.(float64)
		} else {
			resCh := make(chan *Result, 1)
			disp.wp.AddJob(&Job{
				Data:         nodeData,
				BinaryAction: node.BinaryAction,
				Result:       resCh,
			})

			select {
			case res := <-resCh:
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
			resCh := make(chan *Result, 1) //sent request in pull
			disp.wp.AddJob(&Job{
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
