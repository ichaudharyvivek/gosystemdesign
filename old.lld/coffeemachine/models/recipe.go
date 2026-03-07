package models

type Ingredients struct {
	CoffeeBeans int
	Sugar       int
	Milk        int
	Cream       int
}

type Recipe struct {
	Ingredients  *Ingredients
	Instructions string
}

func NewRecipe(in *Ingredients, instr string) *Recipe {
	return &Recipe{in, instr}
}
