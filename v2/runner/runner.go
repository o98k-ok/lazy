package runner

import (
	"sync"
)

type Result[T any] struct {
	Elem T
	Err  error
}

type Runner[T any] interface {
	Parallel(concurrency int) ([]*Result[T], error)
	Run() ([]*Result[T], error)
}

type RunnerImp[T any] struct {
	Tasks []func() (T, error)
}

func NewRunner[T any](fns ...func() (T, error)) Runner[T] {
	return &RunnerImp[T]{
		Tasks: fns,
	}
}

func (r *RunnerImp[T]) Parallel(concurrency int) ([]*Result[T], error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(r.Tasks))

	conchans := make(chan struct{}, concurrency)
	var outputs = make([]*Result[T], len(r.Tasks))
	for i := 0; i < len(r.Tasks); i++ {
		go func(index int) {
			conchans <- struct{}{}
			defer func() {
				defer wg.Done()
				<-conchans
			}()

			res, err := r.Tasks[index]()
			outputs[index] = &Result[T]{
				Elem: res,
				Err:  err,
			}
		}(i)
	}
	wg.Wait()
	return outputs, nil
}

func (r *RunnerImp[T]) Run() ([]*Result[T], error) {
	var outputs = make([]*Result[T], len(r.Tasks))
	for i := 0; i < len(r.Tasks); i++ {
		res, err := r.Tasks[i]()
		outputs[i] = &Result[T]{
			Elem: res,
			Err:  err,
		}
	}
	return outputs, nil
}
