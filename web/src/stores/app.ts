import { defineStore } from "pinia"
import { mande } from "mande"

export interface AppStore {
  styles: AppStyles
  toggles: AppToggles
  values: AppValues
  settings: AppSettings
}
export interface AppStyles {
  lightBorderStyle: string
  lightBgStyle: string
  sidebarStyle: string
  sidebarBrandStyle: string
  sidebarMenuCloseLgStyle: string
  sidebarMenuLabelStyle: string
  sidebarMenuItemStyle: string
  sidebarMenuItemActiveStyle: string
  sidebarMenuItemInactiveStyle: string
  sidebarSubmenuListStyle: string
  navBarItemLabelStyle: string
  navBarItemLabelHoverStyle: string
  navBarItemLabelActiveColorStyle: string
  navBarMenuListUpperLabelStyle: string
  tableTrStyle: string
  tableTrOddStyle: string
  overlayStyle: string
}

export interface AppToggles {
  isFullScreen: boolean
  isSidebarActive: boolean
  isFieldFocusRegistered: boolean
  sidebarLgToggle: boolean
}

export interface AppValues {
  perPage: number
  pageSizes: Array<PageSize>
}

export interface AppSettings {
  avatarType: string
}
export interface AppSetting {
  Name: string
  Value: string
}

export interface PageSize {
  label: string
  value: number
}

const settings = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api") : mande("/api")
const setting = import.meta.env.DEV ? mande("http://192.168.0.41:6135/api") : mande("/api")

export const useApp = defineStore("app", {
  state: () => {
    return {
      styles: {
        lightBorderStyle: "",
        lightBgStyle: "",
        sidebarStyle: "",
        sidebarBrandStyle: "",
        sidebarMenuCloseLgStyle: "",
        sidebarMenuLabelStyle: "",
        sidebarMenuItemStyle: "",
        sidebarMenuItemActiveStyle: "",
        sidebarMenuItemInactiveStyle: "",
        sidebarSubmenuListStyle: "",
        navBarItemLabelStyle: "",
        navBarItemLabelHoverStyle: "",
        navBarItemLabelActiveColorStyle: "",
        navBarMenuListUpperLabelStyle: "",
        tableTrStyle: "",
        tableTrOddStyle: "",
        overlayStyle: "",
      },
      toggles: {
        isFullScreen: false,
        isSidebarActive: false,
        isFieldFocusRegistered: false,
      },
      values: {
        perPage: 1000,
        pageSizes: [
          {
            label: "10",
            value: 10,
          },
          {
            label: "20",
            value: 20,
          },
          {
            label: "50",
            value: 50,
          },
          {
            label: "All",
            value: 0,
          },
        ],
      },
      settings: {
        avatarType: "",
      },
    } as AppStore
  },
  actions: {
    updateSetting(name: string, value: string | boolean | number) {
      switch (name) {
        case "avatarType":
          this.settings.avatarType = value as string
          break
      }
    },
    getSettings() {
      settings.get<{ Name: string; Value: string }[]>("/settings").then((response) => {
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

      setting.post<AppSetting>("/setting", { Name: name, Value: value }).then((response) => {
        this.updateSetting(response.Name, response.Value)
      })
    },
  },
})
