package google

import (
	"sort"
	"strconv"
)

func findHighAccessEmployees(access_times [][]string) []string {
	sort.Slice(access_times, func(i, j int) bool {
		return access_times[i][1] < access_times[j][1]
	})

	userTimeMap := make(map[string][]int)

	for _, strings := range access_times {
		username := strings[0]
		time, err := strconv.Atoi(strings[1])
		if err != nil {
			panic("invalid input")
		}
		userTimeMap[username] = append(userTimeMap[username], time)
	}

	var res []string
	for username, times := range userTimeMap {
		if len(times) < 3 {
			continue
		}

		highAccessUser := false
		for i := 0; i+2 < len(times); i++ {
			if times[i+2]-times[i] < 100 {
				highAccessUser = true
				break
			}
		}
		if highAccessUser {
			res = append(res, username)
		}
	}

	return res

}
