<script setup>
  import { computed } from "vue"
  import MainSection from "@/components/justboil/_MainSection.vue"
  import { useUser } from "@/stores/user"
  import { storeToRefs } from "pinia"

  const userStore = useUser()
  const { darkMode } = storeToRefs(userStore)

  const props = defineProps({
    bg: {
      type: String,
      required: true,
      validator: (value) => ["login", "error"].includes(value),
    },
  })

  const componentClass = computed(() => {
    const bgs = {
      login: "bg-gradient-to-tr from-purple-400 via-pink-500 to-red-500",
      loginDark: "bg-gradient-to-tr from-purple-900 via-pink-900 to-red-900",
      error: "bg-gradient-to-tr from-pink-400 via-red-500 to-yellow-500",
      errorDark: "bg-gradient-to-tr from-pink-900 via-red-900 to-yellow-900",
    }

    const bgKey = darkMode.value ? `${props.bg}Dark` : props.bg

    return bgs[bgKey] ?? ""
  })
</script>

<template>
  <main-section :class="componentClass">
    <slot card-class="w-11/12 md:w-7/12 lg:w-6/12 xl:w-4/12 shadow-2xl" card-rounded="rounded-lg" />
  </main-section>
</template>
