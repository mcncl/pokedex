package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
    appName = "pokedex"
)

type cliCommand struct {
    name string
    description string
    callback func()
}

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
    commands := map[string]cliCommand{
        "help": {
            name: "help",
            description: "Show help",
            callback: func() {
                fmt.Print("help called\n")
            },
        },
        "quit": {
            name: "quit",
            description: "Quit the application",
            callback: commandExit,
        },
    }

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
