package main

import (
	"fmt"
	"math/rand"
)

type (
	Graph struct {
		MaxW, MaxH int
		Nodes      []*Node
	}
	Node struct {
		Adj     []*Node
		ID      int
		Visited bool
		coordinates
	}
	coordinates struct {
		X int
		Y int
	}
)

func newGraph(maxW, maxH int) *Graph {
	return &Graph{
		MaxW:  maxW,
		MaxH:  maxH,
		Nodes: nil,
	}
}

func (g *Graph) AddNode(id int) {
	n := &Node{
		ID:          id,
		coordinates: getRandomCoordinates(g.MaxW, g.MaxH),
	}
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) GetNode(id int) (*Node, error) {
	for _, node := range g.Nodes {
		if node.ID == id {
			return node, nil
		}
	}
	return nil, fmt.Errorf("no node with id = %v", id)
}

func (g *Graph) AddAdj(from, to int) error {
	fNode, err := g.GetNode(from)
	if err != nil {
		return err
	}
	tNode, err := g.GetNode(to)
	if err != nil {
		return err
	}

	fNode.Adj = append(fNode.Adj, tNode)

	return nil
}

func (g *Graph) Fill(count int) {
	g.Nodes = nil
	for i := 1; i <= count; i++ {
		g.AddNode(i)
	}

	for i := 0; i <= count*2; i++ {
		from := rand.Intn(count)
		to := rand.Intn(count)
		if from == to {
			i--
			continue
		}
		g.AddAdj(from, to)
	}

}

func (g *Graph) GetEndAddress() int {
	var ids []int
	for i := 1; i < len(g.Nodes); i++ {
		if len(g.Nodes[i].Adj) != 0 {
			ids = append(ids, g.Nodes[i].ID)
		}
	}

	return ids[rand.Intn(len(ids))]
}
