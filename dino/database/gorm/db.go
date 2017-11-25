package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	// gorm.Model
	id         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	animalType string `gorm:"type:TEXT"`
	nickName   string `gorm:"type:TEXT"`
	zone       int    `gorm:"type:INTEGER"`
	age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//db.DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})
	db.Table("dinos").CreateTable(&animal{})
	a := animal{
		animalType: "Tyrannosaurus rex",
		nickName:   "rex",
		zone:       1,
		age:        11,
	}
	db.Save(&a)

	a = animal{
		animalType: "Velociraptor",
		nickName:   "raptor",
		zone:       2,
		age:        12,
	}
	db.Save(&a)
	/*
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


		result, err := db.Exec("Update dino.animals set age = ? where id=?", 17, 2)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())*/
}
