package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var Db bolt.DB

func main() {
	Db, err := bolt.Open("an.db", 0777, nil)
	if err != nil {
		log.Fatalf("bolt.open err:%v\n", err)
	}

	err = Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quan"))
		if b == nil {
			_, err = tx.CreateBucket([]byte("quan"))
			if err != nil {
				log.Fatalf("tx.createBuffet err:%v", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("db.update err:%v", err)
	}

	Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quan"))
		if b != nil {
			err := b.Put([]byte("1"), []byte("nihao"))
			_ = b.Put([]byte("11"), []byte("hello"))
			if err != nil {
				log.Fatalf("b.put err:%v", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("db.update put err:%v", err)
	}

}
