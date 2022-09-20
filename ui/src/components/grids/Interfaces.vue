<script setup lang="ts">
  import { useDevices } from "@/stores/devices"
  import { useParams } from "@/stores/params"
  import { storeToRefs } from "pinia"

  import Btn from "@/components/elements/Btn.vue"
  import Btns from "@/components/elements/Btns.vue"
  import Control from "@/components/elements/Control.vue"

  const paramsStore = useParams()
  const { InterfaceTypes, VLANs, Statuses } = storeToRefs(paramsStore)

  const deviceStore = useDevices()
  const { allDevices, editingItems, deletingItems, focusedItems } = storeToRefs(deviceStore)

  interface Props {
    checkable?: boolean
    deviceIndex: number
    mode?: string
  }

  const props = withDefaults(defineProps<Props>(), {
    perPage: 10,
    mode: "Informative",
  })

  const iconIFace = (interfaceTypeID: number) => {
    return InterfaceTypes.value.find((item) => item.ID === interfaceTypeID)?.Icon as string
  }

  const rowClass = (interfaceIndex: number) => {
    const rowClass = []
    if (isEditing(interfaceIndex)) {
      rowClass.push("editing")
    }
    if (isDeleting(interfaceIndex)) {
      rowClass.push("deleting")
    }
    if (isFocused(interfaceIndex)) {
      rowClass.push("focused")
    }

    return rowClass.join(" ")
  }

  type Param = {
    ID: number | null
    Label?: string
    Notes?: string | null
    Vendor?: string
    Family?: string
    Version?: string
    Name?: string
    IsServer?: number | boolean
    IsCloud?: number | boolean
    IsOpenSource?: number | boolean
    Icon?: string
  }

  type SelectOption = {
    value: number
    label: string
  }
  const paramToOptions = (params: Array<Param>, labelSwitch: number) => {
    const options = [] as SelectOption[]

    params.forEach((item) => {
      switch (labelSwitch) {
        case 1:
          options.push({ value: item.ID ? item.ID : 0, label: item.Vendor + " " + item.Family + " " + item.Name + " " + item.Version })
          break
        default:
          options.push({ value: item.ID ? item.ID : 0, label: item.Label ? item.Label : "-" })
          break
      }
    })
    options.sort((a, b) => {
      return a.label === b.label ? 0 : a.label < b.label ? -1 : 1
    })
    return options
  }

  const getLabelByID = (params: Array<Param>, label: string, selectedID: number) => {
    let value = "missing"
    if (params.length > 0) {
      let selected = params.find((par) => par.ID === selectedID)
      if (selected) {
        switch (label) {
          case "Label":
            value = selected.Label ? selected.Label : "missing"
            break
          case "Vendor":
            value = selected.Vendor ? selected.Vendor : "missing"
            break
          case "Family":
            value = selected.Family ? selected.Family : "missing"
            break
          case "Version":
            value = selected.Version ? selected.Version : "missing"
            break
          case "Name":
            value = selected.Name ? selected.Name : "missing"
            break
          case "Icon":
            value = selected.Icon ? selected.Icon : "missing"
            break
        }
      }
    }
    return value
  }

  // const hasChanges = (rowIndex: number) => {
  //   return (!isequal(allDevices.value[props.deviceIndex].Interfaces[rowIndex], props.items[rowIndex]))
  // }

  const isEditing = (interfaceIndex: number) => {
    if (allDevices.value[props.deviceIndex].Interfaces[interfaceIndex].ID === null || allDevices.value[props.deviceIndex].Interfaces[interfaceIndex].ID === 0) {
      return true
    }
    return editingItems.value.interfaces.has(props.deviceIndex + "." + interfaceIndex)
  }
  const edit = (interfaceIndex: number) => {
    if (isEditing(interfaceIndex)) {
      deviceStore.$patch((state) => {
        state.editingItems.interfaces.delete(props.deviceIndex + "." + interfaceIndex)
      })
    } else {
      deviceStore.$patch((state) => {
        state.editingItems.interfaces.set(props.deviceIndex + "." + interfaceIndex, allDevices.value[props.deviceIndex].Interfaces[interfaceIndex])
      })
      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces.length; i++) {
        if (i !== interfaceIndex && editingItems.value.interfaces.has(props.deviceIndex + "." + i)) {
          deviceStore.$patch((state) => {
            state.editingItems.interfaces.delete(props.deviceIndex + "." + i)
          })
        }
      }
      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces.length; i++) {
        if (i !== interfaceIndex && focusedItems.value.interfaces.has(props.deviceIndex + "." + i)) {
          deviceStore.$patch((state) => {
            state.focusedItems.interfaces.delete(props.deviceIndex + "." + i)
          })
        }
      }
    }
  }
  const otherEdit = (interfaceIndex: number) => {
    for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces.length; i++) {
      if (i !== interfaceIndex && (editingItems.value.interfaces.has(props.deviceIndex + "." + i) || allDevices.value[props.deviceIndex].Interfaces[i].ID === null || allDevices.value[props.deviceIndex].Interfaces[i].ID === 0)) {
        return true
      }
    }
    return false
  }

  const isFocused = (interfaceIndex: number) => {
    let otherFocus = false
    if (interfaceIndex === 0 && !focusedItems.value.interfaces.has(props.deviceIndex + "." + interfaceIndex)) {
      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces.length; i++) {
        if (i !== interfaceIndex && focusedItems.value.interfaces.has(props.deviceIndex + "." + i)) {
          otherFocus = true
        }
      }
      if (!otherFocus) {
        return true
      }
    }
    return focusedItems.value.interfaces.has(props.deviceIndex + "." + interfaceIndex)
  }
  const focus = (interfaceIndex: number) => {
    if (isFocused(interfaceIndex)) {
      deviceStore.$patch((state) => {
        state.focusedItems.interfaces.delete(props.deviceIndex + "." + interfaceIndex)
      })
    } else {
      deviceStore.$patch((state) => {
        state.focusedItems.interfaces.set(props.deviceIndex + "." + interfaceIndex, allDevices.value[props.deviceIndex].Interfaces[interfaceIndex])
      })
      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces.length; i++) {
        if (i !== interfaceIndex && focusedItems.value.interfaces.has(props.deviceIndex + "." + i)) {
          deviceStore.$patch((state) => {
            state.focusedItems.interfaces.delete(props.deviceIndex + "." + i)
          })
        }
      }
    }
  }
  // const otherFocus = (interfaceIndex: number) => {
  //   for (let i=0; i<allDevices.value[props.deviceIndex].Interfaces.length; i++) {
  //     if (i !== interfaceIndex && (focusedItems.value.interfaces.has(props.deviceIndex +"."+ i) || allDevices.value[props.deviceIndex].Interfaces[i].ID === null || allDevices.value[props.deviceIndex].Interfaces[i].ID === 0)) {
  //       return true
  //     }
  //   }
  //   return false
  // }

  const isDeleting = (interfaceIndex: number) => {
    return deletingItems.value.interfaces.includes(interfaceIndex)
  }
