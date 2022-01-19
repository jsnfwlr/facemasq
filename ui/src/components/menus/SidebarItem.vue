<script setup lang="ts">
  import { computed } from "vue"

  import { useApp } from "@/stores/app"
  import { Item } from "@/data/menu"
  
  import mdIcon from "@/components/elements/MDIcon.vue"

  const appStore = useApp()

  const props = defineProps<{
    item: Item
  }>()

  const emit = defineEmits(["clickMenu"])
  const linkType = computed(() => props.item.Route ? "router-link": "a")
  const itemTo = computed(() => props.item.To || null)
  const itemTarget = computed(() => linkType.value === "a" && props.item.Target ? props.item.Target : null)
  const showLabels = computed(() => appStore.toggles.isSidebarActive || appStore.toggles.isSidebarActive)
  const tooltip = computed(() => props.item.Tooltip ? props.item.Tooltip : null)


  const menuClick = (event: Event) => {
    emit("clickMenu", event, props.item)
  }

  

  const styleActive = ""

  const styleInactive = ""
</script>

<template>
  <li>
    <component :is="linkType" v-slot="vSlot" :to="itemTo" :href="itemTo" :target="itemTarget" class="item" @click="menuClick" :title="tooltip" :class="linkType">
      <div :class="vSlot && vSlot.isExactActive ? 'active' : ''">
        <mdIcon v-if="item.Icon" :icon="item.Icon" class="icon" w="w-10 lg:w-6" h="h-10 lg:h-6" />
        <span v-if="showLabels" class="label">{{ item.Label }}</span>
      </div>
    </component>
  </li>
</template>

<style scoped lang="scss">
  li {
    .item {
      
      div {
        & { 
          @apply flex cursor-pointer py-2; 
        }
        &:hover {
          @apply bg-teal-700 bg-opacity-50;
        }
        .icon {
          @apply flex-none text-gray-900 dark:text-gray-300 mx-6 lg:mx-3;
        }
        .label {
          @apply grow text-gray-800 dark:text-gray-300 text-xl lg:text-lg leading-8;
        }
      }
      &.active {
        & {
          @apply bg-gray-500 bg-opacity-50 dark:bg-gray-700 dark:bg-opacity-50;
        }
        .icon, .label {
          @apply font-bold text-white;
        }
      }
    }
  }

</style>