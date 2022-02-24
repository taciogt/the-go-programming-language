package xkcd

import "fmt"

type ErrNotFound struct {
	Num int
}

func NewErrNotFound(num int) ErrNotFound {
	return ErrNotFound{Num: num}
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("xkcd comic %d not found", e.Num)
}
