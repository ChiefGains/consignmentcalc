package generator

import "fmt"

type category struct {
	tax float32
}

type item struct {
	name  string
	cat   *category
	cost  float32
	price float32
}

type location struct {
	name      string
	address   string
	inventory map[*item]int //stores the item and quantity, and eventually sales data
	items     []*item       //I have this in addition to the map so that items in a location can eventually be sorted
	owed      float32
}

func (l *location) Show() { //shows current information for location
	fmt.Println(l.name)

	fmt.Println("\nCurrent inventory:")

	for _, item := range l.items { //again, this will eventually be sorted
		fmt.Println(item.name, ": ", l.inventory[item])
	}

	fmt.Println("\nOutstanding balance for ", l.name, " :")

	fmt.Println(l.owed)
}

func NewItem(name string, cat *category, cost float32, price float32) *item {
	a := &item{name, cat, cost, price}
	return a
}

func NewLocation(name, address string) *location { //creates a new location to populate with items
	m := make(map[*item]int)
	l := make([]*item, 0)

	a := &location{name, address, m, l, 0}

	return a
}

func NewCategory(tax float32) *category { //in case you want to track data for a certain group of items
	a := &category{tax} //this isn't really necessary for retail consignment in PA, but what the heck

	return a
}

func containsItem(l []*item, x *item) bool { //checks whether a location has been stocked with an item
	contains := false

	for _, i := range l { //if the item is found in the slice, return "true"
		if i == x {
			contains = true
		}
	}
	return contains
}

func (l *location) AddItem(x *item, y int) *location { //adds inventory to the location, unfortunately one item at a time currently
	contains := containsItem(l.items, x) //check to see if the item is currently stocked

	switch contains {
	case true:
		l.inventory[x] += y //increase the inventory count if the item exists at the store
	default:
		l.items = append(l.items, x) //if not, add the item to the list of existing items first
		l.inventory[x] = y
	}

	return l

}

func (l *location) TakeStock(x *item, y int) *location { //sold is original inventory - current inventory
	sold := l.inventory[x] - y
	l.owed += float32(sold) * x.price

	return l
}
