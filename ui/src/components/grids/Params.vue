<script setup lang="ts">
  import { ref, watch, computed } from 'vue'
  import clonedeep from 'lodash.clonedeep'

  import { icons } from '@/data/icons'

  import Btns from '@/components/elements/Btns.vue'
  import Btn from '@/components/elements/Btn.vue'
  import mdIcon from '@/components/elements/MDIcon.vue'
  import Control from '@/components/elements/Control.vue'

  const props = defineProps<{
    mode: string
    items: Array<any>,

  }>()

  const emit = defineEmits(['save-item', 'drop-item', 'delete-item'])

  const edits = ref<Array<any>>([])

  const hasEdits = computed(() => (edits.value.length > 0 || props.items.filter(item => item.ID === null || item.ID === 0)))

  const isEditing = (index: number) => {
    if (props.items[index].ID === null || props.items[index].ID === 0) {
      return true
    }
    return (edits.value.findIndex(item => props.items[index].ID === item.ID) !== -1)
  }

  const editParam = (index: number) => {
    if (isEditing(index)) {
      edits.value = edits.value.filter(item => props.items[index].ID !== item.ID)
    } else {
      edits.value.push(clonedeep(props.items[index]))
    }
  }

  const saveParam = (index: number) => {
    emit('save-item', index)
    editParam(index)
  }

  const dropParam = (index: number) => {
    if (props.items[index].ID === null) {
      emit('drop-item', index)
    } else {
      props.items[index] = edits.value.find(item => props.items[index].ID === item.ID)
      editParam(index)
    }
  }

  const deletes = ref(-1)
  const isDeleting = (index: number) => {
    return (deletes.value === index)
  }
  const deleteParam = (index: number) => {
    deletes.value = index
  }
  const confirmDelete = (index: number) => {
    deletes.value = -1
    emit('delete-item', index)
  }

  const dynamicClass = (item: string, param1: any, param2: any) => {
    if (item === "row") {
      const rowClass = []
      if (param1.ID === null || edits.value.findIndex(item => item.ID === param1.ID) !== -1) {
        rowClass.push('editing')
      }
      if (deletes.value !== -1 && param1.ID === props.items[deletes.value].ID) {
        rowClass.push('deleting')
      }
      // if (isActive(param1)) {
      //   rowClass.push('active')
      // }
      // if (isUnknown(param1)) {
      //   rowClass.push('invading')
      // }
      return rowClass // .join(' ')
    } if (item === "record") {
      const rowClass = []
      rowClass.push('active')
      // if (isActive(param1)) {
      //   rowClass.push('active')
      // }
      // if (isUnknown(param1)) {
      //   rowClass.push('invading')
      // }
      return rowClass // .join(' ')
    } else {
      let colClass = [item]
      colClass.push("unreserved")
      // switch (item) {
      //   case "ip":
      //     if (param1 == 0) {
      //       colClass.push("unreserved")
      //     }
      // }
      return colClass
    }
  }

  const showColumn = (column: string) => {
    const columns = new Map<string, Array<string>>([
        ["ID",              ["categories", "status", "devicetypes", "vlans", "architectures", "locations", "maintainers", "users", "operatingsystems"]],
        ["Notes",           ["categories", "status", "devicetypes", "vlans", "architectures", "locations", "maintainers", "users", "operatingsystems"]],
        ["Label",           ["categories", "status", "devicetypes", "vlans", "architectures", "locations", "maintainers", "users"]],
        ["Icon",            ["categories", "status", "devicetypes"]],
        ["IPv4Mask",        ["vlans"]],
        ["IPv6Mask",        ["vlans"]],
        ["Vendor",          ["operatingsystems"]],
        ["Family",          ["operatingsystems"]],
        ["Name",            ["operatingsystems"]],
        ["Version",         ["operatingsystems"]],
        ["IsOpenSource",    ["operatingsystems"]],
        ["IsServer",        ["operatingsystems"]],
        ["BitSpace",        ["architectures"]],
        ["IsCloud",         ["locations"]],
        ["IsInternal",      ["maintainers"]],
        ["Username",        ["users"]],
        ["Password",        ["users"]],
        ["CanAuthenticate", ["users"]],
        ["AccessLevel",     ["users"]]
    ])
    
    return columns.get(column)?.includes(props.mode)
  }

  const gridMode = computed(() => props.mode)
  const iconOptions = computed(() => {
    //const iconOptions = [{ value: "", label: "", selected: false}]
    const iconOptions = [{ value: "", label: "" }]
    iconOptions.splice(0,1)
    icons.forEach(item => {
      iconOptions.push({ value: item.name, label: item.name })
    })
    return iconOptions
  })

  watch(gridMode, data => {
    edits.value = []
    deletes.value = -1
  })
