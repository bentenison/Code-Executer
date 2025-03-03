// stores/editor.js
import axios from "axios";
import { defineStore } from "pinia";

export const useMainStore = defineStore("main", {
  state: () => ({
    isBlocked: false,
    isDark: false,
  }),

  getters: {
    getCode: (state) => state.code,
    getOutput: (state) => state.output,
    getLanguage: (state) => state.language,
    isRunning: (state) => state.isRunning,
  },

  actions: {
    togglePageBlock() {
      this.isBlocked = !this.isBlocked;
    },
  },
  persist: {
    storage: sessionStorage, // data in sessionStorage is cleared when the page session ends.
  },
});
