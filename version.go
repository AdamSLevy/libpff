package libpff

// #cgo LDFLAGS: -lpff
//
// #include <libpff.h>
import "C"

func Version() string {
	return C.GoString(C.libpff_get_version())
}
