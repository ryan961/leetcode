package algorithm

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
	if len(nums1) == 0 && len(nums2) == 0 {
		return float64(0)
	}

	var (
		index1, index2 = 0, 0
	)

	merge := make([]int, 0, (len(nums1)+len(nums2))/2+1)
	for len(merge) < ((len(nums1)+len(nums2))/2 + 1) {
		// 边界问题
		if index1 == len(nums1) { // mid1边界
			merge = append(merge, nums2[index2])
			index2++
			continue
		}
		if index2 == len(nums2) { // mid2边界
			merge = append(merge, nums1[index1])
			index1++
			continue
		}

		if nums1[index1] < nums2[index2] {
			merge = append(merge, nums1[index1])

			index1++
		} else if nums1[index1] > nums2[index2] {
			merge = append(merge, nums2[index2])

			index2++
		} else {
			if len(merge) == (len(nums1)+len(nums2))/2 { // 最后只需要补一个
				merge = append(merge, nums1[index1])
			} else {
				merge = append(merge, nums1[index1], nums2[index2])
			}

			index1++
			index2++
		}
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(merge[len(merge)-1]+merge[len(merge)-2]) / 2.0
	}
	return float64(merge[len(merge)-1])
}

//leetcode submit region end(Prohibit modification and deletion)
