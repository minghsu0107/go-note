package command

import "fmt"

// concrete command interface
type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

// concrete command inplementation
// tell the reciever what it wants to do 
// without knowing how reciever implements
func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

// concrete command inplementation
func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

// reciever
type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}

func (b *Box) RunAll() {
	b.button1.Execute()
	b.button2.Execute()
}
