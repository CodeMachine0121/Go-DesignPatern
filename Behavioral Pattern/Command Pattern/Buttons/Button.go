package buttons

import commands "Command/Commands"

type Button struct {
	Command commands.Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
