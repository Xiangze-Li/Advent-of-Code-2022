package puzzle13

import (
	"advent2022/base"
	"advent2022/util"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type p struct {
	// blocks [][]string
	pairs [][2]*item
}

type item struct {
	type_  itemType
	value  int
	values []*item
}

func (it item) String() string {
	switch it.type_ {
	case itemTypeInt:
		return fmt.Sprintf("%d", it.value)
	case itemTypeList:
		builder := strings.Builder{}
		builder.WriteByte('[')
		for idx, sub := range it.values {
			builder.WriteString(sub.String())
			if idx < len(it.values)-1 {
				builder.WriteByte(',')
			}
		}
		builder.WriteByte(']')
		return builder.String()
	default:
		return "NIL"
	}
}

func (it *item) Compare(r *item) int {
	if it.type_ == itemTypeInt && r.type_ == itemTypeInt {
		return cmp.Compare(it.value, r.value)
	}
	if it.type_ == itemTypeList && r.type_ == itemTypeList {
		idx := 0
		for idx < len(it.values) && idx < len(r.values) {
			cmp := it.values[idx].Compare(r.values[idx])
			if cmp != 0 {
				return cmp
			}
			idx++
		}
		return cmp.Compare(len(it.values), len(r.values))
	}
	if it.type_ == itemTypeInt {
		return (&item{type_: itemTypeList, values: []*item{it}}).Compare(r)
	}
	return it.Compare(&item{type_: itemTypeList, values: []*item{r}})
}

type itemType int

const (
	_ itemType = iota
	itemTypeInt
	itemTypeList
)

func fromString(s string) (idx int, it *item) {
	it = &item{}
	idx = 0
	pending := 0

	for idx < len(s) {
		idx++
		switch s[idx-1] {
		case '[':
			it.type_ = itemTypeList
			it.values = make([]*item, 0)
			rIdx := idx
			{
				depth := 1
				for rIdx < len(s) {
					if s[rIdx] == '[' {
						depth++
					} else if s[rIdx] == ']' {
						depth--
						if depth == 0 {
							break
						}
					}
					rIdx++
				}
			}
			for idx < rIdx {
				advance, sub := fromString(s[idx:rIdx])
				it.values = append(it.values, sub)
				idx += advance
			}
		case ']':
			it.value = pending
			if idx < len(s) && s[idx] == ',' {
				idx++
			}
			return
		case ',':
			it.value = pending
			return
		default:
			it.type_ = itemTypeInt
			pending = pending*10 + int(s[idx-1]-'0')
		}
	}

	if pending > 0 {
		it.type_ = itemTypeInt
		it.value = pending
	}

	return
}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)
	p.pairs = make([][2]*item, 0, len(blocks))
	for _, block := range blocks {
		_, it0 := fromString(block[0])
		_, it1 := fromString(block[1])
		p.pairs = append(p.pairs, [2]*item{it0, it1})
	}
}

func (p *p) Solve1() any {
	return util.ReduceIndex(p.pairs, func(acc int, idx int, pair [2]*item) int {
		if pair[0].Compare(pair[1]) < 0 {
			return acc + idx + 1
		}
		return acc
	}, 0)
}

func (p *p) Solve2() any {
	packets := make([]*item, 0, 2*len(p.pairs)+2)
	div1 := &item{
		type_: itemTypeList,
		values: []*item{
			{
				type_:  itemTypeList,
				values: []*item{{type_: itemTypeInt, value: 2}},
			},
		},
	}
	div2 := &item{
		type_: itemTypeList,
		values: []*item{
			{
				type_:  itemTypeList,
				values: []*item{{type_: itemTypeInt, value: 6}},
			},
		},
	}

	for _, pair := range p.pairs {
		packets = append(packets, pair[0], pair[1])
	}
	packets = append(packets, div1, div2)

	slices.SortFunc(packets, func(a, b *item) int { return a.Compare(b) })
	idx1, _ := slices.BinarySearchFunc(packets, div1, func(i1, i2 *item) int { return i1.Compare(i2) })
	idx2, _ := slices.BinarySearchFunc(packets, div2, func(i1, i2 *item) int { return i1.Compare(i2) })

	return (idx1 + 1) * (idx2 + 1)
}

func init() {
	base.Register(13, &p{})
}
