<script setup lang="ts">
  import { ref, watch, defineAsyncComponent, computed } from "vue"
  import { useRoute } from "vue-router"

  import { useUser } from "@/stores/user"

  const Standard = defineAsyncComponent(() => import("@/layouts/Standard.vue"))
  const Authentication = defineAsyncComponent(() => import("@/layouts/Authentication.vue"))

  const route = useRoute()
  const userStore = useUser()

  const routeLayout = ref<string|unknown>(Standard)
  const layoutName = ref<string>("Standard")

  const theme = computed(() => userStore.theme() )

  watch(() => route.meta.layout,
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
    }
  )
</script>

<template>
  <main :class="[theme, layoutName]">
    <component :is="routeLayout">
      <slot></slot>
    </component>
  </main>
</template>
