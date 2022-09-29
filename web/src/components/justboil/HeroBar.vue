<script setup>
import { mdiThemeLightDark } from '@mdi/js'
import { computed } from 'vue'
import Level from '@/components/containers/Level.vue'
import JbButton from '@/components/justboil/JbButton.vue'
import { useUser } from '@/stores/user'
import { storeToRefs } from 'pinia'

const userStore = useUser()
const { darkMode } = storeToRefs(userStore)
const darkModeToggle = () => { userStore.toggleDarkMode() }
</script>

<template>
  <section class="bg-white border-t border-b border-gray-100 p-6 dark:bg-gray-900 dark:border-gray-900 dark:text-white">
    <level>
      <h1 class="text-3xl font-semibold leading-tight">
        <slot />
      </h1>
      <jb-button :label="darkMode ? 'Go Light' : 'Go Dark'" :icon="mdiThemeLightDark" :outline="darkMode" @click="darkModeToggle" />

    </level>
  </section>
</template>
