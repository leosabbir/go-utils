package binaryindexedtree

import (
	"fmt"
	"testing"
)

// TestInit tests initialization of BinaryIndexedTree
func TestInit(t *testing.T) {
	bit := BinaryIndexedTree{}
	bit.initialize()
	if bit.frequencies == nil || bit.indexedTree == nil {
		t.Error("freqquencies and indexedTree slices are not initialized.")
	}

	if len(bit.frequencies) != 0 {
		t.Error("Size of frequencies should be initialized to 0.")
	}

	if len(bit.indexedTree) != 1 {
		t.Error("Size of indexedTree should be initialized to 1.")
	}

	if cap(bit.frequencies) != 100 {
		t.Error("Size of frequencies should be initialized to capacity 100.")
	}

	if cap(bit.indexedTree) != 101 {
		t.Error("Size of indexedTree should be initialized to capacity 101.")
	}
} // TestInit

//---------------------------------------------------------------------------------------

func TestAppend(t *testing.T) {
	bit := BinaryIndexedTree{}
	bit.initialize()
	bit.Append(2)
	bit.Append(6)
	if len(bit.frequencies) != 2 || len(bit.indexedTree) != 3 {
		t.Errorf("Frequencies are not properly appended. Expected lenth: 2 and 3 but encountered %d and %d", len(bit.frequencies), len(bit.indexedTree))
	}
	if bit.frequencies[0] != 2 || bit.frequencies[1] != 6 || bit.indexedTree[1] != 2 || bit.indexedTree[2] != 8 {
		t.Errorf("Expecting 2, 6, 2 and 8 but encountered %d, %d, %d and %d", bit.frequencies[0], bit.frequencies[1], bit.indexedTree[1], bit.indexedTree[2])
	}

	bit.Append(6)
	bit.Append(11)
	if len(bit.frequencies) != 4 || len(bit.indexedTree) != 5 {
		t.Errorf("Frequencies are not properly appended. Expected lenth: 4 and 5 but encountered %d and %d", len(bit.frequencies), len(bit.indexedTree))
	}
	if bit.frequencies[0] != 2 || bit.frequencies[1] != 6 || bit.indexedTree[1] != 2 || bit.indexedTree[2] != 8 {
		t.Errorf("Expecting 2, 6, 2 and 8 but encountered %d, %d, %d and %d", bit.frequencies[0], bit.frequencies[1], bit.indexedTree[1], bit.indexedTree[2])
	}
	if bit.frequencies[2] != 6 || bit.frequencies[3] != 11 || bit.indexedTree[3] != 6 || bit.indexedTree[4] != 25 {
		t.Errorf("Expecting 6, 11, 6 and 25 but encountered %d, %d, %d and %d", bit.frequencies[2], bit.frequencies[3], bit.indexedTree[3], bit.indexedTree[4])
	}

	bit.Append(7)
	bit.Append(9)
	if len(bit.frequencies) != 6 || len(bit.indexedTree) != 7 {
		t.Errorf("Frequencies are not properly appended. Expected lenth: 4 and 5 but encountered %d and %d", len(bit.frequencies), len(bit.indexedTree))
	}
	if bit.frequencies[0] != 2 || bit.frequencies[1] != 6 || bit.indexedTree[1] != 2 || bit.indexedTree[2] != 8 {
		t.Errorf("Expecting 2, 6, 2 and 8 but encountered %d, %d, %d and %d", bit.frequencies[0], bit.frequencies[1], bit.indexedTree[1], bit.indexedTree[2])
	}
	if bit.frequencies[2] != 6 || bit.frequencies[3] != 11 || bit.indexedTree[3] != 6 || bit.indexedTree[4] != 25 {
		t.Errorf("Expecting 6, 11, 6 and 25 but encountered %d, %d, %d and %d", bit.frequencies[2], bit.frequencies[3], bit.indexedTree[3], bit.indexedTree[4])
	}
	if bit.frequencies[4] != 7 || bit.frequencies[5] != 9 || bit.indexedTree[5] != 7 || bit.indexedTree[6] != 16 {
		t.Errorf("Expecting 7, 9, 7 and 16 but encountered %d, %d, %d and %d", bit.frequencies[2], bit.frequencies[3], bit.indexedTree[3], bit.indexedTree[4])
	}
} // TestAppend

