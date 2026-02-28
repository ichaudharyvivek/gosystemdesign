package iterator

import "fmt"

// A simple User class which needs to have iterator
type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s. Age: %d\n", u.Name, u.Age)
}

// Collections of User that we will use
type UserCollection struct {
	users []*User
}

func NewUserCollection(users []*User) *UserCollection {
	return &UserCollection{users}
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		index: 0,
		users: u.users,
	}
}
