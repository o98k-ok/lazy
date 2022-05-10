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

func NoError(err error) *JudgeError {
	if err != nil {
		return &JudgeError{err: err}
	}

	return &JudgeError{err: nil}
}

func Equal(expect, got interface{}) *JudgeError {
	if expect == got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect %v got %v", expect, got)}
}

func NotEqual(expect, got interface{}) *JudgeError {
	if expect != got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect not %v got %v", expect, got)}
}

func (j *JudgeError) AndEqual(expect, got interface{}) *JudgeError {
	if j.err != nil {
		return &JudgeError{err: j.err}
	}

	if expect == got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect %v got %v", expect, got)}
}

func (j *JudgeError) AndNotEqual(expect, got interface{}) *JudgeError {
	if j.err != nil {
		return &JudgeError{err: j.err}
	}

	if expect != got {
		return &JudgeError{err: nil}
	}
	return &JudgeError{err: fmt.Errorf("expect not %v got %v", expect, got)}
}
