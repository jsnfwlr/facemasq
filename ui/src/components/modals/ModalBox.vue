<script setup lang="ts">
import { computed } from 'vue'

import Btn from '@/components/elements/Btn.vue'
import Btns from '@/components/elements/Btns.vue'
import Card from '@/components/containers/Card.vue'
import Overlay from '@/components/containers/Overlay.vue'

// const propsX = {
//   title: {
//     type: String,
//     default: null
//   },
//   largeTitle: {
//     type: String,
//     default: null
//   },
//   button: {
//     type: String,
//     default: 'info'
//   },
//   buttonLabel: {
//     type: String,
//     default: 'Done'
//   },
//   hasCancel: Boolean,
//   modelValue: {
//     type: [String, Number, Boolean],
//     default: null
//   }
// }

interface Props {
    title?: string | null;
    largeTitle?: string | null;
    buttonColor?: string;
    buttonLabel?: string;
    hasCancel: boolean;
    modelValue?: string | number | boolean | null;
}

const props = withDefaults(defineProps<Props>(), {
  title: null,
  largeTitle: null,
  buttonColor: 'info',
  buttonLabel: "Done",
  hasCancel: false,
  modelValue: null
})

const emit = defineEmits(['update:modelValue', 'cancel', 'confirm'])

const value = computed({
  get: () => props.modelValue,
  set: value => emit('update:modelValue', value)
})

const confirmCancel = (mode: string) => {
  value.value = false
  emit(mode === "cancel" ? "cancel" : "confirm")
}

const confirm = () => confirmCancel('confirm')

const cancel = () => confirmCancel('cancel')
</script>

<template>
  <overlay v-show="value" @overlay-click="cancel" >
    <card v-show="value" :headingTitle="title" class="shadow-lg w-full max-h-modal md:w-3/5 lg:w-2/5 z-50" headerIcon="Close" @header-icon-click="cancel">
      <div class="space-y-3">
        <h1 v-if="largeTitle" class="text-2xl" >{{ largeTitle }}</h1>
        <slot />
      </div>

      <hr class="my-6 -mx-6 border-t border-gray-100 dark:border-gray-700">
      <btns>
        <btn :label="buttonLabel" :color="buttonColor" @click="confirm" />
        <btn v-if="hasCancel" label="Cancel" :color="buttonColor" outline @click="cancel" />
      </btns>
    </card>
  </overlay>
</template>
