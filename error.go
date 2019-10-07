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

type Error struct {
	ptr *C.libpff_error_t
	str string
}

func NewError() *Error {
	var err Error
	runtime.SetFinalizer(&err, func(err *Error) {
		if err.ptr != nil {
			err.free()
		}
	})
	return &err
}

func (e *Error) free() {
	C.libpff_error_free(&e.ptr)
	e.ptr = nil
}

func (e *Error) Error() string {
	if len(e.str) == 0 && e.ptr != nil {
		defer e.free()
		e.str = string(make([]byte, 50))
		cStr := C.CString(e.str)
		defer C.free(unsafe.Pointer(cStr))
		n := C.libpff_error_sprint(e.ptr, cStr, C.size_t(len(e.str)))
		e.str = C.GoStringN(cStr, n)
	}
	return e.str
}
