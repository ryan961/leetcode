package main

// [4] Median of Two Sorted Arrays ( median-of-two-sorted-arrays )
//Given two sorted arrays nums1 and nums2 of size m and n respectively, return
//the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
//
// Example 1:
//
//
//Input: nums1 = [1,3], nums2 = [2]
//Output: 2.00000
//Explanation: merged array = [1,2,3] and median is 2.
//
//
// Example 2:
//
//
//Input: nums1 = [1,2], nums2 = [3,4]
//Output: 2.50000
//Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
//
//
// Constraints:
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics 数组 二分查找 分治 👍 6739 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid1 := nums1[:len(nums1)/2+1]
	mid2 := nums2[:len(nums2)/2+1]

	var index1, index2 = -1, -1

	merge := make([]int, 0, (len(nums1)+len(nums2))/2+1)
Loop:
	for index1 < len(mid1) {
		index1++

	Internal:
		for index2 < len(mid2) {
			index2++

			if len(merge) == (len(nums1)+len(nums2))/2+1 {
				if (len(nums1)+len(nums2))%2 == 0 {

				}
			}

			if mid1[index1] < mid2[index2] {
				merge = append(merge, mid1[index1])

				continue Loop
			} else if mid1[index1] > mid2[index2] {
				merge = append(merge, mid2[index2])

				continue Internal
			} else {
				merge = append(merge, mid1[index1], mid2[index2])

				continue Loop
			}

		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
