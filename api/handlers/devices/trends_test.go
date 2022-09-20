package devices

// func TestGetDevices(test *testing.T) {
// 	queries := DeviceQueries{
// 		Devices:   `SELECT * FROM Devices;`,
// 		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
// 		Addresses: `SELECT * FROM Addresses ORDER BY InterfaceID ASC, IsPrimary DESC, LastSeen DESC, IsVirtual ASC;`,
// 		Hostnames: `SELECT * FROM Hostnames;`,
// 	}
// 	_, err := getDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"))
// 	if err != nil {
// 		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
// 	}
// }
