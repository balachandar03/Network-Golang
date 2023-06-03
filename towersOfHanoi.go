package main

import "fmt"

func towersOfHanoi(n int, source, auxiliary, destination string) {
	if n > 0 {
		// Move n-1 disks from source to auxiliary
		towersOfHanoi(n-1, source, destination, auxiliary)

		// Move the nth disk from source to destination
		fmt.Printf("Move disk %d from %s to %s\n", n, source, destination)

		// Move the n-1 disks from auxiliary to destination
		towersOfHanoi(n-1, auxiliary, source, destination)
	}
}

func main() {
	n := 3 // Number of disks
	towersOfHanoi(n, "A", "B", "C")
}
