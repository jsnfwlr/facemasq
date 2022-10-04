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
        Label: "dashboard",
        Tooltip: "Test",
      },
      {
        To: "/devices",
        Icon: "AccessPointNetwork",
        Label: "devices",
      },
    ],
  },
  {
    Label: "taxonomy",
    Items: [
      {
        To: "/manage/categories",
        Icon: "Shape",
        Label: "categories",
      },
      {
        To: "/manage/status",
        Icon: "ListStatus",
        Label: "statuses",
      },
      {
        To: "/manage/maintainers",
        Icon: "Account",
        Label: "maintainers",
      },
      {
        To: "/manage/locations",
        Icon: "MapMarker",
        Label: "locations",
      },
      {
        To: "/manage/devicetypes",
        Icon: "Devices",
        Label: "deviceTypes",
      },
      {
        To: "/manage/operatingsystems",
        Icon: "MicrosoftWindowsClassic",
        Label: "osTypes",
      },
      {
        To: "/manage/architectures",
        Icon: "Chip",
        Label: "cpuTypes",
      },
      {
        To: "/manage/vlans",
        Icon: "Vpn",
        Label: "vLANs",
      },
    ],
  },
  {
    Label: "access",
    Items: [
      {
        To: "/admin/users",
        Label: "users",
        Icon: "AccountCircle",
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
        Label: "info",
        Icon: "Information",
      },
    ],
  },
] as Menu

export { primaryMenu, secondaryMenu }
