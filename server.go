package yai

// Run the server.
// This is called from the generated main file.
// If port is non-zero, use that.  Else, read the port from app.conf.
func Run(port int) {

	runStartupHooks()

}


func runStartupHooks() {
	for _, hook := range startupHooks {
		hook()
	}
}

var startupHooks []func()

// OnAppStart Register a function to be run at app startup.
func OnAppStart(f func()) {
	startupHooks = append(startupHooks, f)
}
