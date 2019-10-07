package libpff

// #include <stdlib.h>
// #include <libpff.h>
import "C"
import "unsafe"

type File struct {
	ptr *C.libpff_file_t
}

func Open(filename string, flags int) (*File, error) {
	var f File
	err := NewError()
	if C.libpff_file_initialize(&f.ptr, &err.ptr) != retSuccess {
		return nil, err
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	if C.libpff_file_open(f.ptr, cFilename, C.int(flags), &err.ptr) != retSuccess {
		return nil, err
	}

	return &f, nil
}

func (f *File) Close() error {
	err := NewError()
	if C.libpff_file_close(f.ptr, &err.ptr) != retSuccess {
		return err
	}

	if C.libpff_file_free(&f.ptr, &err.ptr) != retSuccess {
		return err
	}

	f.ptr = nil

	return nil
}
