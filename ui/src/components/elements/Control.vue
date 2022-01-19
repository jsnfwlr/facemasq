<script setup>
  import { computed, ref, onMounted, onBeforeUnmount } from "vue"

  import ControlIcon from "@/components/justboil/ControlIcon.vue"

const props = defineProps({
  name: {
    type: String,
    default: null
  },
  id: {
    type: String,
    default: null
  },
  autocomplete: {
    type: String,
    default: null
  },
  placeholder: {
    type: String,
    default: null
  },
  icon: {
    type: String,
    default: null
  },
  options: {
    type: Array,
    default: null
  },
  type: {
    type: String,
    default: "text"
  },
  modelValue: {
    type: [String, Number, Boolean, Array, Object],
    default: ""
  },
  required: Boolean,
  borderless: Boolean,
  transparent: Boolean,
  ctrlKFocus: Boolean,
  disabled: Boolean,
  column: Boolean,
})

const emit = defineEmits(["update:modelValue", "right-icon-click"])

const computedValue = computed({
  get: () => {
    // if (props.type === "checkbox") {
    //   return (props.modelValue === 1) ? true : false
    // }
    return props.modelValue
  },
  set: value => {
    emit("update:modelValue", value)
  }
})

const inputElClass = computed(() => {
  const base = [
    "px-3 py-2 focus:ring focus:outline-none border-gray-700 rounded",
    "dark:placeholder-gray-400",
    computedType.value !== "checkbox" ? "max-w-full w-full" : "w-6",
    computedType.value === "textarea" ? "h-24" :  computedType.value !== "checkbox" ? "h-12" : "h-6",
    props.borderless ? "border-0" : "border",
    props.transparent ? "bg-transparent" : "bg-white dark:bg-gray-800"
  ]

  if (props.icon) {
    base.push("pl-10")
  }

  return base
})


const isMulti = computed(() => {
  return ((props.options) && (["radio","switch","checkbox"].includes(props.type)))
})



const computedType = computed(() => {
  if (props.options) {
    if (!["radio","switch","checkbox"].includes(props.type)) {
      return "select"
    }
  } 
  return props.type
})

const controlIconH = computed(() => props.type === "textarea" ? "h-full" : "h-12")
const controlIconW = computed(() => "w-6")


const inputEl = ref(null)

// if (props.ctrlKFocus) {
//   const fieldFocusHook = e => {
//     if (e.ctrlKey && e.key === "k") {
//       e.preventDefault()
//       inputEl.value.focus()
//     } else if (e.key === "Escape") {
//       inputEl.value.blur()
//     }
//   }

//   onMounted(() => {
//     if (!store.state.isFieldFocusRegistered) {
//       window.addEventListener("keydown", fieldFocusHook)

//       store.commit("basic", {
//         key: "isFieldFocusRegistered",
//         value: true
//       })
//     } else {
//       // console.error("Duplicate field focus event")
//     }
//   })

//   onBeforeUnmount(() => {
//     window.removeEventListener("keydown", fieldFocusHook)

//     store.commit("basic", {
//       key: "isFieldFocusRegistered",
//       value: false
//     })
//   })
// }
</script>

<template>
  <div v-if="isMulti" class="flex justify-start flex-wrap -mb-3" :class="{'flex-col':column}">
    <label v-for="(value, key) in options" :key="key" :class="type" class="mr-6 mb-3 last:mr-0">
      <input v-model="computedValue" :type="computedType" :name="name" :value="key" :disabled="disabled">
      <span class="check" />
      <span class="control-label">{{ value }}</span>
    </label>
  </div>
  <div v-else class="relative">
    <select v-if="computedType === 'select'" :id="id" v-model="computedValue" :name="name" :class="inputElClass" :disabled="disabled">
      <option v-for="(option, index) in options" :key="index" :value="option.value"> {{ option.label }} </option>
    </select>
    <textarea v-else-if="computedType === 'textarea'" :id="id" v-model="computedValue" :class="inputElClass" :name="name" :placeholder="placeholder" :required="required" :disabled="disabled" />
    <input v-else-if="['radio','switch','checkbox'].includes(computedType)" :id="id" v-model="computedValue" :name="name" :autocomplete="autocomplete" :required="required" :placeholder="placeholder" :type="computedType" :class="inputElClass" :disabled="disabled">
    <input v-else :id="id" ref="inputEl" v-model="computedValue" :name="name" :autocomplete="autocomplete" :required="required" :placeholder="placeholder" :type="computedType" :class="inputElClass" :disabled="disabled">
    <control-icon v-if="icon" :icon="icon" :h="controlIconH" :w="controlIconW" class="ml-3" />
  </div>
</template>