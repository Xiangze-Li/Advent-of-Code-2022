package puzzle16

import (
	"advent2022/base"
	"advent2022/util"
	"cmp"
	"maps"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type p struct {
	rate map[string]int
	tran map[string][]string
	adj  map[[2]string]int
}

type path struct {
	flow  int
	nodes []string
}

func (p path) clone() path {
	return path{p.flow, slices.Clone(p.nodes)}
}

func (p *p) Init(filename string) {
	p.rate = make(map[string]int)
	p.tran = make(map[string][]string)
	p.adj = make(map[[2]string]int)

	lines := util.GetLines(filename)
	re := regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=(\d+); tunnel(?:s)? lead(?:s)? to valve(?:s)? (.+)$`)

	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		name := m[1]
		rate := util.Must(strconv.Atoi(m[2]))
		tran := strings.Split(m[3], ", ")
		if rate > 0 {
			p.rate[name] = rate
		}
		p.tran[name] = tran
	}

	for from, trans := range p.tran {
		conn := util.Associative(trans, func(to string) (string, bool) { return to, true })
		for to := range p.tran {
			if from == to {
				p.adj[[2]string{from, to}] = 0
			} else if conn[to] {
				p.adj[[2]string{from, to}] = 1
			} else {
				p.adj[[2]string{from, to}] = math.MaxInt / 2
			}
		}
	}

	for mid := range p.tran {
		for from := range p.tran {
			for to := range p.tran {
				p.adj[[2]string{from, to}] = min(
					p.adj[[2]string{from, to}],
					p.adj[[2]string{from, mid}]+p.adj[[2]string{mid, to}],
				)
			}
		}
	}
}

func (p *p) dfs(curr string, time int, pathh path, visited map[string]bool) []path {
	paths := []path{pathh}

	for next, rate := range p.rate {
		newTime := time - p.adj[[2]string{curr, next}] - 1
		if visited[next] || newTime <= 0 {
			continue
		}
		newVis := maps.Clone(visited)
		newVis[next] = true
		newPath := pathh.clone()
		newPath.flow += rate * newTime
		newPath.nodes = append(newPath.nodes, next)
		paths = append(paths, p.dfs(next, newTime, newPath, newVis)...)
	}

	return paths
}

func (p *p) Solve1() any {
	paths := p.dfs("AA", 30, path{}, map[string]bool{})
	slices.SortFunc(paths, func(a, b path) int {
		return -cmp.Compare(a.flow, b.flow)
	})
	return paths[0].flow
}

func (p *p) Solve2() any {
	paths := p.dfs("AA", 26, path{}, map[string]bool{})
	var max int

	for _, path := range paths {
		if len(path.nodes) == 0 {
			continue
		}

		vis := util.ToVis(path.nodes)

		for _, pp := range paths {
			flow := path.flow + pp.flow
			if flow <= max || len(pp.nodes) == 0 {
				continue
			}
			if util.Reduce(pp.nodes, func(dup bool, node string) bool { return dup || vis[node] }, false) {
				continue
			}
			max = flow
		}
	}

	return max
}

func init() {
	base.Register(16, &p{})
}
