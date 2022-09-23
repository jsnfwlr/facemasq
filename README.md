<p align="center"><img src="ui/public/images/mask.png" /></p>

<h1 style="color: #00CC99">faceMasq</h1>

A containerised tool with a responsive UI for monitoring devices on a network and exporting PiHole/DNSMasq DNS and DHCP records for the monitored devices.

A blend of DNS/DHCP editor, asset registrar, network monitor, and intruder detector.

# Current Features
- Network scanning
- Port scanning
- Ability to export records as DHCP and DNS configs for DNSMasq and PiHole
- CRUD editor for taxonomy
- Multi-DB support
  - SQLite
  - MySQL
  - MariaDB
  - PostgreSQL
- User Settings
  - Dark mode
  - System-wide pagination options
  - Page-specific pagination options

# Installation and configuration

faceMasq is intended to be run as a Docker container. Details on how to install and configure it can be found in the [Docker Installation](docs/installation/docker.md) Documents.

If you wish to run it as a bare-metal service, details on how to do that can be found in the [Advanced Installation](advanced.md) Documents.

# Roadmap

## Release < v1.0.0

### Codebase
- Migrate all JavaScript to TypeScript
- Swap to websockets for data transfer

### Additional Input sources
- ARPScan to reduce burden on host device's network interfaces
- Log monitoring for DNSMasq
  - Watch for leases from DNSMasq
- Webhook for third-party systems to report devices
  - [DNSMasq](https://etherarp.net/dnsmasq/index.html#run-an-executable-when-a-dhcp-lease-is-created-or-destroyed.)

### Interface for over-riding various App Settings (currently set via env-vars)
- NetScan frequency
- ARPScan frequency

### User Settings
- Allow setting LastSeen cut off
  - Device's that haven't been seen since the cut off time (specified in hours/days/weeks ago) will be hidden from the standard Devices view

### Full device management
- Move the interfaces/addresses from one device to another to merge them
- Add planned devices
  - Possibility to match new devices to planned devices if the MAC vendor matches the brand
- Delete retired/lost/deprecated devices
- Adding/Editing/Deleting device interfaces
- Adding/Editing/Deleting interface addresses
- Adding/Editing/Deleting address hostnames

### Improved management of DNSMasq and PiHole configuration files
- Update files when changes are applied in faceMasq
- Allow configuration of file names, paths, etc
- Allow multiple instances of DNS and DHCP configuration

### History Log
- Log when devices/interfaces had reserved IP addresses assigned/revoked
- Log when devices have their maintainer/location/status/category/architecture/operating-system changed
- Keep historic log of device labels, notes, hostnames, serial, and machine-names

### User Authentication with RBAC
- Roles: Admin, Network Manager, Device Manager, Viewer
- Two factor authentication via TOTP/HOTP
- Two factor authentication via FIDO2

### Device Watching
- Get alerts when watched devices change IP address
- Get alerts when watched devices go offline

### Unknown Device Monitoring
- View network connectivity of an unknown device
- Track IP address changes of unknown device
- Scan ports of unknown devices
- Get alerts for unknown devices via notifications
- Interface for delving in to available information on unknown devices

### Monthly Reporting via Email notifications
- Allow users to get statistical reports for the devices they maintain or watch 
  - Useful for seeing devices with unstable connections to the network
- Allow admins and network managers to get reports from the Unknown Device Monitor
  - Summarise the month's alerts 
  - Provide details about each of the devices (ports, network connectivity, etc)

### Add MQTT support
- Allow integration with Home Assistant as a device tracker
- Allow integration with Node Red

### Implement notifications via a plugin system
- MQTT
- Gotify
- Email


## Release > v1.0.0
### Peripheral tracking
- Manually add monitors, keyboards, mice, speakers, TVs, printers, 3d printers, uninterruptible power supplies, webcams, security cameras and more to computers, chromecasts, etc.
### Network map generation
- Produce network maps that [/r/homelab](https://reddit.com/r/homelab/) would be proud of.

# Contributing

Contributions are always welcome!

Just open a PR.

Contribution Guide to come.

# License

[MIT](https://choosealicense.com/licenses/mit/)

# Notes for considerations

## Adding FIDO2 2FA
- https://github.com/duo-labs/webauthn or https://github.com/fxamacker/webauthn for the API
- https://github.com/MasterKale/SimpleWebAuthn/tree/master/packages/browser for the UI


