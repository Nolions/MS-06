package error_interface

import "fmt"

type err struct {
	code int
}

type NoDataError struct {}

func (n *NoDataError) ErrorCode() int {
	return 500
}

func (n *NoDataError) Error() string {
	return fmt.Sprintf("%d, No Data Error", n.ErrorCode())
}

type StoreDataError struct{}

func (s *StoreDataError) ErrorCode() int {
	return 404
}

func (s *StoreDataError) Error() string {
	return fmt.Sprintf("%d, Store Data Error", s.ErrorCode())
}

func msg(e error) string {
	return e.Error()
}
