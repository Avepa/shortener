package pkg

import "errors"

var (
	ErrorDcumentNotFound = errors.New("document not found")
	ErrorPathIsBusy      = errors.New("the path is busy")
)
