package main

import (
	"encoding/json"
	"fmt"

	"github.com/globalsign/mgo"

	"github.com/nats-io/nats.go"
)

type Book struct {
	ID         string `bson:"id,omitempty"`
	BookName   string `bson:"bookname,omitempty"`
	AuthorName string `bson:"authorname,omitempty"`
	Price      string `bson:"price",omitempty`
}

func main() {
	session, err1 := mgo.Dial("localhost:27017")
	if err1 != nil {
		fmt.Println(err1)
	}
	c := session.DB("Bookdb").C("Booknats")
	fmt.Println("Mongo Dialed : ")

	nc, err1 := nats.Connect(":4222")
	fmt.Println("nats server up")
	fmt.Println(err1)

	b := Book{ID: "3", BookName: "Java", AuthorName: "james goseline", Price: "500"}

	nc.Subscribe("Book2", func(m *nats.Msg) {
		var j *Book
		json.Unmarshal(m.Data, &j)
		c.Insert(&j)
		fmt.Println("Received message: ", string(m.Data))
	})
	bt, _ := json.Marshal(&b)
	nc.Publish("Book2", bt)
	// nc.Publish("Book1", []byte(fmt.Sprintf("%v", b)))
	select {}

}
