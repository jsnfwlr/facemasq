<script setup lang="ts">
  import { Component, computed, ref, onMounted, onBeforeUnmount } from 'vue'

  import NavbarItem from '@/components/menus/NavbarItem.vue'
  import mdIcon from '@/components/elements/MDIcon.vue'

  interface Props {
      hasDivider: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    hasDivider: false,
  })

  const isDropdownActive = ref(false)
  const toggleDropdownIcon = computed(() => isDropdownActive.value ? "ChevronUp" : "ChevronDown")
  const toggle = () => { isDropdownActive.value = !isDropdownActive.value }
  const root = ref<Component>()

  // @TODO: Get this working so that navigation actions close the menus
  // const forceClose = (event: Event) => { if (!root.value.$el.contains(event.target)) { isDropdownActive.value = false } }

  // onMounted(() => {
  //   window.addEventListener('click', forceClose)
  // })

  // onBeforeUnmount(() => {
  //   window.removeEventListener('click', forceClose)
  // })
</script>

<template>
  <navbar-item ref="root" type="block" :has-divider="hasDivider" :active="isDropdownActive" dropdown class="dropdown" @click="toggle" >
    <a class="flex items-center py-0 px-3 bg-gray-100 dark:bg-gray-800 lg:bg-transparent lg:dark:bg-transparent">
      <slot />
      <mdIcon :icon="toggleDropdownIcon" class="hidden lg:inline-flex transition-colors" />
    </a>
    <div class="text-sm border-gray-100 border-b lg:border-b-0 lg:border-gray-200 lg:border-t bg-white lg:absolute lg:top-full lg:left-0 lg:min-w-full lg:z-20 lg:shadow-md lg:rounded-b dark:bg-gray-800 dark:border-gray-700" :class="{'lg:hidden':!isDropdownActive}" >
      <slot name="dropdown" />
    </div>
  </navbar-item>
</template>
