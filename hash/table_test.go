package hash

import (
	"sync"
	"testing"
)

var (
	ht = NewHashTable()    // 基础版散列表
	st = NewSeniorTable(8) // 高级版散列表
)

func TestHashTable_Put(t *testing.T) {
	var err error
	if err = ht.Put("one", 1); err != nil {
		t.Fatal(err)
	}
	if err = ht.Put(1, "one"); err != nil {
		t.Fatal(err)
	}
	if size := ht.Len(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestHashTable_Get(t *testing.T) {
	var (
		v   Value
		err error
	)
	if v, err = ht.Get("one"); err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Errorf("wrong value, expected `1` and got %v", v)
	}

	if v, err = ht.Get(1); err != nil {
		t.Fatal(err)
	}
	if v != "one" {
		t.Errorf("wrong value, expected `one` and got %v", v)
	}
}

func TestHashTable_Delete(t *testing.T) {
	var err error
	if err = ht.Delete("one"); err != nil {
		t.Fatal(err)
	}
	if size := ht.Len(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}
	if err = ht.Delete(1); err != nil {
		t.Fatal(err)
	}
	if size := ht.Len(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestSeniorTable_Put(t *testing.T) {
	var err error
	if err = st.Put("one", 1); err != nil {
		t.Fatal(err)
	}
	if err = st.Put(1, "one"); err != nil {
		t.Fatal(err)
	}
	// eno和one的冲突，会放到一个链表中
	if err = st.Put("eno", -1); err != nil {
		t.Fatal(err)
	}
	if size := st.Len(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestSeniorTable_Get(t *testing.T) {
	var (
		v   Value
		err error
	)
	if v, err = st.Get("one"); err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Errorf("wrong value, expected `1` and got %v", v)
	}

	if v, err = st.Get(1); err != nil {
		t.Fatal(err)
	}
	if v != "one" {
		t.Errorf("wrong value, expected `one` and got %v", v)
	}

	if v, err = st.Get("eno"); err != nil {
		t.Fatal(err)
	}
	if v != -1 {
		t.Errorf("wrong value, expected `-1` and got %v", v)
	}
}

func TestSeniorTable_Delete(t *testing.T) {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 测试并发删除
	go func() {
		defer wg.Done()
		if err = st.Delete("one"); err != nil {
			if err.Error() != "key does not exist" {
				t.Fatal(err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = st.Delete("one"); err != nil {
			if err.Error() != "key does not exist" {
				t.Fatal(err)
			}
		}
	}()
	wg.Wait()
}
