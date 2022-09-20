<script setup lang="ts"></script>

<template>
  <header class="flex items-stretch border-b border-gray-100 dark:border-gray-700">
    <p class="flex items-center py-3 grow font-bold">
      <slot name="lefticon" />
      <slot name="lefttitle" />
      <slot name="centertitle" />
      <slot name="righttitle" />
      <slot name="righticon" />
    </p>
  </header>
</template>
