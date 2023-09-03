package commands

import devices "Command/Devices"

type Command interface {
	Execute()
}

type OnCommand struct {
	Device devices.Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}

//

type OffCommand struct {
	Device devices.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}
