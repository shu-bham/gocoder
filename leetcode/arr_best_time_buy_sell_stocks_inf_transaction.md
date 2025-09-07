# MaxProfitInfTransaction Algorithm Dry Run

## Algorithm Overview
This algorithm finds the maximum profit from unlimited stock transactions. The key insight is that with unlimited transactions, we should capture every profitable price increase by buying at local minima and selling at local maxima.

## Strategy
The algorithm identifies continuous upward trends and treats each trend as a single transaction:
- Buy at the start of an upward trend
- Sell at the end of an upward trend
- When price drops, finalize the current transaction and start looking for the next opportunity

## Sample Input
Let's trace through with: `prices = [7, 1, 5, 3, 6, 4]`

## Initial State
- `totalProfit = 0`
- `buyPos = 0` (index where we "bought")
- `sellPos = 0` (index where we plan to "sell")

## Step-by-Step Execution

| i | prices[i] | prices[i-1] | Comparison | Action | buyPos | sellPos | totalProfit | Notes |
|---|-----------|-------------|------------|--------|---------|---------|-------------|--------|
| 1 | 1 | 7 | 1 < 7 | Price dropped | 1 | 1 | 0 + (7-7) = 0 | Finalize prev transaction (7→7, profit=0), start new at index 1 |
| 2 | 5 | 1 | 5 ≥ 1 | Price rising | 1 | 2 | 0 | Continue upward trend, move sellPos |
| 3 | 3 | 5 | 3 < 5 | Price dropped | 3 | 3 | 0 + (5-1) = 4 | Finalize transaction (1→5, profit=4), start new at index 3 |
| 4 | 6 | 3 | 6 ≥ 3 | Price rising | 3 | 4 | 4 | Continue upward trend, move sellPos |
| 5 | 4 | 6 | 4 < 6 | Price dropped | 5 | 5 | 4 + (6-3) = 7 | Finalize transaction (3→6, profit=3), start new at index 5 |

## Final Step (after loop)
After the loop ends, we need to finalize the last transaction:
- `totalProfit += prices[sellPos] - prices[buyPos]`
- `totalProfit += prices[5] - prices[5] = 4 - 4 = 0`
- **Final totalProfit = 7**

## Detailed Trace with Price Values

| Step | Action | Buy Price | Sell Price | Transaction Profit | Total Profit |
|------|--------|-----------|------------|-------------------|--------------|
| Start | Initialize | - | - | - | 0 |
| i=1 | Drop 7→1, finalize (7,7) | 7 | 7 | 0 | 0 |
| i=2 | Rise 1→5, extend trend | 1 | 5 (planned) | - | 0 |
| i=3 | Drop 5→3, finalize (1,5) | 1 | 5 | 4 | 4 |
| i=4 | Rise 3→6, extend trend | 3 | 6 (planned) | - | 4 |
| i=5 | Drop 6→4, finalize (3,6) | 3 | 6 | 3 | 7 |
| End | Finalize last (4,4) | 4 | 4 | 0 | 7 |

## Alternative Input Example
Let's also trace: `prices = [1, 2, 3, 4, 5]` (continuous upward trend)

| i | prices[i] | prices[i-1] | Comparison | Action | buyPos | sellPos | totalProfit |
|---|-----------|-------------|------------|--------|---------|---------|-------------|
| 1 | 2 | 1 | 2 ≥ 1 | Price rising | 0 | 1 | 0 |
| 2 | 3 | 2 | 3 ≥ 2 | Price rising | 0 | 2 | 0 |
| 3 | 4 | 3 | 4 ≥ 3 | Price rising | 0 | 3 | 0 |
| 4 | 5 | 4 | 5 ≥ 4 | Price rising | 0 | 4 | 0 |

**Final step:** `totalProfit += prices[4] - prices[0] = 5 - 1 = 4`

## Another Example: Zigzag Pattern
`prices = [1, 3, 2, 4]`

| i | prices[i] | prices[i-1] | Comparison | Action | buyPos | sellPos | totalProfit |
|---|-----------|-------------|------------|--------|---------|---------|-------------|
| 1 | 3 | 1 | 3 ≥ 1 | Price rising | 0 | 1 | 0 |
| 2 | 2 | 3 | 2 < 3 | Price dropped | 2 | 2 | 0 + (3-1) = 2 |
| 3 | 4 | 2 | 4 ≥ 2 | Price rising | 2 | 3 | 2 |

**Final step:** `totalProfit += prices[3] - prices[2] = 4 - 2 = 2`
**Total profit = 2 + 2 = 4**

## Key Insights
1. **Greedy Approach:** The algorithm captures every profitable opportunity
2. **Trend Following:** It rides upward trends and cuts losses immediately on downward moves
3. **Optimal Solution:** With unlimited transactions, this greedy approach is optimal
4. **Edge Cases:** Single element arrays return 0, which is handled by the initial check

## Time & Space Complexity
- **Time:** O(n) - single pass through the array
- **Space:** O(1) - only uses constant extra space

## Mathematical Equivalence
This algorithm is equivalent to summing all positive price differences:
```
profit = sum(max(0, prices[i] - prices[i-1]) for i in range(1, len(prices)))
```
But the implementation shown is more intuitive as it explicitly tracks buy/sell positions.