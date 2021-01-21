package internal

// Notifier defines the contract for sending notifications with results
type Notifier interface {
	Notify(*CheckResult) error
}

// CheckResult contains a message explaining the result of a check
type CheckResult struct {
	Message string
}

// Check performs site checks and returns an error
type Check interface {
	DoCheck(Target, Notifier) error
}

// Target represents a site that can be checked
type Target struct {
	Name string
	URL  string
}
