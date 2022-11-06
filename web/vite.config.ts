import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import { resolve } from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: "../dist/web/",
    rollupOptions: {
      output: {
        manualChunks: {
          datefns: ["date-fns"],
          mdi: ["@mdi/js"],
        },
      },
    },
  },
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
      "!": resolve(__dirname),
      "^": resolve(__dirname, "src/stores"),
    },
  },
})
