package files
//package files to handle the loading and writing of object files,
//as well as the marshaling and unmarshaling of json data

import (
	"code/github.com/consignmentcalc/internal/generator"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Files struct {
	Categories []os.FileInfo
	Items      []os.FileInfo
	Locations  []os.FileInfo
}

type MyData struct {
	Categories []*generator.Category
	Items      []*generator.Item
	Locations  []*generator.Location
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//LoadFiles loads the json object files to be unmarshaled, and calls
//the unmarshalFiles func
func LoadFiles() *MyData {
	categories, err := ioutil.ReadDir("internal/files/categories")
	check(err)
	items, err := ioutil.ReadDir("internal/files/items")
	check(err)
	locations, err := ioutil.ReadDir("internal/files/locations")
	check(err)

	a := &Files{categories, items, locations}

	b := unmarshalFiles(a)

	return b
}

//reads all json files in a set of specified directories
//into a new type *myData object
func unmarshalFiles(f *Files) *MyData {
	dirname := "internal/files"

	cat := make([]*generator.Category, 0)
	for _, file := range f.Categories {
		filename := dirname + "/categories/" + file.Name()
		a := &generator.Category{}
		b, err := ioutil.ReadFile(filename)
		check(err)
		err = json.Unmarshal([]byte(b), a)
		check(err)
		cat = append(cat, a)
	}

	it := make([]*generator.Item, 0)
	for _, file := range f.Items {
		filename := dirname + "/items/" + file.Name()
		a := &generator.Item{}
		b, err := ioutil.ReadFile(filename)
		check(err)
		err = json.Unmarshal([]byte(b), a)
		check(err)
		it = append(it, a)
		a.Show()
	}

	loc := make([]*generator.Location, 0)
	for _, file := range f.Locations {
		filename := dirname + "/locations/" + file.Name()
		a := &generator.Location{}
		b, err := ioutil.ReadFile(filename)
		check(err)
		err = json.Unmarshal([]byte(b), a)
		check(err)
		loc = append(loc, a)
	}

	return &MyData{cat, it, loc}
}

//MarshalFiles takes all the objects in a *myData struct and writes them to json
func MarshalFiles(m *MyData) {
	dirname := "internal/files"
	for _, data := range m.Categories {
		file, err := json.MarshalIndent(data, "", " ")
		check(err)
		filename := dirname + "/categories/" + data.Name + ".json"
		err = ioutil.WriteFile(filename, file, 0644)
		check(err)
	}

	for _, data := range m.Items {
		file, err := json.MarshalIndent(data, "", " ")
		check(err)
		filename := dirname + "/items/" + data.Name + ".json"
		err = ioutil.WriteFile(filename, file, 0644)
		check(err)
	}

	for _, data := range m.Locations {
		file, err := json.MarshalIndent(data, "", " ")
		check(err)
		filename := dirname + "/locations/" + data.Name + ".json"
		err = ioutil.WriteFile(filename, file, 0644)
		check(err)
	}

}

//DataFrame create a *MyData object to hold information
func DataFrame() *MyData {
	cat := make([]*generator.Category, 0)
	it := make([]*generator.Item, 0)
	loc := make([]*generator.Location, 0)

	return &MyData{cat, it, loc}
}

//AddLocation adds a location to type *MyData struct
func (m *MyData) AddLocation(l *generator.Location) {
	m.Locations = append(m.Locations, l)
}

//AddItem adds an item to type *MyData struct
func (m *MyData) AddItem(it *generator.Item) {
	m.Items = append(m.Items, it)
}

//AddCategory adds a category to type *MyData struct
func (m *MyData) AddCategory(cat *generator.Category) {
	m.Categories = append(m.Categories, cat)
}
