<script setup lang="ts">
  import { shallowRef, watch, computed } from "vue"
  import { useRoute } from "vue-router"

  import { useUser } from "@/stores/user"

  import Standard from "@/layouts/Standard.vue"
  import Authentication from "@/layouts/Authentication.vue"

  const route = useRoute()
  const userStore = useUser()

  const routeLayout = shallowRef<string | unknown>(Standard)
  const layoutName = shallowRef<string>("Standard")

  const theme = computed(() => userStore.theme())

  watch(
    () => route.meta.layout,
    (layout) => {
      switch (layout || "Standard") {
        case "Authentication":
          routeLayout.value = Authentication
          layoutName.value = "authentication"
          break
        default:
          routeLayout.value = Standard
          layoutName.value = "standard"
          break
      }
    },
  )
</script>

<template>
  <main :class="[theme, layoutName]">
    <component :is="routeLayout">
      <slot></slot>
    </component>
  </main>
</template>
