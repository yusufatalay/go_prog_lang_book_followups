// exercise 7.8
package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"
	"time"
)

// Using the dummy data and the functions that provided by the book

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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // Calculate the column widths and print table
}

// selectiveSort is same with the customSort struct
type selectiveSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// implementing the necessary methods for the sort.Interface
func (x selectiveSort) Len() int           { return len(x.t) }
func (x selectiveSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x selectiveSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var selectiveLess = func(x, y *Track) bool {
	xField := getField(x, fieldName)
	yField := getField(y, fieldName)

	// using x's type is enough since they both will be same

	//	ftype := fmt.Sprint("%v", xReflectType)

	xVal := fmt.Sprintf("%v", xField)
	yVal := fmt.Sprintf("%v", yField)

	if xVal != yVal {
		return xVal < yVal
	}
	return false
}

// getField returns struct's field's value and it's type  by name
func getField(t *Track, field string) interface{} {
	e := reflect.ValueOf(t).Elem()
	varValue := e.FieldByName(field).Interface()
	return varValue
}

func main() {
	fieldName = "Length"
	sort.Sort(selectiveSort{tracks, selectiveLess})
	printTracks(tracks)
}
