//--Summary:
//  Write a program to display server status.
//

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatus(serverMap map[string]int) {
	fmt.Println("You have", len(serverMap), "servers.")
	//fmt.Println()
}

func serverInitialization(serverMap map[string]int, servers []string) {
	for _, server := range servers {
		status, found := serverMap[server]
		if !found {
			fmt.Println("Adding", server, "and turning it Oneline.")
			serverMap[server] = Online

		} else {
			var statusString string
			switch status {
			case Online:
				statusString = "Online"
			case Offline:
				statusString = "Offline"
			case Maintenance:
				statusString = "Maintenance"
			case Retired:
				statusString = "Retired"
			default:
				statusString = "Unknown"
			}
			fmt.Println("You already have the server named", server, "and its status is currently", statusString)
		}
	}
}

// --Requirements:
// * Store the existing slice of servers in a map
// * Default all of the servers to `Online`
// * Perform the following status changes and display server info:
//   - display server info
//   - change `darkstar` to `Retired`
//   - change `aiur` to `Offline`
//   - display server info
//   - change all servers to `Maintenance`
//   - display server info
func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]int)

	serverInitialization(serverMap, servers)

}
