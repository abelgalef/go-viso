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

	for _, item := range DBser.GetTables() {
		fmt.Printf("Field \t Type \t Null \t Key \t Default \t Extra \t \n")
		for _, fields := range item.Schema {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \t \n", fields.Field, fields.Type, fields.Null, fields.Key, fields.DefaultData, fields.Extra)
		}
	}

}
