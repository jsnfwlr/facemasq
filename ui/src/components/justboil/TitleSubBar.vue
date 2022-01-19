<script setup>
import { mdiCog } from '@mdi/js'
import { computed, useSlots } from 'vue'
import Icon from '@/components/justboil/Icon.vue'
import JbButton from '@/components/justboil/JbButton.vue'

defineProps({
  icon: {
    type: String,
    default: null
  },
  title: {
    type: String,
    required: true
  }
})


const slots = useSlots()
const darkMode = true // = computed(() => store.state.darkMode)
</script>

<template>
  <section class="px-6 sm:px-0 mb-6 flex items-center justify-between">
    <div class="flex items-center justify-start">
      <icon v-if="icon" :path="icon" class="mr-3" /> <h1 class="text-2xl">
        {{ title }}
      </h1>
    </div>
    <slot />
  </section>
</template>
