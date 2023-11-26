package compiler

type Compiler struct {
	Backend  string
	Frontend string

	HasDB bool
}
