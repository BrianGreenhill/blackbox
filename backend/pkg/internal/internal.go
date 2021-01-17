package internal

// this will hold our domain definition

// stretch: schedule

// notify about success/failure
type Notifier interface {
	Notify(*CheckResult) error
}

type CheckResult struct {
	Message string
}

// check the target
type Check interface {
	DoCheck(Target, Notifier) error
}

// struct site / target
type Target struct {
	Name string
	URL  string
}
