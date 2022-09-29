<script setup>
import { computed } from 'vue'

const props = defineProps({
  username: {
    type: String,
    default: null
  },
  bg: {
    type: String,
    default: 'bg-gray-100 dark:bg-gray-800'
  },
  api: {
    type: String,
    default: 'api/avataaars'
  }
})

const avatar = computed(() => props.username
  ? `https://avatars.dicebear.com/${props.api}/${props.username.replace(/[^a-z0-9]+/i, '-')}.svg`
  : `https://avatars.dicebear.com/${props.api}/faceMasq`)

const name = computed(() => props.username ? props.username : store.state.userName)
</script>

<template>
  <div>
    <img :src="avatar" :alt="name" class="rounded-full block h-auto w-full max-w-full" :class="bg" >
  </div>
</template>
