<script setup lang="ts">
  import { computed } from "vue"

  import { Trend } from "@/stores/devices"

  import mdIcon from "@/components/elements/MDIcon.vue"
  import Level from "@/components/containers/Level.vue"
  import CardComponent from "@/components/justboil/CardComponent.vue"
  import GrowingNumber from "@/components/justboil/GrowingNumber.vue"
  import TrendPill from "@/components/justboil/TrendPill.vue"

  interface Props {
    icon: string | null
    prefix: string
    suffix: string
    color: string | null
    trend: Trend
  }

  const props = withDefaults(defineProps<Props>(), {
    icon: null,
    color: null,
  })

  const direction = computed(() => {
    if (props.trend.Compare !== 0) {
      if (props.trend.Compare > props.trend.Current) {
        return "down"
      } else if (props.trend.Compare < props.trend.Current) {
        return "up"
      }
    }
    return "equal"
  })

  const percentage = computed(() => {
    if (props.trend.Compare != 0 && props.trend.Compare != props.trend.Current) {
      return Math.round(Math.abs(((props.trend.Current - props.trend.Compare) / props.trend.Compare) * 100)) + "%"
    }
    return null
  })

  const difference = computed(() => {
    return Math.abs(props.trend.Current - props.trend.Compare)
  })
</script>

<template>
  <card-component v-if="trend?.Label !== null">
    <level class="mb-3" mobile>
      <level type="justify-start" class="gap-3">
        <mdIcon v-if="icon" :icon="icon" size="48" w="" h="h-16" :class="color" />
        <h3 class="text-lg leading-tight text-gray-500 dark:text-gray-400" :title="trend.Tooltip !== null ? trend.Tooltip : 'Unique devices detected over the last ' + trend.Label">
          {{ trend.Label }}
        </h3>
      </level>
      <level type="justify-end" class="gap-3">
        <trend-pill v-if="percentage !== null" :trend="percentage" :trend-type="direction" small :title="direction + ' ' + difference + ' compared to previous ' + trend.Label" />
        <h1 class="text-3xl leading-tight font-semibold">
          <growing-number :value="trend.Current" :prefix="prefix" :suffix="suffix" />
        </h1>
      </level>
    </level>
  </card-component>
</template>
