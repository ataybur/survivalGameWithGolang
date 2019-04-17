package main

import (
	"bufio"
	"fmt"
	"funcs"
	"log"
	"os"
	"structs"
)

const (
	REGEX_1  = "There is a ([a-zA-Z ]+) at position ([0-9]+)"
	REGEX_2  = "([a-zA-Z ]+) attack is ([0-9]+)"
	REGEX_3  = "Resources are ([0-9]+) meters away"
	REGEX_4  = "([a-zA-Z ]+) has ([0-9]+) hp"
	REGEX_5  = "([a-zA-Z ]+) is Enemy"
	END_LINE = "\n"
	CONST_1  = "Hero started journey with %d HP!" + END_LINE
	CONST_2  = "Hero defeated %s with %d HP remaining" + END_LINE
	CONST_3  = "Survived" + END_LINE
	CONST_4  = "%s defeated Hero with %d HP remaining" + END_LINE
	CONST_5  = "Hero is Dead!! Last seen at position %d!!" + END_LINE
)

var REGEX_ARR = [5]string{REGEX_1, REGEX_2, REGEX_3, REGEX_4, REGEX_5}

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

func initContext(context *structs.Context) {
	//	if len(context.enemy_map) == 0 {
	//		context.enemy_map = make(map[string]Enemy)
	//	}
	field := context.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
}

func main() {
	file, err := os.Open("lines")
	funcs.LogErr(err)
	defer file.Close()
	var context = structs.Context{}
	context.InitEnemyMap()
	initContext(&context)
	context.InitHero()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		regex, info := funcs.WhichRegexIsAppropiate(line)
		funcs.FillContext(info, regex, &context)
	}
	field := context.Field
	isHeroAlive := true
	fmt.Printf(CONST_1, context.Hero.Hp)
	fmt.Printf("Range is %d"+END_LINE, field.Range_m)
	var lastIndex int
	fighter := funcs.HeroFighter(&context.Hero)

	for i := 1; i <= field.Range_m; i++ {
		enemy, ok := field.Enemy_map[i]
		if ok {
			fmt.Printf("Enemy is %q"+END_LINE, enemy.Species)
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
		fmt.Println(CONST_3)
	} else {
		fmt.Printf(CONST_5, lastIndex)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
