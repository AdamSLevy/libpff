package libpff

// #include <libpff.h>
import "C"

type RecordSet struct {
	ptr *C.libpff_record_set_t
}

func (i *Item) GetNumRecordSets() (int, error) {
	var num C.int
	err := NewError()
	if retSuccess != C.libpff_item_get_number_of_record_sets(i.ptr, &num, &err.ptr) {
		return 0, err
	}
	return int(num), nil
}

func (i *Item) GetRecordSet(index int) (*RecordSet, error) {
	var s RecordSet
	err := NewError()
	if retSuccess != C.libpff_item_get_record_set_by_index(
		i.ptr, C.int(index), &s.ptr, &err.ptr) {
		return nil, err
	}
	return &s, nil
}

func (s *RecordSet) Free() error {
	err := NewError()
	if retSuccess != C.libpff_record_set_free(&s.ptr, &err.ptr) {
		return err
	}
	s.ptr = nil
	return nil
}
