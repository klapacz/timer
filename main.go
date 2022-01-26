package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type result struct {
	result string;
	name string;
}

func run(c chan result, delay int, cmd string, args []string) {
	for {
		out, err := exec.Command(cmd, args...).Output()

		if err != nil {
			log.Fatal(err)
		}

		c<-result{ name: cmd, result: string(out)}

		time.Sleep(time.Duration(delay) * time.Second)
	}
}

// type cmd struct {
// 	cmd string;
// 	args []string;
// }

func main() {
	c := make(chan result)
	results := map[string]string{}
	go run(c, 5, "node", []string{"--no-warnings", "/tmp/szymon/index.js"})
	go run(c, 1, "date", []string{"+%Y-%m-%d %l:%M:%S %p"})

	for {
		msg := <- c
		results[msg.name] = strings.Replace(msg.result, "\n", "", -1)

		final := []string{}
		for _, result := range results {
			final = append(final, result)
		}
		fmt.Println(strings.Join(final, " | "))
	}
}
