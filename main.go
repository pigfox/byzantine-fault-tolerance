package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numNodes     = 7  // Total number of nodes
	byzantineMax = 2  // Max faulty nodes allowed (f)
	valueToAgree = 42 // The correct value honest nodes agree on
)

type Node struct {
	id        int
	byzantine bool
}

func (n *Node) Propose() int {
	if n.byzantine {
		// Byzantine node sends a random (possibly incorrect) value
		fakeValue := rand.Intn(100)
		fmt.Printf("Node %d (Byzantine) proposes: %d\n", n.id, fakeValue)
		return fakeValue
	}
	fmt.Printf("Node %d (Honest) proposes: %d\n", n.id, valueToAgree)
	return valueToAgree
}

func majorityVote(responses map[int]int) (int, bool) {
	counts := make(map[int]int)
	for _, val := range responses {
		counts[val]++
	}

	for val, count := range counts {
		if count >= numNodes-byzantineMax {
			return val, true
		}
	}
	return 0, false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create nodes (with random Byzantine behavior)
	nodes := make([]Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = Node{
			id:        i,
			byzantine: rand.Intn(numNodes) < byzantineMax,
		}
	}

	// Each node proposes a value
	proposals := make(map[int]int)
	for _, node := range nodes {
		proposals[node.id] = node.Propose()
	}

	// Aggregate and decide
	value, ok := majorityVote(proposals)
	if ok {
		fmt.Printf("\n🎉 Consensus reached on value: %d\n", value)
	} else {
		fmt.Println("\n❌ Consensus could not be reached")
	}
}
