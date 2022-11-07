import { createApp } from "vue"
import { createPinia } from "pinia"
// import { createI18n } from "vue-i18n"

import App from "./App.vue"
import router from "./routes"
import "./styles/index.scss"

const app = createApp(App)

// const i18n = createI18n({
//   legacy: false,
//   messages: {
//     en: {
//       message: "WTF",
//       usermenu: {
//         settings: "Settings",
//         setDarkMode: "Go dark",
//         setLightMode: "Go light",
//         logOut: "Log out",
//       },
//       sidebar: {
//         dashboard: "Dashboard",
//         devices: "Devices",
//         taxonomy: "Taxonomy",
//         categories: "Categories",
//         statuses: "Statuses",
//         maintainers: "Maintainers",
//         locations: "Locations",
//         deviceTypes: "Device Types",
//         osTypes: "OS Types",
//         cpuTypes: "CPU Types",
//         vLANs: "VLANs",
//         access: "Access",
//         users: "Users",
//         about: "About",
//         info: "Info",
//       },
//       dashboard: {
//         labels: {},
//         columns: {},
//         tooltips: {},
//       },
//       devicePage: {
//         labels: {},
//         options: {},
//         columns: {},
//         tooltips: {
//           addDevice: "Add device",
//           addInterface: "Add interface",
//           addAddress: "Add address",
//           addHostname: "Add hostname",
//           editDevice: "Edit device",
//           editInterface: "Edit interface",
//           editAddress: "Edit address",
//           editHostname: "Edit hostname",
//           deleteDevice: "Delete device",
//           deleteInterface: "Delete interface",
//           deleteAddress: "Delete address",
//           deleteHostname: "Delete hostname",
//           showFilters: "Show filters",
//           hideFilters: "Hide filters",
//           clearFilters: "Clear filters",
//           expandRow: "Expand row",
//           saveChanges: "Save changes",
//           cancelChanges: "Cancel changes",
//           confirmDelete: "Confirm delete",
//           cancelDelete: "Cancel delete",
//           investigate: "Investigate",
//         },
//       },
//       infoPage: {
//         cards: {
//           SystemDetails: "System Details",
//           InstalledPluggins: "Installed Plugins",
//           LibraryAttribution: "Attribution for included libraries",
//           ImageAttribution: "Attribution for included graphics",
//         },
//         headers: {
//           VersionDetails: "Version Details",
//           ServerConfig: "Server Configuration",
//           ScannerConfig: "Scanner Configuration",
//         },
//         labels: {
//           NetScanFrequency: "Network Scan Frequency",
//           TargetCIDR: "Target Network CIDR",
//           ServerPort: "Port",
//           ReleaseDate: "Release Date",
//         },
//         other: {
//           ReleaseNoteLinkText: "View on GitHub",
//         },
//         attribution: {
//           justboil: "Original Vue UI library created by JustBoil - GitHub.",
//           materialdesignicons: "Icons imported from Templarian/MaterialDesign - GitHub.",
//           logo: "Logo based on Original Mask icons created by mangsaabguru - Flaticon.",
//         },
//       },
//     },
//   },
// })

app.use(router)
app.use(createPinia())
// app.use(i18n)
app.mount("#app")
