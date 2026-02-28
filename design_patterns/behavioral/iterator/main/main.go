package main

import (
	i "design-patterns/behavioral/iterator"
	"fmt"
)

func main() {
	user1 := &i.User{Name: "John", Age: 54}
	user2 := &i.User{Name: "Alice", Age: 23}

	uc := i.NewUserCollection([]*i.User{user1, user2})
	itr := uc.CreateIterator()

	for itr.HasNext() {
		u := itr.Next()
		fmt.Print(u)
	}
}
