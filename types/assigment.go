package types

import (
	"fmt"
	"strings"
)

type Assignee int
type Job int
type Cost int
type Assignment struct {
	Assignee Assignee
	Job      Job
	Cost     Cost
}

type AssignmentMap map[Assignee]Assignment

func (am AssignmentMap) TotalCost() Cost {
	var total Cost = 0
	for _, v := range am {
		total += v.Cost
	}

	return total
}

func (am Assignment) String() string {
	return fmt.Sprintf("assignee(%d)\t=> job(%d)\t: $%d", am.Assignee, am.Job, am.Cost)
}

func (am AssignmentMap) String() string {

	lines := []string{}
	for _, v := range am {
		lines = append(lines, v.String())
	}

	return strings.Join(lines, "\n")
}
