package puzzle20

import (
	"advent2022/base"
	"advent2022/util"
	"slices"
)

type p struct {
	nums []int64
}

type node struct {
	val        int64
	prev, next *node
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.nums = util.ArrayStrToInt64(lines)
}

func mix(nodes []*node) {
	length := len(nodes)
	for _, node := range nodes {
		switch util.Sign(node.val) {
		case 0: // no-op
		case 1:
			v := node.val % int64(length-1)
			tgt := node.next
			for i := int64(1); i < v; i++ {
				tgt = tgt.next
			}
			node.prev.next = node.next
			node.next.prev = node.prev

			tgt.next.prev = node
			node.next = tgt.next
			tgt.next = node
			node.prev = tgt
		case -1:
			v := -node.val % int64(length-1)
			tgt := node.prev
			for i := int64(1); i < v; i++ {
				tgt = tgt.prev
			}
			node.prev.next = node.next
			node.next.prev = node.prev

			tgt.prev.next = node
			node.prev = tgt.prev
			tgt.prev = node
			node.next = tgt
		}
	}

}

func (p *p) Solve1() any {
	var zero *node
	var length = len(p.nums)
	var nodes = make([]*node, 0, length)

	{
		cur := &node{val: p.nums[0]}
		head := cur
		nodes = append(nodes, cur)
		for _, n := range p.nums[1:] {
			next := &node{val: n, prev: cur}
			cur.next = next
			nodes = append(nodes, next)

			if n == 0 {
				zero = next
			}

			cur = next
		}
		cur.next = head
		head.prev = cur
	}

	mix(nodes)

	cur := zero
	keys := []int{1000 % length, 2000 % length, 3000 % length}
	slices.Sort(keys)
	step := 0
	val := int64(0)
	for _, k := range keys {
		for step < k {
			cur = cur.next
			step++
		}
		val += cur.val
	}

	return val
}

func (p *p) Solve2() any {
	var zero *node
	var length = len(p.nums)
	var nodes = make([]*node, 0, length)

	{
		cur := &node{val: p.nums[0] * 811589153}
		head := cur
		nodes = append(nodes, cur)
		for _, n := range p.nums[1:] {
			next := &node{val: n * 811589153, prev: cur}
			cur.next = next
			nodes = append(nodes, next)

			if n == 0 {
				zero = next
			}

			cur = next
		}
		cur.next = head
		head.prev = cur
	}

	for i := 0; i < 10; i++ {
		mix(nodes)
	}

	cur := zero
	keys := []int{1000 % length, 2000 % length, 3000 % length}
	slices.Sort(keys)
	step := 0
	val := int64(0)
	for _, k := range keys {
		for step < k {
			cur = cur.next
			step++
		}
		val += cur.val
	}

	return val
}

func init() {
	base.Register(20, &p{})
}
