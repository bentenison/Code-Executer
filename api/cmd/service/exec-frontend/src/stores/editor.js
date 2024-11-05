// stores/editor.js
import axios from "axios";
import { defineStore } from "pinia";

export const useEditorStore = defineStore("editor", {
  state: () => ({
    code: "",
    output: "",
    language: "javascript",
    isRunning: false,
    questions: null,
  }),

  getters: {
    getCode: (state) => state.code,
    getOutput: (state) => state.output,
    getLanguage: (state) => state.language,
    isRunning: (state) => state.isRunning,
  },

  actions: {
    updateCode(newCode) {
      this.code = newCode;
    },

    setLanguage(newLanguage) {
      this.language = newLanguage;
    },

    async runCode() {
      // this.isRunning = true;
      // try {
      //   const response = await fakeApiRunCode(this.code, this.language);
      //   this.output = response.output;
      // } catch (error) {
      //   this.output = error.message;
      // } finally {
      //   this.isRunning = false;
      // }
    },

    clearOutput() {
      this.output = "";
    },
    getAllQuestions() {
      return new Promise((resolve, reject) => {
        axios
          .post("/broker/getllquestions")
          .then((res) => {
            resolve(res);
            console.log("results:::::", res);
            this.questions = res.data;
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getAllAnswers() {
      return new Promise((resolve, reject) => {
        axios
          .post("/broker/getallanswer")
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getAnswerByID(id) {
      return new Promise((resolve, reject) => {
        axios
          .post(`/broker/getanswer/${id}`)
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getQuestionByID(id) {
      return new Promise((resolve, reject) => {
        axios
          .post(`/broker/getquestion/${id}`)
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    runCode(data) {
      return new Promise((resolve, reject) => {
        axios
          .post("/broker/run", data)
          .then((res) => {
            resolve(res);
            console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    addSubmission(data) {
      return new Promise((resolve, reject) => {
        axios
          .post(`/broker/submission`, data)
          .then((res) => {
            resolve(res);
            console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    encode(code) {
      return new Promise((resolve, reject) => {
        let res = btoa(code);
        resolve(res);
      });
    },
  },
});

// Mock API function
async function fakeApiRunCode(code, language) {
  // Replace with actual API call
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({ output: `Output for ${language}: ${code}` });
    }, 1000);
  });
}
