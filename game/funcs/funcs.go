// funcs project funcs.go
package funcs

import (
	"bufio"
	"fmt"
	"game/structs"
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
