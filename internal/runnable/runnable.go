package runnable

// Runnable represents anything that can be
// spawned and runned as an daemon
type Runnable interface {

	//Starts the Runner function
	//returns the error
	Run() error

	//gives the list of all Env
	//for the process
	Env() []string

	//takes the variables as the list of arguments
	//seperated by = , eg. key=value
	SetEnv(vars ...string) 

	//Calls the clean up logic such as closing the log file
	Flush() error
}
