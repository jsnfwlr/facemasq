<script setup lang="ts">
  import { Item } from '@/data/menu'

  import SidebarItem from '@/components/menus/SidebarItem.vue'

  const props = defineProps<{
    items: Array<Item>
  }>()

  const emit = defineEmits(['clickMenu'])

  const menuClick = (event: Event, item: Item) => {
    emit('clickMenu', event, item)
  }
</script>

<template>
  <ul>
    <sidebar-item v-for="(item, index) in items" :item="item"  @clickMenu="menuClick" :key="index" />
  </ul>
</template>