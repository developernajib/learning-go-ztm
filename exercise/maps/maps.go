//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServerInfo(serverStatus map[string]int) {
	fmt.Println("\nThere are", len(serverStatus), "servers")

	stats := make(map[int]int)

	for _, status := range serverStatus {
		switch status {
		case Online:
			stats[Online]++
		case Offline:
			stats[Offline]++
		case Maintenance:
			stats[Maintenance]++
		case Retired:
			stats[Retired]++
		default:
			panic("Unhandled server status")
		}
	}

	fmt.Println(stats[Online], "servers are Online")
	fmt.Println(stats[Offline], "servers are Offline")
	fmt.Println(stats[Maintenance], "servers are Maintenance")
	fmt.Println(stats[Retired], "servers are Retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}
	displayServerInfo(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	displayServerInfo(serverStatus)

	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}
	displayServerInfo(serverStatus)

	// Extra
	delete(serverStatus, "omicron")
	displayServerInfo(serverStatus)
}
