package main

import "fmt"

func main() {
	user1 := &User{name: "pralask", age: 30}
	user2 := &User{name: "deepak", age: 30}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)

	}

}
