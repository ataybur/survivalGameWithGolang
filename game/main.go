package main

import (
	"bufio"
	"game/funcs"
	"game/structs"
	"log"
	"os"
)

type reg_func_interface func([]string, string, *structs.Context)

type Weapon struct {
	Occurrence
	attackPoint int
}

type Food struct {
	Occurrence
	healhPoint int
}

type Occurrence interface {
}

type OccurrenceMap struct {
	occurrence_map map[int]Occurrence
}

type Being interface {
	appendMap(i int, occurrence Occurrence)
}

type Being2 interface {
	get() Occurrence
}

func (occurrence_map OccurrenceMap) appendMap(i int, occurrence Occurrence) {
	occurrence_map.occurrence_map[i] = occurrence
}

func initContext() structs.Context {
	context := structs.Context{}
	context.InitEnemyMap()
	field := context.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
	context.InitHero()
	return context
}

func main() {
	file, err := os.Open("resources/lines")
	funcs.LogErr(err)
	defer file.Close()
	context := initContext()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		funcs.FillContext(line, &context)
	}
	field := context.Field
	isHeroAlive := true
	funcs.LogHeroStartsJourney(context)
	funcs.LogRangeIs(field)
	var lastIndex int
	fighter := funcs.HeroFighter(&context.Hero)

	for i := 1; i <= field.Range_m; i++ {
		enemy, ok := field.Enemy_map[i]
		if ok {
			funcs.LogEnemyIs(enemy)
			enemy2, ok2 := context.Enemy_map[enemy.Species]
			if ok2 {
				isHeroAlive = fighter(enemy2)
				if !isHeroAlive {
					lastIndex = i
					break
				}
			}
		}
	}
	if isHeroAlive {
		funcs.LogSurvived()
	} else {
		funcs.LogDead(lastIndex)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
