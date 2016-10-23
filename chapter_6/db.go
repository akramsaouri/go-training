package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type Book struct {
	Title       string
	Writer      string
	PublishedAt time.Time
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("biblio")

	err = c.Insert(
		&Book{"Game of thrones", "Random Guy", time.Now()},
		&Book{"Lord of the kings", "Random Guy", time.Now()},
	)
	if err != nil {
		log.Fatal(err)
	}
}
