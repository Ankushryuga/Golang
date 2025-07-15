package main

//question informations: https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

func majorityElement(nums []int) int{
  count := 0
  maj := 0
  for _, v := range nums{
    if count==0{
      maj=v
    }
    if(v==maj){
      count++
    }else{
      count--
    }
  }
  return maj
}
