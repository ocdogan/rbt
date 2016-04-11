package rbt

import (
    "fmt"
)

// ErrNo struct is used for the error code of the error
type ErrNo uintptr

const (
    // ErrNoArgumentNil is used if the function parameter is nil
    ErrNoArgumentNil ErrNo = iota + 1
    // ErrNoArgumentNilWithName is used if the named function parameter is nil
    ErrNoArgumentNilWithName
    // ErrNoEnumeratorModified is used when the tree gets modified while iterating
    ErrNoEnumeratorModified
    // ErrNoIteratorAlreadyRunning is used if the iterator is already running
    ErrNoIteratorAlreadyRunning
    // ErrNoIteratorClosed is used if the iterator is closed
    ErrNoIteratorClosed
    // ErrNoIteratorUninitialized is used if the iterator is uninitialized
    ErrNoIteratorUninitialized
)

var (
    // ErrArgumentNil used if the function parameter is nil
    ErrArgumentNil = NewError(ErrNoArgumentNil) 
    // ErrEnumeratorModified is used when the tree gets modified while iterating
    ErrEnumeratorModified = NewError(ErrNoEnumeratorModified)
    // ErrIteratorAlreadyRunning used if the iterator is already iterating    
    ErrIteratorAlreadyRunning = NewError(ErrNoIteratorAlreadyRunning)
    // ErrIteratorClosed used if the iterator is closed
    ErrIteratorClosed = NewError(ErrNoIteratorClosed)
    // ErrIteratorUninitialized used if the iterator is uninitialized
    ErrIteratorUninitialized = NewError(ErrNoIteratorUninitialized)
)

var errorStr = map[ErrNo]string {
    ErrNoArgumentNil: "Argument cannot be nil.",
    ErrNoArgumentNilWithName: "Argument '%s' cannot be nil.",
    ErrNoEnumeratorModified: "Enumerator has been modified while iterating.",
    ErrNoIteratorAlreadyRunning: "Iterator already running.",
    ErrNoIteratorClosed: "Iteration context closed.",
    ErrNoIteratorUninitialized: "Iteration context uninitialized.",
}

type errorDef struct {
    err ErrNo
    message string
}

// NewError creates a new error with the given error no
func NewError(err ErrNo) error {
	return &errorDef{
        err: err, 
        message: errorStr[err],
    }
}

// NewErrorDetailed creates a new error with the given error no and message
func NewErrorDetailed(err ErrNo, msg string) error {
    if len(msg) == 0 {
        msg = errorStr[err]
    }
	return &errorDef{
        err: err,
        message: msg, 
    }
}

// ArgumentNilError creates a new error with ErrNoArgumentNilWithName error no 
// and named function parameter
func ArgumentNilError(arg string) error {
    if len(arg) == 0 {
        return &errorDef{
            err: ErrNoArgumentNil,
            message: errorStr[ErrNoArgumentNil],
        }
    }
    return &errorDef{
        err: ErrNoArgumentNil,
        message: fmt.Sprintf(errorStr[ErrNoArgumentNilWithName], arg),
    }
}

// Error returns the error message
func (err *errorDef) Error() string {
	return err.message
}

// ErrorNo returns the error no
func (err *errorDef) ErrorNo() ErrNo {
	return err.err
}