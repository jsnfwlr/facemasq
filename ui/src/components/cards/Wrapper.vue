<script setup lang="ts"></script>

<template>
  <div class="bg-white border border-gray-100 dark:bg-gray-900 dark:border-gray-900 md:rounded">
    <slot name="header" />
    <div class="p-6">
      <slot name="body" />
    </div>
  </div>
</template>
