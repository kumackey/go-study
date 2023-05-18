//go:build ignore

package main

/*
#include <stdlib.h>
#include <mylib.h>
// cgo: LDFLAGS: -lmylib
// cgo: CFLAGS: -I/user/local/include
*/
import "C"

func main() {
	s := "index.dat"
	ptr := C.CString(s)
	defer C.free(unsafe.Pointer(ptr))
	C.mylib_update_data(ptr)
}
