package main

import (
	"advent2022/util"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	args := os.Args
	util.Assert(len(args) == 2, "Usage: generate <day>")

	day := util.Must(strconv.Atoi(args[1]))
	util.Assert(day >= 1 && day <= 25, "Day must be between 1 and 25")

	dayName := fmt.Sprintf("%02d", day)
	util.Must(1, os.Mkdir(dayName, 0755))

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		cmd := exec.Command("aoc",
			"download", "-I", "-d", strconv.Itoa(day), "-y2022", "-i", dayName+"/input.txt")
		util.Must(2, cmd.Run())
		util.Must(os.OpenFile(dayName+"/example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)).Close()
	}()

	go func() {
		defer wg.Done()
		lines := util.GetLines("generate/solution.go.tpl")
		file := util.Must(os.OpenFile(dayName+"/solution.go", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644))
		defer file.Close()
		for _, line := range lines {
			line = strings.ReplaceAll(line, "{{.DayString}}", dayName)
			line = strings.ReplaceAll(line, "{{.Day}}", strconv.Itoa(day))
			util.Must(file.WriteString(line + "\n"))
		}
	}()

	go func() {
		defer wg.Done()

		lines := util.GetLines("main.go")

		imStart := slices.Index(lines, "import (")
		util.Assert(imStart != -1, "Could not find import section")

		imEnd := slices.Index(lines[imStart:], ")") + imStart
		util.Assert(imEnd != imStart-1, "Could not find end of import section")

		lines = slices.Insert(lines, imEnd, "\t_ \"advent2022/"+dayName+"\"")

		file := util.Must(os.OpenFile("main.go", os.O_WRONLY|os.O_TRUNC, 0644))
		defer file.Close()
		for _, line := range lines {
			util.Must(file.WriteString(line + "\n"))
		}
	}()

	wg.Wait()
}
