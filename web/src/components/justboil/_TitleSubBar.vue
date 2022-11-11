<script setup>
  import Icon from "@/components/justboil/_Icon.vue"

  defineProps({
    icon: {
      type: String,
      default: null,
    },
    title: {
      type: String,
      required: true,
    },
  })
</script>

<template>
  <section class="px-6 sm:px-0 mb-6 flex items-center justify-between">
    <div class="flex items-center justify-start">
      <icon v-if="icon" :path="icon" class="mr-3" />
      <h1 class="text-2xl">
        {{ title }}
      </h1>
    </div>
    <slot />
  </section>
</template>
