package leetcode_problems

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	k = length-k
	return findKthsmallest(nums ,0,length-1,k)
}

func findKthsmallest(arr []int,l,r,k int)int{
	if l>=r{
		return  arr[l]
	}
	p := partition1(arr,l,r)
	if p>k{
		return findKthsmallest(arr,l,p-1,k)
	}else if p<k{
		return findKthsmallest(arr,p+1,r,k)
	}else{
		return arr[p]
	}
}

func partition1(arr []int,l,r int)int{
	v := arr[l]
	i := l
	j:= l+1
	for ;j<=r;j++{
		if arr[j]<v{
			arr[i+1],arr[j] = arr[j],arr[i+1]
			i++
		}
	}
	arr[l],arr[i] = arr[i],arr[l]
	return i
}

