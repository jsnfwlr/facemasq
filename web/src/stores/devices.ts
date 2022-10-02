import { defineStore } from "pinia"
import { mande } from "mande"
import { useStorage } from "@vueuse/core"
import { format } from "date-fns"

const deviceRecords = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api/records") : mande("/api/records")

export interface Device {
  ID: number | null
  MachineName: string
  Brand: string | null
  Model: string | null
  Purchased: string | null
  Serial: string | null
  IsTracked: boolean
  FirstSeen: string | null
  IsGuest: boolean
  IsOnline: boolean
  Label: string | null
  Notes: string | null
  CategoryID: number
  StatusID: number
  MaintainerID: number
  LocationID: number
  DeviceTypeID: number
  OperatingSystemID: number
  ArchitectureID: number
  Interfaces: Array<Netface>
  Primary: PrimaryConnection
  SortOrder: string
}

export interface Netface {
  ID: number
  MAC: string
  IsPrimary: boolean
  IsVirtual: boolean
  IsOnline: boolean
  Label: string | null
  Notes: string | null
  LastSeen: string
  StatusID: number
  InterfaceTypeID: number
  VlanID: number
  DeviceID: number | null
  Primary: PrimaryConnection
  Addresses: Array<Address>
}

export interface Address {
  ID: number
  IPv4: string | null
  IPv6: string | null
  IsPrimary: boolean
  IsVirtual: boolean
  IsReserved: boolean
  LastSeen: string | null
  Label: string | null
  Notes: string | null
  InterfaceID: number
  Connectivity: Array<Connection> | null
  Hostnames: Array<DomainName>
}

export interface Connection {
  State: boolean
  Time: string
}

export interface DomainName {
  ID: number
  Hostname: string
  IsDNS: boolean
  IsSelfSet: boolean
  Notes: string | null
  AddressID: number | null
}

export interface ChartValues {
  full: Array<TimeLog>
  averaged: Array<TimeLog>
}

export interface TimeLog {
  Time: string
  Addresses: number
}

export interface Trend {
  Label: string
  Current: number
  Compare: number
  Tooltip: string | null
}

export interface ColumnSort {
  column: string
  direction: string
}

export interface DeviceState {
  columnSort: ColumnSort
  allDevices: Array<Device>
  activeDevices: Array<Device>
  unknownDevices: Array<Device>
  devicesOverTime: ChartValues
  trends: Array<Trend>
  lastUnknown: string
  editingItems: EditSets
  focusedItems: FocusSets
  deletingItems: DeleteSets
  investigations: Map<number, Array<Investigation>>
}

export interface PrimaryConnection {
  IPv4: string | null
  IPv6: string | null
  InterfaceTypeID: number
  VlanID: number
  MAC: string | null
  IsReservedIP: boolean | null
  IsVirtualIP: boolean | null
  IsVirtualIFace: boolean | null
}

export interface EditSets {
  devices: Map<string, Device>
  interfaces: Map<string, Netface>
  addresses: Map<string, Address>
  hostnames: Map<string, DomainName>
}

export interface FocusSets {
  devices: Map<string, Device>
  interfaces: Map<string, Netface>
  addresses: Map<string, Address>
  hostnames: Map<string, DomainName>
}

export interface DeleteSets {
  devices: Array<number>
  interfaces: Array<number>
  addresses: Array<number>
  hostnames: Array<number>
}

export interface Investigation {
  AddressID: number
  Connectivity: Array<Connection> | null
}

