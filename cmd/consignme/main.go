package main

import (
	"code/github.com/consignmentcalc/internal/generator"
)

func main() {
	food := generator.NewCategory(.06)
	burgerpalace := generator.NewLocation("address")
	hamburger := generator.NewItem("hamburger", food, 1, 2)
	cheeseburger := generator.NewItem("cheeseburger", food, 1.5, 2.75)

	burgerpalace.AddItem(hamburger, 5)
	burgerpalace.Show()
	burgerpalace.AddItem(hamburger, 3)
	burgerpalace.Show()
	burgerpalace.AddItem(cheeseburger, 4)
	burgerpalace.Show()
}
