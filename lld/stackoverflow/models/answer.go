package models

type Answer struct {
	ID       int
	Votes    int
	Content  string
	Author   User
	Question Question
}
