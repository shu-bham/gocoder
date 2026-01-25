package google

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	empMap := make(map[int]*Employee)
	for _, employee := range employees {
		empMap[employee.Id] = employee
	}

	return calcImportanceOfEmpAndSubs(empMap, id)
}

func calcImportanceOfEmpAndSubs(empMap map[int]*Employee, id int) int {
	employee := empMap[id]
	ans := employee.Importance

	for _, subordinate := range employee.Subordinates {
		ans += calcImportanceOfEmpAndSubs(empMap, subordinate)
	}

	return ans
}

func getImportanceOpt(employees []*Employee, id int) int {
	empMap := make(map[int]*Employee)
	for _, employee := range employees {
		empMap[employee.Id] = employee
	}

	ans := 0
	queue := make([]int, 0)
	queue = append(queue, id)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		ans += empMap[front].Importance
		for _, subordinate := range empMap[front].Subordinates {
			queue = append(queue, subordinate)
		}
	}
	return ans
}
