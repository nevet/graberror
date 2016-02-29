package errpkg

import (
	"github.com/nevet/graberror"
)

type errorCustomHandler struct {
	errMessage string
}

func (handler *errorCustomHandler) SetErrorMessage(msg string) {
	handler.errMessage = msg
}

func (handler *errorCustomHandler) Handle(pkgName, funcName string, args ...interface{}) {
	println("This is custom error tracker: " + handler.errMessage)
}

// ErrorCustomHandlerDemo demostrates how to use GrabError.
func ErrorCustomHandlerDemo() error {
	return graberror.GrabError{
		Message:     "An error that could be handled using custom handler",
		HowToHandle: []graberror.ErrorHandler{&errorCustomHandler{}},
	}
}
