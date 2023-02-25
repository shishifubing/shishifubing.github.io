/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Runtime: 174 ms, faster than 35.29% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8.4 MB, less than 72.03% of Go online submissions for Best Time to Buy and Sell Stock.

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

*/
package main

func maxProfit(prices []int) int {
	profit, index_buy := 0, 0
	for index, price := range prices {
		// changing the buying index if the current price is smaller than the
		// buying price
		if prices[index_buy] > price {
			index_buy = index
		}
		new_profit := price - prices[index_buy]
		if new_profit > profit {
			profit = new_profit
		}
	}
	return profit
}
