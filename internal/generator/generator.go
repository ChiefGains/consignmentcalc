package generator

import "fmt"

type category struct {
	tax float32
}

type item struct {
	name    string
	cat     *category
	cost    float32
	price   float32
	dropped int
	sold    int
}

type location struct {
	name      string
	address   string
	inventory map[*item]*salesData //stores the item and quantity, and eventually sales data
	items     []*item              //I have this in addition to the map so that items in a location can eventually be sorted
	owed      float32
	paid      float32
}

type salesData struct {
	quantity int
	sold     int
}

func (l *location) Show() { //shows current information for location
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

func NewItem(name string, cat *category, cost float32, price float32) *item {
	a := &item{name, cat, cost, price, 0, 0}
	return a
}

func NewLocation(name, address string) *location { //creates a new location to populate with items
	m := make(map[*item]*salesData)
	l := make([]*item, 0)

	a := &location{name, address, m, l, 0, 0}

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

func (l *location) AddItem(x *item, y int) { //adds inventory to the location, unfortunately one item at a time currently
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

func (l *location) TakeStock(x *item, y int) *location { //sold is original inventory - current inventory
	sold := l.inventory[x].quantity - y
	l.owed += float32(sold) * x.price
	l.inventory[x].quantity = y
	l.inventory[x].sold += sold

	x.sold += sold

	return l
}

func (l *location) TakePayment(x float32) {
	l.owed -= x
	l.paid += x
}

func (x *item) Show() {
	fmt.Println(x.name)
	fmt.Println("\nProduction Cost:\n", x.cost)
	fmt.Println("\nWholesale Price:\n", x.price)
	fmt.Println("\nTotal Dropped:\n", x.dropped)
	fmt.Println("\nTotal Sold:\n", x.sold)
}
