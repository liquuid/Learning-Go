package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickName   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "liquuid:foo@/dino")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from dino.animals where id=?", 1)
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
		row := db.QueryRow("select * from dino.animals where age > ?", 10)
		a := animal{}
		err = row.Scan(&a.id, &a.animalType, &a.nickName, &a.zone, &a.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a)

		//insert row
		result, err := db.Exec("Insert into dino.animals (animal_type, nickname, zone, age) values('Carnotaurus', 'Carno', 3, 22)")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	*/

	result, err := db.Exec("Update dino.animals set age = ? where id=?", 17, 2)
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
