package main

import (
	"database/sql"
	"fmt"
	"log"

	"treezoo_backend/db"
)

func main() {
	config := db.LoadConfig("config.yml")
	database := db.OpenConnection(config)
	defer database.Close()

	fetchData(database)
}

func fetchData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM view_parent_child_relations")
	if err != nil {
		log.Fatal("Query error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var childName, childSpecies, parentName, parentSpecies, zooName, zooLocation string
		err := rows.Scan(&childName, &childSpecies, &parentName, &parentSpecies, &zooName, &zooLocation)
		if err != nil {
			log.Fatal("Scan error:", err)
		}
		fmt.Printf("Child: %s, Species: %s, Parent: %s, Parent Species: %s, Zoo: %s, Location: %s\n",
			childName, childSpecies, parentName, parentSpecies, zooName, zooLocation)
	}
}
