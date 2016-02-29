# GrabError

## Concept

GrabError follows this error handling concept:

> The error thrower should be the one who handles the error, not the catcher.

This concept sounds stupid, but actually makes sense. There are many kinds of errors that might be thrown by an API, and they might need different treatmeant since their serverities are different. We might want to log non-critical errors, we might want to monitor some kind of error through datadog, we might want to be warned immediately by some means when some critical error occurs. However, the API user might not understand the serverity of this error, and thus takes less appropriate way to handle the error.

But why don't we leave this job to the API author? Since they are the author, they should be responsible to tell the API user how to handle the error. Acutally some of the authors really do so, but through documentation. Well, not all of the developers read documentations, and even if they read, they don't really buy it.

Here comes GrabError, it asks the API author to tell the user how to handle the error through code, not documentation.

## How To Use It

### As Error Catcher
GrabError implements `error` interface, so the user could handle it just like a normal error.

```go
err := errorFunction()
if err != nil {
	// do error handling as usual
	logging.Info(logTag, err.Error())
}
```

But the plus point is, the error thrower has already defined a series of handling mechanisms for you, and you only need to call `Handle(packageName, functionName)`, then everything is free of charge :)

```go
err := errorFunction()
if err != nil {
	if ge, ok := err.(graberror.GrabError); ok {
		// if this is a GrabError, just call Handle function
		ge.Handle("mainPkg", "testFunc")
	} else {
		// if this is a normal error, handle normally
		logger.Log(logTag, err.Error())
	}
}
```

### As Error Thrower
GrabError exposes `ErrorHandler` interface, so error thrower could customize their error handlers according their needs.

GrabError itself also implements `ErrorHandler` interface, and could be considered as an aggregator of all error handlers. Here is an example of using GrabError:

```go
type errorCustomHandler struct{}

func (errorCustomHandler) SetErrorMessage(string) {}

func (errorCustomHandler) Handle(pkgName, funcName string, args ...interface{}) {
	// your custom error handler	
}

func errorFunction() error {
	return graberror.GrabError {
		Message: "this is a GrabError message",
		HowToHandle: []ErrorHandler{errorCustomHandler{}},
	}
}
```

## Example
Please refer to `graberror/example`.

## License
[Apache License 2.0](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0))
