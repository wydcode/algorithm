package leetcode

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=842 lang=golang
 *
 * [842] 将数组拆分成斐波那契序列
 *
 * https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/description/
 *
 * algorithms
 * Medium (39.32%)
 * Likes:    114
 * Dislikes: 0
 * Total Accepted:    8.9K
 * Total Submissions: 20.7K
 * Testcase Example:  '"123456579"'
 *
 * 给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
 *
 * 形式上，斐波那契式序列是一个非负整数列表 F，且满足：
 *
 *
 * 0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
 * F.length >= 3；
 * 对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
 *
 *
 * 另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
 *
 * 返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。
 *
 *
 *
 * 示例 1：
 *
 * 输入："123456579"
 * 输出：[123,456,579]
 *
 *
 * 示例 2：
 *
 * 输入: "11235813"
 * 输出: [1,1,2,3,5,8,13]
 *
 *
 * 示例 3：
 *
 * 输入: "112358130"
 * 输出: []
 * 解释: 这项任务无法完成。
 *
 *
 * 示例 4：
 *
 * 输入："0123"
 * 输出：[]
 * 解释：每个块的数字不能以零开头，因此 "01"，"2"，"3" 不是有效答案。
 *
 *
 * 示例 5：
 *
 * 输入: "1101111"
 * 输出: [110, 1, 111]
 * 解释: 输出 [11,0,11,11] 也同样被接受。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= S.length <= 200
 * 字符串 S 中只含有数字。
 *
 *
 */

// @lc code=start
var ans []int

func dfs(p int, s string, pre1, pre2, deep int) bool {

	length := len(s)
	if p == length {
		return deep >= 3
	}
	for i := 1; i <= 11; i++ {
		if p+i > length || (s[p] == '0' && i > 1) {
			break
		}

		sub := s[p : p+i]
		num, _ := strconv.Atoi(sub)
		if num > math.MaxInt32 || (deep != 0 && deep != 1 && num > (pre1+pre2)) {
			break
		}

		if deep == 0 || deep == 1 || ((pre1 + pre2) == num) {
			ans = append(ans, num)
			if dfs(p+i, s, pre2, num, deep+1) {
				return true
			}
			ans = ans[:len(ans)-1]
		}

	}
	return false
}

func splitIntoFibonacci(S string) []int {

	ans = nil
	dfs(0, S, 0, 0, 0)

	return ans
}

// @lc code=end
