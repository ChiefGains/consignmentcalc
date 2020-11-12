package files

import (
	"code/github.com/consignmentcalc/internal/generator"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//package to handle the loading and writing of object files,
//as well as the marshaling and unmarshaling of json data

type files struct {
	Categories []os.FileInfo
	Items      []os.FileInfo
	Locations  []os.FileInfo
}

type myData struct {
	Categories []*generator.Category
	Items      []*generator.Item
	Locations  []*generator.Location
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//loads the json object files to be unmarshaled, and calls
//the unmarshalFiles func
func Loadfiles() *myData {
	categories, err := ioutil.ReadDir("internal/files/categories")
	check(err)
	items, err := ioutil.ReadDir("internal/files/items")
	check(err)
	locations, err := ioutil.ReadDir("internal/files/locations")
	check(err)

	a := &files{categories, items, locations}

	b := unmarshalFiles(a)

	return b
}

//reads all json files in a set of specified directories
//into a new type *myData object
func unmarshalFiles(f *files) *myData {
	cat := make([]*generator.Category, 0)
	for _, file := range f.Categories {
		a := &generator.Category{}
		b, _ := ioutil.ReadFile(file.Name())
		_ = json.Unmarshal([]byte(b), a)
		cat = append(cat, a)
	}

	it := make([]*generator.Item, 0)
	for _, file := range f.Items {
		a := &generator.Item{}
		b, _ := ioutil.ReadFile(file.Name())
		_ = json.Unmarshal([]byte(b), a)
		it = append(it, a)
	}

	loc := make([]*generator.Location, 0)
	for _, file := range f.Locations {
		a := &generator.Location{}
		b, _ := ioutil.ReadFile(file.Name())
		_ = json.Unmarshal([]byte(b), a)
		loc = append(loc, a)
	}

	return &myData{cat, it, loc}
}

//takes all the objects in a *myData struct and writes them to json
func MarshalFiles(m *myData) {
	dirname := "code/github.com/consignmentcalc/internal/files"
	for _, data := range m.Categories {
		file, _ := json.MarshalIndent(data, "", " ")
		filename := dirname + "/categories/" + data.Name + ".json"
		_ = ioutil.WriteFile(filename, file, 0644)
	}

	for _, data := range m.Items {
		file, _ := json.MarshalIndent(data, "", " ")
		filename := dirname + "/items/" + data.Name + ".json"
		_ = ioutil.WriteFile(filename, file, 0644)
	}

	for _, data := range m.Locations {
		file, _ := json.MarshalIndent(data, "", " ")
		filename := dirname + "/locations/" + data.Name + ".json"
		_ = ioutil.WriteFile(filename, file, 0644)
	}

}

//create a *myData object to hold information
func DataFrame() *myData {
	cat := make([]*generator.Category, 0)
	it := make([]*generator.Item, 0)
	loc := make([]*generator.Location, 0)

	return &myData{cat, it, loc}
}

//adds a location to type *myData struct
func (m *myData) AddLocation(l *generator.Location) {
	m.Locations = append(m.Locations, l)
}

//adds an item to type *myData struct
func (m *myData) AddItem(it *generator.Item) {
	m.Items = append(m.Items, it)
}

//adds a category to type *myData struct
func (m *myData) AddCategory(cat *generator.Category) {
	m.Categories = append(m.Categories, cat)
}
