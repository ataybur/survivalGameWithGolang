package main

import (
	"game/funcs"
)

func main() {
	lines := funcs.ReadFileIntoLines("resources/lines")
	contextP := funcs.InitContext()
	funcs.FillContext(lines, contextP)
	funcs.Play(contextP)
}
