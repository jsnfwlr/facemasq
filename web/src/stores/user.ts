import { defineStore } from "pinia"
import { format } from "date-fns"
import { mande } from "mande"

const settings = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api/settings/") : mande("/api/settings/")
const setting = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api/setting/") : mande("/api/setting/")

export interface UserState {
  account: User
  settings: Settings
}

export interface User {
  ID: number
  Username: string
  Label: string
  AccessLevel: number
}

export interface Settings {
  darkMode: boolean | null
  lastDismissed: string
  dashboardKnownPageSize: number
  dashboardIntruderPageSize: number
  devicesPageSize: number
  categoriesPageSize: number
  statusesPageSize: number
  deviceTypesPageSize: number
  interfaceTypesPageSize: number
  locationsPageSize: number
  maintainersPageSize: number
  operatingSystemsPageSize: number
  architecturesPageSize: number
  vLansPageSize: number
  usersPageSize: number
}

// const topics = ["details", "params", "devices", "users"]
// const actions = ["read", "write", "delete"]

export interface UserSetting {
  Name: string
  Value: string
}

export const useUser = defineStore("user", {
  state: () => {
    return {
      settings: {
        darkMode: null,
        lastDismissed: "2021-12-21 08:55:56",
        dashboardKnownPageSize: 0,
        dashboardIntruderPageSize: 0,
        devicesPageSize: 0,
        categoriesPageSize: 0,
        statusesPageSize: 0,
        deviceTypesPageSize: 0,
        interfaceTypesPageSize: 0,
        locationsPageSize: 0,
        maintainersPageSize: 0,
        operatingSystemsPageSize: 0,
        architecturesPageSize: 0,
        vLansPageSize: 0,
        usersPageSize: 0,
      },
      account: {
        ID: 1,
        Username: "jason@jsnfwlr.com",
        Label: "Jason Fowler",
        AccessLevel: 1,
      },
    } as UserState
  },
  actions: {
    hasAccess(topic: string, action: string) {
      if (topic === "" || topic !== action) {
        return true
      }
    },
    updateSetting(name: string, value: string | boolean | number) {
      switch (name) {
        case "darkMode":
          this.settings.darkMode = value as boolean
          break
        case "lastDismissed":
          this.settings.lastDismissed = value as string
          break
        case "dashboardKnownPageSize":
          this.settings.dashboardKnownPageSize = value as number
          break
        case "dashboardIntruderPageSize":
          this.settings.dashboardIntruderPageSize = value as number
          break
        case "devicesPageSize":
          this.settings.devicesPageSize = value as number
          break
        case "categoriesPageSize":
          this.settings.categoriesPageSize = value as number
          break
        case "statusesPageSize":
          this.settings.statusesPageSize = value as number
          break
        case "deviceTypesPageSize":
          this.settings.deviceTypesPageSize = value as number
          break
        case "interfaceTypesPageSize":
          this.settings.interfaceTypesPageSize = value as number
          break
        case "locationsPageSize":
          this.settings.locationsPageSize = value as number
          break
        case "maintainersPageSize":
          this.settings.maintainersPageSize = value as number
          break
        case "operatingSystemsPageSize":
          this.settings.operatingSystemsPageSize = value as number
          break
        case "architecturesPageSize":
          this.settings.architecturesPageSize = value as number
          break
        case "vLansPageSize":
          this.settings.vLansPageSize = value as number
          break
        case "usersPageSize":
          this.settings.usersPageSize = value as number
          break
      }
    },
    getSettings() {
      settings.get<{ Name: string; Value: string }[]>(this.account.ID).then((response) => {
        if (response !== null && response.length > 0) {
          response.forEach((elem) => {
            this.updateSetting(elem.Name, elem.Value)
          })
        }
      })
    },
    saveSetting(name: string, value: boolean | string | number) {
      switch (typeof value) {
        case "boolean":
          value = value ? "true" : "false"
          break
        case "number":
          value = value.toString()
          break
      }

      setting.post<UserSetting>(this.account.ID, { Name: name, Value: value }).then((response) => {
        this.updateSetting(response.Name, response.Value)
      })
    },
    isDarkMode() {
      if (this.settings.darkMode === null) {
        return window.matchMedia("(prefers-color-scheme: dark)").matches ? true : false
      }
      return this.settings.darkMode ? true : false
    },
    theme() {
      return this.isDarkMode() ? "dark" : "light"
    },
    toggleDarkMode() {
      this.saveSetting("darkMode", !this.isDarkMode())
    },
    lastDismissedUnknown(date: Date) {
      this.saveSetting("lastDismissed", format(date, "yyyy-MM-dd HH:mm:ss"))
    },
  },
})
