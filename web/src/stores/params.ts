import { defineStore } from "pinia"
import { mande } from "mande"

const params = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api/") : mande("/api/")

export interface ParamState {
  columnSorts: ColumnSorts
  Categories: Array<Category>
  Statuses: Array<Status>
  DeviceTypes: Array<DeviceType>
  InterfaceTypes: Array<InterfaceType>
  Locations: Array<Location>
  Maintainers: Array<Maintainer>
  OperatingSystems: Array<OperatingSystem>
  Architectures: Array<Architecture>
  VLANs: Array<VLAN>
  Users: Array<User>
}

export interface ColumnSorts {
  Categories: ColumnSort
  Statuses: ColumnSort
  DeviceTypes: ColumnSort
  InterfaceTypes: ColumnSort
  Locations: ColumnSort
  Maintainers: ColumnSort
  OperatingSystems: ColumnSort
  Architectures: ColumnSort
  VLANs: ColumnSort
  Users: ColumnSort
}

export interface ColumnSort {
  column: string
  direction: string
}

export interface Category {
  ID: number | null
  Label: string
  Icon: string
  Notes: string | null
  IsLocked: boolean
}
export interface Status {
  ID: number | null
  Label: string
  Icon: string
  Notes: string | null
  IsLocked: boolean
}
export interface DeviceType {
  ID: number | null
  Label: string
  Icon: string
  Notes: string | null
  IsLocked: boolean
}
export interface InterfaceType {
  ID: number | null
  Label: string
  Icon: string
  Notes: string | null
  IsLocked: boolean
}
export interface Location {
  ID: number | null
  Label: string
  IsCloud: number
  Notes: string | null
  IsLocked: boolean
}
export interface Maintainer {
  ID: number | null
  Label: string
  IsInternal: boolean
  Notes: string | null
  IsLocked: boolean
}
export interface OperatingSystem {
  ID: number | null
  Vendor: string
  Family: string
  Name: string
  Version: string
  IsOpenSource: boolean
  IsServer: boolean
  Notes: string | null
  IsLocked: boolean
}
export interface Architecture {
  ID: number | null
  Label: string
  BitSpace: number
  Notes: string | null
  IsLocked: boolean
}
export interface VLAN {
  ID: number | null
  Label: string
  Maskv4: string
  Maskv6: string
  Notes: string | null
  IsLocked: boolean
}
export interface User {
  ID: number | null
  Label: string
  Username: string
  Password: string
  NewPassword: string | null
  AccessLevel: string
  Notes: string | null
  IsLocked: boolean
  CanAuthenticate: boolean
}

export type CategoryArr = Array<Category>
export type StatusArr = Array<Status>
export type DeviceTypeArr = Array<DeviceType>
export type InterfaceTypeArr = Array<InterfaceType>
export type LocationArr = Array<Location>
export type MaintainerArr = Array<Maintainer>
export type OperatingSystemArr = Array<OperatingSystem>
export type VLANArr = Array<VLAN>
export type ArchitectureArr = Array<Architecture>
export type UserArr = Array<User>

