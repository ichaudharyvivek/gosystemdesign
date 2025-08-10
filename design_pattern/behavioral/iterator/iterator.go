/*
NOTE: Iterator pattern is not idiomatic in Go.
Iterator patters basically provides an iterator having hasNext(), getNext() methods for a collection
Diagram:
+-------------------+         +-----------------------------+
|    Iterator       |<--------|   NameIterator              |
|-------------------|         |-----------------------------|
| +HasNext(): bool  |         | -index: int                 |
| +Next(): string   | 		  | -collection: NameCollection |
|					|         | +HasNext()					|
| 					|         | +Next()                     |
+-------------------+    	  +-----------------------------+

+---------------------+
|  NameCollection     |
|---------------------|
| -names: []string    |
| +CreateIterator()   |
+---------------------+
*/
package iterator

// Iterator interface that defined core methods to be implemented by other classes
type Iterator interface {
	HasNext() bool
	Next() *User
}

// Concrete class
// Note: index denotes the next element to be returned by Next()
// For instance when u call Next() for the first time, index=0, so element[0] is returned,
// Then index++ makes index 1, therefore pointing to next element which needs to be returned
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) Next() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}

	return nil
}
