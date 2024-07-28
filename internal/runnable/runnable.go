package runnable

// Runnable represents anything that can be
// spawned and runned as an daemon
type Runnable interface {

	//Starts the Runner function
	//returns the error
	Run() error

	//takes the variables as the list of arguments
	//seperated by = , eg. key=value
	SetEnv(vars []string) 

	//Calls the clean up logic such as closing the log file
	Flush() error

	//kills the process
	Kill() error
}
