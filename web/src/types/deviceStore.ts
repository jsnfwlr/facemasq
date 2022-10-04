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
