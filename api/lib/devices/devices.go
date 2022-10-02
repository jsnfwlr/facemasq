package devices

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/utils"
	"facemasq/models"

	"github.com/uptrace/bun"
)

type DeviceQueries struct {
	Devices   string
	Netfaces  string
	Addresses string
	Hostnames string
}

type DeviceChildren struct {
	Interfaces string
	Addresses  string
}

type InterfaceChildren string

type Connectivity []models.Connections

func (a Connectivity) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }
func (a Connectivity) Len() int           { return len(a) }
func (a Connectivity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetConnectivityData(connTime string) (connections []models.ConnectionGroup, err error) {
	sql := `SELECT time, CASE WHEN addresses IS NULL THEN 204 ELSE addresses END AS addresses FROM (SELECT scans.time, GROUP_CONCAT(address_id, ',') as addresses FROM scans LEFT JOIN histories ON scans.id = histories.scan_id WHERE scans.time > ? GROUP BY scans.time) as scans ORDER BY scans.time ASC;`
	err = db.Conn.NewRaw(sql, connTime).Scan(db.Context, &connections)
	if err != nil {
		logging.Errorf("Error getting connectivity data: %v", err)
		return
	}
	for i := range connections {
		connections[i].AddressList = fmt.Sprintf(",%s,", connections[i].AddressList)
	}
	return
}

func ParseConnectivityData(addresses []int64, connections []models.ConnectionGroup) (connectivity map[int64]Connectivity) {
	connectivity = make(map[int64]Connectivity)
	for _, address := range addresses {
		connectivity[address] = Connectivity{}
		// if address == 39 || address == 13 { // DEBUG
		// 	fmt.Printf("Looking at %d - ", address) // DEBUG
		// } // DEBUG
		for i := range connections {
			conn := models.Connections{Time: connections[i].Time, State: false}
			if strings.Contains(connections[i].AddressList, strconv.Itoa(int(address))) {
				conn.State = true
			}
			// if address == 39 || address == 13 { // DEBUG
			// 	fmt.Printf("%s", utils.Ternary(conn.State, "Y", "N").(string)) // DEBUG
			// } // DEBUG
			connectivity[address] = append(connectivity[address], conn)
		}
		// if address == 39 || address == 13 { // DEBUG
		// 	fmt.Printf("\n") // DEBUG
		// } // DEBUG
	}
	return
}

func GetDevices(queries DeviceQueries, connTime string, includeConnectivity bool) (matchedDevices []models.Device, err error) {
	var devices []models.Device
	var netfaces []models.Interface
	var addresses []models.Address
	var hostnames []models.Hostname
	var connections []models.ConnectionGroup
	var connectivity map[int64]Connectivity

	err = db.Conn.NewRaw(queries.Devices).Scan(db.Context, &devices)
	if err != nil {
		return
	}

	err = db.Conn.NewRaw(queries.Netfaces).Scan(db.Context, &netfaces)
	if err != nil {
		return
	}

	err = db.Conn.NewRaw(queries.Addresses).Scan(db.Context, &addresses)
	if err != nil {
		return
	}

	err = db.Conn.NewRaw(queries.Hostnames).Scan(db.Context, &hostnames)
	if err != nil {
		return
	}

	var addressIDs []int64
	for a := range addresses {
		addressIDs = append(addressIDs, addresses[a].ID)
	}

	if includeConnectivity {
		connections, err = GetConnectivityData(connTime)
		if err != nil {
			return
		}

		connectivity = ParseConnectivityData(addressIDs, connections)
	}

	for a := range addresses {
		if addresses[a].IPv4.Valid {
			sections := strings.Split(addresses[a].IPv4.String, ".")

			for s := range sections {
				var ipnum int
				ipnum, _ = strconv.Atoi(sections[s])
				addresses[a].SortOrder += fmt.Sprintf("%03d", ipnum)
			}
		}
		if includeConnectivity {
			addresses[a].Connectivity = connectivity[addresses[a].ID]
		}

	}

	for d := range devices {
		for n := range netfaces {
			if devices[d].ID == netfaces[n].DeviceID {
				for a := range addresses {
					if netfaces[n].ID == addresses[a].InterfaceID {
						for h := range hostnames {
							if addresses[a].ID == hostnames[h].AddressID {
								addresses[a].Hostnames = append(addresses[a].Hostnames, hostnames[h])
							}
						}
						netfaces[n].Addresses = append(netfaces[n].Addresses, addresses[a])
						if netfaces[n].SortOrder == "" {
							netfaces[n].SortOrder = addresses[a].SortOrder
							netfaces[n].Primary.IPv4 = addresses[a].IPv4.String
							netfaces[n].Primary.IPv6 = addresses[a].IPv6.String
							netfaces[n].Primary.IsVirtualIP = utils.Ternary(addresses[a].IsVirtual.Valid, addresses[a].IsVirtual.Bool, false).(bool)
							netfaces[n].Primary.IsReservedIP = utils.Ternary(addresses[a].IsReserved.Valid, addresses[a].IsReserved.Bool, false).(bool)
						}
					}
				}
				if len(netfaces[n].Addresses) > 0 {
					devices[d].Interfaces = append(devices[d].Interfaces, netfaces[n])
					if devices[d].SortOrder == "" {
						devices[d].SortOrder = netfaces[n].SortOrder
						devices[d].Primary = netfaces[n].Primary
						devices[d].Primary.MAC = netfaces[n].MAC
						devices[d].Primary.InterfaceTypeID = netfaces[n].InterfaceTypeID
						devices[d].Primary.VlanID = netfaces[n].VlanID
						devices[d].Primary.IsVirtualIFace = netfaces[n].IsVirtual
					}
				}
			}
		}
	}

	for d := range devices {
		if len(devices[d].Interfaces) > 0 {
			matchedDevices = append(matchedDevices, devices[d])
		}
	}
	sort.SliceStable(matchedDevices, func(i, j int) bool {
		return matchedDevices[i].SortOrder < matchedDevices[j].SortOrder
	})
	return
}

func GetChangesSince(lastSeen time.Time, includeConnectivity bool) (matchedDevices []models.Device, seenRecently time.Time, err error) {
	var deviceIDs []int64
	var interfaceIDs []int64
	var addressIDs []int64

	var devices []models.Device
	var netfaces []models.Interface
	var addresses []models.Address
	var hostnames []models.Hostname

	var connections []models.ConnectionGroup
	var connectivity map[int64]Connectivity

	err = db.Conn.NewSelect().Model(&addresses).Where("last_seen >= ?", lastSeen).Scan(db.Context)
	if err != nil {
		return
	}
	for a := range addresses {
		interfaceIDs = append(interfaceIDs, addresses[a].InterfaceID)
		addressIDs = append(addressIDs, addresses[a].ID)
	}

	err = db.Conn.NewSelect().Model(&hostnames).Where("address_id IN (?)", bun.In(addressIDs)).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&netfaces).Where("id IN (?)", bun.In(interfaceIDs)).Scan(db.Context)
	if err != nil {
		return
	}

	for i := range netfaces {
		deviceIDs = append(deviceIDs, netfaces[i].DeviceID)
	}

	err = db.Conn.NewSelect().Model(&devices).Where("id IN (?)", bun.In(deviceIDs)).Scan(db.Context)
	if err != nil {
		return
	}

	if includeConnectivity {
		connections, err = GetConnectivityData(lastSeen.Format("2006-01-02 15:04"))
		if err != nil {
			return
		}

		connectivity = ParseConnectivityData(addressIDs, connections)
	}

	for a := range addresses {
		if addresses[a].IPv4.Valid {
			sections := strings.Split(addresses[a].IPv4.String, ".")

			for s := range sections {
				var ipnum int
				ipnum, _ = strconv.Atoi(sections[s])
				addresses[a].SortOrder += fmt.Sprintf("%03d", ipnum)
			}
		}
		if includeConnectivity {
			addresses[a].Connectivity = connectivity[addresses[a].ID]
		}

	}

	seenRecently = lastSeen

	for d := range devices {
		for n := range netfaces {
			if devices[d].ID == netfaces[n].DeviceID {
				for a := range addresses {
					if addresses[a].LastSeen.After(seenRecently) {
						seenRecently = addresses[a].LastSeen
					}
					if netfaces[n].ID == addresses[a].InterfaceID {
						for h := range hostnames {
							if addresses[a].ID == hostnames[h].AddressID {
								addresses[a].Hostnames = append(addresses[a].Hostnames, hostnames[h])
							}
						}
						netfaces[n].Addresses = append(netfaces[n].Addresses, addresses[a])
						if netfaces[n].SortOrder == "" {
							netfaces[n].SortOrder = addresses[a].SortOrder
							netfaces[n].Primary.IPv4 = addresses[a].IPv4.String
							netfaces[n].Primary.IPv6 = addresses[a].IPv6.String
							netfaces[n].Primary.IsVirtualIP = utils.Ternary(addresses[a].IsVirtual.Valid, addresses[a].IsVirtual.Bool, false).(bool)
							netfaces[n].Primary.IsReservedIP = utils.Ternary(addresses[a].IsReserved.Valid, addresses[a].IsReserved.Bool, false).(bool)
						}
					}
				}
				if len(netfaces[n].Addresses) > 0 {
					devices[d].Interfaces = append(devices[d].Interfaces, netfaces[n])
					if devices[d].SortOrder == "" {
						devices[d].SortOrder = netfaces[n].SortOrder
						devices[d].Primary = netfaces[n].Primary
						devices[d].Primary.MAC = netfaces[n].MAC
						devices[d].Primary.InterfaceTypeID = netfaces[n].InterfaceTypeID
						devices[d].Primary.VlanID = netfaces[n].VlanID
						devices[d].Primary.IsVirtualIFace = netfaces[n].IsVirtual
					}
				}
			}
		}
	}

	return
}

