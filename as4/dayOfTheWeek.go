// Enter code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
input is an integer return equivalent day
*/

// TC - O(1), SC - O(1)
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	dayMap := map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}

	fmt.Println(dayMap[N])
}
