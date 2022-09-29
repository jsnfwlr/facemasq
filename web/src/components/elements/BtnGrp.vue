<script setup lang="ts"></script>
<template>
  <div class="button-group">
    <slot />
  </div>
</template>
<style lang="scss">
  .button-group {
    & > button:first-child {
      border-top-right-radius: 0;
      border-bottom-right-radius: 0;
    }
    & > button:not(:last-child):not(:first-child) {
      border-radius: 0;
    }
    & > button:last-child {
      border-top-left-radius: 0;
      border-bottom-left-radius: 0;
    }
  }
</style>
