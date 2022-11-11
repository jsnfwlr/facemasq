<script setup lang="ts">
  import { onMounted } from "vue"
  import layout from "@/layouts/index.vue"
  import { useUser } from "@/stores/user"
  import { useParams } from "@/stores/params"
  import { useApp } from "@/stores/app"

  const userStore = useUser()
  const paramsStore = useParams()
  const appStore = useApp()
  onMounted(() => {
    userStore.getSettings()
    paramsStore.getParams()
    appStore.getSettings()
  })
</script>

<template>
  <layout>
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>
  </layout>
</template>
