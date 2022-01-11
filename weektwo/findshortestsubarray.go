package weektwo

type spanNode struct {
	key   int
	span  int
	first int
	width int
}

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	spanMap := make(map[int]*spanNode)
	spanSlice := make([]*spanNode, 0)
	maxSpan := 1  
	for i := 0; i < len(nums); i++ {
		sn, ok := spanMap[nums[i]]
		if !ok {
			spanMap[nums[i]] = &spanNode{
				key:   nums[i],
				first: i,
				span:  1,
				width: 1,
			}
			spanSlice = append(spanSlice, spanMap[nums[i]])
		} else {
			sn.span++
			sn.width = (i - sn.first) + 1
			if maxSpan < sn.span {
				maxSpan = sn.span
			}
		}
	}

	minWidth := 50001 
	for _, spanNode := range spanSlice {
		if spanNode.span == maxSpan {
			if minWidth > spanNode.width {
				minWidth = spanNode.width
			}
		}
	}

	return minWidth
}
