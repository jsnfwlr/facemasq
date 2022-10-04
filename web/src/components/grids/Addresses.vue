<script setup lang="ts">
  import { ref, onMounted } from "vue"

  import { storeToRefs } from "pinia"
  import cloneDeep from "lodash.clonedeep"

  import { useDevices } from "@/stores/devices"
  import { Address } from "@/types/deviceStore"

  import Btn from "@/components/elements/Btn.vue"
  import Btns from "@/components/elements/Btns.vue"
  import Control from "@/components/elements/Control.vue"

  // import UserAvatar from '@/components/justboil/UserAvatar.vue'

  // const userStore = useUser()
  // const { settings } = storeToRefs(userStore)

  const deviceStore = useDevices()
  const { allDevices, deletingItems, editingItems, focusedItems } = storeToRefs(deviceStore)

  interface Props {
    canEdit: boolean
    items: Array<Address>
    mode?: string
    deviceIndex: number
    interfaceIndex: number
  }

  const props = withDefaults(defineProps<Props>(), {
    perPage: 10,
    mode: "Informative",
  })

  const rowClass = (rowIndex: number) => {
    if (rowIndex > 0) {
      return []
    }
    return []
  }

  const originals = ref<Array<Address>>([])

  onMounted(() => {
    originals.value = cloneDeep(props.items)
  })

  // const paramToOptions = (params: Array<any>, labelSwitch: number) => {
  //   const options = [{ value: "", label: "" }]
  //   options.splice(0, 1)
  //   params.forEach((item) => {
  //     switch (labelSwitch) {
  //       case 1:
  //         options.push({ value: item.ID, label: item.Vendor + " " + item.Family + " " + item.Name + " " + item.Version })
  //         break
  //       default:
  //         options.push({ value: item.ID, label: item.Label })
  //         break
  //     }
  //   })
  //   options.sort((a, b) => {
  //     return a.label === b.label ? 0 : a.label < b.label ? -1 : 1
  //   })
  //   return options
  // }

  const isEditing = (addressIndex: number) => {
    return props.items[addressIndex].ID === null || props.items[addressIndex].ID === 0 ? true : editingItems.value.addresses.has(props.deviceIndex + "." + props.interfaceIndex + "." + addressIndex)
  }
  const edit = (addressIndex: number) => {
    if (isEditing(addressIndex)) {
      deviceStore.$patch((state) => {
        state.editingItems.addresses.delete(props.deviceIndex + "." + props.interfaceIndex + "." + addressIndex)
      })
    } else {
      deviceStore.$patch((state) => {
        state.editingItems.addresses.set(props.deviceIndex + "." + props.interfaceIndex + "." + addressIndex, cloneDeep(props.items[addressIndex]))
      })
      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces[props.interfaceIndex].Addresses.length; i++) {
        if (i !== addressIndex && editingItems.value.addresses.has(props.deviceIndex + "." + props.interfaceIndex + "." + i)) {
          deviceStore.$patch((state) => {
            state.editingItems.addresses.delete(props.deviceIndex + "." + props.interfaceIndex + "." + i)
          })
        }
      }
    }
  }

  const isFocused = (addressIndex: number) => {
    return focusedItems.value.interfaces.has(props.deviceIndex + "." + props.interfaceIndex + "." + addressIndex)
  }
  const focus = (addressIndex: number) => {
    if (isFocused(addressIndex)) {
      deviceStore.$patch((state) => {
        state.focusedItems.interfaces.delete(props.deviceIndex + "." + addressIndex)
      })
    } else {
      deviceStore.$patch((state) => {
        state.focusedItems.interfaces.set(props.deviceIndex + "." + addressIndex, allDevices.value[props.deviceIndex].Interfaces[addressIndex])
      })

      for (let i = 0; i < allDevices.value[props.deviceIndex].Interfaces[props.interfaceIndex].Addresses.length; i++) {
        if (i !== addressIndex && focusedItems.value.addresses.has(props.deviceIndex + "." + props.interfaceIndex + "." + i)) {
          deviceStore.$patch((state) => {
            state.focusedItems.addresses.delete(props.deviceIndex + "." + props.interfaceIndex + "." + i)
          })
        }
      }
    }
  }

  const discard = (addressIndex: number) => {
    deviceStore.Discard([props.deviceIndex, props.interfaceIndex, addressIndex])
  }

  const save = (addressIndex: number) => {
    deviceStore.Save([props.deviceIndex, props.interfaceIndex, addressIndex])
  }

  const isDeleting = (addressIndex: number) => {
    return deletingItems.value.addresses.includes(addressIndex)
  }
  const performDelete = (addressIndex: number) => {
    deviceStore.PerformDelete([props.deviceIndex, props.interfaceIndex, addressIndex])
  }
  const initiateDelete = (addressIndex: number) => {
    deviceStore.InitiateDelete([props.deviceIndex, props.interfaceIndex, addressIndex])
  }
  const cancelDelete = (addressIndex: number) => {
    deviceStore.CancelDelete([props.deviceIndex, props.interfaceIndex, addressIndex])
  }
</script>

<template>
  <div class="datagrid">
    <table>
      <thead>
        <tr>
          <th class="id">ID</th>
          <th class="label">Label</th>
          <th class="other">IPv4</th>
          <th class="other">IPv6</th>
          <th class="checkbox">Primary</th>
          <th class="checkbox">Virtual</th>
          <th class="checkbox">Reserved</th>
          <th class="actions" />
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in items" :key="index" :class="rowClass(index)">
          <td data-label="ID" class="id">
            <div @click="focus(index)">
              {{ row.ID }}
            </div>
          </td>
          <td data-label="Label" class="label">
            <div>
              <div v-if="isEditing(index)">
                <control v-model="row.Label" icon="Tag" />
              </div>
              <div v-else @click="focus(index)">
                {{ row.Label }}
              </div>
            </div>
          </td>
          <td data-label="IPv4" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.IPv4" :disabled="originals[index].IPv4 !== null" />
            </div>
            <div v-else @click="focus(index)">
              {{ row.IPv4 }}
            </div>
          </td>
          <td data-label="IPv6" class="other">
            <div v-if="isEditing(index)">
              <control v-model="row.IPv6" :disabled="originals[index].IPv6 !== null" />
            </div>
            <div v-else @click="focus(index)">
              {{ row.IPv6 }}
            </div>
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
          <td data-label="Reserved" class="checkbox">
            <div v-if="isEditing(index)">
              <control v-model="row.IsReserved" type="checkbox" />
            </div>
            <div v-else @click="focus(index)">
              <control v-model="row.IsReserved" type="checkbox" :disabled="!isEditing(index)" />
            </div>
          </td>
          <td class="actions right">
            <span>
              <btns type="justify-start lg:justify-end" no-wrap>
                <btn v-if="!isDeleting(index) && !isEditing(index)" color="info" icon="Pencil" small @click="edit(index)" :disabled="!canEdit" />

                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click="save(index)" :disabled="!canEdit" />
                <btn v-if="isEditing(index)" color="warning" icon="Eraser" small @click="discard(index)" :disabled="!canEdit" />

                <btn v-if="!isDeleting(index) && !isEditing(index)" color="danger" icon="TrashCan" small @click="initiateDelete(index)" :disabled="!canEdit" />
                <btn v-if="isDeleting(index) && !isEditing(index)" color="danger" icon="Check" small @click="performDelete(index)" :disabled="!canEdit" />
                <btn v-if="isDeleting(index) && !isEditing(index)" color="info" icon="Close" small @click="cancelDelete(index)" :disabled="!canEdit" />
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
