package libpff

// #include <libpff.h>
import "C"
import "runtime"

type RecordSet struct {
	ptr *C.libpff_record_set_t
}

func (i *Item) GetNumRecordSets() (int, error) {
	var num C.int
	err := newError()
	if retSuccess != C.libpff_item_get_number_of_record_sets(i.ptr, &num, &err.ptr) {
		return 0, err
	}
	return int(num), nil
}

func (i *Item) GetRecordSet(index int) (*RecordSet, error) {
	var s RecordSet
	runtime.SetFinalizer(&s, func(s *RecordSet) {
		if s.ptr != nil {
			if err := s.free(); err != nil {
				panic(err)
			}
		}
	})
	err := newError()
	if retSuccess != C.libpff_item_get_record_set_by_index(
		i.ptr, C.int(index), &s.ptr, &err.ptr) {
		return nil, err
	}
	return &s, nil
}
func (s *RecordSet) free() error {
	err := newError()
	if retSuccess != C.libpff_record_set_free(&s.ptr, &err.ptr) {
		return err
	}
	s.ptr = nil
	return nil
}
