# MaxProfitTwoTransaction Algorithm Dry Run

## Algorithm Overview
This algorithm finds the maximum profit from at most two stock transactions (buy-sell pairs). The key insight is to split the problem into two parts:
1. Maximum profit from one transaction up to day i
2. Maximum profit from one transaction from day i onwards

## Sample Input
Let's trace through with: `prices = [3, 3, 5, 0, 0, 3, 1, 4]`

## Step 1: Forward Pass - Calculate `profitIfSoldTodayOrBefore`
This calculates the maximum profit we can get from one transaction if we sell on or before day i.

| Day i | Price | minSoFar (before) | profitIfSoldToday | profitIfSoldTodayOrBefore[i] | minSoFar (after) |
|-------|-------|-------------------|-------------------|------------------------------|------------------|
| 0     | 3     | MaxInt            | -                 | 0                            | 3                |
| 1     | 3     | 3                 | max(0, 3-3) = 0   | max(0, 0) = 0               | 3                |
| 2     | 5     | 3                 | max(0, 5-3) = 2   | max(0, 2) = 2               | 3                |
| 3     | 0     | 3                 | max(0, 0-3) = 0   | max(2, 0) = 2               | 0                |
| 4     | 0     | 0                 | max(0, 0-0) = 0   | max(2, 0) = 2               | 0                |
| 5     | 3     | 0                 | max(0, 3-0) = 3   | max(2, 3) = 3               | 0                |
| 6     | 1     | 0                 | max(0, 1-0) = 1   | max(3, 1) = 3               | 0                |
| 7     | 4     | 0                 | max(0, 4-0) = 4   | max(3, 4) = 4               | 0                |

**Result:** `profitIfSoldTodayOrBefore = [0, 0, 2, 2, 2, 3, 3, 4]`

## Step 2: Backward Pass - Calculate `profitIfBuyTodayOrAfter`
This calculates the maximum profit we can get from one transaction if we buy on or after day i.

| Day i | Price | maxSoFar (before) | profitIfBuyToday | profitIfBuyTodayOrAfter[i] | maxSoFar (after) |
|-------|-------|-------------------|------------------|----------------------------|------------------|
| 7     | 4     | MinInt            | -                | 0                          | 4                |
| 6     | 1     | 4                 | max(0, 4-1) = 3  | max(0, 3) = 3             | 4                |
| 5     | 3     | 4                 | max(0, 4-3) = 1  | max(3, 1) = 3             | 4                |
| 4     | 0     | 4                 | max(0, 4-0) = 4  | max(3, 4) = 4             | 4                |
| 3     | 0     | 4                 | max(0, 4-0) = 4  | max(4, 4) = 4             | 4                |
| 2     | 5     | 4                 | max(0, 4-5) = 0  | max(4, 0) = 4             | 5                |
| 1     | 3     | 5                 | max(0, 5-3) = 2  | max(4, 2) = 4             | 5                |
| 0     | 3     | 5                 | max(0, 5-3) = 2  | max(4, 2) = 4             | 5                |

**Result:** `profitIfBuyTodayOrAfter = [4, 4, 4, 4, 4, 3, 3, 0]`

## Step 3: Find Maximum Combined Profit
Now we find the maximum sum of profits from two non-overlapping transactions:

| Day i | profitIfSoldTodayOrBefore[i] | profitIfBuyTodayOrAfter[i] | Combined Profit |
|-------|------------------------------|----------------------------|-----------------|
| 0     | 0                            | 4                          | 4               |
| 1     | 0                            | 4                          | 4               |
| 2     | 2                            | 4                          | 6               |
| 3     | 2                            | 4                          | 6               |
| 4     | 2                            | 4                          | 6               |
| 5     | 3                            | 3                          | 6               |
| 6     | 3                            | 3                          | 6               |
| 7     | 4                            | 0                          | 4               |

**Maximum Combined Profit:** 6

## Interpretation
The optimal solution involves:
1. **First transaction:** Buy at price 3 (day 0 or 1), sell at price 5 (day 2) → Profit = 2
2. **Second transaction:** Buy at price 0 (day 3 or 4), sell at price 4 (day 7) → Profit = 4
3. **Total profit:** 2 + 4 = 6

## Edge Case Handling
The algorithm handles the edge case where `len(prices) < 4` by returning 0, since you need at least 4 days to complete two transactions (buy-sell-buy-sell).

## Time & Space Complexity
- **Time:** O(n) where n is the length of prices array
- **Space:** O(n) for the two auxiliary arrays