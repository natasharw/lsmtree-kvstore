package storage

import "testing"

func TestMemtableIit(t *testing.T) {
	// [TODO]
	result := MemtableInit()
	if result != nil {
		t.Errorf("New memtable not created")
	}
}

func TestSetInMemtable(t *testing.T) {
	// [TODO]
	result := SetInMemtable([]byte{11}, []byte{1, 2})
	if result != nil {
		t.Errorf("key not set correctly")
	}
}

func TestIsMemtableFull(t *testing.T) {
	m1 := MemtableInit() // this will absolutely not be done here
	r1 := IsMemtableFull(m1)
	if r1 != false {
		t.Errorf("buffer full not returning correct boolean")
	}

	m2 := MemtableInit() // this will absolutely not be done here
	r2 := IsMemtableFull(m2)

	if r2 != true {
		t.Errorf("buffer full not returning correct boolean")
	}
}
