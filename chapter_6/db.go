package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	c := session.DB("go-training").C("biblio")

	deleteAll(c)
	create(c)
	show(c)
	delete(c, "Game of thrones")
	show(c)
	update(c, Book{"Lord of the kings", "Random edited guy", time.Now()})
	show(c)
}

func create(c *mgo.Collection) {
	// make title unique
	err := c.Insert(
		&Book{"Game of thrones", "Random Guy", time.Now()},
		&Book{"Lord of the kings", "Random Guy", time.Now()},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func read(c *mgo.Collection) (results []Book) {
	// TODO: make the read method accept an optional query
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func delete(c *mgo.Collection, title string) {
	err := c.Remove(bson.M{"title": title})
	if err != nil {
		log.Fatal(err)
	}
}

func update(c *mgo.Collection, book Book) {
	err := c.Update(bson.M{"title": book.Title}, book)
	if err != nil {
		log.Fatal(err)
	}
}

func show(c *mgo.Collection) {
	fmt.Println("--------------------")
	for _, result := range read(c) {
		// TODO: add toString method
		fmt.Println(result)
	}
	fmt.Println("--------------------")
}

func deleteAll(c *mgo.Collection) {
	_, err := c.RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

//TODO: user enter book details via console
//TODO: console io
//TODO: add toString method to book
//TODO: bind crud method to Book model
