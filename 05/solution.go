package puzzle05

import (
	"advent2022/base"
	"advent2022/util"
	"fmt"
	"slices"
	"strings"
)

type p struct {
	stacks [][]byte
	ops    [][3]int
}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)

	stackInput := blocks[0]
	maxHeight := len(stackInput) - 1
	numStack := len(strings.Fields(stackInput[maxHeight]))
	p.stacks = make([][]byte, numStack)

	for i := maxHeight - 1; i >= 0; i-- {
		for st := 0; st < numStack; st++ {
			idx := 1 + st*4
			if idx >= len(stackInput[i]) {
				break
			}
			if c := stackInput[i][idx]; c != ' ' {
				p.stacks[st] = append(p.stacks[st], c)
			}
		}
	}

	p.ops = make([][3]int, 0, len(blocks[1]))
	for _, op := range blocks[1] {
		var num, from, to int
		util.Assert(
			util.Must(
				fmt.Sscanf(op, "move %d from %d to %d",
					&num, &from, &to),
			) == 3,
			"invalid op "+op,
		)
		p.ops = append(p.ops, [3]int{num, from, to})
	}
}

func (p *p) Solve1() any {
	stacks := make([][]byte, len(p.stacks))
	for i := 0; i < len(p.stacks); i++ {
		stacks[i] = slices.Clone(p.stacks[i])
	}

	for _, op := range p.ops {
		num, from, to := op[0], op[1]-1, op[2]-1
		var moved []byte

		stacks[from], moved = stacks[from][:len(stacks[from])-num], stacks[from][len(stacks[from])-num:]
		slices.Reverse(moved)
		stacks[to] = append(stacks[to], moved...)
	}

	return util.Reduce(stacks, func(acc string, stack []byte) string {
		return acc + string(stack[len(stack)-1])
	}, "")
}

func (p *p) Solve2() any {
	stacks := make([][]byte, len(p.stacks))
	for i := 0; i < len(p.stacks); i++ {
		stacks[i] = slices.Clone(p.stacks[i])
	}

	for _, op := range p.ops {
		num, from, to := op[0], op[1]-1, op[2]-1
		var moved []byte

		stacks[from], moved = stacks[from][:len(stacks[from])-num], stacks[from][len(stacks[from])-num:]
		stacks[to] = append(stacks[to], moved...)
	}

	return util.Reduce(stacks, func(acc string, stack []byte) string {
		return acc + string(stack[len(stack)-1])
	}, "")
}

func init() {
	base.Register(5, &p{})
}
