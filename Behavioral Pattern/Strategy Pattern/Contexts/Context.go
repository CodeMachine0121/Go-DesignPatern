package Contexts

import "strategy/Strategies"

type Context struct {
	strategy Strategies.Strategy
}

func (c *Context) SetStrategy(strategy Strategies.Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}

func NewContext() *Context {
	return &Context{}
}
