package graberror

// ErrorHandler defines a handler function of an error
type ErrorHandler interface {
	SetErrorMessage(message string)
	Handle(packageName, functionName string, args ...interface{})
}

// GrabError defines a error struct.
type GrabError struct {
	Message     string
	HowToHandle []ErrorHandler
}

// Error implements error interface, this will make GrabError an normal error, so that it could
// be handled (at least) in a normal way.
func (err GrabError) Error() string {
	return err.Message
}

// SetErrorMessage provides an alternative way to set a error message.
func (err GrabError) SetErrorMessage(msg string) {
	err.Message = msg
}

// Handle will go through all provided error handlers and handle the function. args will be
// passed directly into error handlers, so use it if and only if YOU KNOW WHAT YOU ARE DOING.
func (err GrabError) Handle(packageName, functionName string, args ...interface{}) {
	for _, errorHandler := range err.HowToHandle {
		errorHandler.SetErrorMessage(err.Message)
		errorHandler.Handle(packageName, functionName, args...)
	}
}
