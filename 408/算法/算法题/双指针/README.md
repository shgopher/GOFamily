# 双指针

基础知识：常见双指针算法分为三类，同向（即两个指针都相同一个方向移动），背向（两个指针从相同或者相邻的位置出发，背向移动直到其中一根指针到达边界为止），相向（两个指针从两边出发一起向中间移动直到两个指针相遇）

背向双指针：(基本上全是回文串的题)
- Leetcode 409. Longest Palindrome
- Leetcode 125. Valid Palindrome
- Leetcode 5. Longest Palindromic Substring

相向双指针：(以two sum为基础的一系列题)
- Leetcode 1. Two Sum (这里使用的是先排序的双指针算法，不同于hashmap做法)
- Leetcode 167. Two Sum II - Input array is sorted
- Leetcode 15. 3Sum
- Leetcode 16. 3Sum Closest
- Leetcode 18. 4Sum
- Leetcode 454. 4Sum II
- Leetcode 277. Find the Celebrity
- Leetcode 11. Container With Most Water

同向双指针：（个人觉得最难的一类题，可以参考下这里 [TimothyL：- Leetcode 同向双指针/滑动窗口类代码模板](https://zhuanlan.zhihu.com/p/390570255)）
- Leetcode 283. Move Zeroes
- Leetcode 26. Remove Duplicate Numbers in Array
- Leetcode 395. Longest Substring with At Least K Repeating Characters
- Leetcode 340. Longest Substring with At Most K Distinct Characters
- Leetcode 424. Longest Repeating Character Replacement
- Leetcode 76. Minimum Window Substring
- Leetcode 3. Longest Substring Without Repeating Characters
- Leetcode 1004 Max Consecutive Ones III