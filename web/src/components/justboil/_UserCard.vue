<script setup>
  import { ref } from "vue"
  import { mdiCheckDecagram } from "@mdi/js"
  import Level from "@/components/containers/Level.vue"
  import UserAvatar from "@/components/justboil/_UserAvatar.vue"
  import CardComponent from "@/components/justboil/CardComponent.vue"
  import CheckRadioPicker from "@/components/justboil/_CheckRadioPicker.vue"
  import Pill from "@/components/justboil/Pill.vue"

  const userName = "faceMasq" // computed(() => store.state.userName)
  const userSwitchVal = ref([])
</script>

<template>
  <card-component rounded="">
    <level type="justify-around lg:justify-center">
      <user-avatar class="lg:mx-12" />
      <div class="space-y-3 text-center md:text-left lg:mx-12">
        <div class="flex justify-center md:block">
          <check-radio-picker v-model="userSwitchVal" name="sample-switch" type="switch" :options="{ one: 'Notifications' }" />
        </div>
        <h1 class="text-2xl">
          Howdy, <b>{{ userName }}</b>
        </h1>
        <p>Last login <b>12 mins ago</b> from <b>127.0.0.1</b></p>
        <div class="flex justify-center md:block">
          <pill text="Verified" type="info" :icon="mdiCheckDecagram" />
        </div>
      </div>
    </level>
  </card-component>
</template>
