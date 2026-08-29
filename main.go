package main

import (
	"fmt"
	"unsafe"
)

// error code
// 0 no error; successful
// -1 invalid value
// -2 tree not initialized
// -3 insufficient tree size

// root = 1
// left child = index * 2
// right child = index * 2 + 1
var tree []int32
var treeSize int32
var initialized = false

//go:wasmexport treeInit
func initTree(levels int32) unsafe.Pointer {
	for i := range levels + 1 {
		treeSize += 2 ^ i
	}

	tree = make([]int32, treeSize+1)

	initialized = true
	return unsafe.Pointer(&tree[0])
}

//go:wasmexport treeInsert
func insert(key int32) int32 {
	if !initialized {
		return -2
	}

	// 0 및 음수를 에러 코드로 사용하기 위함
	if key <= 0 {
		return -1
	}

	var curr int32 = 1
	for tree[curr] != 0 {
		if key > tree[curr] {
			curr = curr*2 + 1
		} else {
			curr = curr * 2
		}

		if curr > treeSize {
			return -3
		}
	}

	tree[curr] = key

	return 0
}

//go:wasmexport treeSearch
func search(key int32) int32 {
	if !initialized {
		return -2
	}

	if key <= 0 {
		return -1
	}

	var curr int32 = 1
	for tree[curr] != 0 && tree[curr] != key {
		if key > tree[curr] {
			curr = curr*2 + 1
		} else {
			curr = curr * 2
		}
	}

	if tree[curr] == 0 {
		return -1
	}

	return curr
}

// TODO: remove function

func main() {
	fmt.Println("Go wasm loaded")
}
