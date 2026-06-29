package main

// Problem
// Given an integer array nums and an integer k, write a function
// that returns the k most frequent elements in nums.
//
// Example
// Input:
// nums = [1,1,1,2,2,3]
// k = 2
//
// Output (in any order):
// [1, 2]

func topKFrequentElements(nums []int, k int) []int {
	// Count frequency of each number
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// buckets[i] contains all numbers that appear exactly i times
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	// Iterate from highest frequency to lowest
	topK := make([]int, 0, k)
	for count := len(buckets) - 1; count >= 0 && len(topK) < k; count-- {
		for _, num := range buckets[count] {
			topK = append(topK, num)
			if len(topK) == k {
				return topK
			}
		}
	}

	return topK
}
