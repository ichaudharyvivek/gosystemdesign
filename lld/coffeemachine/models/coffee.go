package models

type Coffee struct {
	ID     string
	Name   string
	Price  int
	Type   CoffeeType
	Recipe *Recipe
}

func NewCoffee(id, name string, price int, ct CoffeeType, recipe *Recipe) *Coffee {
	return &Coffee{id, name, price, ct, recipe}
}

func (c *Coffee) GetIngredients() *Ingredients {
	return c.Recipe.Ingredients
}

func (c *Coffee) GetInstructions() string {
	return c.Recipe.Instructions
}