//---------------------------------------------------------------------------------------

func TestGetSum(t *testing.T) {
	bit := BinaryIndexedTree{}
	bit.initialize()
	bit.Append(2)
	bit.Append(6)
	bit.Append(5)
	bit.Append(7)
	bit.Append(13)
	bit.Append(16)
	bit.Append(22)
	bit.Append(66)
	bit.Append(20)
	bit.Append(600)
	bit.Append(2)
	bit.Append(6)

	tables := []struct{ idx, sum int }{
		{0, 2},
		{1, 8},
		{2, 13},
		{3, 20},
		{4, 33},
		{5, 49},
		{6, 71},
		{7, 137},
		{8, 157},
		{9, 757},
		{10, 759},
		{11, 765},
	}

	for _, table := range tables {
		sum, err := bit.GetSum(table.idx)
		if err != nil {
			t.Error(err)
		}
		if sum != table.sum {
			t.Errorf("Cumulative frequency at index %d expected %d, but encountered %d.", table.idx, table.sum, sum)
		}
	}

	_, err := bit.GetSum(-1)
	expected := fmt.Sprintf("idx should have value in between 0 and %d. Encountered: %d", len(bit.frequencies), -1)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
	_, err = bit.GetSum(20)
	expected = fmt.Sprintf("idx should have value in between 0 and %d. Encountered: %d", len(bit.frequencies), 20)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
} // TestGetValue

//---------------------------------------------------------------------------------------

func TestGetRangeSum(t *testing.T) {
	bit := BinaryIndexedTree{}
	bit.initialize()
	bit.Append(2)
	bit.Append(6)
	bit.Append(5)
	bit.Append(7)
	bit.Append(13)
	bit.Append(16)
	bit.Append(22)
	bit.Append(66)
	bit.Append(20)
	bit.Append(600)
	bit.Append(2)
	bit.Append(6)

	tables := []struct{ start, end, sum int }{
		{0, 2, 13},
		{1, 5, 47},
		{2, 2, 5},
		{3, 6, 58},
		{4, 5, 29},
		{5, 7, 104},
	}

	for _, table := range tables {
		sum, err := bit.GetRangeSum(table.start, table.end)
		if err != nil {
			t.Error(err)
		}
		if sum != table.sum {
			t.Errorf("Cumulative frequency from index %d to %d expected %d, but encountered %d.", table.start, table.end, table.sum, sum)
		}
	}

	_, err := bit.GetRangeSum(3, 2)
	expected := fmt.Sprintf("Idices should statisfy startIdx <= endIdx. startIds: %d endIdx: %d", 3, 2)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
	_, err = bit.GetRangeSum(-1, 2)
	expected = fmt.Sprintf("startIdx should be >=0. startIdx: %d", -1)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
	_, err = bit.GetRangeSum(3, 20)
	expected = fmt.Sprintf("endIdx should be <%d. endIdx: %d", len(bit.frequencies), 20)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
} // TestGetRangeSum

//---------------------------------------------------------------------------------------

func TestUpdate(t *testing.T) {
	bit := BinaryIndexedTree{}
	bit.initialize()
	bit.Append(102)
	bit.Append(65)
	bit.Append(57)
	bit.Append(73)
	bit.Append(13)
	bit.Append(6)

	tables := []struct{ idx, updateValue, start, end, sum int }{
		{0, 102, 0, 2, 224},
		{2, 50, 1, 5, 207},
		{0, 0, 2, 2, 50},
		{2, 57, 2, 2, 57},
		{2, 57, 3, 5, 92},
		{5, 10, 4, 5, 23},
	}

	for _, table := range tables {
		err := bit.Update(table.idx, table.updateValue)
		if err != nil {
			t.Error(err)
		}
		sum, err := bit.GetRangeSum(table.start, table.end)
		if err != nil {
			t.Error(err)
		}
		if sum != table.sum {
			t.Errorf("Cumulative frequency from index %d to %d expected %d, but encountered %d.", table.start, table.end, table.sum, sum)
		}
	}

	err := bit.Update(-1, 10)
	expected := fmt.Sprintf("idx should have value in between 0 and %d. Encountered: %d", len(bit.frequencies), -1)
	if err.Error() != expected {
		t.Errorf("Expected: %v Actual: %v", expected, err.Error())
	}
} // TestUpdate
