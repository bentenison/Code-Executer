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
    questionTemplates: null,
    languages: null,
    langArr: [],
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
            // console.log("results:::::", res);
            this.questions = res.data;
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getQuestTemplates() {
      return new Promise((resolve, reject) => {
        axios
          .get("/broker/gettemplates")
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
            this.questionTemplates = res.data;
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getAllLanguages() {
      return new Promise((resolve, reject) => {
        axios
          .get("/broker/getlanguages")
          .then((res) => {
            // console.log("results:::::", res);
            this.languages = res.data;
            resolve(res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    getLanguageID(lang) {
      console.log("language", lang);
      return new Promise((resolve, reject) => {
        var result = this.languages.find((obj) => {
          return obj.name.toLowerCase() === lang;
          // console.log("language", obj);
        });
        resolve(result);
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
    formatCode(data) {
      return new Promise((resolve, reject) => {
        axios
          .post("/broker/formatCode", data)
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    qcQuestion(data) {
      return new Promise((resolve, reject) => {
        axios
          .post("/broker/qcquestion", data)
          .then((res) => {
            resolve(res);
            // console.log("results:::::", res);
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
    changeLanguage(lang) {
      this.langArr = [];
      this.langArr.push(lang);
    },
  },
  persist: {
    storage: sessionStorage, // data in sessionStorage is cleared when the page session ends.
  },
});

// // Mock API function
// async function fakeApiRunCode(code, language) {
//   // Replace with actual API call
//   return new Promise((resolve) => {
//     setTimeout(() => {
//       resolve({ output: `Output for ${language}: ${code}` });
//     }, 1000);
//   });
// }
