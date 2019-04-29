package main

import (
	"game/funcs"
	"game/structs"
)

func main() {
	lines := funcs.ReadFileIntoLines("resources/lines")
	contextP := &structs.Context{}
	contextP.Init()
	contextP.Fill(lines)
	funcs.Play(contextP)
}
