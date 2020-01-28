package list

import (
	"fmt"
	"testing"
)

// 单向循环链表测试
var scl = NewSingleCycleList()

func TestSingleCycleList_Add(t *testing.T) {
	fmt.Println(t.Name())

	err := scl.Add(15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleCycleList_Append(t *testing.T) {
	fmt.Println(t.Name())

	err := scl.Append(12)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleCycleList_Insert(t *testing.T) {
	fmt.Println(t.Name())

	err := scl.Insert(23, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSingleCycleList_Traverse(t *testing.T) {
	fmt.Println(t.Name())
	data := scl.Traverse()
	fmt.Println("data: ", data)
}

func TestSingleCycleList_Set(t *testing.T) {
	fmt.Println(t.Name())
	l := scl.Length()
	if l != 0 {
		err := scl.Set(54, l - 1)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSingleCycleList_Find(t *testing.T) {
	fmt.Println(t.Name())
	l := scl.Length()
	if l != 0 {
		val, err := scl.Find(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("index is", l - 1, ", value is", val)
	}
}

func TestSingleCycleList_Delete(t *testing.T) {
	fmt.Println(t.Name())
	l := scl.Length()
	if l != 0 {
		val, err := scl.Delete(l - 1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("delete value is ", val)
	}
}
