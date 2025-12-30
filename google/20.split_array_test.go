package google_test

import (
	"gocoder/google"
	"testing"
)

func TestSplitArrayBasic1(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	k := 2
	expected := 18
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayBasic2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	expected := 9
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayEqualElements(t *testing.T) {
	nums := []int{1, 4, 4}
	k := 3
	expected := 4
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayOnePartition(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 1
	expected := 15
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayManyPartitions(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	expected := 5
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayLargeNumsOnePartition(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 1
	expected := 60
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayLargeNumsManyPartitions(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 3
	expected := 30
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayMoreComplex(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}
	k := 3
	expected := 60
	result := google.SplitArray(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayBasic1v2(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	k := 2
	expected := 18
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayBasic2v2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	expected := 9
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayEqualElementsv2(t *testing.T) {
	nums := []int{1, 4, 4}
	k := 3
	expected := 4
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayOnePartitionv2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 1
	expected := 15
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayManyPartitionsv2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	expected := 5
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayLargeNumsOnePartitionv2(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 1
	expected := 60
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayLargeNumsManyPartitionsv2(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 3
	expected := 30
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}

func TestSplitArrayMoreComplexv2(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}
	k := 3
	expected := 60
	result := google.SplitArrayv2(nums, k)
	if result != expected {
		t.Errorf("For nums %v and k %d, expected %d, but got %d", nums, k, expected, result)
	}
}
