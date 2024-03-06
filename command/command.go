package command

type CliCommand struct {
    name string
    description string
    callback func()
}

