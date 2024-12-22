// stores/editor.js
import axios from "axios";
import { defineStore } from "pinia";
import { AdminAPI } from "../plugins/connect";
export const useChallengeStore = defineStore("challenges", {
  state: () => ({
    isBlocked: false,
    challenges: null,
    currentQuestion: null,
    challengeQuestions: [],
    selectedLanguage: null,
  }),

  getters: {
    // getCode: (state) => state.code,
    // getOutput: (state) => state.output,
    // getLanguage: (state) => state.language,
    // isRunning: (state) => state.isRunning,
  },

  actions: {
    createChallenges(payload) {
      return new Promise((resolve, reject) => {
        AdminAPI.post("/create-challenge", payload)
          .then((res) => {
            console.log("challenges", res.data);
            this.challenges = res.data;
            resolve(res.data);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    prepareChallenges(payload) {
      return new Promise((resolve, reject) => {
        AdminAPI.post("/prepare-challenge", payload)
          .then((res) => {
            console.log("prepared", res.data);
            // this.challenges = res.data;
            resolve(res);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    fetchChallengeQuests(payload) {
      return new Promise((resolve, reject) => {
        AdminAPI.post("/fetch-questions", payload)
          .then((res) => {
            // console.log("prepared", res.data);
            this.challengeQuestions = res.data;
            resolve(res.data);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    // togglePageBlock() {
    //   this.isBlocked = !this.isBlocked;
    // },
  },
  persist: {
    storage: sessionStorage, // data in sessionStorage is cleared when the page session ends.
  },
});
