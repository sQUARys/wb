package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (toggleOn *ToggleOnCommand) Execute() string {
	return toggleOn.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (toggleOff *ToggleOffCommand) Execute() string {
	return toggleOff.receiver.ToggleOff()
}

type Receiver struct {
}

func (r *Receiver) ToggleOn() string {
	return "On toggle"
}

func (r *Receiver) ToggleOff() string {
	return "Off toggle"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) addCommand(command Command) {
	i.commands = append(i.commands, command)
}
func (i *Invoker) deleteLastCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}
func (i *Invoker) Execute() string {
	var res string
	for _, comm := range i.commands {
		res += comm.Execute() + "\n"
	}
	return res
}

func main() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.addCommand(&ToggleOnCommand{receiver: receiver})
	invoker.addCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()
	fmt.Println(result)
}
