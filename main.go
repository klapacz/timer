package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type result struct {
	result string
	index  int
}

func run(channel chan result, interval int, index int, cmd string) {
	for {
		out, err := exec.Command("sh", "-c", cmd).Output()

		if err != nil {
			log.Fatal(err)
		}

		channel <- result{index: index, result: string(out)}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

type cmd struct {
	Interval int
	Cmd   string
}

type config struct {
	Separator string
	Cmds []cmd
}


func main() {
	conf := `
separator: ' | '
cmds:
  - cmd: node --no-warnings /tmp/szymon/index.js
    interval: 5
  - cmd: date "+%Y-%m-%d %l:%M:%S %p"
    interval: 1
`
	c := config{}
	err := yaml.Unmarshal([]byte(conf), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	channel := make(chan result)
	results := make([]string, len(c.Cmds))

	for i, e := range c.Cmds {
		results[i] = "loading…"
		go run(channel, e.Interval, i, e.Cmd)
	}

	for {
		msg := <-channel
		results[msg.index] = strings.Replace(msg.result, "\n", "", -1)
		fmt.Println(strings.Join(results, c.Separator))
	}
}
