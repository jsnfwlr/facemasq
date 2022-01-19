<script setup lang="ts">
  import { computed, useSlots, ref } from 'vue'

  // defineProps({
  //   label: {
  //     type: String,
  //     default: null
  //   },
  //   help: {
  //     type: String,
  //     default: null
  //   }
  // })

  interface Props {
      help?: string;
      label?: string;
      inline: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    inline: false
  })

  const slots = useSlots()

  const makeSlots = (allSlots: any) => {
    return allSlots.default().map((element: any) => {
        return element
    })
  }

  const slotsLength = ref(makeSlots(slots).length)

  const fieldClass = computed(() => {
    const classes = []
    if (props.inline) {
      classes.push("mr-6", "last:mr-0")
      if (slotsLength.value !== 0) {
        classes.push("flex", "flex-nowrap")
      }
    } else {
      classes.push("mb-6", "last:mb-0")
    }
    return classes
  })

  const labelClass = computed(() => {
    const classes = ["block", "font-bold"]
    if (slotsLength.value !== 0) {
      if (props.inline) {
        classes.push("mr-2")
      } else {
        classes.push("mb-2")

      }
    }
    return classes
  })

  const slotClass = computed(() => {
    const classes = ['control']
    if (slotsLength.value > 1) {
      classes.push('grid grid-cols-1 gap-3')
    }
    if (slotsLength.value === 2) {
      classes.push('md:grid-cols-2')
    }
    return classes
  })

</script>

<template>
  <div :class="fieldClass">
    <label v-if="label" :class="labelClass" >{{ label }}</label>
    <div :class="slotClass">
      <slot />
    </div>
    <div v-if="help" class="text-xs text-gray-500 dark:text-gray-400 mt-1" >
      {{ help }}
    </div>
  </div>
</template>
