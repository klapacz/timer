package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type result struct {
	result string
	index   int
}

func run(c chan result, delay int, index int, cmd string) {
	for {
		out, err := exec.Command("sh", "-c", cmd).Output()

		if err != nil {
			log.Fatal(err)
		}

		c <- result{index: index, result: string(out)}
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

type cmd struct {
	delay int
	cmd   string
}

func main() {
	cmds := []cmd{
		{10, "node --no-warnings /tmp/szymon/index.js"},
		{1, "date \"+%Y-%m-%d %l:%M:%S %p\""},
	}

	c := make(chan result)
	results := make([]string, len(cmds))

	for i, e := range cmds {
		results[i] = "loading…"
		go run(c, e.delay, i, e.cmd)
	}

	for {
		msg := <-c
		results[msg.index] = strings.Replace(msg.result, "\n", "", -1)
		fmt.Println(strings.Join(results, " | "))
	}
}
