package models

type Question struct {
	ID      int
	Title   int
	Votes   int
	Content string
	Author  User
	Answers []*Answer
}
