package assert

import "fmt"

type JudgeError struct {
	err error
}

func (j *JudgeError) Error() string {
	return j.err.Error()
}

func (j *JudgeError) Unwrap() error {
	return j.err
}

func IfNoError(err error) *JudgeError {
	if err != nil {
		return &JudgeError{err: err}
	}

	return &JudgeError{err: nil}
}

func IfEqual(expect, got interface{}) *JudgeError {
	if expect == got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect %v got %v", expect, got)}
}

func IfNotEqual(expect, got interface{}) *JudgeError {
	if expect != got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect not %v got %v", expect, got)}
}

func (j *JudgeError) ElIfEqual(expect, got interface{}) *JudgeError {
	if j.err != nil {
		return &JudgeError{err: j.err}
	}

	if expect == got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect %v got %v", expect, got)}
}

func (j *JudgeError) ElIfNotEqual(expect, got interface{}) *JudgeError {
	if j.err != nil {
		return &JudgeError{err: j.err}
	}

	if expect != got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect not %v got %v", expect, got)}
}
