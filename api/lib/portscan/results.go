package portscan

type ResultSet struct {
	Address string
	Ports   []Port
}

// type Protocol struct {
// 	Open        []int
// 	Filtered    []int
// 	Unavailable []int
// 	Closed      []int
// }

type Port struct {
	Number   int
	State    string
	Protocol string
}
