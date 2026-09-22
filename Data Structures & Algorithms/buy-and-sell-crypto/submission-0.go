func maxProfit(prices []int) int {
// if you buy on the ith day you can only sell on the ith+n day
// you can't compare ith with ith-n
// we can use two indices l,r
// compare
// if profit is below 0 || i-1 profit is bigger then the current one
// r = r - 1
// if prices[l] < prices[r]
// r = r + 1


l := 0
maxProfit := 0

for r := 1 ; r < len(prices); r++ {
    
	profit := prices[r] - prices[l]
	
	if profit < 0 {
		l = r
	} else if profit > maxProfit {
		maxProfit = profit
	}
	
}

return maxProfit

}