</script>

<template>
  <div class="datagrid">
    <table>
      <thead>
        <tr>
          <th class="id">ID</th>
          <th class="label">Label</th>
          <th class="other">MAC</th>
          <th class="other">Interface Type</th>
          <th class="other">VLAN</th>
          <th class="checkbox">Primary</th>
          <th class="checkbox">Virtual</th>
          <th class="other">Status</th>
          <th class="actions" />
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in allDevices[deviceIndex].Interfaces" :key="row.ID" :class="rowClass(index)">
          <td data-label="ID" class="id">
            <div @click="focus(index)">
              {{ row.ID }}
            </div>
          </td>
          <td data-label="Label" class="label">
            <div v-if="isEditing(index)">
              <control v-model="row.Label" icon="Tag" />
            </div>
            <div v-else @click="focus(index)">
              {{ row.Label }}
            </div>
          </td>
          <td data-label="MAC" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.MAC" />
            </div>
            <div v-else @click="focus(index)">
              {{ row.MAC }}
            </div>
          </td>
          <td data-label="Interface Type" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.InterfaceTypeID" :options="paramToOptions(InterfaceTypes, 0)" :icon="iconIFace(row.InterfaceTypeID)" />
            </div>
            <div v-else @click="focus(index)">
              {{ getLabelByID(InterfaceTypes, "Label", row.InterfaceTypeID) }}
            </div>
          </td>
          <td data-label="VLAN" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.VlanID" :options="paramToOptions(VLANs, 0)" :icon="getLabelByID(VLANs, 'Icon', row.VlanID)" />
            </div>
            <div v-else @click="focus(index)">{{ getLabelByID(VLANs, "Label", row.VlanID) }}</div>
          </td>
          <td data-label="Primary" class="checkbox">
            <div v-if="isEditing(index)">
              <control v-model="row.IsPrimary" type="checkbox" />
            </div>
            <div v-else @click="focus(index)">
              <control v-model="row.IsPrimary" type="checkbox" :disabled="!isEditing(index)" />
            </div>
          </td>
          <td data-label="Virtual" class="checkbox">
            <div v-if="isEditing(index)">
              <control v-model="row.IsVirtual" type="checkbox" />
            </div>
            <div v-else @click="focus(index)">
              <control v-model="row.IsVirtual" type="checkbox" :disabled="!isEditing(index)" />
            </div>
          </td>

          <td data-label="Status" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.StatusID" :options="paramToOptions(Statuses, 0)" :icon="getLabelByID(Statuses, 'Icon', row.StatusID)" />
            </div>
            <div v-else @click="focus(index)">{{ getLabelByID(Statuses, "Label", row.StatusID) }}</div>
          </td>

          <td class="actions right">
            <span>
              <btns type="justify-start lg:justify-end" no-wrap>
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="info" icon="Pencil" small @click="edit(index)" :disabled="otherEdit(index)" />

                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click="deviceStore.Save([deviceIndex, index])" />
                <btn v-if="isEditing(index)" color="warning" icon="Eraser" small @click="deviceStore.Discard([deviceIndex, index])" />

                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click="deviceStore.InitiateDelete([deviceIndex, index])" :disabled="otherEdit(index)" />
                <btn v-if="isDeleting(index)" color="danger" icon="Check" small @click="deviceStore.PerformDelete([deviceIndex, index])" />
                <btn v-if="isDeleting(index)" color="info" icon="Close" small @click="deviceStore.CancelDelete([deviceIndex, index])" />
              </btns>
            </span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped lang="scss">
  @import url("@/components/grids/grids.scss");

  .datagrid {
    table {
      thead,
      tbody {
        tr {
          td,
          th {
            &.expand {
              width: 16px;
            }
            &.icon {
              & {
                width: 32px;
              }
              &.reserved {
                span {
                  opacity: 1;
                }
              }
              span {
                opacity: 0.5;
              }
            }
            &.name {
              width: 140px;
            }
            &.description {
              width: 230px;
            }
            &.brand {
              width: 160px;
              text-align: center;
            }
            &.os {
              width: 160px;
              text-align: center;
            }
            &.serial {
              width: 160px;
            }
            &.connectivity {
              width: 200px;
            }
            &.location {
              width: 140px;
              text-align: center;
            }
            &.maintainer {
              width: 140px;
              text-align: center;
            }
            &.mac {
              width: 160px;
            }
            &.ip {
              width: 160px;
            }

            &.firstseen {
              width: 160px;
              text-align: center;
            }
            &.lastseen {
              width: 160px;
              text-align: center;
            }
            &.actions {
              width: 160px;
            }
          }
        }
        &.invading {
          td {
            &,
            small {
              @apply text-red-800 dark:text-red-600;
            }
          }
        }
        &.active {
          td.lastseen {
            small {
              @apply text-teal-800 dark:text-teal-300;
            }
          }
        }
      }
    }
  }
</style>
