package zmq2_test

import (
	zmq "github.com/pebbe/zmq2"

	"fmt"
	"runtime"
)

func Example_test_version() {
	major, minor, _ := zmq.Version()
	fmt.Println("Version major:", major)
	fmt.Println("Version minor == 1|2:", minor == 1 || minor == 2)
	// Output:
	// Version major: 2
	// Version minor == 1|2: true
}

func Example_test_hwm() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_pair_inproc() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_pair_ipc() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_pair_tcp() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_reqrep_inproc() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_reqrep_ipc() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_reqrep_tcp() {

	fmt.Println("Done")
	// Output:
	// Done
}

func Example_test_shutdown_stress() {

	fmt.Println("Done")
	// Output:
	// Done
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
