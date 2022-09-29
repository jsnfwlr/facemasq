<script setup lang="ts">
  import { computed, ref } from "vue"

  import { useUser } from "@/stores/user"
  import { useApp } from "@/stores/app"
  import { storeToRefs } from "pinia"

  import NavbarItem from "@/components/menus/NavbarItem.vue"
  import NavbarItemLabel from "@/components/menus/NavbarItemLabel.vue"
  import NavbarMenu from "@/components/menus/NavbarMenu.vue"
  import mdIcon from "@/components/elements/MDIcon.vue"

  import UserAvatar from "@/components/justboil/UserAvatar.vue"

  const userStore = useUser()
  const appStore = useApp()
  const { account } = storeToRefs(userStore)

  const userName = account.value.Username

  const isNavBarVisible = computed(() => !appStore.toggles.isFullScreen)
  const isSidebarActive = computed(() => appStore.toggles.isSidebarActive)
  const isDark = computed(() => {
    return userStore.isDarkMode()
  })
  const isMenuNavBarActive = ref(false)

  const lightBorderStyle = computed(() => appStore.styles.lightBorderStyle)

  const menuToggleMobileIcon = computed(() => (isSidebarActive.value ? "Backburger" : "Forwardburger"))
  const menuNavBarToggleIcon = computed(() => (isMenuNavBarActive.value ? "Close" : "DotsVertical"))

  const menuToggleMobile = () => {
    appStore.$patch((state) => {
      state.toggles.isSidebarActive = !state.toggles.isSidebarActive
    })
  }

  const darkModeToggle = () => {
    userStore.toggleDarkMode()
  }
  const menuNavBarToggle = () => {
    isMenuNavBarActive.value = !isMenuNavBarActive.value
  }
  const mode = import.meta.env.DEV
</script>

<template>
  <nav v-show="isNavBarVisible" class="topnav top-0 left-0 right-0 fixed flex bg-white h-[64px] lg:h-14 border-b z-30 w-screen transition-position xl:pl-60 lg:w-auto lg:items-stretch dark:bg-gray-900 dark:border-gray-700 ml-0" :class="[lightBorderStyle, isSidebarActive ? 'lg:ml-60' : 'lg:ml-12']">
    <div class="flex-1 items-stretch flex h-[64px] lg:h-14">
      <navbar-item type="flex lg:hidden" @click.prevent="menuToggleMobile">
        <mdIcon :icon="menuToggleMobileIcon" h="h-[48px] lg:h-10" w="w-[48px] lg:w-10" p="p-3 lg:p-0" />
      </navbar-item>

      <navbar-item to="/" class="flex-1 px-3 lg:hidden text-xl lg:text-lg"> <span>face</span><b class="font-black">Masq</b> </navbar-item>
    </div>
    <div class="flex-none items-stretch flex h-[64px] lg:h-14 lg:hidden">
      <navbar-item @click.prevent="menuNavBarToggle">
        <mdIcon :icon="menuNavBarToggleIcon" h="h-[48px] lg:h-10" w="w-[48px] lg:w-10" p="p-3 lg:p-0" />
      </navbar-item>
    </div>
    <div class="absolute w-screen top-[64px] lg:top-14 left-0 bg-white shadow lg:w-auto lg:items-stretch lg:flex lg:grow lg:static lg:border-b-0 lg:overflow-visible lg:shadow-none dark:bg-gray-900 xl:pr-0" :class="[isMenuNavBarActive ? 'block' : 'hidden']">
      <div class="max-h-screen-menu overflow-y-auto lg:overflow-visible lg:flex lg:items-stretch lg:justify-end lg:ml-auto">
        <!-- <navbar-item has-divider class="xl:pr-24">
            <navbar-search />
        </navbar-item> -->

        <navbar-menu has-divider>
          <user-avatar class="w-10 h-10 lg:w-6 lg:h-6 mr-3 inline-flex" :username="userName" />
          <div class="text-xl lg:text-lg leading-8">
            <span>{{ userName }} - {{ mode }}</span>
          </div>

          <template #dropdown>
            <navbar-item to="/settings">
              <navbar-item-label icon="CogOutline" label="Settings" />
            </navbar-item>
            <hr class="hidden lg:block lg:my-2 border-t border-gray-100 dark:border-gray-700" />
            <navbar-item @click.prevent="darkModeToggle">
              <navbar-item-label icon="ThemeLightDark" :label="isDark ? 'Go Light' : 'Go Dark'" />
            </navbar-item>
            <hr class="hidden lg:block lg:my-2 border-t border-gray-100 dark:border-gray-700" />
            <navbar-item to="/auth/logout">
              <navbar-item-label icon="Logout" label="Log Out" />
            </navbar-item>
          </template>
        </navbar-menu>
      </div>
    </div>
  </nav>
</template>

<style lang="scss">
  nav.topnav {
  }
</style>
