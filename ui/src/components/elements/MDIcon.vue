<script setup lang="ts">
import { computed } from 'vue'

import { icons } from '@/data/icons'


interface Props {
  icon: string;
  w: string;
  h: string;
  p: string;
  size?: string | number | null;
}

const props = withDefaults(defineProps<Props>(), {
  w: 'w-6',
  h: 'h-6',
  p: '',
  size: null
})


const spanClass = computed(() => `inline-flex justify-center items-center ${props.w} ${props.h} ${props.p}`)

const svgClass = computed(() => {
  const classVal = [props.icon]
  if (!props.size) {
    classVal.push(props.w, props.h)
  }
  return classVal
})

const svgPath = computed(() => {
  return icons.find(item => item.name == props.icon)?.svg
})

</script>

<template>
  <span :class="spanClass">
    <svg v-if="size" viewBox="0 0 24 24" :width="size" :height="size" class="inline-block" :class="svgClass" >
      <path fill="currentColor" :d="svgPath" />
    </svg>
    <svg v-else viewBox="0 0 24 24" class="inline-block" :class="svgClass" >
      <path fill="currentColor" :d="svgPath" />
    </svg>
  </span>
</template>

