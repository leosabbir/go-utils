package binaryindexedtree

import (
	"fmt"

	"github.com/leosabbir/go-utils/errors"
)

// BinaryIndexedTree represents binary indexed tree data structure
type BinaryIndexedTree struct {
	frequencies []int
	indexedTree []int
} // BinaryIndexedTree

//------------------------------------------------------------------------------

// NewBinaryIndexedTree is a constructor to construct new instance with underlying
// arrays initialized with proper size
func NewBinaryIndexedTree() *BinaryIndexedTree {
	bit := new(BinaryIndexedTree)
	bit.initialize()
	return bit
}

//------------------------------------------------------------------------------

func (tree *BinaryIndexedTree) initialize() {
	tree.frequencies = make([]int, 0, 100)
	tree.indexedTree = make([]int, 1, 101)
	tree.indexedTree[0] = 0
} // init

//------------------------------------------------------------------------------

// Append appends a frequency to the end
func (tree *BinaryIndexedTree) Append(freq int) {
	idx := len(tree.frequencies) + 1
	tree.frequencies = append(tree.frequencies, freq)

	sum := freq
	from := idx - (idx & -idx) + 1
	for from < idx {
		sum += tree.frequencies[from-1]
		from++
	}

	tree.indexedTree = append(tree.indexedTree, sum)
} // Append

//------------------------------------------------------------------------------

// GetSum returns cumulative frequency from beginning to given index of frequencies
func (tree *BinaryIndexedTree) GetSum(idx int) (int, error) {
	if idx < 0 || idx > len(tree.frequencies) {
		msg := fmt.Sprintf("idx should have value in between 0 and %d. Encountered: %d", len(tree.frequencies), idx)
		//log.Fatalf(msg)
		return 0, &errors.ArgError{Msg: msg}
	}
	idx++
	sum := 0
	for idx > 0 {
		sum += tree.indexedTree[idx]
		idx -= idx & -idx
	}
	return sum, nil
} // GetSum

//------------------------------------------------------------------------------

// GetRangeSum computes sum of frequencies from startIdx to endIdx
func (tree *BinaryIndexedTree) GetRangeSum(startIdx, endIdx int) (int, error) {
	if endIdx < startIdx {
		msg := fmt.Sprintf("Idices should statisfy startIdx <= endIdx. startIds: %d endIdx: %d", startIdx, endIdx)
		return 0, &errors.ArgError{Msg: msg}
	}
	if startIdx < 0 {
		msg := fmt.Sprintf("startIdx should be >=0. startIdx: %d", startIdx)
		return 0, &errors.ArgError{Msg: msg}
	}
	if endIdx > len(tree.frequencies) {
		msg := fmt.Sprintf("endIdx should be <%d. endIdx: %d", len(tree.frequencies), endIdx)
		return 0, &errors.ArgError{Msg: msg}
	}
	sum1 := 0
	if startIdx > 0 {
		sum1, _ = tree.GetSum(startIdx - 1)
	}
	sum2, _ := tree.GetSum(endIdx)

	return sum2 - sum1, nil
} // GetRangeSum

//------------------------------------------------------------------------------

// Update the value at given index with given value and
func (tree *BinaryIndexedTree) Update(idx, value int) error {
	if idx < 0 || idx > len(tree.frequencies) {
		msg := fmt.Sprintf("idx should have value in between 0 and %d. Encountered: %d", len(tree.frequencies), idx)
		return &errors.ArgError{Msg: msg}
	}
	diff := value - tree.frequencies[idx]
	n := len(tree.indexedTree)

	tree.frequencies[idx] = value
	idx++
	for idx < n {
		tree.indexedTree[idx] += diff
		idx += (idx & -idx)
	}
	return nil
} // Update
