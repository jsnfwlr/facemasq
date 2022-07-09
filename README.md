# faceMasq

A containerised tool that provides a responsive UI that allows for monitoring devices on the network and exporting PiHole/DNSMasq DNS and DHCP records for the monitored devices.

A blend of DNS/DHCP editor, asset registrar, network monitor, and intruder detector.

## Current Features
- Network scanning
- Port scanning
- Ability to export records as DHCP and DNS configs for DNSMasq and PiHole
- CRUD editor for taxonomy
- SQLite DB support


## Roadmap

### Features

* = already in progress

#### For v1.0.0
- Interface to investigate unknown devices*
- User authentication
- Add MQTT support
  - allow integration with Home Assistant as a device tracker
- Implement notifications via a plugin system
  - MQTT
  - Gotify
  - Email
- Swap to plugins for DB support to allow others to add support for their own requirements*
  - SQLite*
  - MariaDB*
  - PostgreSQL*
  - MySQL

#### Beyond v1.0.0
- Peripheral tracking
- Multi-user access control
- Network map generation

### Codebase
- Convert the JavaScript to Typescript


## Contributing

Contributions are always welcome!

Just open a PR.
## License

[MIT](https://choosealicense.com/licenses/mit/)


## Installation

The objective for v1.0.0 will be to use the Dockerfile to build your own container and the docker-compose.yml to run the official container image.

