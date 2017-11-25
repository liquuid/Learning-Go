package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type animal struct {
	id         int
	animalType string
	nickName   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = createdatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	err = insertsampledata(db)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from animals where id=?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickName, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
	/*
		// query single row
		row := db.QueryRow("select * from animals where age > ?", 10)
		a := animal{}
		err = row.Scan(&a.id, &a.animalType, &a.nickName, &a.zone, &a.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a)

		//insert row
		result, err := db.Exec("Insert into animals (animal_type, nickname, zone, age) values('Carnotaurus', 'Carno', 3, 22)")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	*/

	result, err := db.Exec("Update animals set age = ? where id=?", 17, 2)
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func createdatabase(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS
						animals(id INTEGER PRIMARY KEY AUTOINCREMENT ,
						animal_type TEXT ,
						nickname TEXT ,
						zone INTEGER ,
						age INTEGER)`)
	return err
}

func insertsampledata(db *sql.DB) error {
	_, err := db.Exec(`insert INTO animals (animal_type, nickname,zone,age)
						VALUES ('Tyrannosaurus rex', 'rex', 1, 10),
						('Velociraptor', 'raptor', 2 , 15)`)
	return err
}
