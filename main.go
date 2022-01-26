package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type result struct {
	result string
	index  int
}

type cmd struct {
	Interval int
	Cmd      string
}

type config struct {
	Separator string
	Cmds      []cmd
}

func run(channel chan result, index int, c cmd) {
	for {
		out, err := exec.Command("sh", "-c", c.Cmd).Output()

		if err != nil {
			log.Fatal(err)
		}

		channel <- result{index: index, result: string(out)}
		time.Sleep(time.Duration(c.Interval) * time.Second)
	}
}

// TODO: check if valid config provided
func readConfig(path string) config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	conf := config{}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}
	return conf
}

var configPath string

func main() {
	if len(os.Args) == 2 {
		configPath = os.Args[1]
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		configPath = path.Join(home, ".config/timer.yaml")
	}
	conf := readConfig(configPath)

	channel := make(chan result)
	results := make([]string, len(conf.Cmds))

	for i, e := range conf.Cmds {
		results[i] = "loading…"
		go run(channel, i, e)
	}

	for {
		msg := <-channel
		results[msg.index] = strings.Replace(msg.result, "\n", "", -1)
		fmt.Println(strings.Join(results, conf.Separator))
	}
}
