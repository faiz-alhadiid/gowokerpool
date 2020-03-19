package goworkerpool

// Result ...
type Result interface {
	Value() interface{}
	Error() error
}

// NewResult ...
func NewResult(value interface{}, err error) Result {
	return &resultImpl{
		value,
		err,
	}
}

type resultImpl struct {
	val interface{}
	err error
}

// Value ...
func (r *resultImpl) Value() interface{} {
	return r.val
}

// Value ...
func (r *resultImpl) Error() error {
	return r.err
}

// WorkerAdaptor ...
type WorkerAdaptor interface {
	Init() error
	Execute(in <-chan interface{}, out chan<- Result)
	Close() error
}
