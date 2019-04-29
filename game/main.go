package main

import (
	"game/funcs"
	"game/structs"
)

func main() {
	lines := funcs.ReadFileIntoLines("resources/lines")
	contextP := &structs.Context{}
	contextP.Init()
	funcs.FillContext(lines, contextP)
	funcs.Play(contextP)
}
