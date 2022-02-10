package main

import (
	"fmt"
	"log"

	"github.com/abelgalef/go-viso/pkg/database"
)

func main() {
	DBser, err := database.NewDBService("root", "root", "127.0.0.1", "3306", "recordings", "mysql")
	if err != nil {
		log.Fatal(err)
	}

	if err := DBser.GenerateTableSchema(); err != nil {
		log.Fatal(err)
	}

	tables := DBser.GetTables()

	for _, item := range tables {
		fmt.Printf("Field \t Type \t Null \t Key \t Default \t Extra \t \n")
		for _, fields := range item.Schema {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \t \n", fields.Field, fields.Type, fields.Null, fields.Key, fields.DefaultData, fields.Extra)
		}
	}
	var ins []map[string]interface{}

	in := map[string]interface{}{
		"title":  "let this work",
		"artist": "abel",
		"price":  "14.8",
	}

	ins = append(ins, in)

	vals, er := DBser.InsertRows(tables[0], ins)
	if er != nil {
		log.Fatal(er)
	}

	for _, item := range vals {
		fmt.Println("----------------------------------------------")
		for k, v := range item {
			fmt.Printf("%s \t %v\n", k, v)
		}
	}

	// cons := models.Constraints{
	// 	Sort:          "ID DESC",
	// 	Field:         "ID",
	// 	OperatorValue: ">",
	// 	Value:         "2",
	// }

	// vals, e := DBser.GetRows(tables[0], cons)
	// if e != nil {
	// 	log.Fatal(err)
	// }

	// for _, item := range vals {
	// 	fmt.Println("----------------------------------------------")
	// 	for k, v := range item {
	// 		fmt.Printf("%s \t %s\n", k, v)
	// 	}
	// }

}
