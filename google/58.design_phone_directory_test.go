package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tech Interview Steps for Design Phone Directory:
// 1. Clarifying Questions:
//    - What is the range of maxNumbers? Is it guaranteed to be a positive integer?
//    - Are the phone numbers in a specific range, e.g., from 0 to maxNumbers-1?
//    - What should the Get() method do if no numbers are available? Return -1?
//    - What is the expected behavior for Check() with a number outside the valid range?
//    - What should happen if Release() is called with a number that is already available?
//    - Is concurrent access to the phone directory expected? (i.e., do we need to worry about thread safety?)

// 2. High-Level Approach:
//    - We need a way to keep track of which numbers are used and which are available.
//    - A boolean slice or a hash set (map[int]bool) is good for tracking used numbers. `used[i] = true` means number `i` is taken.
//    - For efficient `Get()` operations, we need a way to quickly find an available number.
//    - A queue or a list of available numbers would be efficient. When a number is released, add it back to the queue. When a number is requested, pull from the queue.
//    - Let's consider using a queue (or a simple slice acting as a queue) for available numbers. For a small `maxNumbers`, we could even just iterate to find the next available spot.

// 3. Data Structure:
//    - A `PhoneDirectory` struct.
//    - A `map[int]bool` or `[]bool` to mark used numbers. This provides O(1) for Check.
//    - A queue (e.g., a slice or `container/list`) to hold recycled numbers for O(1) Get.
//    - A counter for the next available number to be given out sequentially before we start using recycled ones.

// 4. Algorithm Walkthrough (Pseudo-code):
//    - `ConstructorPD(maxNumbers)`:
//        - Initialize `used = make([]bool, maxNumbers)`
//        - Initialize `available = new Queue()`
//        - Initialize `next = 0`
//
//    - `Get()`:
//        - If `available` queue is not empty:
//            - number = available.Dequeue()
//            - used[number] = true
//            - return number
//        - Else if `next < maxNumbers`:
//            - number = next
//            - used[number] = true
//            - next++
//            - return number
//        - Else:
//            - return -1 // No numbers available
//
//    - `Check(number)`:
//        - If `number` is out of range: return false
//        - Return `!used[number]`
//
//    - `Release(number)`:
//        - If `number` is out of range or `!used[number]`:
//             - Do nothing (or log an error)
//        - Else:
//             - used[number] = false
//             - available.Enqueue(number)
//
// 5. Complexity Analysis:
//    - `ConstructorPD`: O(N) to initialize the used array.
//    - `Get`: O(1) on average.
//    - `Check`: O(1).
//    - `Release`: O(1).
//    - Space: O(N) for the used array and the queue of available numbers.

func TestPhoneDirectory(t *testing.T) {
	t.Run("should be able to get a number, check it, and release it", func(t *testing.T) {
		phoneDirectory := ConstructorPD(3) // Numbers are 0, 1, 2

		// Get a number
		num1 := phoneDirectory.Get()
		assert.Contains(t, []int{0, 1, 2}, num1, "Get() should return a valid number")

		// Check the number's availability
		isAvailable := phoneDirectory.Check(num1)
		assert.False(t, isAvailable, "The allocated number should not be available")

		// Release the number
		phoneDirectory.Release(num1)
		isAvailableAfterRelease := phoneDirectory.Check(num1)
		assert.True(t, isAvailableAfterRelease, "The released number should now be available")
	})

	t.Run("should get all numbers and then return -1", func(t *testing.T) {
		maxNumbers := 3
		phoneDirectory := ConstructorPD(maxNumbers)

		// Get all numbers
		numbers := make(map[int]bool)
		for i := 0; i < maxNumbers; i++ {
			num := phoneDirectory.Get()
			assert.NotEqual(t, -1, num, "Should get a valid number")
			numbers[num] = true
		}
		assert.Len(t, numbers, maxNumbers, "Should have allocated all numbers")

		// Try to get one more number
		num := phoneDirectory.Get()
		assert.Equal(t, -1, num, "Should return -1 when no numbers are available")
	})

	t.Run("should reuse released numbers", func(t *testing.T) {
		phoneDirectory := ConstructorPD(2) // Numbers 0, 1

		num1 := phoneDirectory.Get() // Should be 0
		assert.Equal(t, 0, num1)

		num2 := phoneDirectory.Get() // Should be 1
		assert.Equal(t, 1, num2)

		// All numbers are used up
		assert.False(t, phoneDirectory.Check(0))
		assert.False(t, phoneDirectory.Check(1))

		// Release 0
		phoneDirectory.Release(0)
		assert.True(t, phoneDirectory.Check(0), "Number 0 should be available after release")

		// Get should now return the released number 0
		num3 := phoneDirectory.Get()
		assert.Equal(t, 0, num3, "Get should reuse the released number")
	})

	t.Run("should handle releasing an already available number gracefully", func(t *testing.T) {
		phoneDirectory := ConstructorPD(2)

		// Release a number that hasn't been allocated
		phoneDirectory.Release(1)
		assert.True(t, phoneDirectory.Check(1), "Releasing an unallocated number should not change its state")

		// Get a number and release it twice
		num := phoneDirectory.Get()
		phoneDirectory.Release(num)
		assert.True(t, phoneDirectory.Check(num), "Number should be available after first release")
		phoneDirectory.Release(num)
		assert.True(t, phoneDirectory.Check(num), "Re-releasing a number should not make it unavailable")
	})

	t.Run("should handle check for out of range numbers", func(t *testing.T) {
		phoneDirectory := ConstructorPD(3)
		assert.False(t, phoneDirectory.Check(3), "Checking for number equal to maxNumbers should be false")
		assert.False(t, phoneDirectory.Check(-1), "Checking for a negative number should be false")
	})

	t.Run("should follow the example scenario", func(t *testing.T) {
		// Init a phone directory containing 0, 1, and 2.
		pd := ConstructorPD(3)

		// It can be assumed that the first call to get() returns 0, the second returns 1, and the third returns 2.
		assert.Equal(t, 0, pd.Get())
		assert.Equal(t, 1, pd.Get())

		// check(2) returns true.
		assert.True(t, pd.Check(2))

		// get() returns 2, the last available number.
		assert.Equal(t, 2, pd.Get())

		// check(2) returns false, because it has been allocated.
		assert.False(t, pd.Check(2))

		// Assume get() returns -1, because no numbers are available.
		assert.Equal(t, -1, pd.Get())

		// release(2) recycles or releases 2.
		pd.Release(2)

		// check(2) returns true, because it's available now.
		assert.True(t, pd.Check(2))

		// get() returns 2, the number that was just released.
		assert.Equal(t, 2, pd.Get())
	})
}
