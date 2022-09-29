<script setup lang="ts">
import { computed } from 'vue'
import { useApp } from '@/stores/app'
const appStore = useApp()

const isFullScreen = computed(() => appStore.toggles.isFullScreen)
const isSidebarActive = computed(() => appStore.toggles.isSidebarActive)
</script>

<template>
  <section class="px-0 md:px-6" :class="[ isFullScreen ? 'flex h-screen items-center justify-center' : 'py-6', isSidebarActive ? 'lg:ml-60' : 'lg:ml-12' ]" >
    <slot />
  </section>
</template>
