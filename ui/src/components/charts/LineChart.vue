<script setup lang="ts">
  import { ref, watch, computed, onMounted } from "vue"
  import { Chart, LineElement, PointElement, LineController, LinearScale, CategoryScale, Tooltip, Filler, ChartData } from "chart.js"
  import { format, parseISO } from "date-fns"

  import { ChartValues, TimeLog } from "@/stores/devices"

  interface Props {
    colors: Array<string>
    data: ChartValues
    height: number | null
    min: number | null
    max: number | null
  }

  interface ChartPoint {
    x?: number | string | Date | undefined
    y?: number | string | Date | undefined
    r?: number | undefined
    t?: number | string | Date | undefined
  }

  const props = withDefaults(defineProps<Props>(), {
    height: null,
    min: null,
    max: null,
  })

  const root = ref("")

  let chart: Chart

  Chart.register(LineElement, PointElement, LineController, LinearScale, CategoryScale, Tooltip, Filler)

  const parseChartData = (chartData: Array<TimeLog>) => {
    const data = {
      labels: [] as Array<string>,
      points: [] as Array<ChartPoint>,
    }
    chartData.forEach((record) => {
      data.labels.push(format(parseISO(record.Time), "yyyy-MM-dd HH:mm:ss"))
      data.points.push({ x: parseISO(record.Time), y: record.Addresses })
    })
    return data
  }

  const updateData = (colors: Array<string>, chartData: ChartValues) => {
    // const fullData = parseChartData(chartData.full)
    const avgdData = parseChartData(chartData.averaged)

    return {
      labels: avgdData.labels,
      datasets: [
        {
          fill: this,
          borderColor: colors[1],
          borderWidth: 1,
          borderDash: [],
          borderDashOffset: 0.0,
          pointBackgroundColor: colors[1],
          pointBorderColor: "rgba(255,255,255,0)",
          pointHoverBackgroundColor: colors[1],
          pointBorderWidth: 20,
          pointHoverRadius: 4,
          pointHoverBorderWidth: 15,
          pointRadius: 0.5,
          data: avgdData.points,
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
    chart = new Chart(root.value, {
      type: "line",
      data: updateData(props.colors, props.data),
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
            display: true,
            title: {
              display: true,
              text: "Time",
            },
          },
        },
        // plugins: {
        //   legend: {
        //     position: 'top'
        //   }
        // }
      },
    })
  })

  watch(chartData, (data) => {
    if (chart) {
      chart.data = updateData(props.colors, data)
      chart.update()
    }
  })
</script>

<template>
  <canvas ref="root" :style="calcHeight" />
</template>