</script>

<template>

  <div class="params datagrid" :class="hasEdits ? 'edits' : ''">
    <table>
      <thead>
        <tr>
          <th v-if="showColumn('ID')" class="id">ID</th>
          <th v-if="showColumn('Label')" class="label">Label</th>
          <th v-if="showColumn('Icon')" class="icon">Icon</th>
          <th v-if="showColumn('IPv4Mask')" class="other">IPv4Mask</th>
          <th v-if="showColumn('IPv6Mask')" class="other">IPv6Mask</th>
          <th v-if="showColumn('Vendor')" class="other">Vendor</th>
          <th v-if="showColumn('Family')" class="other">Family</th>
          <th v-if="showColumn('Name')" class="other">Name</th>
          <th v-if="showColumn('Version')" class="other">Version</th>
          <th v-if="showColumn('IsOpenSource')" class="checkbox">IsOpenSource</th>
          <th v-if="showColumn('IsInternal')" class="checkbox">IsInternal</th>
          <th v-if="showColumn('IsServer')" class="checkbox">IsServer</th>
          <th v-if="showColumn('BitSpace')" class="other">BitSpace</th>
          <th v-if="showColumn('IsCloud')" class="checkbox">IsCloud</th>
          <th v-if="showColumn('Username')" class="other">Username</th>
          <th v-if="showColumn('Password')" class="other">Password</th>
          <th v-if="showColumn('CanAuthenticate')" class="checkbox">Active</th>
          <th v-if="showColumn('AccessLevel')" class="other">AccessLevel</th>
          <th v-if="showColumn('Notes')" class="notes">Notes</th>
          <th class="actions"></th>
        </tr>
      </thead>
      <tbody v-if="mode === 'categories'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="Icon" class="icon">
            <control v-if="isEditing(index)" type="select" v-model="param.Icon" :options="iconOptions" :icon="param.Icon" />
            <div v-else><mdIcon :icon="param.Icon" :size="32"/></div>
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>
        </tr>
      </tbody>
      <tbody v-if="mode === 'devicetypes'">
        <tr v-for="(param, index) in items" :key="index" :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="Icon" class="icon">
            <control v-if="isEditing(index)" type="select" v-model="param.Icon" :options="iconOptions" :icon="param.Icon" />
            <div v-else><mdIcon :icon="param.Icon" :size="32"/></div>
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>
        </tr>
      </tbody>
      <tbody v-if="mode === 'status'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="Icon" class="icon">
            <control v-if="isEditing(index)" type="select" v-model="param.Icon" :options="iconOptions" :icon="param.Icon" />
            <div v-else><mdIcon :icon="param.Icon" :size="32"/></div>
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'architectures'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="BitSpace" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.BitSpace" :disabled="param.IsLocked" />
            <div v-else>{{ param.BitSpace }}</div>
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'locations'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
              <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="IsCloud" class="checkbox">
            <control type="checkbox" v-model="param.IsCloud" :disabled="!isEditing(index) || param.IsLocked" />
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'vlans'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="IPv4other" class="other">
            {{ param.IPv4Mask }}
          </td>
          <td data-label="IPv6Mask" class="other">
            {{ param.IPv4Mask }}
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'operatingsystems'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Vendor" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.Vendor" :disabled="param.IsLocked" />
            <div v-else>{{ param.Vendor }}</div>
          </td>
          <td data-label="Family" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.Family" :disabled="param.IsLocked" />
            <div v-else>{{ param.Family }}</div>
            
          </td>
          <td data-label="Name" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.Name" :disabled="param.IsLocked" />
            <div v-else>{{ param.Name }}</div>
            
          </td>
          <td data-label="Version" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.Version" :disabled="param.IsLocked" />
            <div v-else>{{ param.Version }}</div>
            
          </td>
          <td data-label="IsOpenSource" class="checkbox">
            <control type="checkbox" v-model="param.IsOpenSource" :disabled="!isEditing(index) || param.IsLocked" />
          </td>
          <td data-label="IsServer" class="checkbox">
            <control type="checkbox" v-model="param.IsServer" :disabled="!isEditing(index) || param.IsLocked" />
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'maintainers'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" :disabled="param.IsLocked" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="IsInternal" class="checkbox">
            <control type="checkbox" v-model="param.IsInternal" :disabled="!isEditing(index) || param.IsLocked" />
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
      <tbody v-if="mode === 'users'">
        <tr v-for="(param, index) in items" :key="index"  :class="dynamicClass('row', param, null)">
          <td data-label="ID" class="id">
            <div>{{ param.ID }}</div>
          </td>
          <td data-label="Label" class="label">
            <control v-if="isEditing(index)" type="text" v-model="param.Label" />
            <div v-else>{{ param.Label }}</div>
          </td>
          <td data-label="Username" class="other">
            <control v-if="isEditing(index)" type="text" v-model="param.Username"  :disabled="!param.CanAuthenticate" />
            <div v-else>{{ param.Username }}</div>
          </td>
          <td data-label="Password" class="other">
            <control v-if="isEditing(index)" type="password" v-model="param.NewPassword" :disabled="!param.CanAuthenticate" />
            <div v-else>●●●●●●●●●●●●●●●●</div>
          </td>
          <td data-label="Active" class="checkbox">
            <control type="checkbox" v-model="param.CanAuthenticate" :disabled="!isEditing(index) || param.IsLocked" />
          </td>
          <td data-label="AccessLevel" class="other">
            <div v-if="isEditing(index)">
              Complex set of permissions goes here
            </div>
            <div v-else>
              {{ param.AccessLevel }}
            </div>
          </td>
          <td data-label="Notes" class="notes">
            <control v-if="isEditing(index)" type="text" v-model="param.Notes" />
            <div v-else>{{ param.Notes }}</div>
          </td>
          <td class="actions right">
            <div>
              <btns type="justify-center lg:justify-end" no-wrap >
                <btn v-if="isEditing(index)" color="success" icon="ContentSave" small @click.prevent="saveParam(index)" />
                <btn v-if="isEditing(index)" color="danger" icon="Eraser" small @click.prevent="dropParam(index)" />
                <btn v-if="!isEditing(index)" color="info" icon="Pencil" small @click.prevent="editParam(index)" />
                <btn v-if="!isEditing(index) && !isDeleting(index)" color="danger" icon="TrashCan" small @click.prevent="deleteParam(index)" :disabled="param.IsLocked" />
                <btn v-if="!isEditing(index) && isDeleting(index)" color="danger" icon="CheckBold" small @click.prevent="confirmDelete(index)" />
              </btns>
            </div>
          </td>

        </tr>
      </tbody>
    </table>
  </div>

</template>

<style scoped lang="scss">

  @import url("@/components/grids/grids.scss");

  .params.datagrid {
    table {
      
      thead, tbody {
        tr {
          td, th {
            &.id {
              @apply lg:w-[80px];
            }

            &.label {
              @apply lg:w-[240px];
            }

            &.icon {
              @apply lg:w-[80px] lg:text-center;
            }

            &.other {
              @apply lg:w-[240px];
            }
            &.checkbox {
              @apply lg:w-[80px] lg:text-center;
            }
            &.actions {
              @apply hidden lg:table-cell lg:w-[140px] lg:text-right text-center;
            }
            &.actions > div {
              @apply w-full lg:w-auto;
            }
          }
        }
      }
    }
    &.edits {
      table {
        
        thead, tbody {
          tr {
            td, th {
              &.id {
                @apply lg:w-[80px];
              }
  
              &.label {
                @apply lg:w-[240px];
              }
  
              &.icon {
                @apply lg:w-[240px];
              }
  
              &.other {
                @apply lg:w-[240px];
              }
              &.checkbox {
                @apply lg:w-[80px];
              }
              &.actions {
                @apply lg:w-[140px];
              }
            }
          }
        }
      }
    }
  }

  


   
</style>

