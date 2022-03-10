package distqueue

import "example.com/go-demo1/interfacelist"

func GetQueueInstFromFactory() interfacelist.IQueueSystem {
	var buff []interfacelist.Item
	var queueImpl interfacelist.IQueueSystem = &ArrayQueue{
		buffer:  buff,
		pointer: 0,
	}
	return queueImpl
}
