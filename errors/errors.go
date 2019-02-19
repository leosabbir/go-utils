package errors

import "fmt"

// ArgError represents argument error
type ArgError struct {
	Msg string
} // argError

//------------------------------------------------------------------------------

// Error gives string message of the error
func (e *ArgError) Error() string {
	return fmt.Sprintf(e.Msg)
} // Error
