// enemy
package structs

type Enemy struct {
	Character
	Species string
}

func (this *Enemy) SetSpecies(species string) {
	this.Species = species
}
