package main

import (
	"code/github.com/consignmentcalc/internal/generator"
)

func main() {
	food := generator.NewCategory(.06)
	burgerpalace := generator.NewLocation("Burger Palace", "address")
	hamburger := generator.NewItem("hamburger", food, 1, 2)
	cheeseburger := generator.NewItem("cheeseburger", food, 1.5, 2.75)

	burgerpalace.AddItem(hamburger, 5)
	burgerpalace.AddItem(cheeseburger, 4)
	burgerpalace.Show()
	burgerpalace.TakeStock(hamburger, 3)
	burgerpalace.TakeStock(cheeseburger, 0)
	burgerpalace.Show()
}
