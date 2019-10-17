package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Employee struct {
	Empid    string   `bson:"empid"`
	Personal Personal `bson:"personal"`
	Profile  Profile  `bson:"profile"`
	Salary   int      `bson:"salary"`
}
type Personal struct {
	Name    string  `bson:"name"`
	Age     string  `bson:"age"`
	Gender  string  `bson:"gender"`
	Address Address `bson:"address"`
}
type Address struct {
	Street     string `bson:"street"`
	District   string `bson:"district"`
	State      string `bson:"state"`
	Postalcode string `bson:"postalcode"`
}
type Profile struct {
	Designation string `bson:"designation"`
	Department  string `bson:"dept"`
}

type m bson.M

func main() {

	session, err1 := mgo.Dial("localhost:27017")
	if err1 != nil {
		fmt.Println(err1)
	}
	c := session.DB("test").C("Employee")
	fmt.Println("Mongo Connected : ")
	// e := details()
	// log.Println(e)
	// er := c.Insert(e)
	// if er != nil {
	// 	panic(er)
	// } else {
	// 	fmt.Println("Inserted")
	// }
	// var results1 []Employee
	// err3 := c.Find(bson.M{}).All(&results1)
	// fmt.Println("Length ", len(results1))

	// if err3 != nil {
	// 	panic(err3)
	// }
	// fmt.Println(results1)

	pipeLine := []m{
		m{"$match": m{"personal.gender": "male"}},

		m{"$project": m{"personal.name": 1, "personal.address.district": 1, "profile.dept": 1, "salary": 1}},
		//m{"$match": m{"salary": m{"$lt": "30000"}}},

		m{"$sort": bson.D{{"personal.name", -1}}},
	}
	iter := c.Pipe(pipeLine).Iter()
	defer iter.Close()
	var res Employee
	for iter.Next(&res) {
		log.Printf("%v", res)
	}
	if iter.Err() != nil {
		log.Println(iter.Err())
	}
}

func details() Employee {
	e := Employee{
		Empid: "127",
		Personal: Personal{Name: "thansika", Age: "23", Gender: "female",
			Address: Address{Street: "anna nagar", District: "salem", State: "tamilnadu", Postalcode: "613801"}},
		Profile: Profile{Designation: "software Testing", Department: "testing"},
		Salary:  25000}
	return e
}

//db.inventory.aggregate([
// {$limit:5},
//{$unwind:"$tags"},
//{$match:{"size.h":{$ne:10}}},//filter
//{$group:{_id:{status:"$status",item:"$item",qty:"$qty",size:"$size.h"}}},//duplicates not allow
//{$group:{_id:"$tags"}}
//{$count:"h"}
//{$sort:{"_id.qty":1}}
// {$project:{
//     item:1,
//     tags:1,
//     info:{               //modify the structure
//         brand:"$item",
//         hsiz:"$size.h",
//         //unary operattaramaniors- $or $and $lt $gt $multiply $type max use in $project function
//         typeItem:{$type:"$item"},
//         typeQty:{$type:"$qty"}
//         },
//     "size.w":1}}
//creats new collection
// {$out:"newCollection"}
//{$sort:{item:-1}}
//{$group:{_id:"$qty",count:{$sum:1}}}
//{$group:{_id:"$item",average:{$avg:"$qty"}}}

//different count methods
//$count;"total",.toArray().length,.count(),.itCount(),find().count()
//]//,{allowDiskUse:true}
//	)
