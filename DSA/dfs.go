package DSA

// 841. Keys and Rooms
func CanVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)

	// DFS function
	var dfs func(room int)
	dfs = func(room int) {
		if visited[room] {
			return
		}
		visited[room] = true
		for _, key := range rooms[room] {
			dfs(key)
		}
	}

	// Start DFS from room 0
	dfs(0)

	// Check if all rooms have been visited
	return len(visited) == len(rooms)
}
