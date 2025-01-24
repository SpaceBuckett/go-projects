package main

import "fmt"

func main() {

	someRandomNumbers := [][]int{{1, 2}, {1, 3}, {2, 3}, {3, 4}, {4, 5}}
	numNodes := 5
	var graph [][]int = make([][]int, numNodes+1)

	for _, edge := range someRandomNumbers {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 1; i <= numNodes; i++ {
		fmt.Printf("%d: %v\n", i, graph[i])
	}

	// making a display of graph

	for i := 0; i <= numNodes; i++ {
		for j := 0; j <= numNodes; j++ {
			if i == 0 {
				fmt.Printf(" %d ", j)
			} else if j == 0 {
				fmt.Printf(" %d ", i)
			} else {
				if contains(graph[i], j) {
					fmt.Printf(" * ")
				} else {
					fmt.Printf(" _ ")
				}
			}

		}
		fmt.Println()
	}
}

func contains(slice []int, number int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}
	return false
}
