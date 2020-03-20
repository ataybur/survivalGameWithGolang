// funcs project funcs.go
package funcs

import (
	"bufio"
	"fmt"
	"game/structs"
	"game/utils"
	"log"
	"os"
)

const (
	END_LINE = "\n"
	CONST_1  = "Hero started journey with %d HP!" + END_LINE
	CONST_2  = "Hero defeated %s with %d HP remaining" + END_LINE
	CONST_3  = "Survived" + END_LINE
	CONST_4  = "%s defeated Hero with %d HP remaining" + END_LINE
	CONST_5  = "Hero is Dead!! Last seen at position %d!!" + END_LINE
)

func Play(context *structs.Context) {
	field := context.Field
	isHeroAlive := true
	LogHeroStartsJourney(context)
	LogRangeIs(context, field)
	var lastIndex int
	fighter := HeroFighter(&context.Hero, context)
	for i, enemy := range field.Enemy_map {
		LogEnemyIs(context, enemy)
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
		LogSurvived(context)
	} else {
		LogDead(context, lastIndex)
	}
}

func ReadFileIntoLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	utils.LogErr(err)
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

func ReadFileIntoLines2(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func fight(context *structs.Context, hero *structs.Hero, enemy structs.Enemy) bool {
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
		LogHeroDefeated(context, enemyName, heroHP)
		result = true
	} else {
		remains := heroHP % enemyAttackP
		if remains != 0 {
			remains -= enemyAttackP
		}
		newHeroHP := heroHP + remains
		multiplier := newHeroHP / enemyAttackP
		multipliedHeroAP := multiplier * heroAttackP
		LogEnemyDefeated(context, enemyName, enemyHP-multipliedHeroAP)
	}
	return result
}

func LogEnemyDefeated(c *structs.Context, enemyName string, enemyHP int) {
	c.Logger.Log(CONST_4, enemyName, enemyHP)
}

func LogHeroDefeated(c *structs.Context, enemyName string, heroHP int) {
	c.Logger.Log(CONST_2, enemyName, heroHP)
}

func LogHeroStartsJourney(c *structs.Context) {
	c.Logger.Log(CONST_1, c.Hero.GetHp())
}

func LogRangeIs(c *structs.Context, f structs.Field) {
	c.Logger.Log("Range is %d"+END_LINE, f.Range_m)
}

func LogEnemyIs(c *structs.Context, e structs.Enemy) {
	c.Logger.Log("Enemy is %q"+END_LINE, e.Species)
}

func LogSurvived(c *structs.Context) {
	c.Logger.Log(CONST_3)
}

func LogDead(c *structs.Context, lastIndex int) {
	c.Logger.Log(CONST_5, lastIndex)
}

func HeroFighter(hero *structs.Hero, c *structs.Context) func(enemy structs.Enemy) bool {
	return func(enemy structs.Enemy) bool {
		return fight(c, hero, enemy)
	}
}
