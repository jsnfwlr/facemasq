<script setup lang="ts">
  import { computed, ref } from "vue"
  import { parse, differenceInWeeks, differenceInHours, differenceInDays, differenceInMinutes } from "date-fns"
  import { storeToRefs } from "pinia"

  import { Device } from "@/stores/devices"
  import { useApp } from "@/stores/app"
  import { useParams } from "@/stores/params"

  import ModalBox from "@/components/justboil/ModalBox.vue"
  import ConnectivityIndicator from "@/components/indicators/Connectivity.vue"
  import mdIcon from "@/components/elements/MDIcon.vue"
  import Paginator from "@/components/grids/extensions/Paginator.vue"

  const paramsStore = useParams()

  const { InterfaceTypes, DeviceTypes } = storeToRefs(paramsStore)

  const props = defineProps({
    perPage: {
      type: Number,
      required: false,
      default: null,
    },
    maxHeight: {
      type: Number,
      required: false,
      default: null,
    },
    items: {
      type: Object,
      required: true,
    },
  })

  const appStore = useApp()

  const iconIFace = (interfaceTypeID: number) => {
    return InterfaceTypes.value.find((item) => item.ID === interfaceTypeID)?.Icon as string
  }

  const perPageCalc = computed(() => {
    if (props.perPage !== null) {
      return props.perPage
    } else if (appStore.values.perPage === 0) {
      return props.items.length
    } else {
      return appStore.values.perPage
    }
  })

  const isModalActive = ref(false)
  const isModalDangerActive = ref(false)

  const itemsPaginated = computed(() => props.items.slice(perPageCalc.value * currentPage.value, perPageCalc.value * (currentPage.value + 1)))

  const currentPage = ref(0)
  const setCurrentPage = (page: number) => {
    currentPage.value = page
  }

  const isUnknown = (client: Device) => {
    return !client.IsTracked
  }

  const isActive = (client: Device) => {
    return client.StatusID == 2
  }

  const dynamicClass = (item: string, param1: Device | number | null) => {
    if (item === "row") {
      const rowClass = []
      if (isActive(param1 as Device)) {
        rowClass.push("active")
      }
      if (isUnknown(param1 as Device)) {
        rowClass.push("invading")
      }
      return rowClass // .join(' ')
    }
    if (item === "record") {
      const rowClass = []
      if (isActive(param1 as Device)) {
        rowClass.push("active")
      }
      if (isUnknown(param1 as Device)) {
        rowClass.push("invading")
      }
      return rowClass // .join(' ')
    } else {
      let colClass = [item]
      switch (item) {
        case "ip":
          if (param1 == 0) {
            colClass.push("unreserved")
          }
      }
      return colClass
    }
  }

  const fuzzyTime = (time: string) => {
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

  const calcMaxHeight = computed(() => {
    if (props.maxHeight !== null) {
      return "max-height: " + (props.maxHeight - 68.5) + "px; overflow: auto;"
    }
    return ""
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

  const getLabelByID = (params: Array<Param>, label: string, selectedID: number) => {
    let value = "missing"
    if (params.length > 0) {
      let selected = params.find((par) => par.ID === selectedID)
      if (selected) {
        switch (label) {
          case "Label":
            value = selected.Label ? selected.Label : "missing"
            break
          case "Vendor":
            value = selected.Vendor ? selected.Vendor : "missing"
            break
          case "Family":
            value = selected.Family ? selected.Family : "missing"
            break
          case "Version":
            value = selected.Version ? selected.Version : "missing"
            break
          case "Name":
            value = selected.Name ? selected.Name : "missing"
            break
          case "Icon":
            value = selected.Icon ? selected.Icon : "missing"
            break
        }
      }
    }
    return value
  }
</script>

<template>
  <div class="summary datagrid">
    <modal-box v-model="isModalActive" title="Sample modal">
      <p>
        Lorem ipsum dolor sit amet
        <b>adipiscing elit</b>
      </p>
      <p>This is sample modal</p>
    </modal-box>

    <modal-box v-model="isModalDangerActive" large-title="Please confirm" button="danger" has-cancel>
      <p>
        Lorem ipsum dolor sit amet
        <b>adipiscing elit</b>
      </p>
      <p>This is sample modal</p>
    </modal-box>
    <div :style="calcMaxHeight">
      <table class="multibody">
        <thead>
          <tr>
            <!-- <th class="expand" /> -->
            <th class="device">Device</th>
            <th class="network">Network</th>
            <th class="details">Details</th>
            <th class="seen">First/Last Seen</th>
            <th class="connectivity" title="the online/offline presence of the device over time">Connectivity</th>
            <th />
          </tr>
        </thead>
        <tbody v-for="client in itemsPaginated" :key="client.Interfaces[0].Addresses[0].ID">
          <tr :class="dynamicClass('row', client)">
            <!-- <td class="expand">
              <div v-if="hasExpand(client)" @click="toggleDeviceState(client.ID)">
                <mdIcon v-if="isDeviceOpen(client.ID)" icon="MinusBox" size="20" />
                <mdIcon v-else icon="PlusBox" size="20" />
              </div>
            </td> -->
            <td data-label="Device" class="device">
              <div class="flex">
                <div class="icon">
                  <mdIcon :icon="getLabelByID(DeviceTypes, 'Icon', client.DeviceTypeID)" size="24" />
                </div>
                <div>
                  <div>{{ client.MachineName }}</div>
                </div>
              </div>
            </td>
            <td data-label="MAC/IP" class="network">
              <div class="flex">
                <div class="icon">
                  <mdIcon :icon="iconIFace(client.Primary.InterfaceTypeID)" size="24" />
                </div>
                <div>
                  <div class="whitespace-nowrap">{{ client.Primary.MAC }}</div>
                  <div :class="dynamicClass('ip', client.Primary.IsReservedIP)">{{ client.Primary.IPv4 }}</div>
                </div>
              </div>
            </td>
            <td data-label="Details" class="details">
              <div class="whitespace-nowrap">{{ dynamicClass("row", client).includes("invading") ? "" : (client.Brand ? client.Brand : "?") + " " + (client.Model ? client.Model : "?") }}</div>
              <div>{{ client.Label }}</div>
            </td>
            <td data-label="First/Last Seen" class="seen">
              <div class="text-gray-500 dark:text-gray-400 whitespace-nowrap first">{{ fuzzyTime(client.FirstSeen) }}</div>
              <div class="text-gray-500 dark:text-gray-400 whitespace-nowrap last">{{ fuzzyTime(client.Interfaces[0].Addresses[0].LastSeen) }}</div>
            </td>
            <td data-label="Connectivity" class="connectivity">
              <div><connectivity-indicator :includeDate="false" :data="client.Interfaces[0].Addresses[0].Connectivity" /></div>
            </td>
            <td />
          </tr>
        </tbody>
      </table>
    </div>
    <paginator v-if="perPage == null || (perPage !== null && items.length > perPage)" class="table-pagination" :numItems="items.length" @changePage="setCurrentPage" />
  </div>
</template>

<style scoped lang="scss">
  @import url("@/components/grids/grids.scss");
  .summary.datagrid {
    table {
      thead,
      tbody {
        tr {
          td,
          th {
            &.device {
              @apply lg:w-[330px];
            }
            &.details {
              @apply lg:w-[330px];
            }

            &.connectivity {
              @apply lg:w-[200px];
            }

            &.network {
              @apply lg:w-[200px];
            }
            td.network .unreserved {
              opacity: 0.5;
            }

            &.seen {
              @apply lg:w-[160px];
              text-align: center;
            }
          }
        }
        &.invading {
          td.seen {
            &,
            .first {
              @apply text-red-800 dark:text-red-600;
            }
          }
        }
        &.active {
          td.seen {
            .last {
              @apply text-teal-800 dark:text-teal-300;
            }
          }
        }
      }
    }
  }

  /* tr:nth-child(odd) td {
      @apply lg:bg-gray-50 lg:dark:bg-gray-800;
    } */
</style>
