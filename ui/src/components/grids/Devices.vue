<script setup lang="ts">
  import { ref } from "vue"
  import { storeToRefs } from "pinia"
  import { parse, differenceInWeeks, differenceInHours, differenceInDays, differenceInMinutes } from "date-fns"
  import clonedeep from "lodash.clonedeep"

  import { useDevices, Device, DomainName, Connection, ColumnSort } from "@/stores/devices"
  import { useParams } from "@/stores/params"
  import { useUser } from "@/stores/user"

  import Btn from "@/components/elements/Btn.vue"
  import Btngrp from "@/components/elements/BtnGrp.vue"
  import Card from "@/components/containers/Card.vue"
  import ConnectivityIndicator from "@/components/indicators/Connectivity.vue"
  import Control from "@/components/elements/Control.vue"
  import Field from "@/components/containers/Field.vue"
  import InterfacesGrid from "@/components/grids/Interfaces.vue"
  import AddressesGrid from "@/components/grids/Addresses.vue"
  import HostnamesGrid from "@/components/grids/Hostnames.vue"

  import mdIcon from "@/components/elements/MDIcon.vue"

  import ModalBox from "@/components/justboil/ModalBox.vue"

  interface MultiRecord {
    DeviceID: number | null
    InterfaceID: number | null
    AddressID: number | null
    MAC: string
    IPv4: string | null
    IsPrimaryMAC: boolean
    IsVirtualMAC: boolean
    IsPrimaryIPv4: boolean
    IsVirtualIPv4: boolean
    IsReservedIPv4: boolean
    InterfaceTypeID: number
    VlanID: number
    Hostnames: Array<DomainName>
    Connectivity: Array<Connection> | null
    LastSeen: string | null
    StatusID: number
  }

  interface Props {
    checkable?: boolean
    items: Array<Device>
    mode?: string
    filtered?: boolean
  }

  interface MaxCols {
    default: number
    administrative: number
  }

  const props = withDefaults(defineProps<Props>(), {
    perPage: 10,
    mode: "Informative",
    filtered: false,
  })

  const maxCols = {
    default: 12,
    administrative: 14,
  } as MaxCols

  const deviceStore = useDevices()
  const { editingItems, deletingItems, focusedItems, investigations } = storeToRefs(deviceStore)

  const userStore = useUser()

  const paramsStore = useParams()
  const { Locations, Maintainers, OperatingSystems, InterfaceTypes, Statuses, Categories, DeviceTypes, Architectures } = storeToRefs(paramsStore)

  const iconIFace = (interfaceTypeID: number) => {
    return InterfaceTypes.value.find((item) => item.ID === interfaceTypeID)?.Icon as string
  }

  const isModalActive = ref(false)
  // const isModalDangerActive = ref(false)

  const investigateID = ref(0)

  const investigate = (id: number) => {
    investigateID.value = id
    isModalActive.value = true
    deviceStore.investigateDevice(id)
    console.log("Do something that allows the user to investigate " + id)
    // this.$router.push('/investigate/' + id)
  }

  const columnClass = (column: string, param1: boolean | null) => {
    let colClass = [column]
    switch (column) {
      case "ip":
        if (param1) {
          colClass.push("reserved")
        }
    }
    return colClass
  }

  const fuzzyTime = (time: string | null) => {
    if (time !== null) {
      const current = new Date()
      const then = parse(time.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date())
      const minutes = differenceInMinutes(current, then)
      const hours = differenceInHours(current, then)
      const days = differenceInDays(current, then)
      const weeks = differenceInWeeks(current, then)

      if (weeks === 0) {
        if (days === 0) {
          if (hours === 0) {
            if (minutes === 0) {
              return "now"
            }
            return minutes + (minutes === 1 ? " minute ago" : " minutes ago")
          }
          return hours + (hours === 1 ? " hour ago" : " hours ago")
        }
        return days + (days === 1 ? " day ago" : " days ago")
      }
      return weeks + (weeks === 1 ? " week ago" : " weeks ago")
    }
    return "-"
  }

  type Param = {
    ID: number | null
    Label?: string
    Notes?: string | null
    Vendor?: string
    Family?: string
    Version?: string
    Name?: string
    IsServer?: number | boolean
    IsCloud?: number | boolean
    IsOpenSource?: number | boolean
    Icon?: string
  }

  const getLabelByID = (params: Array<Param>, label: string, selectedID: number) => {
    let value = ""
    if (params.length > 0) {
      let selected = params.find((par) => par.ID === selectedID)
      if (selected) {
        switch (label) {
          case "Label":
            value = selected.Label ? selected.Label : ""
            break
          case "Vendor":
            value = selected.Vendor ? selected.Vendor : ""
            break
          case "Family":
            value = selected.Family ? selected.Family : ""
            break
          case "Version":
            value = selected.Version ? selected.Version : ""
            break
          case "Name":
            value = selected.Name ? selected.Name : ""
            break
          case "Icon":
            value = selected.Icon ? selected.Icon : ""
            break
        }
      }
    }
    return value
  }

  type SelectOption = {
    value: number
    label: string
  }
  const paramToOptions = (params: Array<Param>, labelSwitch: number) => {
    const options = [] as SelectOption[]

    params.forEach((item) => {
      switch (labelSwitch) {
        case 1:
          options.push({ value: item.ID ? item.ID : 0, label: item.Vendor + " " + item.Family + " " + item.Name + " " + item.Version })
          break
        default:
          options.push({ value: item.ID ? item.ID : 0, label: item.Label ? item.Label : "-" })
          break
      }
    })
    options.sort((a, b) => {
      return a.label === b.label ? 0 : a.label < b.label ? -1 : 1
    })
    return options
  }

  const colCount = (excludeCount: number) => {
    switch (props.mode) {
      case "Administrative":
        return maxCols.administrative - excludeCount
      default:
        return maxCols.default - excludeCount
    }
  }

  // const editAddresses = (deviceIndex: number ) => {
  //   let addresses = [] as Array<Address>
  //   editingItems.value.interfaces.forEach((netFace, index) => {
  //     if (netFace.DeviceID == props.items[deviceIndex].ID) {
  //       addresses = props.items[deviceIndex].Interfaces[index].Addresses
  //     }
  //   })
  //   return addresses
  // }

  /* Open/Close rows with additional interfaces/addresses */
  interface ExpandedRow {
    Devices: Array<number>
    Notes: Array<number>
  }
  const expandedRows = ref<ExpandedRow>({ Devices: [], Notes: [] })
  const isRowExpanded = (deviceIndex: number) => {
    return expandedRows.value.Devices.includes(deviceIndex)
  }
  const toggleRowExpand = (deviceIndex: number) => {
    if (isRowExpanded(deviceIndex)) {
      collapseRow(deviceIndex)
      deviceStore.$patch((state) => {
        state.focusedItems.devices.delete(deviceIndex.toString())
      })
    } else {
      expandRow(deviceIndex)
    }
  }

  const expandRow = (deviceIndex: number) => {
    expandedRows.value.Devices.push(deviceIndex)
  }
  const collapseRow = (deviceIndex: number) => {
    expandedRows.value.Devices = expandedRows.value.Devices.filter((openIndex) => openIndex !== deviceIndex)
  }
  const hasExpand = (deviceIndex: number) => {
    if (props.items[deviceIndex].Interfaces.length > 1 || props.items[deviceIndex].Interfaces[0].Addresses.length > 1) {
      return true
    }
    return false
  }
  /* Open/Close rows with additional interfaces/addresses */

  const rowClass = (deviceIndex: number) => {
    const rowClass = []

    switch (props.items[deviceIndex].StatusID) {
      case 1:
        rowClass.push("invading")
        break
      case 2:
        rowClass.push("active")
        break
      case 3:
        rowClass.push("inactive")
        break
      case 4:
        rowClass.push("planned")
        break
      case 5:
        rowClass.push("deprecated")
        break
      case 6:
        rowClass.push("retired")
        break
      case 7:
        rowClass.push("lost")
        break
    }

    if (props.items[deviceIndex].IsOnline) {
      rowClass.push("online")
    }

    if (isEditing(deviceIndex)) {
      rowClass.push("editing")
    }
    if (isDeleting(deviceIndex)) {
      rowClass.push("deleting")
    }
    if (isFocused(deviceIndex)) {
      rowClass.push("focused")
    }
    return rowClass // .join(' ')
  }

  const edit = (deviceIndex: number) => {
    if (isEditing(deviceIndex)) {
      deviceStore.$patch((state) => {
        state.editingItems.devices.delete(deviceIndex.toString())
      })
    } else {
      deviceStore.$patch((state) => {
        state.editingItems.devices.set(deviceIndex.toString(), clonedeep(props.items[deviceIndex]))
      })
      collapseRow(deviceIndex)
    }
  }
  const isEditing = (deviceIndex: number) => {
    if (props.items[deviceIndex].ID === null || props.items[deviceIndex].ID === 0) {
      return true
    }
    return editingItems.value.devices.has(deviceIndex.toString())
  }
  /*
  const focus = (deviceIndex: number) => {
    if (isEditing(deviceIndex)) {
      deviceStore.$patch((state) => {
        state.editingItems.devices.delete(deviceIndex.toString())
      })
    } else {
      deviceStore.$patch((state) => {
        state.editingItems.devices.set(deviceIndex.toString(), clonedeep(props.items[deviceIndex]))
      })
      collapseRow(deviceIndex)
    }
  }
  */

  const toggleRowFocus = (deviceIndex: number) => {
    if (isFocused(deviceIndex) && !isRowExpanded(deviceIndex)) {
      deviceStore.$patch((state) => {
        state.focusedItems.devices.delete(deviceIndex.toString())
      })
    } else {
      deviceStore.$patch((state) => {
        state.focusedItems.devices.set(deviceIndex.toString(), clonedeep(props.items[deviceIndex]))
      })
    }
  }
  const isFocused = (deviceIndex: number) => {
    if (props.items[deviceIndex].ID === null || props.items[deviceIndex].ID === 0) {
      return true
    }
    return focusedItems.value.devices.has(deviceIndex.toString())
  }

  const isDeleting = (deviceIndex: number) => {
    return deletingItems.value.devices.includes(deviceIndex)
  }

  const sortColumn = ref<ColumnSort>({ column: "SortOrder", direction: "asc" })
  const setSortColumn = (byColumn: string) => {
    if (sortColumn.value.column == byColumn) {
      sortColumn.value.direction = sortColumn.value.direction === "asc" ? "desc" : "asc"
    } else {
      sortColumn.value = { column: byColumn, direction: "asc" }
    }
    deviceStore.setColumnSort(sortColumn.value)
    deviceStore.SortDevices()
  }
  const showSort = (column: string) => {
    if (sortColumn.value.column === column) {
      return sortColumn.value.direction
    }
  }

  const addressesForCurrentInterface = (deviceIndex: number) => {
    const interfaceIndex = getCurrentInterfaceIndex(deviceIndex) === -1 ? 0 : getCurrentInterfaceIndex(deviceIndex)
    let addresses = props.items[deviceIndex].Interfaces[interfaceIndex].Addresses

    // if (getCurrentInterfaceIndex(deviceIndex) !== -1) {
    //   let hasEdit = false
    //   for (let i=0; i<props.items[deviceIndex].Interfaces.length; i++) {
    //     if (editingItems.value.interfaces.has(deviceIndex +"."+ i)) {
    //       hasEdit = true
    //       addresses = props.items[deviceIndex].Interfaces[i].Addresses
    //       break
    //     }
    //   }
    //   if (hasEdit === false) {
    //     for (let i=0; i<props.items[deviceIndex].Interfaces.length; i++) {
    //       if (focusedItems.value.interfaces.has(deviceIndex +"."+ i)) {
    //         addresses = props.items[deviceIndex].Interfaces[i].Addresses
    //         break
    //       }
    //     }
    //   }
    // }
    return addresses
  }

  const getCurrentInterfaceIndex = (deviceIndex: number) => {
    let hasEdit = false
    for (let i = 0; i < props.items[deviceIndex].Interfaces.length; i++) {
      if (editingItems.value.interfaces.has(deviceIndex + "." + i)) {
        return i
      }
    }
    if (hasEdit === false) {
      for (let i = 0; i < props.items[deviceIndex].Interfaces.length; i++) {
        if (focusedItems.value.interfaces.has(deviceIndex + "." + i)) {
          return i
        }
      }
    }
    return -1
  }

  const hostnamesForCurrrentAddress = (deviceIndex: number) => {
    const interfaceIndex = getCurrentInterfaceIndex(deviceIndex) === -1 ? 0 : getCurrentInterfaceIndex(deviceIndex)
    const addressIndex = getCurrentAddressIndex(deviceIndex) === -1 ? 0 : getCurrentAddressIndex(deviceIndex)
    let hostnames = props.items[deviceIndex].Interfaces[interfaceIndex].Addresses[addressIndex].Hostnames

    if (getCurrentInterfaceIndex(deviceIndex) !== -1 && getCurrentAddressIndex(deviceIndex) !== -1) {
      for (let a = 0; a < props.items[deviceIndex].Interfaces[interfaceIndex].Addresses.length; a++) {
        if (editingItems.value.addresses.has(deviceIndex + "." + interfaceIndex + "." + a)) {
          hostnames = props.items[deviceIndex].Interfaces[interfaceIndex].Addresses[a].Hostnames
          break
        }
      }
    }

    return hostnames
  }

  const getCurrentAddressIndex = (deviceIndex: number) => {
    const interfaceIndex = getCurrentInterfaceIndex(deviceIndex) === -1 ? 0 : getCurrentInterfaceIndex(deviceIndex)
    if (getCurrentInterfaceIndex(deviceIndex) !== -1) {
      for (let a = 0; a < props.items[deviceIndex].Interfaces[interfaceIndex].Addresses.length; a++) {
        if (editingItems.value.addresses.has(deviceIndex + "." + interfaceIndex + "." + a)) {
          return a
        }
      }
    }
    return -1
  }
  const allAddresses = (device: Device) => {
    const records = [] as MultiRecord[]
    device.Interfaces.forEach((netface) => {
      netface.Addresses.forEach((address) => {
        if (address.IPv4 !== device.Primary.IPv4) {
          const record = {
            DeviceID: device.ID,
            InterfaceID: netface.ID,
            AddressID: address.ID,
            IPv4: address.IPv4,
            MAC: netface.MAC,
            IsPrimaryMAC: netface.IsPrimary,
            IsVirtualMAC: netface.IsVirtual,
            IsPrimaryIPv4: address.IsPrimary,
            IsVirtualIPv4: address.IsVirtual,
            IsReservedIPv4: address.IsReserved,
            InterfaceTypeID: netface.InterfaceTypeID,
            VlanID: netface.VlanID,
            Hostnames: address.Hostnames,
            LastSeen: address.LastSeen !== null ? address.LastSeen : null,
            Connectivity: address.Connectivity,
            StatusID: netface.StatusID,
          }
          records.push(record)
        }
      })
    })
    return records
  }
