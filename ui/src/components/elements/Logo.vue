<script setup lang="ts">

import { computed } from 'vue'

interface Props {
    title: string | null;
    size: string | null;
    fill: boolean;
    
}

const props = withDefaults(defineProps<Props>(), {
  title: null,
  size: null,
  fill: false,
})

const computedClass = computed(() => {
    if (props.size === null) {
        return 'w-auto'
    }
})

const computedSize = computed(() => {
    if (props.size === null) {
        return "100%"
    } 
    return props.size
})

</script>

<template>
    <img alt="" src="/favicon.png" :class="computedClass" :width="computedSize" :height="computedSize" />
</template>