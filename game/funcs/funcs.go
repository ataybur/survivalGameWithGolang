// funcs project funcs.go
package funcs

import (
	"fmt"
	"game/structs"
	"log"
	"regexp"
	"strconv"
	"strings"
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

var func_map = map[string]reg_func_interface{
	REGEX_1: reg1,
	REGEX_2: reg2,
	REGEX_3: reg3,
	REGEX_4: reg4,
	REGEX_5: reg5,
}

type reg_func_interface func([]string, string, *structs.Context)

func isStringMatches(line, regex string) []string {
	r, err := regexp.Compile(regex)
	LogErr(err)
	result := r.FindStringSubmatch(line)
	return result
}

func WhichRegexIsAppropiate(line string) (string, []string) {
	result := ""
	var line_result []string
	for _, REGEX := range REGEX_ARR {
		line_result = isStringMatches(line, REGEX)
		if len(line_result) != 0 {
			result = REGEX
			break
		}
	}
	fmt.Println("1: " + result)
	line_result_log := strings.Join(line_result, "")
	fmt.Println("2: " + line_result_log)
	return result, line_result
}

func FillContext(info []string, regex string, context *structs.Context) {
	fmt.Println()
	fmt.Println(regex)
	reg_funct := func_map[regex]
	reg_funct(info, regex, context)
}

func fight(hero *structs.Hero, enemy structs.Enemy) bool {
	result := false
	heroAttackP := hero.AttackPoint
	enemyAttackP := enemy.AttackPoint
	enemyHP := enemy.Hp
	heroHP := hero.Hp
	remains := enemyHP % heroAttackP
	if remains != 0 {
		remains -= heroAttackP
	}
	newEnemyHP := enemyHP + remains
	multiplier := newEnemyHP / heroAttackP
	multipliedEnemyAP := multiplier * enemyAttackP
	enemyName := enemy.Species
	if heroHP > multipliedEnemyAP {
		heroHP -= multipliedEnemyAP
		hero.SetHp(heroHP)
		fmt.Printf(CONST_2, enemyName, heroHP)
		result = true
	} else {
		remains := heroHP % enemyAttackP
		if remains != 0 {
			remains -= enemyAttackP
		}
		newHeroHP := heroHP + remains
		multiplier := newHeroHP / enemyAttackP
		multipliedHeroAP := multiplier * heroAttackP
		fmt.Printf(CONST_4, enemyName, enemyHP-multipliedHeroAP)
	}
	return result
}
func HeroFighter(hero *structs.Hero) func(enemy structs.Enemy) bool {
	return func(enemy structs.Enemy) bool {
		return fight(hero, enemy)
	}
}

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func reg1(info []string, regex string, context *structs.Context) {
	enemyName := info[1]
	position := info[2]
	positionInt := getInteger(position)
	enemytemp := getEnemyFromContext(enemyName, context)
	saveEnemyToField(positionInt, enemytemp, context)
	fmt.Printf("%q %q\n", enemyName, position)
}

func reg2(info []string, regex string, context *structs.Context) {
	character := info[1]
	attackPoint := info[2]
	fmt.Printf("%q %q\n", character, attackPoint)
	attackPointInt := getInteger(info[2])
	updateAttackPoint(character, attackPointInt, context)
}
func reg3(info []string, regex string, context *structs.Context) {
	fmt.Printf("%q\n", info[1])
	rangeInt := getInteger(info[1])
	field := &context.Field
	field.SetRangeM(rangeInt)
}
func reg4(info []string, regex string, context *structs.Context) {
	character := info[1]
	hp := info[2]
	hpInt := getInteger(hp)
	fmt.Printf("%q %q\n", character, hp)
	updateHealthPoint(character, hpInt, context)
}
func reg5(info []string, regex string, context *structs.Context) {
	species := info[1]
	fmt.Printf("%q\n", species)
	enemymap := context.Enemy_map
	enemytemp, _ := enemymap[species]
	enemytemp.SetSpecies(species)
	context.Enemy_map[species] = enemytemp
}

func getInteger(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		result = 0
		LogErr(err)
	}
	return result
}

func isCharacterHero(characterName string) bool {
	return characterName == "Hero"
}

func updateHeroPoint(character string, newPointInt int, context *structs.Context, putPoint func(*structs.Hero, int)) {
	herotemp := &context.Hero
	putPoint(herotemp, newPointInt)
	context.SetHero(*herotemp)
}

func updateEnemyPoint(character string, newPointInt int, context *structs.Context, putPoint func(*structs.Enemy, int)) {
	enemymap := &context.Enemy_map
	enemytemp, ok := (*enemymap)[character]
	if !ok {
		enemytemp = structs.Enemy{}
	}
	enemytempP := &enemytemp
	enemytempP.SetSpecies(character)
	putPoint(enemytempP, newPointInt)
	//enemytemp.hp = hpInt
	enemy := structs.Enemy{}
	enemy.SetAttackPoint(enemytemp.AttackPoint)
	enemy.SetHp(enemytemp.Hp)
	enemy.SetSpecies(enemytemp.Species)
	context.Enemy_map[character] = enemy

}

func updateAttackPoint(character string, attackPointInt int, context *structs.Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, attackPointInt, context, PutHeroAttackPoint)
	} else {
		updateEnemyPoint(character, attackPointInt, context, PutEnemyAttackPoint)
	}
}

func updateHealthPoint(character string, hpInt int, context *structs.Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, hpInt, context, PutHeroHP)
	} else {
		updateEnemyPoint(character, hpInt, context, PutEnemyHP)
	}
}

func PutHeroHP(m *structs.Hero, hpInt int) {
	fmt.Printf("PutHP %q\n", hpInt)
	m.SetHp(hpInt)
}
func PutHeroAttackPoint(m *structs.Hero, attackPointInt int) {
	fmt.Printf("PutAttackPoint %q\n", attackPointInt)
	m.SetAttackPoint(attackPointInt)
}

func PutEnemyHP(m *structs.Enemy, hpInt int) {
	fmt.Printf("PutHP %q\n", hpInt)
	m.SetHp(hpInt)
}

func PutEnemyAttackPoint(m *structs.Enemy, attackPointInt int) {
	fmt.Printf("PutAttackPoint %q\n", attackPointInt)
	m.SetAttackPoint(attackPointInt)
}

func getEnemyFromContext(enemyName string, context *structs.Context) structs.Enemy {
	enemy, ok := context.Enemy_map[enemyName]
	if !ok {
		enemy = structs.Enemy{}
		enemy.SetSpecies(enemyName)
	}
	return enemy
}

func saveEnemyToField(position int, enemy structs.Enemy, context *structs.Context) {
	field := &context.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
	field.Enemy_map[position] = enemy
}
