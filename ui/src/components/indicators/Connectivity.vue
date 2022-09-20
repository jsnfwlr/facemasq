<script setup lang="ts">
  import { computed } from "vue"
  import { parse, format, differenceInMinutes } from "date-fns"

  import { Connection } from "@/stores/devices"

  const props = defineProps<{
    data: Array<Connection> | null
    includeDate: boolean
  }>()

  type Band = {
    state: boolean
    scans: number
    from: Date
    to?: Date
    duration?: string
    width?: number
  }

  const bar = computed(() => {
    const bands: Array<Band> = []
    if (props.data !== null && typeof props.data[0] !== "undefined") {
      let last: Band

      last = {
        state: props.data[0].State,
        scans: 1,
        from: parse(props.data[0].Time.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date()),
      }

      for (let i = 1; i < props.data.length; i++) {
        if (props.data[i].State === last.state) {
          last.scans++
          continue
        }

        last.to = parse(props.data[i].Time.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date())
        last.duration = differenceInMinutes(last.to, last.from) + " minutes"
        last.width = (100 / props.data.length) * last.scans

        bands.push(last)
        last = {
          state: props.data[i].State,
          scans: 1,
          from: parse(props.data[i].Time.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date()),
        }
      }
      last.to = parse(props.data[props.data.length - 1].Time.replace("T", " ").replace("Z", ""), "yyyy-MM-dd HH:mm:ss", new Date())
      last.duration = differenceInMinutes(last.to, last.from) + " minutes"
      last.width = (100 / props.data.length) * last.scans
      bands.push(last)
      return bands
    }
    bands.push({ state: false, scans: 1, from: new Date() })
    return bands
  })
</script>

<template>
  <div class="bar">
    <div v-for="(band, index) in bar" :key="index" :class="band.state ? 'online band' : 'offline band'" :style="' width: ' + band.width + '%'" :title="(band.state ? 'online' : 'offline') + ' for ' + band.duration + ' from ' + format(band.from, includeDate ? 'EEE HH:mm' : 'HH:mm')" />
  </div>
</template>

<style scoped lang="scss">
  .bar {
    & {
      height: 1rem;
      width: 100%;
      opacity: 0.5;
      min-width: 60px;
    }

    &:hover {
      opacity: 1;
    }

    .band {
      & {
        display: inline-block;
      }
      &.online {
        background: hsl(160, 84%, 39%);
        height: 1rem;
      }
      &.offline {
        background: hsl(345, 100%, 39%);
        height: 1rem;
      }
    }
  }
</style>
