import { defineStore } from "pinia"

export interface AppStore {
  styles: AppStyles
  toggles: AppToggles
  values: AppValues
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

export interface PageSize {
  label: string
  value: number
}

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
        perPage: 10,
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
    } as AppStore
  },
  actions: {},
})
