package base

import "fmt"

type Puzzle interface {
	Init(filename string)
	Solve1() int64
	Solve2() int64
}

var puzzles = make(map[int]Puzzle)

func Register(day int, puzzle Puzzle) {
	if _, exist := puzzles[day]; exist {
		panic(fmt.Errorf("Duplicate registration for puzzle %d", day))
	}
	puzzles[day] = puzzle
}

func Get(day int) (Puzzle, error) {
	puzzle, exist := puzzles[day]
	if !exist {
		return nil, fmt.Errorf("No puzzle registered for day %d", day)
	}
	return puzzle, nil
}
