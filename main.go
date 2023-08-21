package main

import (
	"database/sql"
	"embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "459svelte",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	// invokes user function 'createDatabase' below.
	if err := createDatabase(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readData() [][]string {
	csvFileName := "QandA.csv"
	csvInput, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalf("open %s failed: %s", csvFileName, err)
	}
	defer csvInput.Close()

	csvReader := csv.NewReader(csvInput)
	csvReader.FieldsPerRecord = -1

	fmt.Printf("Reading %s\n", csvFileName)
	allRecords, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("csvReader.ReadAll failed: %s", err)
	}

	fmt.Printf("%d records read from csv file %s:\n",
		len(allRecords), csvFileName)
	fmt.Println("--------------------------------------------------")

	for i := 0; i < len(allRecords); i++ {
		record := allRecords[i]
		fmt.Printf("question: %s answer: %s \n", record[0], record[1])
	}
	fmt.Printf("Finished reading file %s \n", csvFileName)
	return allRecords
}

func createDatabase() error {

	// expected directory name for the db
	dbDirName := "tutorDB"

	// expected file name for the db
	dbFileName := "QandA.db"

	// Check if the tutorDB directory already exists
	// If os.Stat returns information about the directory.
	// IsNotExist(err) = if the file exists, then it should return an error.

	// If one of these is false (that is the directory doesn't exist), then we write and structure the data base.
	// However, if there is already something there, then the statements would be false and this function just returns that the directory already exists.
	if _, err := os.Stat(dbDirName); os.IsNotExist(err) {

		fmt.Println("--------------------------------------------------")
		fmt.Printf("Creating database directory %s and database file %s \n",
			dbDirName, dbFileName)

		// The data from the dbDirName is read in.
		allRecords := readData()
		fmt.Printf("Input data allRecords has %d question/answer pairs\n", len(allRecords))

		// '0755 is a base-8 code that reads and writes permissions.'
		err := os.Mkdir(dbDirName, 0755) // read and write permissions
		if err != nil {
			log.Fatalf("os.Mkdir failed: %s", err)
		}

		// appending file name onto the directory name for access.
		fn := filepath.Join(dbDirName, dbFileName)

		// open the db.
		db, err := sql.Open("sqlite", fn)
		if err != nil {
			log.Fatalf("sql.Open failed: %s", err)
		}

		defer db.Close()

		// let qat be the name of the question-and-answer table
		stmt, err := db.Prepare(`create table if not exists qat(id integer, question text, answer text)`)
		if err != nil {
			log.Fatalf("dbPrepare create table failed: %s", err)
		}

		if _, err = stmt.Exec(); err != nil {
			log.Fatalf("stmt.Exec create table failed: %s", err)
		}

		// This goes through every record in the read in file and places them in the table.
		for id := 0; id < len(allRecords); id++ {
			record := allRecords[id]
			question := strings.Trim(record[0], "'")
			answer := strings.Trim(record[1], "'")

			fmt.Printf("defining database id %d question %s and answer %s\n",
				id, question, answer)

			stmt, err = db.Prepare("insert into qat(id, question, answer) values(?, ?, ?)")
			if err != nil {
				log.Fatalf("db.Prepare insert statement failed: %s", err)
			}

			_, err = stmt.Exec(strconv.Itoa(id), question, answer)
			if err != nil {
				log.Fatalf("db.Exec populate table failed: %s", err)
			}
		}

		fmt.Printf("Finished creating database file %s \n", dbFileName)
		return nil
	} else {
		fmt.Printf("Database directory %s already exists\n", dbDirName)
		return nil
	}
}
