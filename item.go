package libpff

// #include <libpff.h>
import "C"
import "runtime"

// Item is a *libpff_item_t.
type Item struct {
	ptr *C.libpff_item_t
}

func newItem() *Item {
	var i Item
	runtime.SetFinalizer(&i, func(i *Item) {
		if i.ptr != nil {
			if err := i.free(); err != nil {
				panic(err)
			}
		}
	})
	return &i
}
func (i *Item) free() error {
	err := newError()
	if retSuccess != C.libpff_item_free(&i.ptr, &err.ptr) {
		return err
	}
	i.ptr = nil
	return nil
}

// RootItem returns a new Item representing the root item for f.
func (f *File) RootItem() (*Item, error) {
	i := newItem()
	err := newError()
	if retSuccess != C.libpff_file_get_root_item(f.ptr, &i.ptr, &err.ptr) {
		return nil, err
	}
	return i, nil
}

// RootFolder returns a new Item representing the root folder for f.
func (f *File) RootFolder() (*Item, error) {
	i := newItem()
	err := newError()
	if retSuccess != C.libpff_file_get_root_folder(f.ptr, &i.ptr, &err.ptr) {
		return nil, err
	}
	return i, nil
}

// GetIdentifier returns the unique identifier for i.
func (i *Item) GetIdentifier() (uint32, error) {
	var id C.uint32_t
	err := newError()
	if retSuccess != C.libpff_item_get_identifier(i.ptr, &id, &err.ptr) {
		return 0, err
	}
	return uint32(id), nil
}

func (i *Item) GetNumEntries() (uint32, error) {
	var num C.uint32_t
	err := newError()
	if retSuccess != C.libpff_item_get_number_of_entries(i.ptr, &num, &err.ptr) {
		return 0, err
	}
	return uint32(num), nil
}
