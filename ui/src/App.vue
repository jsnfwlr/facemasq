<script setup>
  import layout from "@/layouts/index.vue"
</script>

<template>
  <layout>
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>
  </layout>
</template>
