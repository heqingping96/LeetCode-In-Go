package leetcode_problems

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
// 说明:
//
//
// 返回的下标值（index1 和 index2）不是从零开始的。
// 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//
//
// 示例:
//
// 输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
// Related Topics 数组 双指针 二分查找

//解法一：二分法
func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	arr := []int{}
	for i := 0;i<length;i++{
		element:= target-numbers[i]
		for j := 0;j <length;j++{
			if numbers[j]==element&&i<j{
				arr = []int{i+1,j+1}
				return arr
			}
		}
	}
	return arr
}

//解法二
//暴力法，双重循环
func twoSum2(numbers []int, target int) []int {
	length := len(numbers)
	arr := []int{}
	for i := 0;i<length;i++{
		for j := 0;j <length;j++{
			if numbers[i]+numbers[j]==target&&i<j{
				arr = []int{i+1,j+1}
				return arr
			}
		}
	}
	return arr
}

//方法3 对撞指针
func twoSum3(numbers []int, target int) []int {
	l,r := 0,len(numbers)-1
	arr := []int{}
	for l<r{
		if numbers[l]+numbers[r]== target{
			arr = []int{l+1,r+1}
			return arr
		}else if numbers[l]+numbers[r]> target{
			r--
		}else{
			l++
		}
	}
	return arr
}
