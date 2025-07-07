package main
import (
    // "errors"
    "fmt"
    )
    
func ErrorImplementation(nums[] int)(int, error){
    defer func(){
        if r:=recover(); r!=nil{
            fmt.Println("Recovered from panic")
        }
    }()
    
    for i:=0;i<=len(nums);i++{
        if i==len(nums){
            panic("Array out of bound array")
        }
        fmt.Println(nums[i])
    }
    return 0, nil
}


func main(){
    nums := []int{1,2,3,4,5}
    val, err := ErrorImplementation(nums)
    
    fmt.Println("val", val)
    fmt.Println("err", err)

}
