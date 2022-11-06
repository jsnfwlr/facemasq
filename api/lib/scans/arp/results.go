package arpscan

import (
	"database/sql"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/volatiletech/null"
)

type Result struct {
	ScanID    int64
	Hostname  string
	IPv4      string
	IPv6      null.String
	MAC       string
	FirstSeen string
	LastSeen  string
	ScanCount int
	Notes     string
}

func (record Result) Store() (err error) {
	var netface models.Interface
	var address models.Address
	var sqlres sql.Result

	sql := `SELECT * FROM interfaces WHERE mac = ?;`
	err = db.Conn.NewRaw(sql, record.MAC).Scan(db.Context, &netface)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logging.Error("%v", err)
			return
		}
		sql = `INSERT INTO Devices (first_seen, notes, machine_name, label, is_online, status_id) VALUES (?, ?, "", "Unknown Device", true, 1)`
		sqlres, err = db.Conn.Exec(sql, record.LastSeen, record.Notes)
		if err != nil {
			logging.Error("Error creating device: %v", err)
			return
		}
		lastDeviceID, _ := sqlres.LastInsertId()

		// sql = `SELECT ID FROM Device WHERE Label = ?`
		// err = db.Conn.Get(&device, sql, record.MAC)
		// if err != nil {
		// 	return
		// }

		sql = `INSERT INTO interfaces (mac, device_id, notes, last_seen, is_online, status_id) VALUES (?, ?, ?, ?, true, 1);`
		sqlres, err = db.Conn.Exec(sql, record.MAC, lastDeviceID, record.Notes, record.LastSeen)
		if err != nil {
			logging.Error("Error creating interface: %v", err)
			return
		}
		lastInterfaceID, _ := sqlres.LastInsertId()

		sql = `INSERT INTO addresses (ipv4, ipv6, interface_id, last_seen, notes) VALUES (?, ?, ?, ?, ?);`
		sqlres, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, lastInterfaceID, record.LastSeen, record.Notes)
		if err != nil {
			logging.Error("Error creating address: %v", err)
			return
		}
		lastAddressID, _ := sqlres.LastInsertId()

		if record.Hostname != "" {
			sql = `INSERT INTO hostnames (hostname, address_id, notes) VALUES (?,?, ?);`
			_, err = db.Conn.Exec(sql, record.Hostname, lastAddressID, record.Notes)
			if err != nil {
				logging.Error("Error creating hostname: %v", err)
				return
			}
		}

		sql = `INSERT INTO histories (address_id, scan_id) VALUES (?,?);`
		_, err = db.Conn.Exec(sql, lastAddressID, record.ScanID)
		if err != nil {
			logging.Error("Error creating histories: %v", err)
			return
		}
		return
	}

	sql = `SELECT * FROM addresses WHERE ipv4 = ? AND interface_id = ?`
	err = db.Conn.NewRaw(sql, record.IPv4, netface.ID).Scan(db.Context, &address)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logging.Error("%v", err)
			return
		}

		sql = `INSERT INTO addresses (ipv4, ipv6, interface_id, last_seen, notes) VALUES (?, ?, ?, ?, ?);`
		sqlres, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, netface.ID, record.LastSeen, record.Notes)
		if err != nil {
			logging.Error("Error adding address: %v", err)
			return
		}
		lastAddressID, _ := sqlres.LastInsertId()

		if record.Hostname != "" {
			sql = `INSERT INTO hostnames (hostname, address_id, notes) VALUES (?, ?, ?);`
			_, err = db.Conn.Exec(sql, record.Hostname, lastAddressID, record.Notes)
			if err != nil {
				logging.Error("Error adding hostname: %v", err)
				return
			}
		}

		sql = `INSERT INTO histories (address_id, scan_id) VALUES (?, ?);`
		_, err = db.Conn.Exec(sql, lastAddressID, record.ScanID)
		if err != nil {
			logging.Error("Error adding histories: %v", err)
			return
		}

		sql = `UPDATE addresses SET label = NULL WHERE label = ?;`
		_, err = db.Conn.Exec(sql, record.MAC)
		if err != nil {
			return
		}
		return
	}

	sql = `UPDATE interfaces SET last_seen = ? WHERE id = ?;`
	_, err = db.Conn.Exec(sql, record.LastSeen, netface.ID)
	if err != nil {
		logging.Error("Error recording interfaces lastseen: %v", err)
		return
	}

	sql = `UPDATE addresses SET last_seen = ? WHERE id = ?;`
	_, err = db.Conn.Exec(sql, record.LastSeen, address.ID)
	if err != nil {
		logging.Error("Error recording address lastseen: %v", err)
		return
	}

	sql = `INSERT INTO histories (address_id, scan_id) VALUES (?,?);`
	_, err = db.Conn.Exec(sql, address.ID, record.ScanID)
	if err != nil {
		if err.Error() != "UNIQUE constraint failed: histories.address_id, histories.scan_id" {
			logging.Error("Error recording new histories: %v", err)
		}
		return
	}
	return
}
