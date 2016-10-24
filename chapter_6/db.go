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
	delete(c)
	show(c)
	update(c)
	show(c)
}

func create(c *mgo.Collection) {
	var title, writer string

	fmt.Println("> title: ")
	fmt.Scanf("%s", &title)
	fmt.Println("> writer: ")
	fmt.Scanf("%s", &writer)

	err := c.Insert(
		&Book{title, writer, time.Now()},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book inserted")
}

func read(c *mgo.Collection) (results []Book) {
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func delete(c *mgo.Collection) {
	var title string
	fmt.Println("> title:")
	fmt.Scanf("%s", &title)
	err := c.Remove(bson.M{"title": title})
	if err != nil {
		fmt.Println("Book not found.")
	} else {
		fmt.Println("Book deleted.")
	}
}

func update(c *mgo.Collection) {
	var title, newTitle, newWriter string
	fmt.Println("> title:")
	fmt.Scanf("%s", &title)
	fmt.Println("> new title:")
	fmt.Scanf("%s", &newTitle)
	fmt.Println("> new writer:")
	fmt.Scanf("%s", &newWriter)

	book := Book{newTitle, newWriter, time.Now()}
	err := c.Update(bson.M{"title": title}, book)
	if err != nil {
		fmt.Println("Book not found")
	} else {
		fmt.Println("Book updated")
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
//TODO: make title unique
//TODO: make the read method accept an optional query
//TODO : PublishedAt bug
