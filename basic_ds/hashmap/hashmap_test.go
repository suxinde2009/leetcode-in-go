package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {

	hashMap := NewHashMap(16)

	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap: ", hashMap.Capiticy(), "len: ", hashMap.Len())

	hashMap.Range(func(key string, value interface{}) {
		fmt.Printf("'%v'='%v", key, value)
	})
	fmt.Println("-----")

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	hashMap.Delete(key)
	fmt.Println("after delete cap: ", hashMap.Capiticy(), "len: ", hashMap.Len())

	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get '%v' not found\n'", key)
	}

}
