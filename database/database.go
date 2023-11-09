package database

import "github.com/cetinboran/gojson/gojson"

func DBStart() gojson.Database {
	ScraperDB := gojson.CreateDatabase("YAVUZLAR", "./")

	// Table1
	Collection := gojson.CreateTable("Collection")
	Collection.AddProperty("collectionId", "int", "PK")
	Collection.AddProperty("Searched", "string", "")
	Collection.AddProperty("Findings", "[]string", "")

	// Adds the table to the database
	ScraperDB.AddTable(&Collection)

	ScraperDB.CreateFiles()

	return ScraperDB
}
