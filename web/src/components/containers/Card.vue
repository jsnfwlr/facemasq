<script setup lang="ts">
  import { ref, computed, useSlots } from "vue"
  import mdIcon from "@/components/elements/MDIcon.vue"

  interface Props {
    headingTitle: string | null
    icon?: string | null
    headerIcon?: string | null
    headerIconTitle?: string
    hideHeaderIcon?: boolean
    rounded?: string
    hasTable?: boolean
    empty?: boolean
    form?: boolean
    hoverable?: boolean
  }

  const props = withDefaults(defineProps<Props>(), {
    headingTitle: null,
    icon: null,
    headerIcon: null,
    headerIconTitle: "more options",
    hideHeaderIcon: false,
    rounded: "md:rounded",
    hasTable: false,
    empty: false,
    form: false,
    hoverable: false,
  })

  const emit = defineEmits(["header-icon-click", "submit"])

  const slots = useSlots()

  const hasFilters = computed(() => slots.filter)

  const is = computed(() => (props.form ? "form" : "section"))

  const componentClass = computed(() => {
    const base = [props.rounded]

    if (props.hoverable) {
      base.push("hover:shadow-lg transition-shadow duration-500")
    }

    return base
  })

  const headerIconClick = () => {
    emit("header-icon-click")
  }

  const submit = (event: Event) => {
    emit("submit", event)
  }

  const showFilters = ref(false)
</script>

<template>
  <component :is="is" :class="componentClass" class="cardContainer bg-white border border-gray-100 dark:bg-gray-900 dark:border-gray-900" @submit="submit">
    <header v-if="headingTitle" class="flex items-stretch border-b border-gray-200 dark:border-gray-700">
      <h2 class="flex items-center py-3 grow font-bold" :class="[icon ? 'px-4' : 'px-6']">
        <mdIcon v-if="icon" :icon="icon" class="mr-3 text-teal-500" :size="18" />
        {{ headingTitle }}
      </h2>
      <div v-if="hasTable && hasFilters" :class="showFilters ? 'filters-wrapper active' : 'filters-wrapper'">
        <a href="#" class="flex items-center py-3 px-4 justify-center ring-blue-700 focus:ring" aria-label="Toggle the filter inputs" :title="showFilters ? 'Hide filters' : 'Show filters'" @click.prevent="showFilters = !showFilters"> <mdIcon :icon="showFilters ? 'ChevronRight' : 'ChevronLeft'" :size="18" /> <mdIcon icon="Filter" :size="18" /> </a>
        <div v-show="showFilters" class="filters">
          <slot name="filter" />
        </div>
      </div>
      <a v-if="headerIcon && !hideHeaderIcon" href="#" class="flex items-center py-3 px-4 justify-center ring-blue-700 focus:ring" :aria-label="headerIconTitle" :title="headerIconTitle" @click.prevent="headerIconClick">
        <mdIcon v-if="headerIcon" :icon="headerIcon" :size="18" />
      </a>
    </header>
    <div v-if="empty" class="text-center py-24 text-gray-500 dark:text-gray-400">
      <p>Nothing's here…</p>
    </div>
    <div v-else :class="{ 'p-6': !hasTable }">
      <slot />
    </div>
  </component>
</template>

<style lang="scss">
  .cardContainer {
    header {
      & {
        height: 60px;
      }
      h2 {
        font-size: 16px;
      }
      .filters-wrapper {
        & {
          @apply flex;
        }
        &.active {
          background: hsla(0, 0%, 0%, 0.2);
        }
        .filters {
          & {
            height: 40px;
            margin-top: 10px;
            @apply flex mx-0 lg:mx-5;
          }
          label {
            line-height: 40px;
          }
          button {
            height: 28px;
            margin-top: 6px;
          }
          .control {
            min-width: 150px;
          }
        }
      }
    }
  }
</style>
