
<script setup lang="ts">
  import { computed, onMounted } from 'vue'
  import { useApp } from '@/stores/app'
  import { useUser } from '@/stores/user'
  import { useParams } from '@/stores/params'

  import Navbar from '@/components/menus/Navbar.vue'
  import Sidebar from '@/components/menus/Sidebar.vue'

  const appStore = useApp()
  const userStore = useUser()
  const paramsStore = useParams()

  const isFullScreen = computed(() => appStore.toggles.isFullScreen)
  const isSidebarActive = computed(() => appStore.toggles.isSidebarActive)
  

  onMounted(() => {
    userStore.getSettings()
    paramsStore.getParams()
  })
</script>

<template>
  <navbar />
  <sidebar />
  <article class="px-0 md:px-6" :class="[ isFullScreen ? 'flex h-screen items-center justify-center' : 'py-6', isSidebarActive ? 'lg:ml-60' : 'lg:ml-12' ]" >
    <slot />
  </article>
</template>
