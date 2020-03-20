// context
package structs

import (
	"fmt"
	"game/utils"
	"strings"
)

type Logger struct {
	log string
}

func (l *Logger) Log(_const string, args ...interface{}) {
	l.log += fmt.Sprintf(_const, args...)
}

type Context struct {
	Hero      Hero
	Field     Field
	Enemy_map map[string]Enemy
	Logger    *Logger
}

func (this Context) GetLog() string {
	return this.Logger.log
}

func (this *Context) SetHero(i Hero) {
	this.Hero = i
}

func (this *Context) InitEnemyMap() {
	this.Enemy_map = make(map[string]Enemy)
}

func (this *Context) InitHero() {
	this.Hero = Hero{}
}

func (this *Context) Init() {
	this.InitEnemyMap()
	field := this.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
	this.InitHero()
	this.Logger = new(Logger)
}

func (this *Context) Fill(lines []string) {
	for _, line := range lines {
		if line != "" {
			regex, info := whichRegexIsAppropiate(line)
			fmt.Printf("Line:%s, Regex: %s, Info: %v \n", line, regex, info)
			reg_funct := func_map[regex]
			reg_funct(info, regex, this)
		}
	}
}

func whichRegexIsAppropiate(line string) (string, []string) {
	result := ""
	var line_result []string
	for _, REGEX := range REGEX_ARR {
		line_result = isStringMatches(line, REGEX)
		if len(line_result) != 0 {
			result = REGEX
			break
		}
	}
	line_result_log := strings.Join(line_result, "")
	fmt.Println(line_result_log)
	return result, line_result
}

func reg1(info []string, regex string, context *Context) {
	enemyName := info[1]
	position := info[2]
	positionInt := utils.GetInteger(position)
	enemytemp := getEnemyFromContext(enemyName, context)
	saveEnemyToField(positionInt, enemytemp, context)
	fmt.Println()
	fmt.Printf("%q %q\n", enemyName, position)
}

func reg2(info []string, regex string, context *Context) {
	character := info[1]
	attackPoint := info[2]
	fmt.Println()
	fmt.Printf("%q %q\n", character, attackPoint)
	attackPointInt := utils.GetInteger(info[2])
	updateAttackPoint(character, attackPointInt, context)
}
func reg3(info []string, regex string, context *Context) {
	fmt.Println()
	fmt.Printf("%q\n", info[1])
	rangeInt := utils.GetInteger(info[1])
	field := &context.Field
	field.SetRangeM(rangeInt)
}
func reg4(info []string, regex string, context *Context) {
	character := info[1]
	hp := info[2]
	hpInt := utils.GetInteger(hp)
	fmt.Println()
	fmt.Printf("%q %q\n", character, hp)
	updateHealthPoint(character, hpInt, context)
}
func reg5(info []string, regex string, context *Context) {
	species := info[1]
	fmt.Println()
	fmt.Printf("%q\n", species)
	enemymap := context.Enemy_map
	enemytemp, _ := enemymap[species]
	enemytemp.SetSpecies(species)
	context.Enemy_map[species] = enemytemp
}

func getEnemyFromContext(enemyName string, context *Context) Enemy {
	enemy, ok := context.Enemy_map[enemyName]
	if !ok {
		enemy = Enemy{}
		enemy.SetSpecies(enemyName)
	}
	return enemy
}

func saveEnemyToField(position int, enemy Enemy, context *Context) {
	field := &context.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
	field.Enemy_map[position] = enemy
}

func updateAttackPoint(character string, attackPointInt int, context *Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, attackPointInt, context, PutAttackPoint)
	} else {
		updateEnemyPoint(character, attackPointInt, context, PutAttackPoint)
	}
}

func updateHealthPoint(character string, hpInt int, context *Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, hpInt, context, PutHP)
	} else {
		updateEnemyPoint(character, hpInt, context, PutHP)
	}
}

func PutHP(m CharacterI, hpInt int) {
	fmt.Printf("PutHP %q\n", hpInt)
	m.SetHp(hpInt)
}
func PutAttackPoint(m CharacterI, attackPointInt int) {
	fmt.Printf("PutAttackPoint %q\n", attackPointInt)
	m.SetAttackPoint(attackPointInt)
}

func isCharacterHero(characterName string) bool {
	return characterName == "Hero"
}

func updateHeroPoint(character string, newPointInt int, context *Context, putPoint func(CharacterI, int)) {
	herotemp := &context.Hero
	putPoint(herotemp, newPointInt)
	context.SetHero(*herotemp)
}

func updateEnemyPoint(character string, newPointInt int, context *Context, putPoint func(CharacterI, int)) {
	enemymap := &context.Enemy_map
	enemytemp, ok := (*enemymap)[character]
	if !ok {
		enemytemp = Enemy{}
	}
	enemytempP := &enemytemp
	enemytempP.SetSpecies(character)
	putPoint(enemytempP, newPointInt)
	enemy := Enemy{}
	enemy.SetAttackPoint(enemytemp.GetAttackPoint())
	enemy.SetHp(enemytemp.GetHp())
	enemy.SetSpecies(enemytemp.Species)
	context.Enemy_map[character] = enemy

}
