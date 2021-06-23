package main

import (
	"fmt"
	"zhangh1220/Algorithms-Week1/mergetwolists"
	"zhangh1220/Algorithms-Week1/plusone"
	"zhangh1220/Algorithms-Week1/subarrysum"
	"zhangh1220/Algorithms-Week1/mycirculardeque"
)

func main() {
	fmt.Println("加一:")
	plusone.Test()
	fmt.Println("合并两个有序链表")
	mergetwolists.Test()
	fmt.Println("循环双端队列")
	mycirculardeque.Test()
	fmt.Println("和为k的子数组")
	subarrysum.Test()
}