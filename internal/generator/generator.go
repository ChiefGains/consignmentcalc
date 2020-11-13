package generator

import "fmt"

//Category is a struct that holds data about what
//category an *Item belongs to, mostly for data
//analysis
type Category struct {
	Name string
	Tax  float32
}

//Item that is being sold for consignment
type Item struct {
	Name    string
	Cat     *Category
	Cost    float32
	Price   float32
	Dropped int
	Sold    int
}

//Location is a struct for places where
//consignment objects are being sold
type Location struct {
	Name      string                `json:"Name"`
	Address   string                `json:"Address"`
	Inventory map[string]*SalesData `json:"Inventory"` //stores the item quantity and sales
	Items     []string              `json:"Items"`     //I have this in addition to the map so that items in a location can eventually be sorted
	Owed      float32               `json:"Owed"`
	Paid      float32               `json:"Paid"`
}

//SalesData - exactly what it sounds like
type SalesData struct {
	Quantity int `json:"Quantity"`
	Sold     int `json:"Sold"`
}

//Show displays current information about a type *Location struct
func (l *Location) Show() {
	fmt.Println(l.Name)

	fmt.Println("\nCurrent inventory:")
	for _, item := range l.Items { //again, this will eventually be sorted
		fmt.Println(item, ": ", l.Inventory[item].Quantity)
	}

	fmt.Println("\nTotal Sales:")
	for _, item := range l.Items {
		fmt.Println(item, ": ", l.Inventory[item].Sold)
	}

	fmt.Println("\nOutstanding balance for ", l.Name, " :")
	fmt.Println(l.Owed)

	fmt.Println("\nTotal Paid for ", l.Name, " :")
	fmt.Println(l.Paid)
}

//NewItem creates a new type *Item struct to drop at locations
func NewItem(name string, cat *Category, cost float32, price float32) *Item {
	a := &Item{name, cat, cost, price, 0, 0}
	return a
}

//NewLocation creates a new type *Location struct to populate with type *Item
func NewLocation(name, address string) *Location {
	m := make(map[string]*SalesData)
	l := make([]string, 0)

	a := &Location{name, address, m, l, 0, 0}

	return a
}

//NewCategory allows the user to create a new
//category to label items with and keep track
//of things like sales tax or sales by category
func NewCategory(s string, tax float32) *Category {
	a := &Category{s, tax} //this isn't really necessary for retail consignment in PA, but what the heck

	return a
}

//containsItem checks whether a location has been stocked with an item
//and returns a bool
func containsItem(l []string, x string) bool {
	contains := false

	for _, i := range l { //if the item is found in the slice, return "true"
		if i == x {
			contains = true
		}
	}
	return contains
}

//AddItem does what you might expect, adds an item and its
//quantity to a specified location's inventory
func (l *Location) AddItem(x *Item, y int) {
	contains := containsItem(l.Items, x.Name) //check to see if the item is currently stocked

	switch contains {
	case true:
		l.Inventory[x.Name].Quantity += y //increase the inventory count if the item exists at the store
	default:
		l.Items = append(l.Items, x.Name) //if not, add the item to the list of existing items first

		a := &SalesData{0, 0} //instantiate sales data for item at this location
		l.Inventory[x.Name] = a
		l.Inventory[x.Name].Quantity = y
	}

	x.Dropped += y

}

//TakeStock allows a user to enter in the current inventory
//and compare it to what was originally dropped off at that
//location, giving total number sold
func (l *Location) TakeStock(x *Item, y int) *Location { //sold is original inventory - current inventory
	sold := l.Inventory[x.Name].Quantity - y
	l.Owed += float32(sold) * x.Price
	l.Inventory[x.Name].Quantity = y
	l.Inventory[x.Name].Sold += sold

	x.Sold += sold

	return l
}

//TakePayment allows the user to enter in a payment
//given by a location, which is automatically subtracted
//from the total amount owed
func (l *Location) TakePayment(x float32) {
	l.Owed -= x
	l.Paid += x
}

//Show displays information about type *Item
func (x *Item) Show() {
	fmt.Println(x.Name)
	fmt.Println("\nProduction Cost:\n", x.Cost)
	fmt.Println("\nWholesale Price:\n", x.Price)
	fmt.Println("\nTotal Dropped:\n", x.Dropped)
	fmt.Println("\nTotal Sold:\n", x.Sold)
}
