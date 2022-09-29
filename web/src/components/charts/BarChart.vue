<script setup lang="ts">
  import { ref, watch, computed, onMounted } from "vue"
  import { Chart, BarElement, BarController, LinearScale, CategoryScale, Tooltip, ChartData } from "chart.js"
  import { format, parseISO } from "date-fns"

  import { TimeLog } from "@/stores/devices"

  const props = defineProps<{
    color: string
    data: Array<TimeLog>
    height: number | null
  }>()

  const root = ref()

  let chartjs: Chart

  Chart.register(BarElement, BarController, LinearScale, CategoryScale, Tooltip)

  const updateData = (color: string, chartData: Array<TimeLog>) => {
    const labels = [] as Array<string>
    const points = [] as Array<number>
    chartData.forEach((record) => {
      labels.push(format(parseISO(record.Time), "yyyy-MM-dd HH:mm:ss"))
      points.push(record.Addresses)
    })

    return {
      labels: labels,
      datasets: [
        {
          backgroundColor: color,
          barPercentage: 1,
          minBarLength: 2,
          borderColor: color,
          borderWidth: 0,
          borderDash: [],
          borderDashOffset: 0.0,
          data: points,
        },
      ],
    } as ChartData
  }

  const calcHeight = computed(() => {
    if (props.height !== null) {
      return "height: " + (props.height - 59.5) + "px;"
    }
    return ""
  })

  const chartData = computed(() => props.data)

  onMounted(() => {
    chartjs = new Chart(root.value, {
      type: "bar",
      data: updateData(props.color, props.data),
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          y: {
            display: true,
            title: {
              display: true,
              text: "Devices",
            },
          },
          x: {
            display: false,
            title: {
              display: true,
              text: "Time",
            },
          },
        },
        plugins: {
          legend: {
            display: false,
          },
        },
      },
    })
  })

  watch(chartData, (data) => {
    if (chartjs) {
      chartjs.data = updateData(props.color, data)
      chartjs.update()
    }
  })
</script>

<template>
  <canvas ref="root" :style="calcHeight" />
</template>
