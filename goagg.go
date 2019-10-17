// MongoDB Aggregate Functionality is Implemented With MGO's Pipe Feature
// You can do all sorts of stuff with the Aggregate tools. Check out MongoDB's Documentation.
// The following is a brief example.
// It displays a property's address, noteCount sorted by State, City where the property doc array field "notes" has more than 2 entries

package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// just for brevity, bson.M type is map[string]interface{}

// Note: bson.D type used in the sort step, defines a slice of elements, so the order is preserved (unlike a map)
type Person struct {
	ID        string `bson:"_id"`
	State     string `bson:"st"`
	City      string `bson:"city"`
	Address   string `bson:"address"`
	NoteCount int    `bson:"notecount"`
}

type Details struct {
	ID string `bson:"_id"`

	City      string `bson:"city"`
	Address   string `bson:"address"`
	NoteCount int    `bson:"notecount"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	db := session.DB("Detailsdb")

	collection := db.C("Person")
	//p := Person{ID: "1", State: "Tamilnadu", City: "Chennai", Address: "12,guru street", NoteCount: 3}
	//p1 := Person{ID: "2", State: "Kerala", City: "Alapuzha", Address: "13,shiva street", NoteCount: 4}
	//var res []Person
	//collection.Find(bson.M{}).All(res)
	//fmt.Println(result)
	//collection.Insert(p1)
	pipeLine := []bson.M{
		bson.M{"$project": bson.M{"_id": 1, "city": 1, "address": 1, "notecount": 1}}, // output address, city, st, notecount
		// 	m{"$match": m{"notecount": m{"$gt": 2}}},                                     // keep docs with more than 2 notes
		// 	m{"$sort": bson.D{{"st", -1}, {"city", 1}}},                                  // sort results by state, city - see note above
	}
	results := []bson.M{}
	collection.Pipe(pipeLine).All(&results)
	var resarr []Details
	for i := 0; i < len(results); i++ {
		D := Details{
			ID: results[i]["_id"].(string),

			City:      results[i]["city"].(string),
			Address:   results[i]["address"].(string),
			NoteCount: results[i]["notecount"].(int),
		}
		fmt.Println(D)
		resarr = append(resarr, D)
	}
	// for i := 0; i < len(resarr); i++ {
	// 	fmt.Println(resarr[i])
	// 	fmt.Println("----------------------------------")
	// }
	// defer iter.Close()

	// for iter.Next(&result) {
	// 	log.Printf("%+v", result)
	// }
	// if iter.Err() != nil {
	// 	log.Println(iter.Err())
	// }
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
//         //unary operators- $or $and $lt $gt $multiply $type max use in $project function
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
