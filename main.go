package main

import (
	"bufio"
	"fmt"
	"github.com/mcncl/pokedex/command"
	"github.com/mcncl/pokedex/help"
	"os"
	"strings"
	"time"
)

const (
	appName = "pokedex"
)

func readInput() *bufio.Reader {
	input := bufio.NewReader(os.Stdin)
	return input
}

func commandExit() {
	fmt.Print("Quitting Pokedex...\n")
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func main() {
	for {
		input := readInput()
		fmt.Printf("%s > ", appName)
		text, _ := input.ReadString('\n')
		for _, command := range commands {
			if strings.TrimSpace(text) == command.name {
				command.callback()
			}
		}
	}
}
