<script setup>
  import mdIcon from "@/components/elements/MDIcon.vue"

  defineProps({
    icon: {
      type: String,
      default: null,
    },
    h: {
      type: String,
      default: null,
    },
  })
</script>

<template>
  <mdIcon :icon="icon" w="w-10" :h="h" class="absolute top-0 left-0 z-10 pointer-events-none text-gray-500 dark:text-gray-400" />
</template>
