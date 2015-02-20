package taskrunner

import (
	"10gen.com/mci/model/host"
)

// Interface responsible for finding all hosts that are ready to run a new
// task
type HostFinder interface {
	// Find any hosts that are available to run a task
	FindAvailableHosts() ([]host.Host, error)
}

// Implementation of HostFinder that fetches the hosts from the database
type DBHostFinder struct{}

// Find all hosts available to have a task run on them.  Fetches hosts from the
// database whose status is "running" and who have no task currently being run
// on them.
func (self *DBHostFinder) FindAvailableHosts() ([]host.Host, error) {
	// find and return any hosts not currently running a task
	availableHosts, err := host.Find(host.IsAvailableAndFree)
	if err != nil {
		return nil, err
	}
	return availableHosts, nil
}
