package error_interface

import "fmt"

func ExampleStoreDataError_Error() {
	fmt.Println(msg(new(StoreDataError)))
	// output:
	// 404, Store Data Error
}

func ExampleNoDataError_Error() {
	fmt.Println(msg(new(NoDataError)))
	// output:
	// 500, No Data Error
}
