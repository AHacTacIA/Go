func maxSubArray(nums []int) int {
    var s,smax int = 0, nums[0]
    for i:=0;i<len(nums);i++{
        if s<0{
            s=0
        }
        s+=nums[i]
        
        if s>smax{
            smax=s
        }     
        
    }
    
    return smax
}
