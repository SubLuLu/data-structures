package list

import (
	"fmt"
	"testing"
)

// 双链表测试
var dl = NewDoubleList()

func TestDoubleList_Add(t *testing.T) {
	fmt.Println(t.Name())

	err := dl.Add(15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleList_Append(t *testing.T) {
	fmt.Println(t.Name())

	err := dl.Append(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleList_Insert(t *testing.T) {
	fmt.Println(t.Name())

	err := dl.Insert(23, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleList_Traverse(t *testing.T) {
	fmt.Println(t.Name())
	data := dl.Traverse()
	fmt.Println("data: ", data)
}

func TestDoubleList_Set(t *testing.T) {
	fmt.Println(t.Name())
	l := dl.Length()
	if l != 0 {
		err := dl.Set(54, l - 1)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDoubleList_Reverse(t *testing.T) {
	fmt.Println(t.Name())
	data := dl.Reverse()
	fmt.Println("data: ", data)
}

func TestDoubleList_Find(t *testing.T) {
	fmt.Println(t.Name())
	l := dl.Length()
	if l != 0 {
		val, err := dl.Find(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("index is", l - 1, ", value is", val)
	}
}

func TestDoubleList_Delete(t *testing.T) {
	fmt.Println(t.Name())
	l := dl.Length()
	if l != 0 {
		val, err := dl.Delete(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("delete value is ", val)
	}
}
