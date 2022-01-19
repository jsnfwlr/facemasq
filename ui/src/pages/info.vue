<script setup lang="ts">
  import Card from '@/components/containers/Card.vue'
  import Logo from '@/components/elements/Logo.vue'

  const details = {
    release: {
      version: "1.0.0",
      date: "2022-04-01",
      notes: "https://github.com/jsnfwlr/facemasq/releases/tag/v1.0.0"
    },
    env: {
      port: "6135",
      target: "192.168.0.1/22",
      frequency: "1m",
      tz: "Australia/Perth"
    }
  }
</script>

<template>
  <section class="info-page">
    <card icon="FaceMask" headingTitle="System Details" class="mb-6 xl:w-1/2 xl:mx-auto">
      <div class="flex">
        <logo size="256" class="logo" />
        <div class="details-list">
          <h1>FaceMasq {{ details.release.version }}</h1>

          <h2>Version Details</h2>
          <dl>
            <dt>Release Date</dt>
            <dd>{{ details.release.date }}</dd>
          </dl>
          <dl>
            <dt>Release Notes</dt>
            <dd><a :href="details.release.notes">View on GitHub</a></dd>
          </dl>

          <h2>Server Config</h2>
          <dl>
            <dt>Port</dt>
            <dd>{{ details.env.port }}</dd>
          </dl>

          <h2>Scanner Config</h2>
          <dl>
            <dt>Frequency</dt>
            <dd>{{ details.env.frequency }}</dd>
          </dl>
          <dl>
            <dt>Scan Target</dt>
            <dd>{{ details.env.target }}</dd>
          </dl>
        </div>
      </div>
    </card>

    <card icon="ToyBrick" headingTitle="Installed Plugins" class="mb-6 xl:w-1/2 xl:mx-auto">
      <i>Plugin system not yet implemented</i>
    </card>

    <div class="grid grid-cols-1 gap-6 mb-6 xl:grid-cols-2 xl:w-1/2 xl:mx-auto">
      <card icon="CardAccountDetails" headingTitle="Attribution for included libraries" class="mb-6">
        <p>Original <a href="https://github.com/justboil/admin-one-vue-tailwind/tree/8366cf5167981851eaa56eebbbe549e6290cded2">Vue UI library created by JustBoil - GitHub</a>.</p>
        <p>&nbsp;</p>
      </card>

      <card icon="CardAccountDetails" headingTitle="Attribution for included graphics" class="mb-6">
        <p>Icons imported from <a href="https://github.com/Templarian/MaterialDesign" title="Material Desing Icons">Templarian/MaterialDesign - GitHub</a>.</p>
        <p>Logo based on Original <a href="https://www.flaticon.com/free-icons/mask" title="mask icons">Mask icons created by mangsaabguru - Flaticon</a>.</p>
      </card>
    </div>
  </section>

</template>

<style lang="scss">
.info-page {
  .logo {
    @apply ml-4 mr-10;
  }
  .details-list {
    & {
      @apply border-l border-gray-200 dark:border-gray-700 pl-6 w-full;
    }
    h1 { 
      font-size: 1.75rem;
      font-weight: 700;
      text-align: center;

    }
    h2 {
      font-size: 1.5rem;
      font-weight: 400;
      @apply mt-6;
    }    

    dl {
      & { 
        @apply flex mt-3;
      }
      dt { 
        width: 200px;
      }
      dd { 
        width: calc(100% - 200px);
      }
    }
    
  }
  p + p { 
    @apply mt-3;
  }
}
</style>