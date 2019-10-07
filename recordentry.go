package libpff

// #include <stdlib.h>
// #include <libpff.h>
import "C"
import "unsafe"

type RecordEntry struct {
	ptr *C.libpff_record_entry_t
}

func (s *RecordSet) GetNumEntries() (int, error) {
	var num C.int
	err := newError()
	if retSuccess != C.libpff_record_set_get_number_of_entries(s.ptr, &num, &err.ptr) {
		return 0, err
	}
	return int(num), nil
}

func (s *RecordSet) GetRecordEntry(index int) (*RecordEntry, error) {
	var e RecordEntry
	err := newError()
	if retSuccess != C.libpff_record_set_get_entry_by_index(
		s.ptr, C.int(index), &e.ptr, &err.ptr) {
		return nil, err
	}
	return &e, nil
}

func (s *RecordSet) GetRecordEntryByType(
	entryType, valueType uint32, flags uint8) (*RecordEntry, error) {
	var e RecordEntry
	err := newError()
	if retSuccess != C.libpff_record_set_get_entry_by_type(
		s.ptr, C.uint32_t(entryType), C.uint32_t(valueType),
		&e.ptr, C.uint8_t(flags), &err.ptr) {
		return nil, err
	}
	return &e, nil
}

func (s *RecordSet) GetRecordEntryByName(
	name string, valueType uint32, flags uint8) (*RecordEntry, error) {
	var e RecordEntry
	cName := (*C.uchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(cName))
	err := newError()
	if retSuccess != C.libpff_record_set_get_entry_by_utf8_name(
		s.ptr, cName, C.size_t(len(name)), C.uint32_t(valueType),
		&e.ptr, C.uint8_t(flags), &err.ptr) {
		return nil, err
	}
	return &e, nil
}

func (e *RecordEntry) Free() error {
	err := newError()
	if retSuccess != C.libpff_record_entry_free(&e.ptr, &err.ptr) {
		return err
	}
	e.ptr = nil
	return nil
}

func (e *RecordEntry) GetEntryType() (uint32, error) {
	var eType C.uint32_t
	err := newError()
	if retSuccess != C.libpff_record_entry_get_entry_type(e.ptr, &eType, &err.ptr) {
		return 0, err
	}
	return uint32(eType), nil
}

func (e *RecordEntry) GetValueType() (uint32, error) {
	var vType C.uint32_t
	err := newError()
	if retSuccess != C.libpff_record_entry_get_value_type(e.ptr, &vType, &err.ptr) {
		return 0, err
	}
	return uint32(vType), nil
}

func (e *RecordEntry) GetDataSize() (int, error) {
	var size C.size_t
	err := newError()
	if retSuccess != C.libpff_record_entry_get_data_size(e.ptr, &size, &err.ptr) {
		return 0, err
	}
	return int(size), nil
}

func (e *RecordEntry) GetData() ([]byte, error) {
	size, sizeErr := e.GetDataSize()
	if sizeErr != nil {
		return nil, sizeErr
	}
	data := C.CBytes(make([]byte, size))
	defer C.free(data)
	err := newError()
	if retSuccess != C.libpff_record_entry_get_data(
		e.ptr, (*C.uchar)(data), C.size_t(size), &err.ptr) {
		return nil, err
	}
	return C.GoBytes(data, C.int(size)), nil
}
