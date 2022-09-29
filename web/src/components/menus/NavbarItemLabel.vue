<script setup lang="ts">
  import mdIcon from "@/components/elements/MDIcon.vue"

  interface Props {
    isDesktopIconOnly?: boolean
    icon: string
    label: string
  }

  const props = withDefaults(defineProps<Props>(), {
    isDesktopIconOnly: false,
  })
</script>

<template>
  <mdIcon :icon="icon" class="transition-colors" w="w-10 lg:w-6" h="h-10 lg:h-6" />
  <span class="px-2 transition-colors text-xl lg:text-lg leading-8" :class="{ 'lg:hidden': props.isDesktopIconOnly }">{{ label }}</span>
</template>