</script>

<template>
  <div class="devices datagrid">
    <modal-box v-model="isModalActive" title="edit.Label">
      <div v-for="(address, index) in investigations.get(investigateID)" :key="'investigation-address-connectivity-id:' + index">
        {{ address.AddressID }}
        <connectivity-indicator :data="address.Connectivity" :includeDate="true" />
      </div>
    </modal-box>
    <table :class="props.mode" class="multibody cardify">
      <thead>
        <tr>
          <th class="expand" />
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('MachineName')]" class="name sortable" @click="setSortColumn('MachineName')">Machine Name</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('Maintainer')]" class="maintainer sortable" @click="setSortColumn('Maintainer')">Maintainer</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('Location')]" class="location sortable" @click="setSortColumn('Location')">Location</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('Label')]" class="description sortable" @click="setSortColumn('Label')">Description</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('SortOrder')]" class="ip sortable" @click="setSortColumn('SortOrder')">IP</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('MAC')]" class="mac sortable" @click="setSortColumn('MAC')">MAC</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('FirstSeen')]" class="firstseen sortable" @click="setSortColumn('FirstSeen')">First Seen</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('LastSeen')]" class="lastseen sortable" @click="setSortColumn('LastSeen')">Last Seen</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('OS')]" class="os sortable" @click="setSortColumn('OS')">Operating System</th>
          <th :class="[deviceStore.canChangeSort ? '' : 'stealth', showSort('Brand')]" class="brand sortable" @click="setSortColumn('Brand')">Make/Model</th>
          <th class="connectivity" title="the online/offline presence of the device over time">Connectivity</th>
          <th v-if="userStore.hasAccess('details', 'read')" class="actions" />
        </tr>
      </thead>
      <tbody v-for="(row, index) in items" :key="row.Interfaces[0].Addresses[0].ID">
        <tr :class="rowClass(index)" @click="toggleRowFocus(index)">
          <td v-if="hasExpand(index)" class="expand" @click="toggleRowExpand(index)">
            <div>
              <mdIcon v-if="isRowExpanded(index)" icon="ChevronDownBox" size="20" />
              <mdIcon v-else icon="ChevronRightBox" size="20" />
            </div>
          </td>
          <td v-else class="expand" />
          <td :data-label="row.Label" class="whitespace-nowrap name">
            <div><mdIcon :icon="getLabelByID(DeviceTypes, 'Icon', row.DeviceTypeID)" size="14" />{{ row.MachineName }}</div>
          </td>
          <td data-label="Maintainer" class="maintainer">
            <div>{{ rowClass(index).includes("invading") ? "-" : getLabelByID(Maintainers, "Label", row.MaintainerID) }}</div>
          </td>
          <td data-label="Location" class="location">
            <div>{{ rowClass(index).includes("invading") ? "-" : getLabelByID(Locations, "Label", row.LocationID) }}</div>
          </td>
          <td data-label="Description" class="description">
            <div class="text-gray-500 dark:text-gray-400 whitespace-nowrap">{{ row.Label }}</div>
          </td>
          <td data-label="IP" class="ip">
            <div :class="columnClass('ip', row.Primary.IsReservedIP)">{{ row.Primary.IPv4 }}</div>
          </td>
          <td data-label="MAC" class="mac">
            <div><mdIcon :icon="iconIFace(row.Primary.InterfaceTypeID)" size="10" />{{ row.Primary.MAC }}</div>
          </td>
          <td data-label="First Seen" class="firstseen">
            <div>{{ fuzzyTime(row.FirstSeen) }}</div>
          </td>
          <td data-label="Last Seen" class="lastseen">
            <div>
              <mdIcon :title="getLabelByID(Statuses, 'Label', row.StatusID)" :icon="getLabelByID(Statuses, 'Icon', row.StatusID)" size="14" /><span :title="row.Interfaces[0].LastSeen">{{ fuzzyTime(row.Interfaces[0].LastSeen) }}</span>
            </div>
          </td>
          <td data-label="Operating System" class="os">
            <div>{{ rowClass(index).includes("invading") ? "-" : getLabelByID(OperatingSystems, "Family", row.OperatingSystemID) + " " + getLabelByID(OperatingSystems, "Version", row.OperatingSystemID) }}</div>
          </td>
          <td data-label="Make+Model" class="brand">
            <div>{{ rowClass(index).includes("invading") && row.Brand === null ? "" : (row.Brand ? row.Brand : "?") + " " + (row.Model ? row.Model : "?") }}</div>
          </td>
          <td data-label="Connectivity" class="connectivity">
            <div><connectivity-indicator :includeDate="false" :data="row.Interfaces[0].Addresses[0].Connectivity" /></div>
          </td>
          <td class="actions right" v-if="userStore.hasAccess('devices', 'write')">
            <div class="flex flex-end">
              <btn v-if="rowClass(index).includes('invading') && userStore.hasAccess('details', 'read')" color="danger" title="Investigate" icon="Magnify" :outline="true" small @click="investigate(index)" />

              <btngrp>
                <btn v-if="!filtered && !isEditing(index) && !isDeleting(index) && userStore.hasAccess('devices', 'write')" color="info" icon="Pencil" small @click="edit(index)" />
                <btn v-if="!filtered && isEditing(index) && userStore.hasAccess('devices', 'write')" color="success" icon="ContentSave" small @click="deviceStore.Save([index])" />
                <btn v-if="!filtered && isEditing(index) && userStore.hasAccess('devices', 'write')" color="warning" icon="Eraser" small @click="deviceStore.Discard([index])" />
                <btn v-if="!filtered && !isEditing(index) && !isDeleting(index) && userStore.hasAccess('devices', 'delete')" color="danger" icon="TrashCan" small @click="deviceStore.InitiateDelete([index])" />
                <btn v-if="!filtered && !isEditing(index) && isDeleting(index) && userStore.hasAccess('devices', 'delete')" color="danger" icon="Check" small @click="deviceStore.PerformDelete([index])" />
                <btn v-if="!filtered && !isEditing(index) && isDeleting(index) && userStore.hasAccess('devices', 'delete')" color="info" icon="Close" small @click="deviceStore.CancelDelete([index])" />
              </btngrp>
            </div>
          </td>
        </tr>
        <tr v-show="isRowExpanded(index)" v-for="(record, index2) in allAddresses(row)" :class="rowClass(index)" class="address-expansion" :key="'expanded-addresses-row:' + row.ID + '-address:' + index2">
          <td class="expand" />
          <td class="name" />
          <td class="maintainer" />
          <td class="location" />
          <td class="description" />
          <td class="ip">
            <div :class="columnClass('ip', record.IsReservedIPv4)">{{ record.IPv4 }}</div>
          </td>
          <td class="mac">
            <div><mdIcon :icon="iconIFace(record.InterfaceTypeID)" size="10" />{{ record.MAC }}</div>
          </td>
          <td class="firstseen" />
          <td class="lastseen">
            <div><mdIcon :title="getLabelByID(Statuses, 'Label', record.StatusID)" :icon="getLabelByID(Statuses, 'Icon', record.StatusID)" size="14" />{{ fuzzyTime(record.LastSeen) }}</div>
          </td>
          <td class="os" />
          <td class="brand" />
          <td class="connectivity">
            <div><connectivity-indicator :includeDate="false" :data="record.Connectivity" /></div>
          </td>
          <td v-if="userStore.hasAccess('details', 'read')" class="actions right"></td>
        </tr>
        <tr v-show="isEditing(index)" class="editor">
          <td />
          <td :colspan="colCount(1)">
            <card headingTitle="Device Details" icon="Pencil" form>
              <div class="grid grid-cols-1 gap-6 lg:grid-cols-3 xl:grid-cols-8">
                <field spacing label="Machine Name"><control v-model="row.MachineName" icon="CardAccountDetails" /></field>
                <field spacing label="Maintainer"><control v-model="row.MaintainerID" :options="paramToOptions(Maintainers, 0)" icon="Account" /></field>
                <field spacing label="Location"><control v-model="row.LocationID" :options="paramToOptions(Locations, 0)" icon="HelpCircle" /></field>
                <field spacing label="Label"><control v-model="row.Label" icon="Tag" /></field>
                <field spacing label="Device Type"><control v-model="row.DeviceTypeID" :options="paramToOptions(DeviceTypes, 0)" :icon="getLabelByID(DeviceTypes, 'Icon', row.DeviceTypeID)" /></field>
                <field spacing label="Category"><control v-model="row.CategoryID" :options="paramToOptions(Categories, 0)" :icon="getLabelByID(Categories, 'Icon', row.CategoryID)" /></field>
                <field spacing label="Status"><control v-model="row.StatusID" :options="paramToOptions(Statuses, 0)" :icon="getLabelByID(Statuses, 'Icon', row.StatusID)" /></field>
                <field spacing label="Is Tracked"><control v-model="row.IsTracked" type="checkbox" /></field>
              </div>
              <div class="grid grid-cols-1 gap-6 lg:grid-cols-3 xl:grid-cols-8">
                <field spacing label="Brand"><control v-model="row.Brand" /></field>
                <field spacing label="Model"><control v-model="row.Model" /></field>
                <field spacing label="Serial"><control v-model="row.Serial" /></field>
                <field spacing label="Purchased"><control v-model="row.Purchased" type="date" /></field>
                <field spacing label="Operating System"><control v-model="row.OperatingSystemID" :options="paramToOptions(OperatingSystems, 1)" icon="MicrosoftWindowsClassic" /></field>
                <field spacing label="CPU Architecture"><control v-model="row.ArchitectureID" :options="paramToOptions(Architectures, 0)" icon="Chip" /></field>
                <field spacing label="Is Guest  "><control v-model="row.IsGuest" type="checkbox" /></field>
              </div>
              <field spacing label="Notes"><control v-model="row.Notes" icon="NoteText" /></field>
              <!-- <field spacing label="With help line" help="Do not enter the leading zero" >
                <control v-model="row.Notes" type="tel" placeholder="Your phone number" />
              </field> -->
            </card>
            <div class="grid grid-cols-1 gap-6 lg:grid-cols-3 xl:grid-cols-3 mt-6">
              <card headingTitle="Interfaces" icon="Pencil" table headerIcon="PlusBox" @header-icon-click="deviceStore.Add([index])" class="interfaces">
                <interfaces-grid :perPage="200" :items="row.Interfaces" :deviceIndex="index" :editingDevice="isEditing(index)" />
              </card>
              <card headingTitle="Addresses" icon="Pencil" table :hideHeaderIcon="getCurrentInterfaceIndex(index) === -1" headerIcon="PlusBox" @header-icon-click="deviceStore.Add([index, getCurrentInterfaceIndex(index)])" class="addresses">
                <addresses-grid :perPage="200" :items="addressesForCurrentInterface(index)" :deviceIndex="index" :interfaceIndex="getCurrentInterfaceIndex(index)" :canEdit="getCurrentInterfaceIndex(index) !== -1" />
              </card>
              <card headingTitle="Hostnames" icon="Pencil" table :hideHeaderIcon="getCurrentAddressIndex(index) === -1" headerIcon="PlusBox" @header-icon-click="deviceStore.Add([index, getCurrentInterfaceIndex(index), getCurrentAddressIndex(index)])" class="hostnames">
                <hostnames-grid :perPage="200" :items="hostnamesForCurrrentAddress(index)" :deviceIndex="index" :interfaceIndex="getCurrentInterfaceIndex(index)" :addressIndex="getCurrentAddressIndex(index)" :canEdit="getCurrentAddressIndex(index) !== -1" />
              </card>
            </div>
          </td>
        </tr>
        <tr class="noterow" style="display: none">
          <td :colspan="userStore.hasAccess('details', 'read') ? '14' : '12'" class="notes">{{ row.Notes }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="scss">
  @import url("@/components/grids/grids.scss");

  button + .button-group {
    margin-left: 1rem;
  }
  .flex-end {
    justify-content: flex-end;
  }
  .devices.datagrid {
    & > table {
      & > thead,
      & > tbody {
        & > tr {
          & > td,
          & > th {
            & > & {
              @apply hidden;
            }
            &.expand {
              @apply lg:table-cell lg:w-[16px];
            }
            &.name {
              @apply flex lg:table-cell lg:w-[140px];
            }
            &.description {
              @apply flex lg:table-cell lg:w-[230px];
            }
            &.brand {
              @apply lg:table-cell lg:w-[160px] lg:text-center;
            }
            &.os {
              @apply lg:table-cell lg:w-[160px] lg:text-center;
            }
            &.serial {
              @apply lg:table-cell lg:w-[160px];
            }
            &.connectivity {
              @apply flex lg:table-cell lg:w-[200px];
            }
            &.location {
              @apply lg:table-cell lg:w-[140px] lg:text-center;
            }
            &.maintainer {
              @apply lg:table-cell lg:w-[140px] lg:text-center;
            }
            &.mac {
              @apply flex lg:table-cell lg:w-[160px];
            }
            &.ip {
              @apply flex lg:table-cell lg:w-[160px];
            }

            &.firstseen {
              @apply flex lg:table-cell lg:w-[160px];
            }
            &.lastseen {
              @apply flex lg:table-cell lg:w-[160px];
            }
            &.actions {
              @apply flex lg:table-cell lg:w-[160px] text-center lg:text-right;
            }
            &.actions > div {
              @apply w-full lg:w-auto;
            }
          }
        }
      }
      &.cardify {
        & {
          @apply bg-[#f7fdfb] dark:bg-[#222f39];
        }
        & > tbody {
          & {
            @apply block lg:table-row-group;
          }
          & {
            @apply mt-16 lg:mt-0 border-y-gray-100 dark:border-y-gray-900;
          }
          &:first-of-type {
            @apply mt-0 lg:mt-0 border-b-0 dark:border-b-0;
          }
          & > tr {
            & {
              grid-template-columns: [col0] 1fr [col1] 1fr [col2] 1fr [col3] 1fr [col4] 1fr [col5] 1fr [col6] 1fr [col7] 1fr [col8] 1fr [coln];
              grid-template-rows: [row0] 4rem [row1] 3.5rem [row2] 3.5rem [row3] 3.5rem [row4] 3.5rem [row5] 3.5rem [rown];
              @apply sm:border-0;
            }
            & > td:not(.actions):not(.connectivity) {
              &:before {
                @apply block lg:hidden text-sm font-thin;
                content: attr(data-label);
                white-space: nowrap;
                position: absolute;
                top: 4px;
                left: 1rem;
                opacity: 0.7;
              }
            }

            & > td {
              & {
                @apply block lg:table-cell relative lg:static;
              }

              &.name {
                grid-area: row0 / col0 / row1 / col6;
                @apply pt-6 lg:pt-2 text-2xl lg:text-base font-bold	lg:font-normal;
              }
              &.description,
              &.serial {
                @apply hidden lg:table-cell lg:pt-2;
              }
              &.maintainer {
                grid-area: row4 / col0 / row5 / col4;
                @apply pt-6 lg:pt-2;
              }
              &.location {
                grid-area: row4 / col4 / row5 / coln;
                @apply pt-6 lg:pt-2;
              }
              &.ip {
                & {
                  grid-area: row1 / col0 / row2 / col4;
                  @apply pt-6 lg:pt-2;
                }
                div {
                  & {
                    opacity: 0.6;
                  }
                  &.reserved {
                    opacity: 1;
                  }
                }
              }
              &.mac {
                grid-area: row1 / col4 / row2 / coln;
                @apply pt-6 lg:pt-2;
              }
              &.firstseen {
                grid-area: row2 / col0 / row3 / col5;
                @apply pt-6 lg:pt-2;
              }
              &.lastseen {
                grid-area: row2 / col4 / row3 / coln;
                @apply pt-6 lg:pt-2;
              }
              &.brand {
                grid-area: row3 / col0 / row4 / col4;
                @apply pt-6 lg:pt-2;
              }
              &.os {
                grid-area: row3 / col4 / row4 / coln;
                @apply pt-6 lg:pt-2;
              }
              &.connectivity {
                grid-area: row5 / col0 / rown / coln;
                @apply pt-4 lg:pt-2;
              }
              &.actions {
                grid-area: row0 / col6 / row1 / coln;
                @apply text-right;
              }

              &.expand {
                @apply hidden lg:table-cell lg:pt-2 lg:pb-0;
              }
            }
            /*
            &.address-expansion {
              background: red;
            }
            */
            &.editor {
              & {
                @apply block lg:table-row;
              }
              & > td:first-child {
                @apply hidden lg:table-cell;
              }
              & > td {
                @apply block lg:table-cell;
              }
              .hostnames {
                @apply lg:max-w-[800px];
              }
            }
          }
        }
      }

      tr.online > td.lastseen div {
        color: hsl(160, 84%, 39%);
      }
      tr:not(.online) > td.lastseen div {
        color: hsl(345, 100%, 39%);
      }

      tr.invading {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-red-800 dark:text-red-500;
        }
      }

      tr.active {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-teal-800 dark:text-teal-500;
        }
      }

      tr.planned {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-lime-800 dark:text-lime-500;
        }
      }

      tr.deprecated {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-sky-800 dark:text-sky-500;
        }
      }

      tr.retired {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-violet-800 dark:text-violet-500;
        }
      }

      tr.lost {
        // &.online > td.lastseen div,
        & > td.name div {
          @apply text-orange-800 dark:text-orange-500;
        }
      }
    }
  }
</style>
