export type Menu = Array<Section>

export interface Section {
  Label: string
  Items: Array<Item>
}

export interface Item {
  To: string // The address or route the item should link to
  Route?: boolean // If true, the item will be created as a vue route, if false an anchor-link
  Target?: string // The target attribute of the item
  Label: string // The text label of the item
  Icon: string // The icon label of the item
  Tooltip?: string // The tooltip to be shown when the item is hovered over
}

const primaryMenu = [
  {
    Label: "",
    Items: [
      {
        To: "/",
        Icon: "ViewDashboardVariant",
        Label: "Dashboard",
        Tooltip: "Test",
      },
      {
        To: "/devices",
        Icon: "AccessPointNetwork",
        Label: "Devices",
      },
    ],
  },
  {
    Label: "taxonomy",
    Items: [
      {
        To: "/manage/categories",
        Icon: "Shape",
        Label: "Categories",
      },
      {
        To: "/manage/status",
        Icon: "ListStatus",
        Label: "Statuses",
      },
      {
        To: "/manage/maintainers",
        Icon: "Account",
        Label: "Maintainers",
      },
      {
        To: "/manage/locations",
        Icon: "MapMarker",
        Label: "Locations",
      },
      {
        To: "/manage/devicetypes",
        Icon: "Devices",
        Label: "DeviceTypes",
      },
      {
        To: "/manage/operatingsystems",
        Icon: "MicrosoftWindowsClassic",
        Label: "OS Types",
      },
      {
        To: "/manage/architectures",
        Icon: "Chip",
        Label: "CPU Types",
      },
      {
        To: "/manage/vlans",
        Icon: "Vpn",
        Label: "VLANs",
      },
    ],
  },
  {
    Label: "access",
    Items: [
      {
        To: "/admin/users",
        Icon: "AccountCircle",
        Label: "Users",
      },
    ],
  },
]

const secondaryMenu = [
  {
    Label: "about",
    Items: [
      {
        To: "/about/info",
        Icon: "Information",
        Label: "Info",
      },
    ],
  },
] as Menu

export { primaryMenu, secondaryMenu }
