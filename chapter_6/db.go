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

	choice := -1
	for choice != 0 {
		fmt.Println("1: create")
		fmt.Println("2: read")
		fmt.Println("3: update")
		fmt.Println("4: delete")
		fmt.Println("0: quit")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("Creating new book:")
			var title, writer string
			fmt.Println("> title: ")
			fmt.Scanf("%s", &title)
			fmt.Println("> writer: ")
			fmt.Scanf("%s", &writer)
			create(c, Book{title, writer, time.Now()})
		case 2:
			fmt.Println("Reading all books:")
			show(c)
		case 3:
			fmt.Println("Updating a book:")
			var title, newTitle, newWriter string
			fmt.Println("> title:")
			fmt.Scanf("%s", &title)
			fmt.Println("> new title:")
			fmt.Scanf("%s", &newTitle)
			fmt.Println("> new writer:")
			fmt.Scanf("%s", &newWriter)
			update(c, title, Book{newTitle, newWriter, time.Now()})
		case 4:
			fmt.Println("Deleting a book:")
			var title string
			fmt.Println("> title:")
			fmt.Scanf("%s", &title)
			delete(c, title)
		}
		fmt.Println("----------------")
	}
}

func create(c *mgo.Collection, book Book) {
	err := c.Insert(book)
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

func delete(c *mgo.Collection, title string) {

	err := c.Remove(bson.M{"title": title})
	if err != nil {
		fmt.Println("Book not found.")
	} else {
		fmt.Println("Book deleted.")
	}
}

func update(c *mgo.Collection, title string, book Book) {
	err := c.Update(bson.M{"title": title}, book)
	if err != nil {
		fmt.Println("Book not found")
	} else {
		fmt.Println("Book updated")
	}
}

func show(c *mgo.Collection) {
	for _, result := range read(c) {
		fmt.Println(result)
	}
}

func deleteAll(c *mgo.Collection) {
	_, err := c.RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

//TODO: add toString method to book
//TODO: bind crud method to Book model
//TODO: make title unique
//TODO: make the read method accept an optional query
//TODO : PublishedAt bug
// TODO: add toString method
