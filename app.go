package main

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Test search function returns a greeting for the given name
// func (a *App) Search(query string) string {
// 	return fmt.Sprintf("Hello %s, It's showtime!", query)
// }

// // This function queries the data base in the createDatabase() function.
// func (a *App) Search(query string) *sql.Rows {

// 	// List out the directories.
// 	dbDirName := "tutorDB"
// 	dbFileName := "QandA.db"

// 	fn := filepath.Join(dbDirName, dbFileName)

// 	// Open the db through 'sqlite.'
// 	db, err := sql.Open("sqlite", fn)
// 	if err != nil {
// 		log.Fatalf("sql.Open failed: %s", err)
// 	}

// 	defer db.Close()

// 	// Return everything from the Q&A table.
// 	// Gotta remember to use triple quotes since, well, the original entries are surrounded in quotes.
// 	theQuery := "select * from qat;"
// 	// fmt.Println(theQuery)

// 	rows, err := db.Query(theQuery)

// 	//////////////////////////////////////////////////
// 	// Everything below this seems to crash my app...
// 	//////////////////////////////////////////////////

// 	// This code below seems to be crashing the Wails application.
// 	if err != nil {
// 		log.Fatalf("db.Query failed: %s", err)
// 	}

// 	// This code definitely crashes my app.
// 	defer rows.Close()

// 	var question string
// 	var answer string

// 	// Returns a printed string here.
// 	for rows.Next() {
// 		var id int
// 		var question string
// 		var answer string
// 		rows.Scan(&id, &question, &answer)
// 		fmt.Printf("%s: %s\n", question, answer)
// 	}

// 	// return fmt.Sprintf("%s: %s\n", question, query)
// 	return rows
// }

// func (a *App) Search(query string) string {
// 	dbDirName := "tutorDB"
// 	dbFileName := "QandA.db"

// 	fn := filepath.Join(dbDirName, dbFileName)

// 	db, err := sql.Open("sqlite", fn)
// 	if err != nil {
// 		return "asd"
// 	}
// 	defer db.Close()

// 	theQuery := "select * from qat WHERE question = 'type';"

// 	rows, err := db.Query(theQuery)
// 	if err != nil {
// 		return "asd"
// 	}

// 	var question string
// 	var answer string

// 	// Returns a printed string here.
// 	for rows.Next() {
// 		var id int
// 		var question string
// 		var answer string
// 		rows.Scan(&id, &question, &answer)
// 		fmt.Printf("%s: %s\n", question, answer)
// 	}

// 	// Defer closing the rows
// 	defer rows.Close()

// 	return fmt.Sprintf("%s: %s\n", question, answer)

// 	// return rows, nil
// }

// func (a *App) Search(query string) string {
// 	dbDirName := "tutorDB"
// 	dbFileName := "QandA.db"

// 	fn := filepath.Join(dbDirName, dbFileName)

// 	db, err := sql.Open("sqlite", fn)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	defer db.Close()

// 	theQuery := "select * from qat;"

// 	rows, err := db.Query(theQuery)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	defer rows.Close() // Close the rows immediately after opening

// 	var question string
// 	var answer string

// 	// Accumulate results in a buffer
// 	var resultBuffer bytes.Buffer

// 	for rows.Next() {
// 		var id int
// 		rows.Scan(&id, &question, &answer)
// 		resultBuffer.WriteString(fmt.Sprintf("%s: %s\n", question, answer))
// 	}

// 	return resultBuffer.String()
// }

////////////////////////
// Notable Attempt
////////////////////////

func (a *App) Search(query string) string {

	// Because the folder Wales actually builds from is build\bin, we have to backtrack to the root directory to get to tutorDb\QandA.
	relPath := "..\\..\\tutorDB\\QandA.db"
	fn, _ := filepath.Abs(relPath)

	db, err := sql.Open("sqlite", fn)
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	theQuery := "SELECT * FROM qat WHERE question = ?;"

	rows, err := db.Query(theQuery, query)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	var resultBuffer strings.Builder // Use strings.Builder for better performance

	for rows.Next() {
		var id int
		var question, answer string
		err := rows.Scan(&id, &question, &answer)
		if err != nil {
			return "Error scanning rows."
		}
		resultBuffer.WriteString(fmt.Sprintf("%s: %s\n", question, answer))
	}

	return resultBuffer.String()
}
