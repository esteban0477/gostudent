package main

import ( "fmt")

type customErr struct {
    message string
}

/* With the next method, I'm implementing the interface error
type error interface {
    Error() string
} */

func (m *customErr) Error() string {
    return fmt.Sprintf("This is my own error with message: %v", m.message)
}

func main() {
    error1 := customErr{
            message: "fail to do something",    
    }
    foo(&error1)
}

func foo(e error) {
	// Println function is designed to call Error() method automatically when 
	// the value is an error hence any string customized on the Error() String 
	// Method, returns our custom message.
	fmt.Println(e)
}
