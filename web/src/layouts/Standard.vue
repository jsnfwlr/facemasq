<script setup lang="ts">
  import { computed } from "vue"
  import { useApp } from "@/stores/app"

  import Navbar from "@/components/menus/Navbar.vue"
  import Sidebar from "@/components/menus/Sidebar.vue"

  const appStore = useApp()

  const isFullScreen = computed(() => appStore.toggles.isFullScreen)
  const isSidebarActive = computed(() => appStore.toggles.isSidebarActive)
</script>

<template>
  <navbar />
  <sidebar />
  <article class="px-0 md:px-6" :class="[isFullScreen ? 'flex h-screen items-center justify-center' : 'py-6', isSidebarActive ? 'lg:ml-60' : 'lg:ml-12']">
    <slot />
  </article>
</template>
