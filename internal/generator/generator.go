package generator

import "fmt"

type category struct {
	tax float32
}

type item struct {
	cat   category
	cost  float32
	price float32
}

type location struct {
	address   string
	inventory map[item]int
	items     []item
}

func (l *location) show() {
	for _, item := range l.items {
		fmt.Println(l.inventory[item])
	}
}

func newItem(cat category, cost float32, price float32) *item {
	a := &item{cat, cost, price}
	return a
}

func newLocation(s string) *location {
	m := make(map[item]int)
	l := make([]item)

	a := &location{s, m, l}

	return a
}

func containsItem(l []item, x item) bool {
	contains := false

	for _, i := range l {
		if i == x {
			contains = true
		}
	}
	return contains
}

func (l *location) addItem(x item, y int) *location {
	contains := containsItem(l.items, x)

	switch contains {
	case true:
		l.inventory[x] += y
	default:
		l.items = append(l.items, x)
		l.inventory[x] = y
	}

}
