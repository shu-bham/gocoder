package google

// PhoneDirectory struct represents a phone directory
type PhoneDirectory struct {
	used       []bool
	recycled   []int
	nextNum    int
	maxNumbers int
}

// Constructor initializes your data structure here
//
//	@param maxNumbers - The maximum numbers that can be stored in the phone directory.
//	@return a new PhoneDirectory object
func ConstructorPD(maxNumbers int) *PhoneDirectory {
	return &PhoneDirectory{
		used:       make([]bool, maxNumbers),
		recycled:   make([]int, 0),
		nextNum:    0,
		maxNumbers: maxNumbers,
	}
}

// Get - Provide a number which is not currently in use.
//
//	@return - The number that has been allocated.
func (pd *PhoneDirectory) Get() int {
	// If there are recycled numbers, use them first.
	if len(pd.recycled) > 0 {
		number := pd.recycled[0]
		pd.recycled = pd.recycled[1:]
		pd.used[number] = true
		return number
	}

	// If there are no recycled numbers, allocate a new one.
	if pd.nextNum < pd.maxNumbers {
		number := pd.nextNum
		pd.used[number] = true
		pd.nextNum++
		return number
	}

	// No numbers available.
	return -1
}

// Check - Check if a number is available or not.
func (pd *PhoneDirectory) Check(number int) bool {
	if number < 0 || number >= pd.maxNumbers {
		return false
	}
	return !pd.used[number]
}

// Release - Recycle or release a number.
func (pd *PhoneDirectory) Release(number int) {
	if number < 0 || number >= pd.maxNumbers || !pd.used[number] {
		return
	}

	pd.used[number] = false
	pd.recycled = append(pd.recycled, number)
}
