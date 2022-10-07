# Quick Start

## Core Configuration

### Environment Variables

| Variable            | Purpose                                                                                                                                                                  | Options         | Default                   | Example          |
|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------|---------------------------|------------------|
| `PASSWORDSALT`      | The password salt used when hashing all the passwords stored in the database                                                                                             |                 | `NULL`                    | `changeme`       |
| `PORT`              | The port you want the faceMasq to run on within the container - useful if you want the container on the host network or are running the program as a bare-metal service. |                 | `6135`                    | `8080`           |
| `ADMINPASSWORD`     | The password to use as the admin password when you first run faceMasq. Note: this can be removed once your database is setup.                                            |                 | `NULL`                    |                  |
| `BUNDEBUG`          | The verbosity of the database debug messages. Useful for troubleshooting and development. `0` is off, `1` is kinda-verbose, `2` is very verbose.                         | [`0,1,2`]       | `0`                       |                  |
| `VERBOSE`           | Enable verbose messages in the logs. Useful for troubleshooting and development. Assigning this variable any value activates the verbose logging                         |                 | `NULL`                    | `y`              |
| `NETMASK`           | The network mask, written in slash notation, that you want to scan. Leaving this blank disables all scanning processes.                                                  |                 | `NULL`                    | `192.168.0.0/24` |
| `DHCPFILENAME`      | The filename you want the DHCP lease reservations exported to.                                                                                                           |                 | `01.dhcp.conf`            |                  |
| `DNSFILENAME`       | The filename you want the DNS records exported to.                                                                                                                       |                 | `02.dns.conf`             |                  |
| `ARPSCAN_FREQUENCY` |                                                                                                                                                                          | Seconds         | `60`                      |                  |
| `NETSCAN_FREQUENCY` |                                                                                                                                                                          | Seconds         | `60`                      |                  |
| `NETSCAN_TIMEOUT`   |                                                                                                                                                                          | Milliseconds    | `2000`                    |                  |
| `PORTSCAN`          |                                                                                                                                                                          | [`true,false`]  | `false`                   |                  |
| `PORTSCAN_WIDTH`    |                                                                                                                                                                          | [`narrow,wide`] | `narrow`                  |                  |
| `DBCONNSTR`         | See [Database Configuration](#database-configuration) for details                                                                                                        |                 | `sqlite://network.sqlite` |                  |
| `CHART_FREQUENCY`   | `Deprecated - changed to meta-setting`                                                                                                                                   |                 |                           |                  |
|                     |                                                                                                                                                                          |                 |                           |                  |

## Database Configuration

### SQLite

### MySQL

### Volumes

| Volume      | Purpose | Path         |
|-------------|---------|--------------|
| `export`    |         | `/export`    |
| `data`      |         | `/data`      |
| `templates` |         | `/templates` |
