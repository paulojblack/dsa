package ds

import (
	"fmt"
	"testing"
)

func setupDLL() DLL {
	dll := CreateDLL()
	sample := []int{15, 2, 1, 5, 19, 2, 3, 4, 192, 4}

	for _, v := range sample {
		dll.prepend(v)
	}

	return dll
}

func TestDLLAppend(t *testing.T) {
	dll := setupDLL()

	err := dll.append(1000)

	if err != nil {
		t.Errorf("failed to append with error %v", err)
	}

	if dll.length != 11 {
		t.Errorf("incorrect length, received %v", 11)
	}

	if v, _ := dll.getAt(10); v != 1000 {
		t.Errorf("wrong value at last index, expected 1000 received %v", v)
	}

}

func TestDLLInsertAt(t *testing.T) {
	dll := setupDLL()

	dll.insertAt(1000, 1)

	if dll.length != 11 {
		t.Errorf("failed to increment length in append")
	}

	if dll.head.Next.Val != 1000 {
		t.Errorf("incorrect head.Next: %v", dll.head.Next)
	}

	if dll.head.Next.Prev != dll.head {
		t.Errorf("incorrect prev: %v", dll.head.Next.Prev)
	}

	err := dll.insertAt(99, 11)

	if err != nil {
		t.Errorf("Err %v while inserting at end of DLL", err)
	}

	if dll.length != 12 {
		t.Errorf("failed to insert at end of dll")
	}

	if v, _ := dll.getAt(11); v != 99 {
		t.Errorf("did not find value 99 at last index, instead: %v", v)
	}

}

func TestDLLInsertAtHead(t *testing.T) {
	dll := setupDLL()

	fmt.Println(dll.length)

	dll.insertAt(1000, 0)

	if dll.length != 11 {
		t.Errorf("failed to increment length in append")
	}

	if dll.head.Val != 1000 {
		t.Errorf("wrong value at head")
	}

	if dll.head.Next.Val != 4 {
		t.Errorf("incorrect head.Next: %v", dll.head.Next)
	}
}

func TestDLLPrepend(t *testing.T) {
	// testArr := []int{15, 2, 1, 5, 19, 2, 3, 4, 192, 4}

	dll := CreateDLL()

	dll.prepend(1)

	if dll.length != 1 {
		t.Errorf("failed to prepend first element to DLL")
	}

	dll.prepend(2)

	if dll.length != 2 {
		t.Errorf("failed to prepend second element to DLL")
	}

	if dll.head.Next == nil || dll.head.Next.Val != 1 {
		t.Errorf("incorrect Next assignment in prepend of second DLL element")
	}

	if dll.head.Next.Prev != dll.head {
		t.Errorf("incorrect Prev assignment during prepnd of second DLL element")
	}

	fmt.Println("correctly prepended two elements to DLL")
}

func TestDLLGet(t *testing.T) {
	dll := setupDLL()

	val, err := dll.getAt(4)

	if err != nil {
		t.Errorf("error getting expected value 4 %v", err)
	}

	if val != 2 {
		t.Errorf("incorrectly received %v instead of 2", val)
	}

	val, err = dll.getAt(9)

	if err != nil {
		t.Errorf("error getting expected value 15 %v", err)
	}

	if val != 15 {
		t.Errorf("incorrectly received %v instead of 15", val)
	}

	_, err = dll.getAt(10)

	if err == nil {
		t.Errorf("did not correctly throw error when attempting to access index 10 beyond length %v", dll.length)
	}
}
