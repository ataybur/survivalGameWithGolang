// funcs project funcs.go
package funcs

import (
	"bufio"
	"fmt"
	"game/structs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isStringMatches(line, regex string) []string {
	r, err := regexp.Compile(regex)
	LogErr(err)
	result := r.FindStringSubmatch(line)
	return result
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

func Play(context *structs.Context) {
	field := context.Field
	isHeroAlive := true
	LogHeroStartsJourney(context)
	LogRangeIs(field)
	var lastIndex int
	fighter := HeroFighter(&context.Hero)
	for i, enemy := range field.Enemy_map {
		LogEnemyIs(enemy)
		enemy2, ok2 := context.Enemy_map[enemy.Species]
		if ok2 {
			isHeroAlive = fighter(enemy2)
			if !isHeroAlive {
				lastIndex = i
				break
			}
		}

	}
	if isHeroAlive {
		LogSurvived()
	} else {
		LogDead(lastIndex)
	}
}

func InitContext() *structs.Context {
	context := structs.Context{}
	context.InitEnemyMap()
	field := context.Field
	if len(field.Enemy_map) == 0 {
		field.InitEnemyMap()
	}
	context.InitHero()
	return &context
}

func ReadFileIntoLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	LogErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func FillContext(lines []string, context *structs.Context) {
	for _, line := range lines {
		regex, info := whichRegexIsAppropiate(line)
		fmt.Println()
		reg_funct := func_map[regex]
		reg_funct(info, regex, context)
	}
}

func fight(hero *structs.Hero, enemy structs.Enemy) bool {
	result := false
	heroAttackP := hero.GetAttackPoint()
	enemyAttackP := enemy.GetAttackPoint()
	enemyHP := enemy.GetHp()
	heroHP := hero.GetHp()
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

func LogHeroStartsJourney(c *structs.Context) {
	fmt.Printf(CONST_1, c.Hero.GetHp())
}

func LogRangeIs(f structs.Field) {
	fmt.Printf("Range is %d"+END_LINE, f.Range_m)
}

func LogEnemyIs(e structs.Enemy) {
	fmt.Printf("Enemy is %q"+END_LINE, e.Species)
}

func LogSurvived() {
	fmt.Println(CONST_3)
}

func LogDead(lastIndex int) {
	fmt.Printf(CONST_5, lastIndex)
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
	fmt.Println()
	fmt.Printf("%q %q\n", enemyName, position)
}

func reg2(info []string, regex string, context *structs.Context) {
	character := info[1]
	attackPoint := info[2]
	fmt.Println()
	fmt.Printf("%q %q\n", character, attackPoint)
	attackPointInt := getInteger(info[2])
	updateAttackPoint(character, attackPointInt, context)
}
func reg3(info []string, regex string, context *structs.Context) {
	fmt.Println()
	fmt.Printf("%q\n", info[1])
	rangeInt := getInteger(info[1])
	field := &context.Field
	field.SetRangeM(rangeInt)
}
func reg4(info []string, regex string, context *structs.Context) {
	character := info[1]
	hp := info[2]
	hpInt := getInteger(hp)
	fmt.Println()
	fmt.Printf("%q %q\n", character, hp)
	updateHealthPoint(character, hpInt, context)
}
func reg5(info []string, regex string, context *structs.Context) {
	species := info[1]
	fmt.Println()
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

func updateHeroPoint(character string, newPointInt int, context *structs.Context, putPoint func(structs.CharacterI, int)) {
	herotemp := &context.Hero
	putPoint(herotemp, newPointInt)
	context.SetHero(*herotemp)
}

func updateEnemyPoint(character string, newPointInt int, context *structs.Context, putPoint func(structs.CharacterI, int)) {
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
	enemy.SetAttackPoint(enemytemp.GetAttackPoint())
	enemy.SetHp(enemytemp.GetHp())
	enemy.SetSpecies(enemytemp.Species)
	context.Enemy_map[character] = enemy

}

func updateAttackPoint(character string, attackPointInt int, context *structs.Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, attackPointInt, context, PutAttackPoint)
	} else {
		updateEnemyPoint(character, attackPointInt, context, PutAttackPoint)
	}
}

func updateHealthPoint(character string, hpInt int, context *structs.Context) {
	if isCharacterHero(character) {
		updateHeroPoint(character, hpInt, context, PutHP)
	} else {
		updateEnemyPoint(character, hpInt, context, PutHP)
	}
}

func PutHP(m structs.CharacterI, hpInt int) {
	fmt.Printf("PutHP %q\n", hpInt)
	m.SetHp(hpInt)
}
func PutAttackPoint(m structs.CharacterI, attackPointInt int) {
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
