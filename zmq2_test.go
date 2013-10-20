package zmq2_test

import (
	zmq "github.com/pebbe/zmq2"

	"errors"
	"fmt"
	"runtime"
)


func Example_test_hwm() {


	// Output:
}


func Example_test_pair_inproc() {


	// Output:
}


func Example_test_pair_ipc() {


	// Output:
}


func Example_test_pair_tcp() {


	// Output:
}


func Example_test_reqrep_inproc() {


	// Output:
}


func Example_test_reqrep_ipc() {


	// Output:
}


func Example_test_reqrep_tcp() {


	// Output:
}


func Example_test_shutdown_stress() {


	// Output:
}

func checkErr(err error) bool {
	if err != nil {
		_, filename, lineno, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("%v:%v: %v\n", filename, lineno, err)
		} else {
			fmt.Println(err)
		}
		return true
	}
	return false
}
