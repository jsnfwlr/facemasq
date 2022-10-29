package port

type addressPortRecord struct {
	Address string
	Ports   []portDetails
}

// type Protocol struct {
// 	Open        []int
// 	Filtered    []int
// 	Unavailable []int
// 	Closed      []int
// }

type portDetails struct {
	Number   int64
	State    string
	Protocol string
}
