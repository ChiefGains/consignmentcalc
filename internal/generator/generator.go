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
	inventory map[*item]int
	items     []*item
	owed      float32
}

func (l *location) Show() {
	fmt.Println(l.name)

	fmt.Println("\nCurrent inventory:")

	for _, item := range l.items {
		fmt.Println(item.name, ": ", l.inventory[item])
	}

	fmt.Println("\nOutstanding balance for ", l.name, " :")

	fmt.Println(l.owed)
}

func NewItem(name string, cat *category, cost float32, price float32) *item {
	a := &item{name, cat, cost, price}
	return a
}

func NewLocation(name, address string) *location {
	m := make(map[*item]int)
	l := make([]*item, 0)

	a := &location{name, address, m, l, 0}

	return a
}

func NewCategory(tax float32) *category {
	a := &category{tax}

	return a
}

func containsItem(l []*item, x *item) bool {
	contains := false

	for _, i := range l {
		if i == x {
			contains = true
		}
	}
	return contains
}

func (l *location) AddItem(x *item, y int) *location {
	contains := containsItem(l.items, x)

	switch contains {
	case true:
		l.inventory[x] += y
	default:
		l.items = append(l.items, x)
		l.inventory[x] = y
	}

	return l

}

func (l *location) TakeStock(x *item, y int) *location {
	sold := l.inventory[x] - y
	l.owed += float32(sold) * x.price

	return l
}
