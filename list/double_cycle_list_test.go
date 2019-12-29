package list

import (
	"fmt"
	"testing"
)

// 双向循环链表测试

var dcl = NewDoubleCycleList()

func TestDoubleCycleList_Add(t *testing.T) {
	fmt.Println(t.Name())

	err := dcl.Add(15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleCycleList_Append(t *testing.T) {
	fmt.Println(t.Name())

	err := dcl.Append(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleCycleList_Insert(t *testing.T) {
	fmt.Println(t.Name())

	err := dcl.Insert(23, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleCycleList_Traverse(t *testing.T) {
	fmt.Println(t.Name())
	data := dcl.Traverse()
	fmt.Println("data: ", data)
}

func TestDoubleCycleList_Set(t *testing.T) {
	fmt.Println(t.Name())
	l := dcl.Length()
	if l != 0 {
		err := dcl.Set(54, l - 1)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDoubleCycleList_Reverse(t *testing.T) {
	fmt.Println(t.Name())
	data := dcl.Reverse()
	fmt.Println("data: ", data)
}

func TestDoubleCycleList_Find(t *testing.T) {
	fmt.Println(t.Name())
	l := dcl.Length()
	if l != 0 {
		val, err := dcl.Find(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("index is", l - 1, ", value is", val)
	}
}

func TestDoubleCycleList_Delete(t *testing.T) {
	fmt.Println(t.Name())
	l := dcl.Length()
	if l != 0 {
		val, err := dcl.Delete(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("delete value is ", val)
	}
}