func GetSpecificAddressConnectivityData(addressID int64, connTime string) (connections []models.Connections, err error) {
	sql := `SELECT scans.time, CASE WHEN address_id IS NULL THEN false ELSE true END as state FROM scans LEFT JOIN histories ON histories.scan_id = scans.id AND histories.address_id = ? WHERE scans.time > ? ORDER BY scans.id DESC;`
	err = db.Conn.NewRaw(sql, addressID, connTime).Scan(db.Context, &connections)
	if err != nil {
		logging.Errorf("Error getting connectivity data: %v", err)
		return
	}
	sort.Sort(Connectivity(connections))
	return
}

func GetDeviceChildren(devicePK int64) (children DeviceChildren, err error) {
	logging.Verbosity = 3
	err = db.Conn.NewRaw(`SELECT GROUP_CONCAT(interfaces.id) as interfaces, GROUP_CONCAT(addresses.id) as addresses FROM devices JOIN interfaces ON devices.id = interfaces.device_id JOIN addresses on interfaces.id = addresses.interface_id WHERE devices.id = ?`, devicePK).Scan(db.Context, &children)
	return
}

func GetIntefaceChildren(interfacePK int64) (children string, err error) {
	logging.Verbosity = 3
	err = db.Conn.NewRaw(`SELECT GROUP_CONCAT(addresses.id) as addresses FROM devices JOIN addresses on interfaces.id = addresses.interface_id WHERE interfaces.id = ?`, interfacePK).Scan(db.Context, &children)
	return
}

