package db

import "example.com/go-demo1/interfacelist"

func GetDBInstFromFactory() interfacelist.ISortedDataStore {
	var temp PriorityQueue
	return &PriorityQueueDBStore{
		queueArr: temp,
		lookup:   make(map[string]interfacelist.Item),
		isInit:   false,
	}
}
