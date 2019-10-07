package libpff

// #include <libpff.h>
import "C"

type Item struct {
	ptr *C.libpff_item_t
	id  *uint32
}

func (f *File) RootItem() (*Item, error) {
	var i Item
	err := NewError()
	if C.libpff_file_get_root_item(f.ptr, &i.ptr, &err.ptr) != retSuccess {
		return nil, err
	}
	return &i, nil
}

func (i *Item) Free() error {
	err := NewError()
	if C.libpff_item_free(&i.ptr, &err.ptr) != retSuccess {
		return err
	}
	i.ptr = nil
	return nil
}

func (i *Item) GetIdentifier() (uint32, error) {
	if i.id == nil {
		i.id = new(uint32)
		err := NewError()
		if C.libpff_item_get_identifier(i.ptr, (*C.uint32_t)(i.id), &err.ptr) !=
			retSuccess {
			return 0, err
		}
	}
	return *i.id, nil
}
