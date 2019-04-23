// field
package structs

type Field struct {
	Range_m   int
	Enemy_map map[int]Enemy
}

func (this *Field) InitEnemyMap() {
	this.Enemy_map = make(map[int]Enemy)
}

func (this *Field) SetRangeM(i int) {
	this.Range_m = i
}
