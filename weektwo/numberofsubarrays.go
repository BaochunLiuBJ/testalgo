package weektwo

func numberOfSubarrays(nums []int, k int) int {
	// sums[i] = num[0] + num[1] + ...+ num[i-1]
	// sums[0] = 0, sums[1] = sum[0] + num[0]
	sums := make([]int, len(nums) + 1)
	sums[0] = 0
	for i := 1; i <= len(nums) + 1; i++ {
		sums[i] = sums[i-1] + nums[i-1] % 2 
	}

	counts := make(map[int]int)
	var result int = 0  

	counts[sums[0]] = 1 
	for i:= 1; i <= len(nums); i++ {
		if sums[i] - k >= 0 {
			result = result + counts[sums[i] - k ]
		}
		counts[sums[i]] = counts[sums[i]] + 1 
	}
	return  result 
}

/*

   0  1  2  3  4  5  6  7  8  9  10      index
  [2, 2, 2, 1, 2, 2, 1, 2, 2, 2]  		nums[]      
  [0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2] 	sums[]

  counts[0] = 4
  counts[1] = 3
  counts[2] = 4 
   
   0, 0, 0, 0, 0, , 0, 4, 8, 12, 16     result 
 
 when index = 7
   sum[7] - 2 (k=2) >=0 
   result = count[sum[7] - k] = 4
when index = 8
   sum[8] - k >=0
   result = result + count[sum[8] -k ]
          = 4 + 4 
		  =8
*/