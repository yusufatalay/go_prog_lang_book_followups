package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Comic struct has the same keys for xkcd comics
type Comic struct {
	Index      int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

const baseURL = "http://xkcd.com/"
const urlpostfix = "/info.0.json"
const fileext = ".json"

// fileExists checks if given file exists adn returns accordingly
func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getComic(index int) (*Comic, error) {

	resp, err := http.Get(baseURL + strconv.Itoa(index) + urlpostfix)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		log.Fatalf("\n[!]That was all of the comics\n")
		resp.Body.Close()
		errorTheEnd := errors.New("No more comics left to search")
		return nil, errorTheEnd
	}

	var comic Comic

	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &comic, nil

}

// createTable function will execute a sqlite command to create this comictable
func createTable(db *sql.DB) {
	createComicTableSQL := "CREATE TABLE comic (indexComic INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, date TEXT, title TEXT,transcript TEXT)"

	// using log package instead of fmt because of the timestamp ability of log
	log.Println("[!]Creating comic table")

	// preparing the sql statement above to execute
	statement, err := db.Prepare(createComicTableSQL)
	log.Println("after db.Prep")
	if err != nil {
		log.Println("there is an error")
		log.Fatalf("%v", err)
	}
	// execute the sql statement
	log.Println("before exec")
	statement.Exec()

	log.Println("[+]Comic table created")
}

// insertComic function inserts the given props of the comic to the given database
func insertComic(db *sql.DB, date, title, transcript string) {
	log.Println("[!]Inserting the comic record")

	// sql command for inserting a record to the database
	insertComicSQL := "INSERT INTO comic(date, title, transcript) VALUES(?, ?, ?)"
	// prepare the command above for execution
	statement, err := db.Prepare(insertComicSQL)

	if err != nil {
		log.Fatalf("%v", err)
	}

	_, err = statement.Exec(date, title, transcript)

	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println("[+]Record inserted to the database")
}

func displayAllComics(db *sql.DB) {

	row, err := db.Query("SELECT * FROM comic ORDER BY indexComic ASC")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	defer row.Close()

	// iterate for all rows in the table
	for row.Next() {
		var indexComic int
		var date string
		var title string
		var transcript string

		row.Scan(&indexComic, &date, &title, &transcript)

		fmt.Printf("\n\n%s\n%d\t%s\n%s\n", date, indexComic, title, transcript)
	}
}

func displaySelectedComic(db *sql.DB, substring string) {
	queryString := "SELECT * FROM comic WHERE title LIKE " + "'%" + substring + "%'" + " OR transcript LIKE " + "'%" + substring + "%'" + " ORDER BY indexComic ASC"
	row, err := db.Query(queryString)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	defer row.Close()
	for row.Next() {
		var indexComic int
		var date string
		var title string
		var transcript string

		row.Scan(&indexComic, &date, &title, &transcript)

		fmt.Printf("\n\n%s\n%d\t%s\n%s\n", date, indexComic, title, transcript)
	}
}

func main() {

	// creating the database file

	//	log.Printf("[!]Creating database file\n")
	//	dbfile, err := os.Create("comics.db")
	//
	//	if err != nil {
	//		fmt.Errorf("%v\n", err)
	//		return
	//	}
	//
	//	dbfile.Close()
	//	log.Printf("[+]Database file created --> comics.db\n")
	//
	//	// open the file as a sqlite database file
	comicsDB, _ := sql.Open("sqlite3", "./comics.db")
	//
	//	// close the databse file after everything is done
	//
	//	// creating a table
	//	createTable(comicsDB)
	//
	//	// insert the found comics
	//	index := 1
	//	for {
	//		comic, err := getComic(index)
	//
	//		if err != nil {
	//			log.Fatalf("%v", err)
	//			break
	//		}
	//
	//		insertComic(comicsDB, comic.Day+"/"+comic.Month+"/"+comic.Year, comic.Title, comic.Transcript)
	//		index++
	//	}
	displayAllComics(comicsDB)
	defer comicsDB.Close()
}
