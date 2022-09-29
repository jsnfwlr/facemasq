<script setup lang="ts">
  import { onMounted, computed } from "vue"

  import { useParams } from "@/stores/params"
  import { useRoute } from "vue-router"

  import Card from "@/components/containers/Card.vue"
  import ParamGrid from "@/components/grids/Params.vue"

  const route = useRoute()

  const paramsStore = useParams()

  const getIcon = computed(() => {
    switch (route.params.param) {
      case "categories":
        return "Shape"
      case "status":
        return "ListStatus"
      case "maintainers":
        return "Account"
      case "locations":
        return "MapMarker"
      case "devicetypes":
        return "Devices"
      case "operatingsystems":
        return "MicrosoftWindowsClassic"
      case "architectures":
        return "Chip"
      case "vlans":
        return "Vpn"
      case "users":
        return "AccountCircle"
      default:
        return "HelpCircle"
    }
  })

  const getTitle = computed(() => {
    switch (route.params.param) {
      case "categories":
        return "Categories"
      case "status":
        return "Statuses"
      case "maintainers":
        return "Maintainers"
      case "locations":
        return "Locations"
      case "devicetypes":
        return "Device Types"
      case "operatingsystems":
        return "OS Types"
      case "architectures":
        return "CPU Types"
      case "vlans":
        return "Virtual LANs"
      case "users":
        return "Users"
      default:
        return ""
    }
  })

  const routeParam = computed(() => {
    return route.params.param as string
  })

  const addItem = () => {
    switch (route.params.param) {
      case "categories":
        paramsStore.NewCategory()
        break
      case "status":
        paramsStore.NewStatus()
        break
      case "maintainers":
        paramsStore.NewMaintainer()
        break
      case "locations":
        paramsStore.NewLocation()
        break
      case "devicetypes":
        paramsStore.NewDeviceType()
        break
      case "operatingsystems":
        paramsStore.NewOperatingSystem()
        break
      case "architectures":
        paramsStore.NewArchitecture()
        break
      case "vlans":
        paramsStore.NewVLAN()
        break
      case "users":
        paramsStore.NewUser()
        break
    }
  }

  const dropItem = (index: number) => {
    switch (route.params.param) {
      case "categories":
        paramsStore.DropCategory(index)
        break
      case "status":
        paramsStore.DropStatus(index)
        break
      case "maintainers":
        paramsStore.DropMaintainer(index)
        break
      case "locations":
        paramsStore.DropLocation(index)
        break
      case "devicetypes":
        paramsStore.DropDeviceType(index)
        break
      case "operatingsystems":
        paramsStore.DropOperatingSystem(index)
        break
      case "architectures":
        paramsStore.DropArchitecture(index)
        break
      case "vlans":
        paramsStore.DropVLAN(index)
        break
      case "users":
        paramsStore.DropUser(index)
        break
    }
  }

  const deleteItem = (index: number) => {
    switch (route.params.param) {
      case "categories":
        paramsStore.DeleteCategory(index)
        break
      case "status":
        paramsStore.DeleteStatus(index)
        break
      case "maintainers":
        paramsStore.DeleteMaintainer(index)
        break
      case "locations":
        paramsStore.DeleteLocation(index)
        break
      case "devicetypes":
        paramsStore.DeleteDeviceType(index)
        break
      case "operatingsystems":
        paramsStore.DeleteOperatingSystem(index)
        break
      case "architectures":
        paramsStore.DeleteArchitecture(index)
        break
      case "vlans":
        paramsStore.DeleteVLAN(index)
        break
      case "users":
        paramsStore.DeleteUser(index)
        break
    }
  }

  const saveItem = (index: number) => {
    switch (route.params.param) {
      case "categories":
        paramsStore.SaveCategory(index)
        break
      case "status":
        paramsStore.SaveStatus(index)
        break
      case "maintainers":
        paramsStore.SaveMaintainer(index)
        break
      case "locations":
        paramsStore.SaveLocation(index)
        break
      case "devicetypes":
        paramsStore.SaveDeviceType(index)
        break
      case "operatingsystems":
        paramsStore.SaveOperatingSystem(index)
        break
      case "architectures":
        paramsStore.SaveArchitecture(index)
        break
      case "vlans":
        paramsStore.SaveVLAN(index)
        break
      case "users":
        paramsStore.SaveUser(index)
        break
    }
  }

  onMounted(() => {
    paramsStore.getParams()
  })
</script>

<template>
  <card :icon="getIcon" :headingTitle="getTitle" has-table class="mb-6" headerIcon="PlusBox" @header-icon-click="addItem">
    <param-grid :mode="routeParam" @save-item="saveItem" @drop-item="dropItem" @deleteItem="deleteItem" />
  </card>
</template>
