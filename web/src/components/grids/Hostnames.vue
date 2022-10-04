<script setup lang="ts">
  import { ref, onMounted } from "vue"

  import { storeToRefs } from "pinia"
  import cloneDeep from "lodash.clonedeep"

  import { useDevices } from "@/stores/devices"
  import { DomainName } from "@/types/deviceStore"

  import Btn from "@/components/elements/Btn.vue"
  import Btns from "@/components/elements/Btns.vue"
  import Control from "@/components/elements/Control.vue"

  // import UserAvatar from '@/components/justboil/UserAvatar.vue'

  const deviceStore = useDevices()
  const { deletingItems, editingItems } = storeToRefs(deviceStore)

  interface Props {
    canEdit?: boolean
    items?: Array<DomainName>
    mode?: string
    deviceIndex: number
    interfaceIndex: number
    addressIndex: number
  }

  const props = withDefaults(defineProps<Props>(), {
    perPage: 10,
    mode: "Informative",
    items: () => {
      return []
    },
  })

  const rowClass = (rowIndex: number) => {
    if (rowIndex > 0) {
      return []
    }
    return []
  }

  const originals = ref<Array<DomainName>>([])

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

  const isEditing = (hostnameIndex: number) => {
    return props.items[hostnameIndex].ID === null || props.items[hostnameIndex].ID === 0 ? true : editingItems.value.hostnames.has(props.deviceIndex + "." + props.interfaceIndex + "." + props.addressIndex + "." + hostnameIndex)
  }
  const edit = (hostnameIndex: number) => {
    if (isEditing(hostnameIndex)) {
      deviceStore.$patch((state) => {
        state.editingItems.hostnames.delete(props.deviceIndex + "." + props.interfaceIndex + "." + props.addressIndex + "." + hostnameIndex)
      })
    } else {
      deviceStore.$patch((state) => {
        state.editingItems.hostnames.set(props.deviceIndex + "." + props.interfaceIndex + "." + props.addressIndex + "." + hostnameIndex, cloneDeep(props.items[hostnameIndex]))
      })
    }
  }

  const isDeleting = (hostnameIndex: number) => {
    return deletingItems.value.hostnames.includes(hostnameIndex)
  }
</script>

<template>
  <div class="datagrid">
    <table>
      <thead>
        <tr>
          <th class="id">ID</th>
          <th class="label hostname">Hostname</th>
          <th class="checkbox dns">DNS</th>
          <th class="checkbox selfset">Self Set</th>
          <th class="actions" />
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in items" :key="index" :class="rowClass(index)">
          <td data-label="ID" class="id">
            <div>
              {{ row.ID }}
            </div>
          </td>
          <td data-label="Hostname" class="label hostname">
            <div>
              <div v-if="isEditing(index)">
                <control v-model="row.Hostname" icon="Tag" />
              </div>
              <div v-else>
                {{ row.Hostname }}
              </div>
            </div>
          </td>
          <td data-label="DNS" class="checkbox dns">
            <div>
              <control v-model="row.IsDNS" type="checkbox" :disabled="!isEditing(index)" />
            </div>
          </td>
          <td data-label="Self Set" class="checkbox selfset">
            <div>
              <control v-model="row.IsSelfSet" type="checkbox" :disabled="!isEditing(index)" />
            </div>
          </td>
          <td class="actions right">
            <span>
              <btns type="justify-start lg:justify-end" no-wrap>
                <btn v-if="!isDeleting(index) && !isEditing(index)" color="info" icon="Pencil" small @click="edit(index)" :disabled="!canEdit" />

                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click="deviceStore.Save([props.deviceIndex, props.interfaceIndex, props.addressIndex, index])" :disabled="!canEdit" />
                <btn v-if="isEditing(index)" color="warning" icon="Eraser" small @click="deviceStore.Discard([props.deviceIndex, props.interfaceIndex, props.addressIndex, index])" :disabled="!canEdit" />

                <btn v-if="!isDeleting(index) && !isEditing(index)" color="danger" icon="TrashCan" small @click="deviceStore.InitiateDelete([props.deviceIndex, props.interfaceIndex, props.addressIndex, index])" :disabled="!canEdit" />
                <btn v-if="isDeleting(index) && !isEditing(index)" color="danger" icon="Check" small @click="deviceStore.PerformDelete([props.deviceIndex, props.interfaceIndex, props.addressIndex, index])" :disabled="!canEdit" />
                <btn v-if="isDeleting(index) && !isEditing(index)" color="info" icon="Close" small @click="deviceStore.CancelDelete([props.deviceIndex, props.interfaceIndex, props.addressIndex, index])" :disabled="!canEdit" />
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

            &.hostname {
              width: 350px;
            }

            &.dns {
              width: 60px;
              text-align: center;
            }
            &.selfset {
              width: 70px;
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
