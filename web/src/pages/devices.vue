<script setup lang="ts">
  import { onMounted, onBeforeUnmount, computed, ref } from "vue"
  import { storeToRefs } from "pinia"
  import { useDevices, Device } from "@/stores/devices"
  import { useParams } from "@/stores/params"

  import { parse, differenceInMinutes } from "date-fns"

  import Btn from "@/components/elements/Btn.vue"
  import Card from "@/components/containers/Card.vue"
  import Control from "@/components/elements/Control.vue"
  import DeviceGrid from "@/components/grids/Devices.vue"
  import Field from "@/components/containers/Field.vue"

  // const userStore = useUser()
  const paramsStore = useParams()
  const deviceStore = useDevices()

  const { Maintainers, Statuses, Locations, DeviceTypes, Categories, InterfaceTypes } = storeToRefs(paramsStore)

  const { allDevices, editingItems } = storeToRefs(deviceStore)

  let timer: ReturnType<typeof setInterval>

  const pageRefreshInterval = 30

  const add = () => {
    deviceStore.Add([])
  }

  const autoUpdate = computed(() => {
    return editingItems.value.devices.size === 0
  })

  const updateDevices = () => {
    if (autoUpdate.value) {
      deviceStore.getAll()
    }
  }

  onMounted(() => {
    updateDevices()
    timer = setInterval(() => {
      updateDevices()
    }, pageRefreshInterval * 1000)
  })

  onBeforeUnmount(() => {
    clearInterval(timer)
  })

  const filteredDevices = computed(() => {
    let results = allDevices.value
    if (filters.value.interfaceTypeID !== null) {
      let filteredResults = Array<Device>()
      results.forEach((device) => {
        let added = false
        device.Interfaces.forEach((netFace) => {
          if (!added && netFace.InterfaceTypeID === filters.value.interfaceTypeID) {
            added = true
            filteredResults.push(device)
          }
        })
      })
      results = filteredResults
    }
    if (filters.value.lastseenID !== null) {
      let filteredResults = Array<Device>()
      const current = new Date()
      results.forEach((device) => {
        let lastSeen = null as number | null
        device.Interfaces.forEach((netFace) => {
          if (netFace.LastSeen !== null) {
            const then = parse(netFace.LastSeen.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date())
            const minutes = differenceInMinutes(current, then)
            if (lastSeen === null || minutes < lastSeen) {
              lastSeen = minutes
            }
          }
        })
        if (lastSeen !== null) {
          switch (filters.value.lastseenID) {
            case 1:
              if (lastSeen <= 2) {
                filteredResults.push(device)
              }
              break
            case 2:
              if (lastSeen <= 60) {
                filteredResults.push(device)
              }
              break
            case 3:
              if (lastSeen <= 1440) {
                filteredResults.push(device)
              }
              break
            case 4:
              if (lastSeen <= 10080) {
                filteredResults.push(device)
              }
              break
            case 5:
              if (lastSeen > 2) {
                filteredResults.push(device)
              }
              break
            case 6:
              if (lastSeen > 60) {
                filteredResults.push(device)
              }
              break
            case 7:
              if (lastSeen > 1440) {
                filteredResults.push(device)
              }
              break
            case 8:
              if (lastSeen > 10080) {
                filteredResults.push(device)
              }
              break
          }
        }
      })
      results = filteredResults
    }
    Object.keys(filters.value).forEach((key) => {
      if (filters.value[key as keyof typeof filters.value] !== null) {
        if (!["lastseenID", "interfaceTypeID"].includes(key)) {
          results = results.filter((device) => device[(key.charAt(0).toUpperCase() + key.slice(1)) as keyof typeof device] === filters.value[key as keyof typeof filters.value])
        }
      }
    })
    // if (results.length === allDevices.value.length) {
    //   results = results.filter(device => device.IsTracked && device.StatusID !== 5 && device.StatusID !== 6 && device.StatusID !== 7)
    // }
    return results
  })

  const filters = ref({
    maintainerID: null as number | null,
    statusID: null as number | null,
    locationID: null as number | null,
    categoryID: null as number | null,
    deviceTypeID: null as number | null,
    interfaceTypeID: null as number | null,
    lastseenID: null as number | null,
    isTracked: null as boolean | null,
  })
  const resetFilters = () => {
    Object.keys(filters.value).forEach((key) => {
      filters.value[key as keyof typeof filters.value] = null
    })
  }
  const filterdTitle = computed(() => {
    let title = "All Devices (" + allDevices.value.length + ")"
    if (filteredDevices.value.length !== allDevices.value.length) {
      title = "Devices (" + filteredDevices.value.length + " of " + allDevices.value.length + ")"
    }
    return title
  })

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

  const lastSeenOptions = [
    { value: null, label: "Any" },
    { value: 1, label: "< Now" },
    { value: 2, label: "< Hour" },
    { value: 3, label: "< Day" },
    { value: 4, label: "< Week" },
    { value: 5, label: "> Now" },
    { value: 6, label: "> Hour" },
    { value: 7, label: "> Day" },
    { value: 8, label: "> Week" },
  ]
</script>

<template>
  <card :icon="autoUpdate ? 'MonitorCellphone' : 'Pause'" :headingTitle="filterdTitle" has-table class="mb-6" headerIcon="PlusBox" @header-icon-click="add" headerIconTitle="Add device">
    <!-- <device-grid :perPage="200" :items="allDevices" mode="administrative" @discard="discardItem" @delete="deleteItem" @save="saveItem" @edit="editItem" /> -->
    <template v-slot:filter>
      <field inline label="Status">
        <control v-model="filters.statusID" :options="paramToOptions(Statuses, 0)" />
      </field>
      <field inline label="Maintainer">
        <control v-model="filters.maintainerID" :options="paramToOptions(Maintainers, 0)" />
      </field>
      <field inline label="Location">
        <control v-model="filters.locationID" :options="paramToOptions(Locations, 0)" />
      </field>
      <field inline label="Category">
        <control v-model="filters.categoryID" :options="paramToOptions(Categories, 0)" />
      </field>
      <field inline label="Type">
        <control v-model="filters.deviceTypeID" :options="paramToOptions(DeviceTypes, 0)" />
      </field>
      <field inline label="Connection">
        <control v-model="filters.interfaceTypeID" :options="paramToOptions(InterfaceTypes, 0)" />
      </field>
      <field inline label="Last Seen">
        <control v-model="filters.lastseenID" :options="lastSeenOptions" />
      </field>
      <field inline label="Tracked">
        <control v-model="filters.isTracked" type="checkbox" />
      </field>
      <btn color="danger" title="Clear filters" icon="CloseBox" :outline="false" small @click="resetFilters" />
    </template>
    <device-grid :perPage="200" :items="filteredDevices" mode="administrative" />
  </card>
</template>

<style lang="scss">
  .filters {
    & {
      display: flex;
    }
    .control {
      select {
      }
    }
  }
</style>
