// character
package structs

type CharacterI interface {
	GetAttackPoint() int
	GetHp() int
	SetHp(int)
	SetAttackPoint(int)
}

type Character struct {
	attackPoint int
	hp          int
}

func (this Character) GetHp() int {
	return this.hp
}

func (this Character) GetAttackPoint() int {
	return this.attackPoint
}

func (this *Character) SetHp(hp int) {
	this.hp = hp
}

func (this *Character) SetAttackPoint(attackPoint int) {
	this.attackPoint = attackPoint
}
