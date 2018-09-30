package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Character interface {
	GetHp() int
	GetAttackPoint() int
	SetHp(hp int)
	SetAttackPoint(attackPoint int)
}

type HeroInterface interface {
	Character
}

type EnemyInterface interface {
	Character
	GetSpecies() string
	SetSpecies(species string)
}

type Hero struct {
	Hp          int
	AttackPoint int
}

func (this *Hero) GetHp() int {
	return (*this).Hp
}

func (this *Hero) GetAttackPoint() int {
	return (*this).AttackPoint
}

func (this *Hero) SetHp(hp int) {
	(*this).Hp = hp
}

func (this *Hero) SetAttackPoint(attackPoint int) {
	(*this).AttackPoint = attackPoint
}

type Enemy struct {
	Hp          int
	AttackPoint int
	Species     string
}

func (this *Enemy) GetSpecies() string {
	return (*this).Species
}

func (this *Enemy) SetSpecies(species string) {
	(*this).Species = species
}

func (this *Enemy) GetHp() int {
	return (*this).Hp
}

func (this *Enemy) GetAttackPoint() int {
	return (*this).AttackPoint
}

func (this *Enemy) SetHp(hp int) {
	(*this).Hp = hp
}

func (this *Enemy) SetAttackPoint(attackPoint int) {
	(*this).AttackPoint = attackPoint
}

type reg_func_interface func([]string, string, *Context)

type Weapon struct {
	Occurrence
	attackPoint int
}

type Food struct {
	Occurrence
	healhPoint int
}

type Field struct {
	range_m        int
	enemy_map      map[int]EnemyInterface
	occurrence_map OccurrenceMap
}

type Occurrence interface {
}

type Context struct {
	hero      HeroInterface
	field     Field
	enemy_map map[string]EnemyInterface
}

type OccurrenceMap struct {
	occurrence_map map[int]Occurrence
}

func isStringMatches(line, regex string) []string {
	r, err := regexp.Compile(regex)
	logErr(err)
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
	fmt.Println("1: " + result)
	line_result_log := strings.Join(line_result, "")
	fmt.Println("2: " + line_result_log)
	return result, line_result
}

type Being interface {
	appendMap(i int, occurrence Occurrence)
}

type Being2 interface {
	get() Occurrence
}

func (e Enemy) get() Enemy {
	return e
}

func (occurrence_map OccurrenceMap) appendMap(i int, occurrence Occurrence) {
	occurrence_map.occurrence_map[i] = occurrence
}

func initContext(context *Context) {
	//	if len(context.enemy_map) == 0 {
	//		context.enemy_map = make(map[string]Enemy)
	//	}
	if len((*context).field.enemy_map) == 0 {
		(*context).field.enemy_map = make(map[int]EnemyInterface)
	}
}
func saveEnemyToField(position int, enemy EnemyInterface, context *Context) {
	(*context).field.enemy_map[position] = enemy
}

func getEnemyFromContext(enemyName string, context *Context) EnemyInterface {
	enemy, ok := (*context).enemy_map[enemyName]
	if !ok {
		fmt.Printf("LAA %q \n", enemyName)
		enemy = new(Enemy)
		enemy.SetSpecies(enemyName)
	}
	return enemy
}

func getInteger(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		result = 0
		logErr(err)
	}
	return result
}

func reg1(info []string, regex string, context *Context) {
	enemyName := info[1]
	position := info[2]
	positionInt := getInteger(position)
	enemytemp := getEnemyFromContext(enemyName, context)
	saveEnemyToField(positionInt, enemytemp, context)
	fmt.Printf("%q %q\n", enemyName, position)
}

func isCharacterHero(characterName string) bool {
	return characterName == "Hero"
}

func updateAttackPoint(character string, attackPointInt int, context *Context) {
	//	if isCharacterHero(character) {
	//		herotemp := context.hero
	//		putPoint(herotemp, attackPointInt)
	//		//herotemp.attackPoint = attackPointInt
	//		context.hero = herotemp
	//	} else {
	//		enemytemp, ok := context.enemy_map[character]
	//		if !ok {
	//			enemytemp = Enemy{}
	//		}
	//		enemytemp.species = character
	//		putPoint(enemytemp, attackPointInt)
	//		//enemytemp.attackPoint = attackPointInt
	//		context.enemy_map[character] = enemytemp
	//	}
	updatePoint(character, attackPointInt, context, PutAttackPoint)
}

func PutHP(m Character, hpInt int) {
	fmt.Printf("PutHP %q\n", hpInt)
	m.SetHp(hpInt)
}
func PutAttackPoint(m Character, attackPointInt int) {
	fmt.Printf("PutAttackPoint %q\n", attackPointInt)
	m.SetAttackPoint(attackPointInt)
}

