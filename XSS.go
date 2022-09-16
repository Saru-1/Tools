package main

import (
	"bufio"
	"fmt"
	"github.com/Saru-1/XSS/scanning"
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
