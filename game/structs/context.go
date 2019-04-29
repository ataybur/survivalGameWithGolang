// context
package structs

type Context struct {
	Hero      Hero
	Field     Field
	Enemy_map map[string]Enemy
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
}
