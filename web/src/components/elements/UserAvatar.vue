<script setup lang="ts">
  import { computed } from "vue"
  import md5 from "crypto-js/md5"

  import { storeToRefs } from "pinia"
  import { useApp } from "@/stores/app"
  import { useUser } from "@/stores/user"

  const appStore = useApp()
  const { settings } = storeToRefs(appStore)

  const userStore = useUser()
  const { account } = storeToRefs(userStore)

  interface Props {
    username?: string
    bg?: string
  }

  const props = withDefaults(defineProps<Props>(), {
    username: "faceMasq",
    bg: "bg-gray-100 dark:bg-gray-800",
  })

  // appStore.saveSetting("avatarType", "https://www.gravatar.com/avatar/")
  // appStore.saveSetting("avatarType", "https://avatars.dicebear.com/api/avataaars/")

  const avatar = computed(() => {
    if (settings.value.avatarType.includes("gravatar")) {
      return settings.value.avatarType + md5(account.value.Username) + "?d=" + "https%3A%2F%2Favatars.dicebear.com%2Fapi%2Finitials%2F" + account.value.Username + ".png"
    } else if (settings.value.avatarType.includes("dicebear")) {
      return settings.value.avatarType + md5(account.value.Username) + ".svg" // .replace(/[^a-z0-9]+/i, "-") + ".svg"
    }
    return "-"
  })
</script>

<template>
  <div>
    <img :src="avatar" :alt="props.username" class="rounded-full block h-auto w-full max-w-full" :class="bg" />
  </div>
</template>
