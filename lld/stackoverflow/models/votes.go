package models

type Vote int

const (
	UpVote Vote = iota
	Downvote
)

func (v Vote) String() string {
	return [...]string{"upvote", "downvote"}[v]
}
