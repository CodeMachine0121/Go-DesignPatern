package interfaces

type InterfaceNode interface {
	Print(string)
	Clone() InterfaceNode
}
