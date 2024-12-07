import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
// import MonacoEditorPlugin from 'vite-plugin-monaco-editor';
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // new MonacoEditorPlugin({
    //   languages: [
    //     "javascript",
    //     "python",
    //     "java",
    //     "html",
    //     "css",
    //     "json",
    //     "typescript",
    //   ],
    // }),
  ],
  resolve: {
    alias: {
      "monaco-editor": "monaco-editor/esm/vs/editor/editor.api.js",
    },
  },
  build: {
    rollupOptions: {
      external: ["monaco-editor"],
    },
  },
  server: {
    port: 8080,
    proxy: {
      "/server": {
        target: "http://localhost:8003",
        ws: true,
        secure: false,
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/server/, ""),
      },
      "/creatorapi": {
        target: "http://localhost:8005",
        ws: true,
        secure: false,
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/creatorapi/, ""),
      },
    },
  },
  productionSourceMap: false,
});
