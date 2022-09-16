package main

import (
	"bufio"
	"fmt"
	"m/scanning"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		target = sc.Text()
		targets = append(targets, target)
	}
	for i := range targets {
		_, _ := scanning.Scan(targets[i], options, strconv.Itoa(i))
	}
}
