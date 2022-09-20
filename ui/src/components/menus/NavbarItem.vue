<script setup lang="ts">
  import { computed } from "vue"

  interface Props {
    href?: string | null
    to?: string | null
    type?: string
    activeColor?: string
    hasDivider?: boolean
    isDesktopIconOnly?: boolean
    dropdown?: boolean
    active?: boolean
  }

  const props = withDefaults(defineProps<Props>(), {
    href: null,
    to: null,
    type: "flex",
    activeColor: "text-teal-500",
    hasDivider: false,
    isDesktopIconOnly: false,
    dropdown: false,
    active: false,
  })

  const is = computed(() => {
    if (props.href) {
      return "a"
    }

    if (props.to) {
      return "router-link"
    }

    return "div"
  })

  const componentClass = computed(() => {
    const base = [props.type, "items-center", "grow-0", "shrink-0", "relative", "cursor-pointer", "hover:text-teal-500", props.active ? props.activeColor : "text-black dark:text-white dark:hover:text-gray-400"]

    if (props.type === "block") {
      base.push("lg:flex")
    }

    if (!props.dropdown) {
      base.push("py-2", "px-3")
    } else {
      base.push("p-0", "lg:py-2", "lg:px-3")
    }

    if (props.hasDivider) {
      base.push("lg:border-r", "lg:border-gray-100", "lg:dark:border-gray-800")
    }

    if (props.isDesktopIconOnly) {
      base.push("lg:w-16", "lg:justify-center")
    }

    return base
  })

  const activeClass = computed(() => {
    return is.value === "router-link" ? props.activeColor : null
  })
</script>

<template>
  <component :is="is" :class="componentClass" :to="to" :href="href" :exact-active-class="activeClass">
    <slot />
  </component>
</template>
