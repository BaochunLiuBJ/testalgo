package weekone 

func plusOne(digits []int) []int {
	needExtend := false  
	for i := len(digits) - 1 ; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1 
			break
		} else {
			digits[i] = 0
			if i == 0 {
				needExtend = true 
			}
		}	
	} 	
	if !needExtend {
		return digits
	} else {
		newDigits := make([]int, len(digits) + 1 ) 	
		newDigits[0] = 1
		for i := 0 ; i < len(digits) - 1; i++ {
			newDigits[i+1] = digits[i]
		}
		return newDigits 
	}		
}