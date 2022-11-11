<script setup lang="ts">
  import { computed, ref } from "vue"

  import { useApp } from "@/stores/app"
  import { useUser } from "@/stores/user"

  import Btn from "@/components/elements/Btn.vue"
  import Btns from "@/components/elements/Btns.vue"
  import Level from "@/components/containers/Level.vue"

  const appStore = useApp()
  const userStore = useUser()

  interface Props {
    numItems: number
    perPage?: number
  }

  const props = withDefaults(defineProps<Props>(), {
    numItems: 1,
    perPage: 0,
  })

  const currentPage = ref(0)
  const currentPageHuman = computed(() => currentPage.value + 1)
  const numPages = computed(() => Math.ceil(props.numItems / perPageCalc.value))
  const pageSizes = computed(() => appStore.values.pageSizes)

  const pagesList = computed(() => {
    const pagesList = []
    for (let i = 0; i < numPages.value; i++) {
      pagesList.push(i)
    }
    return pagesList
  })

  const perPageCalc = computed(() => {
    if (props.perPage !== 0) {
      return props.perPage
    } else {
      return props.numItems
    }
  })

  const emit = defineEmits(["changePage", "setPageSize"])

  const setPageSize = (size: number) => {
    if (size === 0) {
      setCurrentPage(0)
    } else {
      let target = Math.round((appStore.values.perPage * currentPage.value + 1) / size) - 1 < 0 ? 0 : Math.round((appStore.values.perPage * currentPage.value + 1) / size) - 1
      setCurrentPage(target)
    }
    emit("setPageSize", size)
  }
  const setCurrentPage = (page: number) => {
    currentPage.value = page
    emit("changePage", page)
  }

  const isDark = computed(() => {
    return userStore.isDarkMode()
  })
</script>
<template>
  <level>
    <btns>
      <btn v-for="page in pagesList" :key="page" :active="page === currentPage" :label="page + 1" :outline="isDark" small @click="setCurrentPage(page)" />
    </btns>
    <div class="flex">
      <span class="block mr-3 mt-1">Items Per Page:</span>
      <btns>
        <btn v-for="(size, index) in pageSizes" :key="index" :active="size.value === appStore.values.perPage" :label="size.label" :outline="isDark" small @click="setPageSize(size.value)" />
      </btns>
    </div>
    <small>Page {{ currentPageHuman }} of {{ numPages }}</small>
  </level>
</template>

<style scoped lang="scss">
  .table-pagination {
    @apply px-6 py-3 border-t border-gray-100 dark:border-gray-700;
  }
</style>
