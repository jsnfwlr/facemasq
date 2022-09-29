<script setup>
  import CardComponent from "@/components/justboil/CardComponent.vue"
  import GrowingNumber from "@/components/justboil/GrowingNumber.vue"
  import Icon from "@/components/justboil/Icon.vue"
  import Level from "@/components/containers/Level.vue"
  import TrendPill from "@/components/justboil/TrendPill.vue"

  defineProps({
    number: {
      type: Number,
      default: 0,
    },
    icon: {
      type: String,
      default: null,
    },
    prefix: {
      type: String,
      default: null,
    },
    suffix: {
      type: String,
      default: null,
    },
    label: {
      type: String,
      default: null,
    },
    color: {
      type: String,
      default: null,
    },
    trend: {
      type: String,
      default: null,
    },
    trendType: {
      type: String,
      default: null,
    },
  })
</script>

<template>
  <card-component>
    <level v-if="trend" class="mb-3" mobile>
      <trend-pill :trend="trend" :trend-type="trendType" small />
    </level>
    <level mobile>
      <div>
        <h3 class="text-lg leading-tight text-gray-500 dark:text-gray-400">
          {{ label }}
        </h3>
        <h1 class="text-3xl leading-tight font-semibold">
          <growing-number :value="number" :prefix="prefix" :suffix="suffix" />
        </h1>
      </div>
      <icon v-if="icon" :path="icon" size="48" w="" h="h-16" :class="color" />
    </level>
  </card-component>
</template>
