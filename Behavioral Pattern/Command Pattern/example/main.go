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

	offButtonjiera := &buttons.Button{Command: OffCommand}

	offButtonjiera.Press()
	onButton.Press()

}
