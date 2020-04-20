package storage

import "testing"

func TestMemtableInit(t *testing.T) {
	result := MemtableInit()
	if result != nil {
		t.Errorf("New memtable not created")
	}
}

func TestSetInMemtable(t *testing.T) {
	result := InsertToMemtable("foo", []byte{1, 2})
	if result != nil {
		t.Errorf("key not set correctly")
	}
}

func TestIsMemtableFull(t *testing.T) {
	m1 := MemtableInit() // [TODO] do not instantiate here
	r1 := IsMemtableFull(m1)
	if r1 != false {
		t.Errorf("buffer full not returning correct boolean")
	}

	m2 := MemtableInit() // [TODO] do not instantiate here
	r2 := IsMemtableFull(m2)

	if r2 != true {
		t.Errorf("buffer full not returning correct boolean")
	}
}
