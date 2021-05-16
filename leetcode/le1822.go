package main

func arraySign(nums []int) int {
	lens := len(nums)
	res :=0
	for i:=0;i<lens;i++{
		if nums[i]==0{
			return 0
		}else if nums[i]<0{
			res+=1
		}
	}
	if res&0x01 ==0{
		return 1
	}
	return -1
}


