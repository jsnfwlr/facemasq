<script setup lang="ts">
import { computed, ref } from 'vue'
import { storeToRefs } from 'pinia'

import { useApp } from '@/stores/app'
import { useUser } from '@/stores/user'

import Btn from '@/components/elements/Btn.vue'
import Btns from '@/components/elements/Btns.vue'
import Level from '@/components/containers/Level.vue'

const userStore = useUser()
const appStore = useApp()
const { settings } = storeToRefs(userStore)

interface Props {
    numItems: number;
    perPage: number | null;
}

const props = withDefaults(defineProps<Props>(), {
  numItems: 1,
  perPage: null
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
  if (props.perPage !== null) {
    return props.perPage
  } else if (appStore.values.perPage === 0)  {
    return props.numItems
  } else {
    return appStore.values.perPage
  }
})

const emit = defineEmits(["changePage"])

const setPageSize = (size: number) => {
  if (size === 0) {
    setCurrentPage(0)   
  } else {
    let target = ((Math.round(((appStore.values.perPage * currentPage.value) + 1) / size) - 1) < 0) ? 0 : (Math.round(((appStore.values.perPage * currentPage.value) + 1) / size) - 1)
    setCurrentPage(target)
  }
  appStore.$patch((state) => { state.values.perPage = size })
}
const setCurrentPage = (page: number) => {
    currentPage.value = page
    emit('changePage', page)
}

const isDark = computed(() => { return userStore.isDarkMode() })

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
