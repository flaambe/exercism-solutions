package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)

	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("records not sequential")
		}

		if r.ID < r.Parent {
			return nil, errors.New("parent id larger than self id")
		}

		if r.ID == r.Parent && r.ID != 0 {
			return nil, errors.New("non zero id has self as parent")
		}

		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID == r.Parent {
			continue
		}

		nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
	}

	return nodes[0], nil
}
