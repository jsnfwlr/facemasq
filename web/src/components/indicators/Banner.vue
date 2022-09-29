<script setup>
  import { ref, computed, useSlots } from "vue"
  import { storeToRefs } from "pinia"

  import { useUser } from "@/stores/user"
  import { colorsBg, colorsBorders, colorsOutline } from "@/colors.js"

  import mdIcon from "@/components/elements/MDIcon.vue"
  import Btn from "@/components/elements/Btn.vue"

  const userStore = useUser()
  const { darkMode } = storeToRefs(userStore)

  const props = defineProps({
    icon: {
      type: String,
      default: null,
    },
    outline: Boolean,
    color: {
      type: String,
      required: true,
    },
  })
  const emit = defineEmits(["dismiss"])

  const componentClass = computed(() => (props.outline ? colorsOutline[props.color] : [colorsBg[props.color], colorsBorders[props.color]]))
  const isDismissed = ref(false)
  const dismiss = () => {
    isDismissed.value = true
    emit("dismiss", isDismissed)
  }

  const slots = useSlots()
  const hasRight = computed(() => slots.right)
  const hasCenter = computed(() => slots.center)
</script>

<template>
  <div v-if="!isDismissed" :class="componentClass" class="px-3 py-6 md:py-3 mx-6 md:mx-0 mb-6 last:mb-0 border rounded transition-colors duration-150 banner">
    <div class="justify-between items-center parent block md:flex">
      <div class="flex items-center justify-center left grow-0 shrink-0">
        <div class="flex flex-col md:flex-row items-center">
          <mdIcon v-if="icon" :icon="icon" w="w-10 md:w-5" h="h-10 md:h-5" size="24" class="md:mr-2" />
          <span class="text-center md:text-left"><slot /></span>
        </div>
      </div>
      <div v-if="hasCenter" class="flex items-center justify-center center grow w-full">
        <slot name="center" />
      </div>
      <div v-if="hasRight" class="flex items-center justify-center right grow-0 shrink-0 mb-6 md:mb-0 md:mr-3">
        <slot name="right" />
      </div>
      <div class="flex items-center justify-center icon shrink-0 grow-0">
        <slot name="icon">
          <btn icon="Close" :outline="outline || (darkMode && ['white', 'light'].indexOf(color) < 0)" small @click="dismiss" />
        </slot>
      </div>
    </div>
  </div>
</template>
