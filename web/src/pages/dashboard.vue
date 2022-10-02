<script setup lang="ts">
  // Libraries
  import { computed, onMounted, onBeforeUnmount } from "vue"
  import { storeToRefs } from "pinia"
  import { parse, subMinutes } from "date-fns"

  // Stores & Data
  import { useUser } from "@/stores/user"
  import { useDevices } from "@/stores/devices"

  // Components
  import Chart from "@/components/charts/LineChart.vue"
  import TrendIndicator from "@/components/indicators/Trend.vue"
  import ThingsGrid from "@/components/grids/Things.vue"
  import Card from "@/components/containers/Card.vue"
  import Banner from "@/components/indicators/Banner.vue"
  import Btn from "@/components/elements/Btn.vue"

  const userStore = useUser()
  const deviceStore = useDevices()

  const { settings } = storeToRefs(userStore)
  const { allDevices, trends, devicesOverTime, lastUnknown } = storeToRefs(deviceStore)

  let timer: ReturnType<typeof setInterval>

  const updateDevices = () => {
    deviceStore.getAll()
    deviceStore.getCharts()
    deviceStore.getTrends()
  }

  onMounted(() => {
    updateDevices()
    timer = setInterval(() => {
      updateDevices()
    }, 0.5 * 60 * 1000)
  })

  onBeforeUnmount(() => {
    clearInterval(timer)
  })

  const fakeTrends = [
    { Label: "Historic", Compare: 0, Current: 0, Tooltip: null },
    { Label: "Concurrent", Compare: 0, Current: 0, Tooltip: null },
    { Label: "30 days", Compare: 0, Current: 0, Tooltip: null },
    { Label: "7 days", Compare: 0, Current: 0, Tooltip: null },
    { Label: "24 hours", Compare: 0, Current: 0, Tooltip: null },
    { Label: "30 minutes", Compare: 0, Current: 0, Tooltip: null },
  ]

  const getMin = computed(() => {
    if (devicesOverTime.value.full.length == 0) {
      return 0
    }
    const counts = devicesOverTime.value.full.map((x) => x.Addresses)
    const min = Math.min(...counts)
    return min - 10 >= 0 ? min - 10 : 0
  })

  const getMax = computed(() => {
    if (devicesOverTime.value.full.length == 0) {
      return 100
    }
    const counts = devicesOverTime.value.full.map((x) => x.Addresses)
    const max = Math.max(...counts)
    return max + 10
  })

  const newUnknown = computed(() => {
    return allDevices.value.length < 0 && settings.value.lastDismissed && lastUnknown.value ? settings.value.lastDismissed < lastUnknown.value : false
  })
  const dismissUnknown = () => {
    userStore.lastDismissedUnknown(new Date())
  }
  const showUnknown = () => {
    console.log("Navigate to unknown")
  }

  // Devices seen in the last 30 minutes
  const recentDevices = computed(() => {
    return allDevices.value.filter((device) => device.Interfaces[0].Addresses[0].LastSeen !== null && parse(device.Interfaces[0].Addresses[0].LastSeen.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date()) > subMinutes(new Date(), 30))
  })
</script>

<template>
  <banner v-if="newUnknown" color="danger" icon="AlertOctagon">
    New unknown devices have been found
    <template v-slot:icon>
      <btn icon="Magnify" small @click="showUnknown" label="Investigate" class="mr-1" />
      <btn icon="Close" small @click="dismissUnknown" />
    </template>
  </banner>
  <div v-if="trends.length > 0" class="grid grid-cols-2 gap-6 lg:grid-cols-6 mb-6">
    <trend-indicator v-for="(trend, index) in trends" :trend="trend" color="text-teal-500" icon="Radar" :key="'trends-' + index" :prefix="''" :suffix="''" />
  </div>
  <div v-else class="grid grid-cols-2 gap-6 lg:grid-cols-6 mb-6">
    <trend-indicator v-for="(trend, index) in fakeTrends" :trend="trend" color="text-teal-500" icon="Radar" :key="'fake-trends-' + index" :prefix="''" :suffix="''" />
  </div>
  <div class="grid grid-cols-1 xl:grid-cols-2 gap-6 mb-6">
    <div>
      <card headingTitle="Devices over time" icon="Finance" class="mb-6">
        <div v-if="devicesOverTime">
          <chart :colors="['#00D1B2', '#3399CC']" :data="devicesOverTime" :height="625.5" :min="getMin" :max="getMax" />
        </div>
      </card>
    </div>
    <div>
      <card icon="MonitorCellphone" :headingTitle="'Recent Devices (' + recentDevices.length + ')'" has-table class="mb-6">
        <things-grid :maxHeight="625.5" :items="recentDevices" />
      </card>
    </div>
  </div>
</template>
