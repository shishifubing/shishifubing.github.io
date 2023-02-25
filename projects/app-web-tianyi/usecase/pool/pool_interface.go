package usePool

import (
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
)

type (
	Handler      func(job *entity.Job) error
	ErrorHandler func(job *entity.Job, err error) error
)

type Interactor interface {
	Pool() Pool
}

type Pool interface {
	// define a job
	Job(name string, handler Handler, errorHandler ErrorHandler)
	// start the pool
	Start()
	// close the pool
	Close() error
}
