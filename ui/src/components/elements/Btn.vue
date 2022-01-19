<script setup lang="ts">
import { computed } from 'vue'

import { getButtonColor } from '@/colors.js'

import Icon from '@/components/elements/MDIcon.vue'

interface Props {
    label: string | number | null;
    icon: string | null;
    href: string | null;
    target: string | null;
    to: string | Object | null;
    type: string | null;
    color: string;
    as: string | null;
    small: boolean;
    outline: boolean;
    active: boolean;
    disabled: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  label: null,
  icon: null,
  href: null,
  target: null,
  to: null,
  type: null,
  color: 'white',
  as: null,
  small: false,
  outline: false,
  active: false,
  disabled: false
})


const is = computed(() => {
  if (props.as) {
    return props.as
  }

  if (props.to) {
    return 'router-link'
  }

  if (props.href) {
    return 'a'
  }

  return 'button'
})

const computedType = computed(() => {
  if (is.value === 'button') {
    return props.type ?? 'button'
  }

  return null
})

const labelClass = computed(() => props.small && props.icon ? 'px-1' : 'px-2')

const componentClass = computed(() => {
  const base = [
    'inline-flex',
    'cursor-pointer',
    'justify-center',
    'items-center',
    'whitespace-nowrap',
    'focus:outline-none',
    'transition-colors',
    'focus:ring',
    'duration-150',
    'border',
    'rounded',
    props.active ? 'ring ring-black dark:ring-white' : 'ring-blue-700',
    props.small ? 'p-1' : 'p-2',
    getButtonColor(props.color, props.outline, !props.disabled)
  ]

  if (props.disabled) {
    base.push('cursor-not-allowed', props.outline ? 'opacity-50' : 'opacity-70')
  }

  return base
})
</script>

<template>
  <component :is="is" :class="componentClass" :href="href" :type="computedType" :to="to" :target="target" :disabled="disabled">
    <icon v-if="icon" :icon="icon" />
    <span v-if="label" :class="labelClass" >{{ label }}</span>
  </component>
</template>
