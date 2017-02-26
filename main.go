package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// oddExec is return 3n + 1
func oddExec(n int) int {
	return 3*n + 1
}

// evenExec is return n / 2
func evenExec(n int) int {
	return n / 2
}

func main() {
	var fp *os.File
	var err error

	if len(os.Args) != 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		t := scanner.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			break
		}

		cnt := 0
		for {
			if n == 1 {
				break
			}
			if n%2 == 0 {
				n = evenExec(n)
			} else {
				n = oddExec(n)
			}
			cnt++
		}
		fmt.Printf("%d\n", cnt)
	}
}