export const useParams = defineStore("params", {
  state: () => {
    return {
      columnSorts: {
        Categories: {
          column: "ID",
          direction: "asc",
        },
        Statuses: {
          column: "ID",
          direction: "asc",
        },
        DeviceTypes: {
          column: "ID",
          direction: "asc",
        },
        InterfaceTypes: {
          column: "ID",
          direction: "asc",
        },
        Locations: {
          column: "ID",
          direction: "asc",
        },
        Maintainers: {
          column: "ID",
          direction: "asc",
        },
        OperatingSystems: {
          column: "ID",
          direction: "asc",
        },
        Architectures: {
          column: "ID",
          direction: "asc",
        },
        VLANs: {
          column: "ID",
          direction: "asc",
        },
        Users: {
          column: "ID",
          direction: "asc",
        },
      },
      Categories: [],
      Statuses: [],
      DeviceTypes: [],
      InterfaceTypes: [],
      Locations: [],
      Maintainers: [],
      OperatingSystems: [],
      Architectures: [],
      VLANs: [],
      Users: [],
    } as ParamState
  },
  actions: {
    getParams() {
      params.get<ParamState>("/params").then((response) => {
        Object.keys(response).forEach((key) => {
          switch (key) {
            case "Categories":
              this.Categories = response.Categories
              break
            case "Statuses":
              this.Statuses = response.Statuses
              break
            case "DeviceTypes":
              this.DeviceTypes = response.DeviceTypes
              break
            case "InterfaceTypes":
              this.InterfaceTypes = response.InterfaceTypes
              break
            case "Maintainers":
              this.Maintainers = response.Maintainers
              break
            case "Locations":
              this.Locations = response.Locations
              break
            case "OperatingSystems":
              this.OperatingSystems = response.OperatingSystems
              break
            case "Architectures":
              this.Architectures = response.Architectures
              break
            case "VLANs":
              this.VLANs = response.VLANs
              break
            case "Users":
              this.Users = response.Users
              break
          }
        })
      })
    },

    setColumnSort(mode: string, colSort: ColumnSort) {
      switch (mode) {
        case "Category":
          this.columnSorts.Categories = colSort
          break
        case "Status":
          this.columnSorts.Statuses = colSort
          break
        case "DeviceType":
          this.columnSorts.DeviceTypes = colSort
          break
        case "InterfaceType":
          this.columnSorts.InterfaceTypes = colSort
          break
        case "Location":
          this.columnSorts.Locations = colSort
          break
        case "Architecture":
          this.columnSorts.Architectures = colSort
          break
        case "OperatingSystem":
          this.columnSorts.OperatingSystems = colSort
          break
        case "User":
          this.columnSorts.Users = colSort
          break
        case "Maintainer":
          this.columnSorts.Maintainers = colSort
          break
        case "VLAN":
          this.columnSorts.VLANs = colSort
          break
      }
    },
    SortCategories() {
      if (this.columnSorts.Categories.column !== "ID") {
        alert("SortCategories: " + this.columnSorts.Categories.column)
      }
      this.Categories.sort((a: Category, b: Category) => {
        switch (this.columnSorts.Categories.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Categories.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Categories.direction === "asc" ? 1 : -1
        }
      })
    },
    SortStatuses() {
      if (this.columnSorts.Statuses.column !== "ID") {
        alert("SortStatuses: " + this.columnSorts.Statuses.column)
      }
      this.Statuses.sort((a: Status, b: Status) => {
        switch (this.columnSorts.Statuses.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Statuses.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Statuses.direction === "asc" ? 1 : -1
        }
      })
    },
    SortDeviceTypes() {
      if (this.columnSorts.DeviceTypes.column !== "ID") {
        alert("SortDeviceTypes: " + this.columnSorts.DeviceTypes.column)
      }
      this.DeviceTypes.sort((a: DeviceType, b: DeviceType) => {
        switch (this.columnSorts.DeviceTypes.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.DeviceTypes.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.DeviceTypes.direction === "asc" ? 1 : -1
        }
      })
    },
    SortInterfaceTypes() {
      if (this.columnSorts.InterfaceTypes.column !== "ID") {
        alert("SortInterfaceTypes: " + this.columnSorts.InterfaceTypes.column)
      }
      this.InterfaceTypes.sort((a: InterfaceType, b: InterfaceType) => {
        switch (this.columnSorts.InterfaceTypes.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.InterfaceTypes.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.InterfaceTypes.direction === "asc" ? 1 : -1
        }
      })
    },
    SortLocations() {
      if (this.columnSorts.Locations.column !== "ID") {
        alert("SortLocations: " + this.columnSorts.Locations.column)
      }
      this.Locations.sort((a: Location, b: Location) => {
        switch (this.columnSorts.Locations.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Locations.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Locations.direction === "asc" ? 1 : -1
        }
      })
    },
    SortArchitectures() {
      if (this.columnSorts.Architectures.column !== "ID") {
        alert("SortArchitectures: " + this.columnSorts.Architectures.column)
      }
      this.Architectures.sort((a: Architecture, b: Architecture) => {
        switch (this.columnSorts.Architectures.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Architectures.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Architectures.direction === "asc" ? 1 : -1
        }
      })
    },
    SortOperatingSystems() {
      if (this.columnSorts.OperatingSystems.column !== "ID") {
        alert("SortOperatingSystems: " + this.columnSorts.OperatingSystems.column)
      }
      this.OperatingSystems.sort((a: OperatingSystem, b: OperatingSystem) => {
        switch (this.columnSorts.OperatingSystems.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.OperatingSystems.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.OperatingSystems.direction === "asc" ? 1 : -1
        }
      })
    },
    SortUsers() {
      if (this.columnSorts.Users.column !== "ID") {
        alert("SortUsers: " + this.columnSorts.Users.column)
      }
      this.Users.sort((a: User, b: User) => {
        switch (this.columnSorts.Users.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Users.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Users.direction === "asc" ? 1 : -1
        }
      })
    },
    SortMaintainers() {
      if (this.columnSorts.Maintainers.column !== "ID") {
        alert("SortMaintainers: " + this.columnSorts.Maintainers.column)
      }
      this.Maintainers.sort((a: Maintainer, b: Maintainer) => {
        switch (this.columnSorts.Maintainers.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.Maintainers.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.Maintainers.direction === "asc" ? 1 : -1
        }
      })
    },
    SortVLANs() {
      if (this.columnSorts.VLANs.column !== "ID") {
        alert("SortVLANs: " + this.columnSorts.VLANs.column)
      }
      this.VLANs.sort((a: VLAN, b: VLAN) => {
        switch (this.columnSorts.VLANs.column) {
          default:
            if ((a.ID ?? 0) < (b.ID ?? 0)) {
              return this.columnSorts.VLANs.direction === "asc" ? -1 : 1
            }
            return this.columnSorts.VLANs.direction === "asc" ? 1 : -1
        }
      })
    },

    NewCategory() {
      this.Categories.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
    },
    SaveCategory(index: number) {
      const isNew = this.Categories[index].ID === null || this.Categories[index].ID === 0
      params.post("/categories", this.Categories[index]).then((response) => {
        this.Categories[index].ID = (response as Category).ID
        if (isNew) {
          const move = this.Categories.splice(index, 1)
          move.forEach((item) => {
            this.Categories.push(item)
          })
        }
      })
    },
    DropCategory(index: number) {
      this.Categories.splice(index, 1)
    },
    DeleteCategory(index: number) {
      params.delete("/categories/" + this.Categories[index].ID).then(() => {
        this.Categories.splice(index, 1)
      })
    },

    NewStatus() {
      this.Statuses.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
    },
    SaveStatus(index: number) {
      const isNew = this.Statuses[index].ID === null || this.Statuses[index].ID === 0
      params.post("/statuses", this.Statuses[index]).then((response) => {
        this.Statuses[index].ID = (response as Status).ID
        if (isNew) {
          const move = this.Statuses.splice(index, 1)
          move.forEach((item) => {
            this.Statuses.push(item)
          })
        }
      })
    },
    DropStatus(index: number) {
      this.Statuses.splice(index, 1)
    },
    DeleteStatus(index: number) {
      params.delete("/statuses/" + this.Statuses[index].ID).then(() => {
        this.Statuses.splice(index, 1)
      })
    },

    NewMaintainer() {
      this.Maintainers.splice(0, 0, { ID: null, Label: "", Notes: null, IsInternal: false, IsLocked: false })
    },
    SaveMaintainer(index: number) {
      const isNew = this.Maintainers[index].ID === null || this.Maintainers[index].ID === 0

      params.post("/maintainers", this.Maintainers[index]).then((response) => {
        this.Maintainers[index].ID = (response as Maintainer).ID
        if (isNew) {
          const move = this.Maintainers.splice(index, 1)
          move.forEach((item) => {
            this.Maintainers.push(item)
          })
        }
      })
    },
    DropMaintainer(index: number) {
      this.Maintainers.splice(index, 1)
    },
    DeleteMaintainer(index: number) {
      params.delete("/maintainers/" + this.Maintainers[index].ID).then(() => {
        this.Maintainers.splice(index, 1)
      })
    },

    NewLocation() {
      this.Locations.splice(0, 0, { ID: null, Label: "", IsCloud: 0, Notes: null, IsLocked: false })
    },
    SaveLocation(index: number) {
      const isNew = this.Locations[index].ID === null || this.Locations[index].ID === 0

      params.post("/locations", this.Locations[index]).then((response) => {
        this.Locations[index].ID = (response as Location).ID
        if (isNew) {
          const move = this.Locations.splice(index, 1)
          move.forEach((item) => {
            this.Locations.push(item)
          })
        }
      })
    },
    DropLocation(index: number) {
      this.Locations.splice(index, 1)
    },
    DeleteLocation(index: number) {
      params.delete("/locations/" + this.Locations[index].ID).then(() => {
        this.Locations.splice(index, 1)
      })
    },

    NewDeviceType() {
      this.DeviceTypes.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
    },
    SaveDeviceType(index: number) {
      const isNew = this.DeviceTypes[index].ID === null || this.DeviceTypes[index].ID === 0

      params.post("/deviceTypes", this.DeviceTypes[index]).then((response) => {
        this.DeviceTypes[index].ID = (response as DeviceType).ID
        if (isNew) {
          const move = this.DeviceTypes.splice(index, 1)
          move.forEach((item) => {
            this.DeviceTypes.push(item)
          })
        }
      })
    },
    DropDeviceType(index: number) {
      this.DeviceTypes.splice(index, 1)
    },
    DeleteDeviceType(index: number) {
      params.delete("/deviceTypes/" + this.DeviceTypes[index].ID).then(() => {
        this.DeviceTypes.splice(index, 1)
      })
    },

    NewOperatingSystem() {
      this.OperatingSystems.splice(0, 0, { ID: null, Vendor: "", Family: "", Name: "", Version: "", IsOpenSource: false, IsServer: false, Notes: null, IsLocked: false })
    },
    SaveOperatingSystem(index: number) {
      const isNew = this.OperatingSystems[index].ID === null || this.OperatingSystems[index].ID === 0
      params.post("/operatingSystems", this.OperatingSystems[index]).then((response) => {
        this.OperatingSystems[index].ID = (response as OperatingSystem).ID
        if (isNew) {
          const move = this.OperatingSystems.splice(index, 1)
          move.forEach((item) => {
            this.OperatingSystems.push(item)
          })
        }
      })
    },
    DropOperatingSystem(index: number) {
      this.OperatingSystems.splice(index, 1)
    },
    DeleteOperatingSystem(index: number) {
      params.delete("/operatingSystems/" + this.OperatingSystems[index].ID).then(() => {
        this.OperatingSystems.splice(index, 1)
      })
    },

    NewArchitecture() {
      this.Architectures.splice(0, 0, { ID: null, Label: "", BitSpace: 0, Notes: null, IsLocked: false })
    },
    SaveArchitecture(index: number) {
      const isNew = this.Architectures[index].ID === null || this.Architectures[index].ID === 0

      params.post("/architectures", this.Architectures[index]).then((response) => {
        this.Architectures[index].ID = (response as Architecture).ID
        if (isNew) {
          const move = this.Architectures.splice(index, 1)
          move.forEach((item) => {
            this.Architectures.push(item)
          })
        }
      })
    },
    DropArchitecture(index: number) {
      this.Architectures.splice(index, 1)
    },
    DeleteArchitecture(index: number) {
      params.delete("/architectures/" + this.Architectures[index].ID).then(() => {
        this.Architectures.splice(index, 1)
      })
    },

    NewVLAN() {
      this.VLANs.splice(0, 0, { ID: null, Label: "", Maskv4: "", Maskv6: "", Notes: null, IsLocked: false })
    },
    SaveVLAN(index: number) {
      const isNew = this.VLANs[index].ID === null || this.VLANs[index].ID === 0

      params.post("/vlans", this.VLANs[index]).then((response) => {
        this.VLANs[index].ID = (response as VLAN).ID
        if (isNew) {
          const move = this.VLANs.splice(index, 1)
          move.forEach((item) => {
            this.VLANs.push(item)
          })
        }
      })
    },
    DropVLAN(index: number) {
      this.VLANs.splice(index, 1)
    },
    DeleteVLAN(index: number) {
      params.delete("/vlans/" + this.VLANs[index].ID).then(() => {
        this.VLANs.splice(index, 1)
      })
    },

    NewUser() {
      this.Users.splice(0, 0, { ID: null, Label: "", NewPassword: null, Username: "", Password: "", AccessLevel: "", Notes: null, CanAuthenticate: true, IsLocked: false })
    },
    SaveUser(index: number) {
      const isNew = this.Users[index].ID === null || this.Users[index].ID === 0

      params.post("/users", this.Users[index]).then((response) => {
        this.Users[index].ID = (response as User).ID
        if (isNew) {
          const move = this.Users.splice(index, 1)
          move.forEach((item) => {
            this.Users.push(item)
          })
        }
      })
    },
    DropUser(index: number) {
      this.Users.splice(index, 1)
    },
    DeleteUser(index: number) {
      params.delete("/users/" + this.Users[index].ID).then(() => {
        this.Users.splice(index, 1)
      })
    },
  },
})
