package libpff

// #include <stdlib.h>
// #include <libpff.h>
import "C"
import (
	"runtime"
	"unsafe"
)

const (
	retSuccess = 1
	retFail    = -1
)

// Error is a *libpff_error_t.
type Error struct {
	ptr *C.libpff_error_t
	str string
}

// newError creates a new Error with a finalizer set that ensures that any
// underlying allocated resources are freed.
func newError() *Error {
	var e Error
	runtime.SetFinalizer(&e, func(e *Error) {
		if e.ptr != nil {
			e.free()
		}
	})
	return &e
}

func (e *Error) free() {
	C.libpff_error_free(&e.ptr)
	e.ptr = nil
}

func (e *Error) Error() string {
	if len(e.str) == 0 && e.ptr != nil {
		e.str = string(make([]byte, 50))
		cStr := C.CString(e.str)
		defer C.free(unsafe.Pointer(cStr))
		n := C.libpff_error_sprint(e.ptr, cStr, C.size_t(len(e.str)))
		e.str = C.GoStringN(cStr, n)
	}
	return e.str
}