export const useDevices = defineStore("devices", {
  state: () => {
    return {
      columnSort: {
        column: "SortOrder",
        direction: "asc",
      },
      // allDevices: [],
      activeDevices: [],
      unknownDevices: [],
      devicesOverTime: {
        full: [],
        averaged: [],
      },
      trends: [],
      lastUnknown: "2021-12-31 08:55:00",
      investigations: new Map<number, Array<Investigation>>(),
      editingItems: {
        devices: new Map<string, Device>(),
        interfaces: new Map<string, Netface>(),
        addresses: new Map<string, Address>(),
        hostnames: new Map<string, DomainName>(),
      },
      focusedItems: {
        devices: new Map<string, Device>(),
        interfaces: new Map<string, Netface>(),
        addresses: new Map<string, Address>(),
        hostnames: new Map<string, DomainName>(),
      },
      deletingItems: {
        devices: [],
        interfaces: [],
        addresses: [],
        hostnames: [],
      },
      allDevices: useStorage("allDevices", []) as unknown as Array<Device>,
    } as DeviceState
  },
  getters: {
    canChangeSort: (state) => {
      return state.editingItems.devices.size === 0 && state.editingItems.interfaces.size === 0 && state.editingItems.addresses.size === 0 && state.editingItems.hostnames.size === 0
    },
    edItems: (state) => {
      return state.editingItems.devices.size + ", " + state.editingItems.interfaces.size + ", " + state.editingItems.addresses.size + ", " + state.editingItems.hostnames.size
    },
  },
  actions: {
    getTrends() {
      deviceRecords.get<Array<Trend>>("/trends").then((response) => {
        this.trends = response
        // for (let i = 0; i < response.length; i++) {
        //   for (let j = 0; j < this.trends.length; j++) {
        //     if (response[i].Label === this.trends[j].Label) {
        //       this.trends[j].Current = response[i].Current
        //       this.trends[j].Compare = response[i].Compare
        //       break
        //     }
        //   }
        // }
      })
    },
    getCharts() {
      deviceRecords.get<ChartValues>("/chart").then((response) => {
        // this.devicesOverTime = []
        // for (let i = 1400; i < response.length; i++) {
        //   this.devicesOverTime.push(response[i])
        // }
        this.devicesOverTime = response
      })
    },
    getAll() {
      deviceRecords.get<Array<Device>>("/all").then((response) => {
        if (this.canChangeSort) {
          this.allDevices = response
          this.SortDevices()
        }
      })
    },
    investigateDevice(deviceID: number) {
      const addresses = [] as Array<number>
      this.allDevices
        .find((dev) => dev.ID === deviceID)
        ?.Interfaces.forEach((Netface) => {
          Netface.Addresses.forEach((address) => {
            addresses.push(address.ID)
          })
        })
      deviceRecords.post("/investigate", addresses).then((response) => {
        this.investigations.set(deviceID, response as Array<Investigation>)
      })
    },
    Add(indexes: Array<number>) {
      const hostname = { ID: 0, Hostname: "", IsDNS: false, IsSelfSet: false, Notes: null, AddressID: 0 } as DomainName
      const address = { ID: 0, IPv4: null, IPv6: null, IsPrimary: false, IsVirtual: false, IsReserved: false, LastSeen: format(new Date(), "yyyy-MM-dd") + "T" + format(new Date(), "HH:mm:ss") + "Z", Label: "", Notes: null, InterfaceID: 0, Connectivity: null, Hostnames: [hostname] } as Address
      const primary = { IPv4: null, IPv6: null, InterfaceTypeID: 0, VlanID: 0, MAC: null, IsReservedIP: null, IsVirtualIP: null, IsVirtualIFace: null } as PrimaryConnection
      const iFace = { ID: 0, MAC: "", IsPrimary: true, IsVirtual: false, Label: "", Notes: null, InterfaceTypeID: 1, VlanID: 1, LastSeen: format(new Date(), "yyyy-MM-dd") + "T" + format(new Date(), "HH:mm:ss") + "Z", StatusID: 1, DeviceID: 0, Primary: primary, Addresses: [address] } as Netface
      const device = {
        ID: 0,
        Label: "",
        MachineName: "",
        FirstSeen: format(new Date(), "yyyy-MM-dd") + "T" + format(new Date(), "HH:mm:ss") + "Z",
        CategoryID: 1,
        StatusID: 1,
        MaintainerID: 1,
        LocationID: 1,
        DeviceTypeID: 1,
        OperatingSystemID: 1,
        ArchitectureID: 1,
        Notes: null,
        Brand: null,
        Model: null,
        Purchased: null,
        Serial: null,
        IsGuest: false,
        IsTracked: true,
        SortOrder: "",
        Interfaces: [iFace],
        Primary: primary,
      } as Device

      switch (indexes.length) {
        case 0: // Device []
          this.allDevices.splice(0, 0, device)
          this.editingItems.devices.set("0", this.allDevices[0])
          this.editingItems.interfaces.set("0", this.allDevices[0].Interfaces[0])
          this.editingItems.addresses.set("0", this.allDevices[0].Interfaces[0].Addresses[0])
          this.editingItems.hostnames.set("0", this.allDevices[0].Interfaces[0].Addresses[0].Hostnames[0])
          break
        case 1: // Interface [device]
          this.allDevices[indexes[0]].Interfaces.splice(0, 0, iFace)
          this.editingItems.interfaces.set("0", this.allDevices[0].Interfaces[0])
          this.editingItems.addresses.set("0", this.allDevices[0].Interfaces[0].Addresses[0])
          this.editingItems.hostnames.set("0", this.allDevices[0].Interfaces[0].Addresses[0].Hostnames[0])
          break
        case 2: // Address [device, interface]
          if (indexes[1] !== -1) {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.splice(0, 0, address)
            this.editingItems.addresses.set("0", this.allDevices[0].Interfaces[0].Addresses[0])
            this.editingItems.hostnames.set("0", this.allDevices[0].Interfaces[0].Addresses[0].Hostnames[0])
          }
          break
        case 3: // Hostname [device, interface, address]
          if (indexes[2] !== -1) {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.splice(0, 0, hostname)
            this.editingItems.hostnames.set("0", this.allDevices[0].Interfaces[0].Addresses[0].Hostnames[0])
          }
          break
      }
    },
    setColumnSort(colSort: ColumnSort) {
      if (this.editingItems.devices.size === 0 && this.editingItems.interfaces.size === 0 && this.editingItems.addresses.size === 0 && this.editingItems.hostnames.size === 0) {
        this.columnSort = colSort
      }
    },
    SortDevices() {
      this.allDevices.sort((a: Device, b: Device) => {
        switch (this.columnSort.column) {
          case "MachineName":
            if (a.MachineName === b.MachineName) {
              return 0
            } else if (a.MachineName < b.MachineName) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1

          case "Label":
            if (a.Label === b.Label) {
              return 0
            } else if ((a.Label ?? "") < (b.Label ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          case "MAC":
            if (a.Interfaces[0].MAC === b.Interfaces[0].MAC) {
              return 0
            } else if ((a.Interfaces[0].MAC ?? "") < (b.Interfaces[0].MAC ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1

          case "FirstSeen":
            if (a.FirstSeen === b.FirstSeen) {
              return 0
            } else if ((a.FirstSeen ?? "") < (b.FirstSeen ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1

          case "LastSeen":
            if (a.Interfaces[0].LastSeen === b.Interfaces[0].LastSeen) {
              return 0
            } else if ((a.Interfaces[0].LastSeen ?? "") < (b.Interfaces[0].LastSeen ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          case "Serial":
            if (a.Serial === b.Serial) {
              return 0
            } else if ((a.Serial ?? "") < (b.Serial ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1

          case "Location":
            if (a.LocationID === b.LocationID) {
              return 0
            } else if (a.LocationID < b.LocationID) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          case "Maintainer":
            if (a.MaintainerID === b.MaintainerID) {
              return 0
            } else if (a.MaintainerID < b.MaintainerID) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          case "OS":
            if (a.OperatingSystemID === b.OperatingSystemID) {
              return 0
            } else if (a.OperatingSystemID < b.OperatingSystemID) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          case "Brand":
            if (a.Brand === b.Brand) {
              if (a.Model === b.Model) {
                return 0
              } else if ((a.Model ?? "") < (b.Model ?? "")) {
                return this.columnSort.direction === "asc" ? -1 : 1
              }
            } else if ((a.Brand ?? "") < (b.Brand ?? "")) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
          default:
            if (a.SortOrder === b.SortOrder) {
              return 0
            } else if (a.SortOrder < b.SortOrder) {
              return this.columnSort.direction === "asc" ? -1 : 1
            }
            return this.columnSort.direction === "asc" ? 1 : -1
        }
      })
    },
    Discard(indexes: Array<number>) {
      switch (indexes.length) {
        case 1: // Device
          if (this.allDevices[indexes[0]].ID === null || this.allDevices[indexes[0]].ID === 0) {
            this.allDevices.splice(indexes[0], 1)
          } else {
            if (this.editingItems.devices.has(indexes[0].toString())) {
              this.allDevices[indexes[0]] = this.editingItems.devices.get(indexes[0].toString()) as Device
              this.editingItems.devices.delete(indexes[0].toString())
            }
          }
          for (let i = 0; i < this.allDevices[indexes[0]].Interfaces.length; i++) {
            if (this.allDevices[indexes[0]].Interfaces[i].ID === null || this.allDevices[indexes[0]].Interfaces[i].ID === 0) {
              this.allDevices[indexes[0]].Interfaces.splice(i, 1)
            } else {
              if (this.editingItems.interfaces.has(indexes[0] + "." + i)) {
                this.allDevices[indexes[0]].Interfaces[i] = this.editingItems.interfaces.get(indexes[0] + "." + i) as Netface
                this.editingItems.interfaces.delete(indexes[0] + "." + i)
              }
            }
            for (let a = 0; a < this.allDevices[indexes[0]].Interfaces[i].Addresses.length; a++) {
              if (this.allDevices[indexes[0]].Interfaces[i].Addresses[a].ID === null || this.allDevices[indexes[0]].Interfaces[i].Addresses[a].ID === 0) {
                this.allDevices[indexes[0]].Interfaces[i].Addresses.splice(a, 1)
              } else {
                if (this.editingItems.addresses.has(indexes[0] + "." + i + "." + a)) {
                  this.allDevices[indexes[0]].Interfaces[i].Addresses[a] = this.editingItems.addresses.get(indexes[0] + "." + i + "." + a) as Address
                  this.editingItems.addresses.delete(indexes[0] + "." + i + "." + a)
                }
              }
              if (typeof this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames !== "undefined" && this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames !== null) {
                for (let h = 0; h < this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames.length; h++) {
                  if (this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames[h].ID === null || this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames[h].ID === 0) {
                    this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames.splice(h, 1)
                  } else {
                    if (this.editingItems.hostnames.has(indexes[0] + "." + i + "." + a + "." + h)) {
                      this.allDevices[indexes[0]].Interfaces[i].Addresses[a].Hostnames[h] = this.editingItems.hostnames.get(indexes[0] + "." + i + "." + a + "." + h) as DomainName
                      this.editingItems.hostnames.delete(indexes[0] + "." + i + "." + a + "." + h)
                    }
                  }
                }
              }
            }
          }

          break
        case 2: // Interface
          if (this.allDevices[indexes[0]].Interfaces[indexes[1]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].ID === 0) {
            this.allDevices[indexes[0]].Interfaces.splice(indexes[1], 1)
          } else {
            if (this.editingItems.interfaces.has(indexes[0] + "." + indexes[1])) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]] = this.editingItems.interfaces.get(indexes[0] + "." + indexes[1]) as Netface
              this.editingItems.interfaces.delete(indexes[0] + "." + indexes[1])
            }
          }
          for (let a = 0; a < this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.length; a++) {
            if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].ID === 0) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.splice(a, 1)
            } else {
              if (this.editingItems.addresses.has(indexes[0] + "." + indexes[1] + "." + a)) {
                this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a] = this.editingItems.addresses.get(indexes[0] + "." + indexes[1] + "." + a) as Address
                this.editingItems.addresses.delete(indexes[0] + "." + indexes[1] + "." + a)
              }
            }
            if (typeof this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames !== "undefined" && this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames != null) {
              for (let h = 0; h < this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames.length; h++) {
                if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames[h].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames[h].ID === 0) {
                  this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames.splice(h, 1)
                } else {
                  if (this.editingItems.hostnames.has(indexes[0] + "." + indexes[1] + "." + a + "." + h)) {
                    this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].Hostnames[h] = this.editingItems.hostnames.get(indexes[0] + "." + indexes[1] + "." + a + "." + h) as DomainName
                    this.editingItems.hostnames.delete(indexes[0] + "." + indexes[1] + "." + a + "." + h)
                  }
                }
              }
            }
          }

          break
        case 3: // Address
          if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID === 0) {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.splice(indexes[2], 1)
          } else {
            if (this.editingItems.addresses.has(indexes[0] + "." + indexes[1] + "." + indexes[2])) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]] = this.editingItems.addresses.get(indexes[0] + "." + indexes[1] + "." + indexes[2]) as Address
              this.editingItems.addresses.delete(indexes[0] + "." + indexes[1] + "." + indexes[2])
            }
          }
          if (typeof this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames !== "undefined" && this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames != null) {
            for (let h = 0; h < this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.length; h++) {
              if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[h].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[h].ID === 0) {
                this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.splice(h, 1)
              } else {
                if (this.editingItems.hostnames.has(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + h)) {
                  this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[h] = this.editingItems.hostnames.get(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + h) as DomainName
                  this.editingItems.hostnames.delete(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + h)
                }
              }
            }
          }
          break
        case 4: // Hostname
          if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID === 0) {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.splice(indexes[3], 1)
          } else {
            if (this.editingItems.hostnames.has(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + indexes[3])) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]] = this.editingItems.hostnames.get(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + indexes[3]) as DomainName
              this.editingItems.hostnames.delete(indexes[0] + "." + indexes[1] + "." + indexes[2] + "." + indexes[3])
            }
          }
          break
      }
    },
    PerformDelete(indexes: Array<number>) {
      switch (indexes.length) {
        case 1: // Device
          if (!this.deletingItems.devices.includes(indexes[0])) {
            console.log("delete Device ID: " + this.allDevices[indexes[0]].ID)
          }
          break
        case 2: // Interface
          if (!this.deletingItems.interfaces.includes(indexes[0])) {
            console.log("delete Interface ID: " + this.allDevices[indexes[0]].Interfaces[indexes[1]].ID)
          }
          break
        case 3: // Address
          if (!this.deletingItems.addresses.includes(indexes[0])) {
            console.log("delete Address ID: " + this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID)
          }
          break
        case 4: // Hostname
          if (!this.deletingItems.hostnames.includes(indexes[0])) {
            console.log("delete Hostname ID: " + this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID)
          }
          break
      }
    },
    InitiateDelete(indexes: Array<number>) {
      switch (indexes.length) {
        case 1: // Devuce
          if (!this.deletingItems.devices.includes(indexes[0])) {
            this.deletingItems.devices.push(indexes[0])
          }
          break
        case 2: // Interface
          if (!this.deletingItems.interfaces.includes(indexes[1])) {
            this.deletingItems.interfaces.push(indexes[1])
          }
          break
        case 3: // Address
          if (!this.deletingItems.addresses.includes(indexes[2])) {
            this.deletingItems.addresses.push(indexes[2])
          }
          break
        case 4: // Hostname
          if (!this.deletingItems.hostnames.includes(indexes[3])) {
            this.deletingItems.hostnames.push(indexes[3])
          }
          break
      }
    },
    CancelDelete(indexes: Array<number>) {
      switch (indexes.length) {
        case 1: // Device
          this.deletingItems.devices.splice(indexes[0], 1)
          break
        case 2: // Interface
          this.deletingItems.interfaces.splice(indexes[1], 1)
          break
        case 3: // Address
          this.deletingItems.addresses.splice(indexes[2], 1)
          break
        case 4: // Hostname
          this.deletingItems.hostnames.splice(indexes[3], 1)
          break
      }
    },
    Save(indexes: Array<number>) {
      return new Promise((resolve, reject) => {
        let isNew = false
        if (indexes.length === 1) {
          // Device
          isNew = this.allDevices[indexes[0]].ID === null || this.allDevices[indexes[0]].ID === 0
          // @FIXME: if isNew, run tests on all the interfaces, addresses, and hostnames before attempting to save
          const deviceOnly = {
            ID: this.allDevices[indexes[0]].ID,
            MachineName: this.allDevices[indexes[0]].MachineName,
            Brand: this.allDevices[indexes[0]].Brand,
            Model: this.allDevices[indexes[0]].Model,
            Purchased: this.allDevices[indexes[0]].Purchased,
            Serial: this.allDevices[indexes[0]].Serial,
            IsTracked: this.allDevices[indexes[0]].IsTracked,
            IsGuest: this.allDevices[indexes[0]].IsGuest,
            IsOnline: this.allDevices[indexes[0]].IsOnline,
            Label: this.allDevices[indexes[0]].Label,
            FirstSeen: this.allDevices[indexes[0]].FirstSeen,
            Notes: this.allDevices[indexes[0]].Notes,
            CategoryID: this.allDevices[indexes[0]].CategoryID,
            StatusID: this.allDevices[indexes[0]].StatusID,
            MaintainerID: this.allDevices[indexes[0]].MaintainerID,
            LocationID: this.allDevices[indexes[0]].LocationID,
            DeviceTypeID: this.allDevices[indexes[0]].DeviceTypeID,
            OperatingSystemID: this.allDevices[indexes[0]].OperatingSystemID,
            ArchitectureID: this.allDevices[indexes[0]].ArchitectureID,
          }
          deviceRecords.post("/device", deviceOnly).then((response) => {
            this.allDevices[indexes[0]].ID = (response as Device).ID
            if (isNew) {
              this.allDevices[indexes[0]].Interfaces.forEach((_, n) => {
                this.allDevices[indexes[0]].Interfaces[n].DeviceID = (response as Device).ID
                this.Save([indexes[0], n]).then(() => {
                  const move = this.allDevices.splice(indexes[0], 1)
                  move.forEach((item) => {
                    this.allDevices.push(item)
                  })
                  resolve(response)
                })
              })
            } else {
              resolve(response)
            }
          })
          this.editingItems.devices.delete(indexes[0].toString())
        } else if (indexes.length === 2) {
          //Interface
          isNew = this.allDevices[indexes[0]].Interfaces[indexes[1]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].ID === 0
          const interfaceOnly = {
            ID: this.allDevices[indexes[0]].Interfaces[indexes[1]].ID,
            MAC: this.allDevices[indexes[0]].Interfaces[indexes[1]].MAC,
            IsPrimary: this.allDevices[indexes[0]].Interfaces[indexes[1]].IsPrimary,
            IsVirtual: this.allDevices[indexes[0]].Interfaces[indexes[1]].IsVirtual,
            IsOnline: this.allDevices[indexes[0]].Interfaces[indexes[1]].IsOnline,
            LastSeen: this.allDevices[indexes[0]].Interfaces[indexes[1]].LastSeen,
            Label: this.allDevices[indexes[0]].Interfaces[indexes[1]].Label,
            Notes: this.allDevices[indexes[0]].Interfaces[indexes[1]].Notes,
            StatusID: this.allDevices[indexes[0]].Interfaces[indexes[1]].StatusID,
            InterfaceTypeID: this.allDevices[indexes[0]].Interfaces[indexes[1]].InterfaceTypeID,
            VlanID: this.allDevices[indexes[0]].Interfaces[indexes[1]].VlanID,
            DeviceID: this.allDevices[indexes[0]].ID,
          }
          deviceRecords.post("/interface", interfaceOnly).then((response) => {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].ID = (response as Netface).ID
            if (isNew) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.forEach((_, a) => {
                this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[a].InterfaceID = (response as Netface).ID
                this.Save([indexes[0], indexes[1], a]).then(() => {
                  const move = this.allDevices[indexes[0]].Interfaces.splice(indexes[1], 1)
                  move.forEach((item) => {
                    this.allDevices[indexes[0]].Interfaces.push(item)
                  })
                  resolve(response)
                })
              })
            } else {
              resolve(response)
            }
          })
          this.editingItems.interfaces.delete(indexes[0].toString() + "." + indexes[1].toString())
        } else if (indexes.length === 3) {
          //Address
          isNew = this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID === 0

          const addressOnly = {
            ID: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID,
            IPv4: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].IPv4,
            IPv6: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].IPv6,
            IsPrimary: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].IsPrimary,
            IsVirtual: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].IsVirtual,
            IsReserved: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].IsReserved,
            LastSeen: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].LastSeen,
            Label: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Label,
            Notes: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Notes,
            InterfaceID: this.allDevices[indexes[0]].Interfaces[indexes[1]].ID,
          }
          deviceRecords.post("/address", addressOnly).then((response) => {
            this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID = (response as Address).ID
            if (isNew) {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.forEach((_, h) => {
                this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[h].AddressID = (response as Address).ID
                this.Save([indexes[0], indexes[1], indexes[2], h]).then(() => {
                  const move = this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.splice(indexes[2], 1)
                  move.forEach((item) => {
                    this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses.push(item)
                  })
                  resolve(response)
                })
              })
            } else {
              resolve(response)
            }
          })
          this.editingItems.addresses.delete(indexes[0].toString() + "." + indexes[1].toString() + "." + indexes[2].toString())
        } else if (indexes.length === 4) {
          //Hostname
          if (this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].Hostname !== null && this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].Hostname !== "") {
            isNew = this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID === null || this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID === 0
            const hostnameOnly = {
              ID: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID,
              Hostname: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].Hostname,
              IsDNS: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].IsDNS,
              IsSelfSet: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].IsSelfSet,
              Notes: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].Notes,
              AddressID: this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].ID,
            }
            deviceRecords.post("/hostname", hostnameOnly).then((response) => {
              this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames[indexes[3]].ID = (response as DomainName).ID
              if (isNew) {
                const move = this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.splice(indexes[3], 1)
                move.forEach((item) => {
                  this.allDevices[indexes[0]].Interfaces[indexes[1]].Addresses[indexes[2]].Hostnames.push(item)
                })
                resolve(response)
              } else {
                resolve(response)
              }
            })
            this.editingItems.hostnames.delete(indexes[0].toString() + "." + indexes[1].toString() + "." + indexes[2].toString() + "." + indexes[3].toString())
          }
        } else {
          reject("incorrect number of indexes")
        }
      })
    },
    Fake() {
      /*
        ,
        NewHostname(addressID: number) {
  
        },
        SaveHostname(addressID: number, index: number) {
  
        },
        DeleteHostname(addressID: number, index: number) {
  
        },
        DropHostname(addressID: number, index: number) {
          this.allDevices.forEach((device, d) => {
            device.Interfaces.forEach((iFace, i) => {
              iFace.Addresses.forEach((address, a) => {
                if (address.ID === addressID) {
                  this.allDevices[d].Interfaces[i].Addresses[a].Hostnames.splice(index, 1)
                }
              })
            })
          })
        },
        NewAddress(iFaceID: number) {
  
        },
        SaveAddress(iFaceID: number, index: number) {
  
        },
        DeleteAddress(iFaceID: number, index: number) {
  
        },
        DropAddress(iFaceID: number, index: number) {
          this.allDevices.forEach((device, d) => {
            device.Interfaces.forEach((iFace, i) => {
              if (iFace.ID === iFaceID) {
                this.allDevices[d].Interfaces[i].Addresses.splice(index, 1)
              }
  
            })
          })
        },
        NewIFace(deviceID: number) {
  
        },
        SaveIFace(deviceID: number, index: number) {
  
        },
        DeleteIFace(deviceID: number, index: number) {
  
        },
        DropIFace(deviceID: number, index: number) {
          this.allDevices.forEach((device, d) => {
            if (device.ID === deviceID) {
              this.allDevices[d].Interfaces.splice(index, 1)
            }
          })
        }
      */
    },
  },
})
