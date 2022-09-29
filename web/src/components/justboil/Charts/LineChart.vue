<script setup>
import { ref, watch, computed, onMounted } from 'vue'
import {
  Chart,
  LineElement,
  PointElement,
  LineController,
  LinearScale,
  CategoryScale,
  Tooltip
} from 'chart.js'

const props = defineProps({
  color: {
    type: String,
    required: true
  },
  data: {
    type: Object,
    required: true
  },
  height: {
    type: Number,
    default: null
  }

})

const root = ref(null)

let chart

Chart.register(LineElement, PointElement, LineController, LinearScale, CategoryScale, Tooltip)

const updateData = (color, chartData) => {
  const labels = []
  const points = []
  chartData.forEach((record, i) => {
    labels.push(record.Time)
    points.push(record.Addresses)
  })

  return {
    labels: labels,
    datasets: [
      {
        fill: false,
        borderColor: color,
        borderWidth: 1,
        borderDash: [],
        borderDashOffset: 0.0,
        pointBackgroundColor: color,
        pointBorderColor: 'rgba(255,255,255,0)',
        pointHoverBackgroundColor: color,
        pointBorderWidth: 20,
        pointHoverRadius: 4,
        pointHoverBorderWidth: 15,
        pointRadius: 1,
        tension: 0.5,
        cubicInterpolationMode: 'monotone',
        data: points
      }

    ]
  }
}

const calcHeight = computed(() => {
  if (props.height !== null) {
    return "height: " +  (props.height - 59.5) + "px;"
  }
})

const chartData = computed(() => props.data)

onMounted(() => {
  chart = new Chart(root.value, {
    type: 'line',
    data: updateData(props.color, props.data),
    options: {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        y: {
          display: true,
          title: {
            display: true,
            text: 'Devices'
          }
        },
        x: {
          display: false,
          title: {
            display: true,
            text: 'Time'
          }
        }
      },
      plugins: {
        legend: {
          display: false
        }
      }
    }
  })
})

watch(chartData, data => {
  if (chart) {
    chart.data = updateData(props.color, data)
    chart.update()
  }
})
</script>

<template>
  <canvas ref="root" :style="calcHeight" />
</template>
