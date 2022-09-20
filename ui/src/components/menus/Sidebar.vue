<script setup lang="ts">
  import { computed, watch } from "vue"
  import { useRoute } from "vue-router"

  import { useApp } from "@/stores/app"
  import { primaryMenu, secondaryMenu, Item } from "@/data/menu"

  import SidebarList from "@/components/menus/SidebarList.vue"
  import NavbarItem from "@/components/menus/NavbarItem.vue"
  import mdIcon from "@/components/elements/MDIcon.vue"

  const appStore = useApp()
  const showSidebar = computed(() => !appStore.toggles.isFullScreen)
  const isSidebarActive = computed(() => appStore.toggles.isSidebarActive)
  const showLabels = computed(() => appStore.toggles.isSidebarActive || appStore.toggles.isSidebarActive)
  const sidebarToggle = () => {
    appStore.$patch((state) => {
      state.toggles.isSidebarActive = !state.toggles.isSidebarActive
    })
  }
  const menuClick = (event: Event, item: Item) => {
    console.log(event, item)
  }

  const route = useRoute()
  watch(
    () => route.fullPath,
    () => {
      appStore.$patch((state) => {
        state.toggles.isSidebarActive = false
      })
    },
  )
</script>

<template>
  <aside v-show="showSidebar" id="aside" class="fixed top-[64px] lg:top-0 z-40 h-screen bg-white transition-position lg:left-0 dark:bg-gray-900" :class="[isSidebarActive ? 'lg:block w-full lg:w-60 left-0' : 'lg:block w-0 lg:w-12 -left-60']">
    <div class="flex-row w-full bg-white dark:bg-gray-900 text-white flex-1 h-24 lg:h-14 items-center hidden lg:flex">
      <navbar-item active-color="text-black dark:text-white" @click="sidebarToggle">
        <mdIcon icon="Menu" class="cursor-pointer" size="24" />
      </navbar-item>
      <navbar-item to="/" active-color="text-black dark:text-white" class="flex-1 px-3 text-xl lg:text-lg"> <span>face</span><b class="font-black">Masq</b> </navbar-item>
    </div>
    <div class="menus">
      <section>
        <div v-for="(section, index) in primaryMenu" :key="'primary-menu-' + index">
          <hr v-if="!showLabels && section.Label.length > 0" />
          <p v-else-if="section.Label.length > 0" :key="`a-${index}`" class="p-3 text-lg lg:text-base uppercase text-gray-500 dark:text-gray-400">
            {{ section.Label }}
          </p>
          <sidebar-list :key="`b-${index}`" :items="section.Items" @menu-click="menuClick" />
        </div>
      </section>
      <section>
        <div v-for="(section, index) in secondaryMenu" :key="'secondary-menu-' + index">
          <hr v-if="!showLabels && section.Label.length > 0" />
          <p v-else-if="section.Label.length > 0" :key="`a-${index}`" class="p-3 text-lg lg:text-base uppercase text-gray-500 dark:text-gray-400">
            {{ section.Label }}
          </p>
          <sidebar-list :key="`b-${index}`" :items="section.Items" @menu-click="menuClick" />
        </div>
      </section>
    </div>
  </aside>
</template>

<style lang="scss">
  aside {
    // & {
    //   @apply ;
    // }
    hr {
      @apply dark:border-gray-700;
    }
    .menus {
      @apply flex flex-col justify-between border-t border-r dark:border-gray-700;
      height: calc(100% - 45.5px);
    }
  }
</style>
