<script setup lang="ts">
import { computed, useSlots, watch } from 'vue'

const slots = useSlots()

interface Props {
    mobile: boolean;
    type?: string;
    fill: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  mobile: false,
  fill: false,
  type: 'justify-between'
})

const parentMobileClass = ['flex']
const parentStdClass = ['block', 'md:flex']

const parentClass = computed (() => {
    const parentBaseClasses = [props.type, 'items-center', 'parent']
    return (props.mobile) ? parentBaseClasses.concat(parentMobileClass) : parentBaseClasses.concat(parentStdClass)
})
const childStdClasses = ['flex', 'items-center', 'justify-center', 'child'].concat(props.fill ? ['shrink-1', 'grow-1', 'w-full'] : ['shrink-0', 'grow-0'])

const compStdClasses = props.fill ? ['w-full'] : ['']

const childClass = (allSlots: any, index: number) => {
    if (allSlots.default().length === 4) {
        console.log(allSlots.default(), allSlots)
    }
    return (!props.mobile && allSlots.default().length > index + 1) ? childStdClasses.concat(['mb-6', 'md:mb-0']) : childStdClasses
}

const makeSlots = (allSlots: any) => {
    return allSlots.default().map((element: any) => {
        return element
    })
}

</script>


<template>
    <div :class="parentClass">
        <div v-for="(element,index) in makeSlots($slots)" :class="childClass($slots, element)">
            <component :is="element" :class="compStdClasses"/>
        </div>
    </div>
</template>
