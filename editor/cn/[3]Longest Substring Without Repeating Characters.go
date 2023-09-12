package algorithm

import "strings"

// [3] Longest Substring Without Repeating Characters ( longest-substring-without-repeating-characters )
//Given a string s, find the length of the longest substring without repeating
//characters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 9552 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var longest, current string

	for k, v := range s {
		if strings.Contains(s[:k], string(v)) {
			current = current[strings.Index(current, string(v))+1:]
		}

		current += string(v)

		if len(current) >= len(longest) {
			longest = current
		}
	}

	return len(longest)
}

//leetcode submit region end(Prohibit modification and deletion)
