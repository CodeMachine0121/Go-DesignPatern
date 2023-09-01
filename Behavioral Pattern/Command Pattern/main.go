package main

import (
	buttons "Command/Buttons"
	commands "Command/Commands"
	devices "Command/Devices"
)

func main() {

	Light := &devices.Light{}

	OnCommand := &commands.OnCommand{
		Device: Light,
	}

	OffCommand := &commands.OffCommand{
		Device: Light,
	}

	onButton := &buttons.Button{Command: OnCommand}
	offButton := &buttons.Button{Command: OffCommand}

	offButton.Press()
	onButton.Press()

}
