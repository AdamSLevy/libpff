package libpff

// #include <libpff.h>
import "C"

// Item is a *libpff_item_t.
type Item struct {
	ptr *C.libpff_item_t
}

// RootItem returns a new Item representing the root item for f.
//
// Callers are responsible for calling Free on the returned Item.
func (f *File) RootItem() (*Item, error) {
	var i Item
	err := NewError()
	if retSuccess != C.libpff_file_get_root_item(f.ptr, &i.ptr, &err.ptr) {
		return nil, err
	}
	return &i, nil
}

// RootFolder returns a new Item representing the root folder for f.
//
// Callers are responsible for calling Free on the returned Item.
func (f *File) RootFolder() (*Item, error) {
	var i Item
	err := NewError()
	if retSuccess != C.libpff_file_get_root_folder(f.ptr, &i.ptr, &err.ptr) {
		return nil, err
	}
	return &i, nil
}

// Free the underlying resources associated with i.
//
// This must be called before i is garbage collected or else memory leaks will
// occur.
func (i *Item) Free() error {
	err := NewError()
	if retSuccess != C.libpff_item_free(&i.ptr, &err.ptr) {
		return err
	}
	i.ptr = nil
	return nil
}

func (i *Item) GetIdentifier() (uint32, error) {
	var id C.uint32_t
	err := NewError()
	if retSuccess != C.libpff_item_get_identifier(i.ptr, &id, &err.ptr) {
		return 0, err
	}
	return uint32(id), nil
}

func (i *Item) GetNumEntries() (uint32, error) {
	var num C.uint32_t
	err := NewError()
	if retSuccess != C.libpff_item_get_number_of_entries(i.ptr, &num, &err.ptr) {
		return 0, err
	}
	return uint32(num), nil
}
