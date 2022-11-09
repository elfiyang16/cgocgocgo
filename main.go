package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"
)

/*
<stdlib.h> -> call malloc and free
#cgo CFLAGS: -g -Wall compiles the C files with the gcc options -g(enable debug symbols) and -Wall(enable all warnings) enabled
*/

func main() {
	name := C.CString("Gopher") //takes a go string and returns a pointer to a C char ->*C.char
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)
	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	ptr := C.malloc(C.sizeof_char * 1024) // alloc 1024 chars for butter use later
	// Cgo ensures that in the event that the malloc fails, the entire program will crash
	// Hence no explict error handling here
	defer C.free(unsafe.Pointer(ptr))
	// Because C.malloc returns an object of type unsafe.Pointer, so we need to cast it to ptr to char before passing as arg
	//size := C.greet(name, year, (*C.char)(ptr))
	size := C.greet(&g, (*C.char)(ptr))
	b := C.GoBytes(ptr, size) // convert the C buffer to a go []byte
	// byte slice returned does not share memory with the bytes we allocated using malloc
	fmt.Println(string(b))
}
