package main
//Question Infromation: https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int){
  i := m-1
  j := n-1
  k := m+n-1
  for j>=0{
    if i>=0 && nums1[i]>nums2[j]{
      nums1[k]=nums1[i]
      k--
      i--
    }else{
      nums1[k]=nums2[j]
      k--
      j--
    }
  }
}
