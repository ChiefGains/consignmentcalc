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
	Categories []generator.Category
	Items      []generator.Item
	Locations  []generator.Location
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Loadfiles() *files { //loads the json item files to be unmarshaled
	categories, err := ioutil.ReadDir("internal/files/categories")
	check(err)
	items, err := ioutil.ReadDir("internal/files/items")
	check(err)
	locations, err := ioutil.ReadDir("internal/files/locations")
	check(err)

	a := &files{categories, items, locations}

	return a
}

func UnmarshalFiles(f *files) *myData {
	cat := make([]generator.Category)
	for _, file := range f.Categories {
		a := &generator.Category{}
		_ = json.Unmarshal([]byte(file), a)
		cat = append(cat, a)
	}

	it := make([]generator.Item)
	for _, file := range f.Items {
		a := &generator.Item{}
		_ = json.Unmarshal([]byte(file), a)
		it = append(it, a)
	}

	loc := make([]generator.Location)
	for _, file := range f.Locations {
		a := &generator.Location{}
		_ = json.Unmarshal([]byte(file), a)
		loc = append(loc, a)
	}
}
