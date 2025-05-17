package task

type State int

const (
	Pending State = iota
	Scheduled
	Running 
	Completed
	Fail
)