package main

import (
	"code/github.com/consignmentcalc/internal/generator"
)

func main() {

	food := generator.NewCategory("food", .06)
	hamburger := generator.NewItem("Hamburger", food, 1, 2)
	cheeseburger := generator.NewItem("Cheeseburger", food, 1.25, 2.5)

	BurgerPalace := generator.NewLocation("Burger Palace", "1000 Hamburgler Blvd")

	BurgerPalace.AddItem(hamburger, 5)
	BurgerPalace.AddItem(cheeseburger, 8)

	BurgerPalace.TakeStock(hamburger, 2)
	BurgerPalace.TakeStock(cheeseburger, 3)
	BurgerPalace.AddItem(hamburger, 3)
	BurgerPalace.AddItem(cheeseburger, 5)

	BurgerPalace.TakeStock(hamburger, 1)
	BurgerPalace.TakeStock(cheeseburger, 4)
	BurgerPalace.Show()

	BurgerPalace.TakePayment(20)
	BurgerPalace.Show()
	hamburger.Show()
	cheeseburger.Show()

}
