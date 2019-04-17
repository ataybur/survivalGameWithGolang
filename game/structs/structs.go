// structs project structs.go
package structs

type Character struct {
	AttackPoint int
	Hp          int
}

type Hero struct {
	Character
}

type Enemy struct {
	Character
	Species string
}

type Context struct {
	Hero      Hero
	Field     Field
	Enemy_map map[string]Enemy
}

type Field struct {
	Range_m   int
	Enemy_map map[int]Enemy
}

func (this *Enemy) SetSpecies(species string) {
	this.Species = species
}

func (this *Character) SetHp(hp int) {
	this.Hp = hp
}

func (this *Character) SetAttackPoint(attackPoint int) {
	this.AttackPoint = attackPoint
}

func (this *Field) InitEnemyMap() {
	this.Enemy_map = make(map[int]Enemy)
}

func (this *Context) SetHero(i Hero) {
	this.Hero = i
}

func (this *Field) SetRangeM(i int) {
	this.Range_m = i
}

func (this *Context) InitEnemyMap() {
	this.Enemy_map = make(map[string]Enemy)
}

func (this *Context) InitHero() {
	this.Hero = Hero{}
}
