package main

import (
	"github.com/chiefgains/consignmentcalc/internal/files"
	"github.com/chiefgains/consignmentcalc/internal/generator"
)

func main() {

	initialize()

	//readstuff()

}

func initialize() {

	food := generator.NewCategory("food", .06)
	retail := generator.NewCategory("retail", .06)

	hamburger := generator.NewItem("Hamburger", food, 1, 2)
	cheeseburger := generator.NewItem("Cheeseburger", food, 1.2, 2.5)
	toy := generator.NewItem("Toy", retail, .25, 1.5)

	BurgerPalace := generator.NewLocation("Burger Palace", "Park Place")
	KrustyBurger := generator.NewLocation("Krusty Burger", "Springfield")

	BurgerPalace.AddItem(hamburger, 4)
	BurgerPalace.AddItem(cheeseburger, 5)
	BurgerPalace.AddItem(toy, 8)

	KrustyBurger.AddItem(hamburger, 10)
	KrustyBurger.AddItem(cheeseburger, 20)
	KrustyBurger.AddItem(toy, 25)

	df := files.DataFrame()
	df.AddCategory(food)
	df.AddCategory(retail)
	df.AddItem(hamburger)
	df.AddItem(cheeseburger)
	df.AddItem(toy)
	df.AddLocation(BurgerPalace)
	df.AddLocation(KrustyBurger)

	//KrustyBurger.Show()
	//BurgerPalace.Show()

	//df.Items[0].Show()
	//df.Locations[0].Show()

	files.MarshalFiles(df)

}

func readstuff() {
	df := files.LoadFiles()

	for _, item := range df.Items {
		item.Show()
	}
}
