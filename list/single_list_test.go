package list

import (
	"fmt"
	"testing"
)

// 单链表测试
var sl = NewSingleList()

func TestSingleList_Add(t *testing.T) {
	fmt.Println(t.Name())

	err := sl.Add(15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleList_Append(t *testing.T) {
	fmt.Println(t.Name())

	err := sl.Append(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleList_Insert(t *testing.T) {
	fmt.Println(t.Name())

	err := sl.Insert(23, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleList_Traverse(t *testing.T) {
	fmt.Println(t.Name())
	data := sl.Traverse()
	fmt.Println("data: ", data)
}

func TestSingleList_Set(t *testing.T) {
	fmt.Println(t.Name())
	l := sl.Length()
	if l != 0 {
		err := sl.Set(54, l - 1)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSingleList_Find(t *testing.T) {
	fmt.Println(t.Name())
	l := sl.Length()
	if l != 0 {
		val, err := sl.Find(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("index is", l - 1, ", value is", val)
	}
}

func TestSingleList_Delete(t *testing.T) {
	fmt.Println(t.Name())
	l := sl.Length()
	if l != 0 {
		val, err := sl.Delete(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("delete value is ", val)
	}
}