//func (m *Character) PutHP(hpInt int) {
//	m.hp = hpInt
//}
//func (m *Character) PutAttackPoint(attackPointInt int) {
//	m.attackPoint = attackPointInt
//}

func updatePoint(character string, newPointInt int, context *Context, putPoint func(Character, int)) {
	if isCharacterHero(character) {
		herotemp := context.hero
		putPoint(herotemp, newPointInt)
		//herotemp.hp = hpInt
		context.hero = herotemp
	} else {
		enemytemp, ok := context.enemy_map[character]
		if !ok {
			enemytemp = new(Enemy)
		}
		enemytemp.SetSpecies(character)
		putPoint(enemytemp, newPointInt)
		//enemytemp.hp = hpInt
		enemy := new(Enemy)
		enemy.SetAttackPoint(enemytemp.GetAttackPoint())
		enemy.SetHp(enemytemp.GetHp())
		enemy.SetSpecies(enemytemp.GetSpecies())
		context.enemy_map[character] = enemy
	}
}

func updateHealthPoint(character string, hpInt int, context *Context) {
	//	if isCharacterHero(character) {
	//		herotemp := context.hero
	//		herotemp.hp = hpInt
	//		context.hero = herotemp
	//	} else {
	//		enemytemp, ok := context.enemy_map[character]
	//		if !ok {
	//			enemytemp = Enemy{}
	//		}
	//		enemytemp.species = character
	//		enemytemp.hp = hpInt
	//		context.enemy_map[character] = enemytemp
	//	}
	updatePoint(character, hpInt, context, PutHP)
}

func reg2(info []string, regex string, context *Context) {
	character := info[1]
	attackPoint := info[2]
	fmt.Printf("%q %q\n", character, attackPoint)
	attackPointInt := getInteger(info[2])
	updateAttackPoint(character, attackPointInt, context)
}
func reg3(info []string, regex string, context *Context) {
	fmt.Printf("%q\n", info[1])
	rangeInt := getInteger(info[1])
	context.field.range_m = rangeInt
}
func reg4(info []string, regex string, context *Context) {
	character := info[1]
	hp := info[2]
	hpInt := getInteger(hp)
	fmt.Printf("%q %q\n", character, hp)
	updateHealthPoint(character, hpInt, context)
}
func reg5(info []string, regex string, context *Context) {
	species := info[1]
	fmt.Printf("%q\n", species)

	enemytemp, ok := context.enemy_map[species]
	if !ok {
		enemytemp = new(Enemy)
	}
	enemytemp.SetSpecies(species)
	context.enemy_map[species] = enemytemp
}

var func_map = map[string]reg_func_interface{
	REGEX_1: reg1,
	REGEX_2: reg2,
	REGEX_3: reg3,
	REGEX_4: reg4,
	REGEX_5: reg5,
}

func fillContext(info []string, regex string, context *Context) {
	fmt.Println()
	fmt.Println(regex)
	reg_funct := func_map[regex]
	reg_funct(info, regex, context)
}

func fight(hero *HeroInterface, enemy EnemyInterface) bool {
	result := false
	heroAttackP := (*hero).GetAttackPoint()
	enemyAttackP := enemy.GetAttackPoint()
	enemyHP := enemy.GetHp()
	heroHP := (*hero).GetHp()
	remains := enemyHP % heroAttackP
	if remains != 0 {
		remains -= heroAttackP
	}
	newEnemyHP := enemyHP + remains
	multiplier := newEnemyHP / heroAttackP
	multipliedEnemyAP := multiplier * enemyAttackP
	enemyName := enemy.GetSpecies()
	if heroHP > multipliedEnemyAP {
		heroHP -= multipliedEnemyAP
		(*hero).SetHp(heroHP)
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
func heroFighter(hero *HeroInterface) func(enemy EnemyInterface) bool {
	return func(enemy EnemyInterface) bool {
		return fight(hero, enemy)
	}
}

func main() {
	file, err := os.Open("lines")
	logErr(err)
	defer file.Close()
	var context = Context{}
	context.enemy_map = make(map[string]EnemyInterface)
	initContext(&context)
	context.hero = new(Hero)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		regex, info := whichRegexIsAppropiate(line)
		fillContext(info, regex, &context)
	}
	field := context.field
	isHeroAlive := true
	fmt.Printf(CONST_1, context.hero.GetHp())
	fmt.Printf("Range is %d"+END_LINE, field.range_m)
	var lastIndex int
	fighter := heroFighter(&context.hero)

	for i := 1; i <= field.range_m; i++ {
		enemy, ok := field.enemy_map[i]
		if ok {
			fmt.Printf("Enemy is %q"+END_LINE, enemy.GetSpecies())
			enemy2, ok2 := context.enemy_map[enemy.GetSpecies()]
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
