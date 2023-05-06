package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var Db bolt.DB

func main() {
	Db, err := bolt.Open("an.db", 0777, nil)
	if err != nil {
		log.Fatalf("bolt.open err:%v\n", err)
	}

	err = Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quan"))
		if b != nil {
			data := b.Get([]byte("1"))
			fmt.Printf("%s\n", data)
			a := b.Get([]byte("11"))
			fmt.Printf("%s\n", a)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Db.view err:%v", err)
	}
}