// @TODO: Swap to soft delete
func DeleteDevice(devicePK int64) (err error) {
	var children DeviceChildren
	var ids []int64
	children, err = GetDeviceChildren(devicePK)
	if err != nil {
		return
	}
	ids, err = utils.CSVtoInt64s(children.Addresses)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.History)(nil)).Where(`address_id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Hostname)(nil)).Where(`address_id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Address)(nil)).Where(`id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}

	ids, err = utils.CSVtoInt64s(children.Interfaces)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Interface)(nil)).Where(`id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}

	_, err = db.Conn.NewDelete().Model((*models.Device)(nil)).Where(`id  = ?`, devicePK).Exec(db.Context)
	if err != nil {
		return
	}
	return
}

// @TODO: Swap to soft delete
func DeleteInterface(interfacePK int64) (err error) {
	var addresses string
	var ids []int64
	addresses, err = GetIntefaceChildren(interfacePK)
	if err != nil {
		return
	}
	ids, err = utils.CSVtoInt64s(addresses)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.History)(nil)).Where(`address_id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Hostname)(nil)).Where(`address_id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Address)(nil)).Where(`id in (?)`, bun.In(ids)).Exec(db.Context)
	if err != nil {
		return
	}

	_, err = db.Conn.NewDelete().Model((*models.Interface)(nil)).Where(`id  = ?`, interfacePK).Exec(db.Context)
	if err != nil {
		return
	}
	return
}

// @TODO: Swap to soft delete
func DeleteAddress(addressPK int64) (err error) {
	_, err = db.Conn.NewDelete().Model((*models.History)(nil)).Where(`address_id = ?`, addressPK).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Hostname)(nil)).Where(`address_id = ?`, addressPK).Exec(db.Context)
	if err != nil {
		return
	}
	_, err = db.Conn.NewDelete().Model((*models.Address)(nil)).Where(`address_id = ?`, addressPK).Exec(db.Context)
	if err != nil {
		return
	}
	return
}
