package db

import (
	"container/heap"

	"example.com/go-demo1/interfacelist"
)

type PriorityQueue []*interfacelist.Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].TimeStamp < pq[j].TimeStamp
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*interfacelist.Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *interfacelist.Item, id string, timestamp uint64) {
	item.Id = id
	item.TimeStamp = timestamp
	heap.Fix(pq, item.Index)
}

func (pq *PriorityQueue) top() interface{} {
	return (*pq)[0]
}

type PriorityQueueDBStore struct {
	queueArr PriorityQueue
	lookup   map[string]interfacelist.Item
	isInit   bool
}

func (db *PriorityQueueDBStore) Extract(numberOfItem int, minTimestamp uint64) []interfacelist.Item {
	var retList []interfacelist.Item
	for db.queueArr.Len() > 0 && numberOfItem > 0 {
		//
		item := heap.Pop(&db.queueArr).(*interfacelist.Item)
		if item.TimeStamp < minTimestamp {
			break
		}
		retList = append(retList, *item)
		numberOfItem--
	}
	return retList
}

func (db *PriorityQueueDBStore) Get(id string) interfacelist.Item {
	return db.lookup[id]
}

func (db *PriorityQueueDBStore) Insert(item interfacelist.Item) bool {
	if !db.isInit {
		db.isInit = true
		db.queueArr = make(PriorityQueue, 0)
		heap.Init(&db.queueArr)
		db.lookup = make(map[string]interfacelist.Item)
	}
	heap.Push(&db.queueArr, &item)
	db.lookup[item.Id] = item
	return true
}
