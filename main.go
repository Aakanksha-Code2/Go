package main

import (
	"database/sql"
	"fmt"

	"github.com/aakanksha/Crud/models"
	"github.com/aakanksha/Crud/stores"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:<password>@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("error in sql open")
		panic(err.Error())
	} else {
		fmt.Println("connected successfully")
	}
	emp1, err := stores.Insert(models.Emp{2, "aakanksha2", "ak@gmail.com", "sde"}, db)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("inserted ", emp1)
	var idupdate int
	fmt.Println("enter id number you want to update")
	fmt.Scanln(&idupdate)
	err = stores.Update(models.Emp{1, "aakankshalovely", "lovely@gmail.com", "sde"}, db)
	if err != nil {
		fmt.Println("error", err)
	}
	//fmt.Println("updated ", empU)

	var iddelete int
	fmt.Println("enter id number you want to deleted")
	fmt.Scanln(&iddelete)

	err = stores.Delete(iddelete, db)
	if err != nil {
		fmt.Println("error", err)
	}
	//	fmt.Println("deleted ", empd)

	var idget int
	fmt.Println("enter id number you want to get")
	fmt.Scanln(&idget)

	// empg, err := getbyid(idget, db)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }
	// fmt.Printf("get by id %v \n", empg)
	defer db.Close()

}
