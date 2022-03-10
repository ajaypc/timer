package distqueue

import (
	"math"

	"net/http"

	"example.com/go-demo1/interfacelist"
)

type ArrayQueue struct {
	buffer  []interfacelist.Item
	pointer int
}

func (arr *ArrayQueue) Publish(items []interfacelist.Item) {
	arr.buffer = append(arr.buffer, items...)
}
func (arr *ArrayQueue) SubscribeAndExecute(numberOfItem int) []interfacelist.Item {
	maxLen := int(math.Min(float64(len(arr.buffer)), float64(arr.pointer+numberOfItem)))
	retValue := arr.buffer[arr.pointer:maxLen]
	//TODO: Parallelize this
	for _, item := range retValue {
		_, err := http.Get(item.Url)
		//Just requeue for now
		if err != nil {
			arr.Publish([]interfacelist.Item{item})
		}
	}
	arr.pointer += maxLen
	return retValue
}
