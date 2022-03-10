package interfacelist

type Item struct {
	Id        string
	Url       string
	TimeStamp uint64
	Index     int
}

type ISortedDataStore interface {
	Insert(item Item) bool
	Extract(numberOfItem int, minTimstamp uint64) []Item
	Get(id string) Item
}

type IQueueSystem interface {
	Publish(item []Item)
	SubscribeAndExecute(numberOfItem int) []Item
}
