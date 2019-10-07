package libpff

// #include <stdlib.h>
// #include <libpff.h>
import "C"
import "unsafe"

// File is a *libpff_file_t.
type File struct {
	ptr *C.libpff_file_t
}

// Open initializes and opens a new *File with the given filename and flags.
func Open(filename string, flags int) (*File, error) {
	var f File
	err := NewError()
	if retSuccess != C.libpff_file_initialize(&f.ptr, &err.ptr) {
		return nil, err
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	if retSuccess != C.libpff_file_open(f.ptr, cFilename, C.int(flags), &err.ptr) {
		return nil, err
	}

	return &f, nil
}

// Close closes and frees f.
func (f *File) Close() error {
	err := NewError()
	if retSuccess != C.libpff_file_close(f.ptr, &err.ptr) {
		return err
	}

	if retSuccess != C.libpff_file_free(&f.ptr, &err.ptr) {
		return err
	}

	f.ptr = nil

	return nil
}
