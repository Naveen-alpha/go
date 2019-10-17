package main

import (
	"fmt"

	config "./config"
	"github.com/globalsign/mgo"
)

func main() {
	session, _ := config.GetMongoDB()
	c := session.DB("learn_mongodb_golang").C("product")
	job := &mgo.MapReduce{
		Map:    "function() { emit(this.name, this.price) }",
		Reduce: "function(key, values) { return Array.sum(values)}",

		Out: nil,
	}

	var result []struct {
		Name  string `bson:"_id"`
		Value int
	}
	_, err := c.Find(nil).MapReduce(job, &result)
	if err != nil {
		panic(err)
	}
	for _, item := range result {
		fmt.Println(fmt.Sprintf("%s %d", item.Name, item.Value))

	}

	// db.mr.mapReduce(
	// 	function() {
	// 		for (var idx = 0; idx < this.items.length; idx++) {
	// 			var key = this.items[idx].sku;
	// 			var value = { count: 1, qty: this.items[idx].qty};
	// 			emit(key, value);
	// 		}
	// 	 },
	// 	function(keySKU, countObjVals) {
	// 	  reducedVal = { count: 0, qty: 0 };
	// 		 for (var idx = 0; idx < countObjVals.length; idx++) {
	// 		  reducedVal.count += countObjVals[idx].count;
	// 		  reducedVal.qty += countObjVals[idx].qty;
	// 	  }

	// 	  return reducedVal;
	//    },

	// 	  {
	// 		out: "map_reduce_example",
	// 		query: { ord_date:
	// 				   { $gt: new Date('01/01/2012') }
	// 			   }
	//                        function (key, reducedVal) {

	//                        reducedVal.avg = reducedVal.qty/reducedVal.count;

	//                        return reducedVal;
	//                      }

}
