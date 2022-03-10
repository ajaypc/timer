// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"fmt"

	"example.com/go-demo1/db"
	"example.com/go-demo1/distqueue"
	"example.com/go-demo1/interfacelist"
)

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	var sortDataImpl interfacelist.ISortedDataStore = db.GetDBInstFromFactory()
	// POST /timer, <url>
	sortDataImpl.Insert(
		interfacelist.Item{
			Id:        "id-1",
			Url:       "www.google.com",
			TimeStamp: 1646935649,
		},
	)
	sortDataImpl.Insert(
		interfacelist.Item{
			Id:        "id-2",
			Url:       "www.yahoo.com",
			TimeStamp: 1546935649,
		},
	)

	// GET /timer/<id>
	fmt.Println(sortDataImpl.Get("id-2"))
	items := sortDataImpl.Extract(2, 0)
	fmt.Println(items)
	var queueImpl interfacelist.IQueueSystem = distqueue.GetQueueInstFromFactory()
	queueImpl.Publish(items)
	subItems := queueImpl.SubscribeAndExecute(2)
	fmt.Println(subItems)
}
