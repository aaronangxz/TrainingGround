package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	prerequisiteMap := make(map[int][]int)

	// Iterate through all prerequisites and store in map
	for _, prerequisite := range prerequisites {
		prerequisiteMap[prerequisite[1]] = append(prerequisiteMap[prerequisite[1]], prerequisite[0])
	}

	result := make([]bool, len(queries))

	var dfs func(course, prereq int, visited map[int]bool) bool
	dfs = func(course, prereq int, visited map[int]bool) bool {
		// Visited, return immediately
		if visited[course] {
			return false
		}
		visited[course] = true

		// DFS through all of the pre-requisites
		// When either there is a match or it is searched before
		// Considered OK
		for _, val := range prerequisiteMap[course] {
			if prereq == val || dfs(val, prereq, visited) {
				return true
			}
		}
		return false
	}

	for index, query := range queries {
		course := query[1]
		prereq := query[0]
		visited := make(map[int]bool)
		if dfs(course, prereq, visited) {
			result[index] = true
		}
	}
	return result
}
