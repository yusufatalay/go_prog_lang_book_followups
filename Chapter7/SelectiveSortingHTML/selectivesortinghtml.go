// exercise 7.9
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"sort"
	"time"
)

// Using the dummy data and the functions that provided by the book

// Track is holding the attributes of a track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

// fieldName is a global variable for the sorting key
var fieldName string

// fieldNames will be holding all the struct field's names
var fieldNames []string

// TemplateStruct will hold necessary values for the html template
type TemplateStruct struct {
	Tracks    []*Track
	FieldList []string
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// SelectiveSort is same with the customSort struct
type SelectiveSort struct {
	t        []*Track
	lessFunc func(x, y *Track) bool
}

// implementing the necessary methods for the sort.Interface
func (x SelectiveSort) Len() int           { return len(x.t) }
func (x SelectiveSort) Less(i, j int) bool { return x.lessFunc(x.t[i], x.t[j]) }
func (x SelectiveSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var selectiveLess = func(x, y *Track) bool {
	xField := getField(x, fieldName)
	yField := getField(y, fieldName)

	// using x's type is enough since they both will be same
	xVal := fmt.Sprintf("%v", xField)
	yVal := fmt.Sprintf("%v", yField)

	if xVal != yVal {
		return xVal < yVal
	}
	return false
}

// getField returns struct's field's value and it's type  by name
func getField(T *Track, field string) interface{} {
	e := reflect.ValueOf(T).Elem()
	varValue := e.FieldByName(field).Interface()
	return varValue
}
func getFieldNames(T *Track) {
	if len(fieldNames) > 0 {
		return
	}
	e := reflect.ValueOf(T).Elem()

	for i := 0; i < e.NumField(); i++ {
		fieldNames = append(fieldNames, e.Type().Field(i).Name)
	}
}

func printTracks(w http.ResponseWriter, r *http.Request) {

	// populate the fieldNames
	getFieldNames(tracks[0])
	ts := TemplateStruct{Tracks: tracks, FieldList: fieldNames}

	t, _ := template.ParseFiles("template.html")

	key := r.URL.Query()
	if key.Get("key") == "" {
		fieldName = "Title"
	} else {
		fieldName = key.Get("key")
		key.Set("key", "")
	}

	sort.Sort(SelectiveSort{tracks, selectiveLess})
	fmt.Println(t.Execute(w, ts))

}

func main() {
	http.HandleFunc("/", printTracks)
	http.ListenAndServe(":8000", nil)
}
