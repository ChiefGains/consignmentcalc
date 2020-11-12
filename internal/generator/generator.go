package generator

import "fmt"

type Category struct { //mostly for data analysis
	name string
	tax  float32
}

type Item struct { //item that is being sold for consignment
	name    string
	cat     *Category
	cost    float32
	price   float32
	dropped int
	sold    int
}

type Location struct { //place where consignment items are being sold
	name      string
	address   string
	inventory map[*Item]*salesData //stores the item quantity and sales
	items     []*Item              //I have this in addition to the map so that items in a location can eventually be sorted
	owed      float32
	paid      float32
}

type salesData struct {
	quantity int
	sold     int
}

func (l *Location) Show() { //shows current information for location
	fmt.Println(l.name)

	fmt.Println("\nCurrent inventory:")
	for _, item := range l.items { //again, this will eventually be sorted
		fmt.Println(item.name, ": ", l.inventory[item].quantity)
	}

	fmt.Println("\nTotal Sales:")
	for _, item := range l.items {
		fmt.Println(item.name, ": ", l.inventory[item].sold)
	}

	fmt.Println("\nOutstanding balance for ", l.name, " :")
	fmt.Println(l.owed)

	fmt.Println("\nTotal Paid for ", l.name, " :")
	fmt.Println(l.paid)
}

func NewItem(name string, cat *Category, cost float32, price float32) *Item {
	a := &Item{name, cat, cost, price, 0, 0}
	return a
}

func NewLocation(name, address string) *Location { //creates a new location to populate with items
	m := make(map[*Item]*salesData)
	l := make([]*Item, 0)

	a := &Location{name, address, m, l, 0, 0}

	return a
}

func NewCategory(s string, tax float32) *Category { //in case you want to track data for a certain group of items
	a := &Category{s, tax} //this isn't really necessary for retail consignment in PA, but what the heck

	return a
}

func containsItem(l []*Item, x *Item) bool { //checks whether a location has been stocked with an item
	contains := false

	for _, i := range l { //if the item is found in the slice, return "true"
		if i == x {
			contains = true
		}
	}
	return contains
}

func (l *Location) AddItem(x *Item, y int) { //adds inventory to the location, unfortunately one item at a time currently
	contains := containsItem(l.items, x) //check to see if the item is currently stocked

	switch contains {
	case true:
		l.inventory[x].quantity += y //increase the inventory count if the item exists at the store
	default:
		l.items = append(l.items, x) //if not, add the item to the list of existing items first

		a := &salesData{0, 0} //instantiate sales data for item at this location
		l.inventory[x] = a
		l.inventory[x].quantity = y
	}

	x.dropped += y

}

func (l *Location) TakeStock(x *Item, y int) *Location { //sold is original inventory - current inventory
	sold := l.inventory[x].quantity - y
	l.owed += float32(sold) * x.price
	l.inventory[x].quantity = y
	l.inventory[x].sold += sold

	x.sold += sold

	return l
}

func (l *Location) TakePayment(x float32) {
	l.owed -= x
	l.paid += x
}

func (x *Item) Show() {
	fmt.Println(x.name)
	fmt.Println("\nProduction Cost:\n", x.cost)
	fmt.Println("\nWholesale Price:\n", x.price)
	fmt.Println("\nTotal Dropped:\n", x.dropped)
	fmt.Println("\nTotal Sold:\n", x.sold)
}
