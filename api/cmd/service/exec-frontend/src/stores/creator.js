// stores/editor.js
import axios from "axios";
import { defineStore } from "pinia";

export const useCreatorStore = defineStore("creator", {
  state: () => ({
    questions: [],
    qcQuestions: null,
    prompt: null,
    filteredQuestion: [],
    concepts: [],
  }),

  getters: {},

  actions: {
    addQuestions(qts) {
      //   console.log("recieved the qts");
      this.questions = qts;
      return new Promise((resolve, reject) => {
        axios.defaults.baseURL = "/creatorapi";
        axios
          .post("/creator/addquestions", qts)
          .then((res) => {
            // console.log("result is ", res);
            resolve(res.data);
          })
          .catch((err) => {
            axios.defaults.baseURL = "/server";
            reject(err);
          });
      });
    },
    getAllQuestions() {
      //   console.log("recieved the qts");
      // this.questions = qts;
      return new Promise((resolve, reject) => {
        axios.defaults.baseURL = "/creatorapi";
        axios
          .post("/creator/getallquestions")
          .then((res) => {
            // console.log("result is from creator ", res);
            this.filteredQuestion = res.data;
            resolve(res.data);
          })
          .catch((err) => {
            axios.defaults.baseURL = "/server";
            reject(err);
          });
      });
    },
    getAllConcepts() {
      //   console.log("recieved the qts");
      // this.questions = qts;
      return new Promise((resolve, reject) => {
        axios.defaults.baseURL = "/creatorapi";
        axios
          .get("/creator/languageConcepts")
          .then((res) => {
            // console.log("result is from creator ", res);
            this.concepts = res.data;
            resolve(res.data);
          })
          .catch((err) => {
            axios.defaults.baseURL = "/server";
            reject(err);
          });
      });
    },
    queryWithFilters(filters) {
      return new Promise((resolve, reject) => {
        axios.defaults.baseURL = "/creatorapi";
        let objectLength = Object.keys(filters).length;
        // console.log(objectLength);
        let query = "?";
        let count = 0;
        for (const [key, value] of Object.entries(filters)) {
          if (count !== objectLength - 1) {
            query = query + key + "=" + value + "&&";
            count++;
          } else {
            query = query + key + "=" + value;
          }
          // console.log(key, value);
        }
        axios
          .get(`/creator/query${query}`)
          .then((res) => {
            resolve(res.data);
          })
          .catch((err) => {
            axios.defaults.baseURL = "/server";
            reject(err);
          });
      });
    },

    // setLanguage(newLanguage) {
    //   this.language = newLanguage;
    // },

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
    // getAllQuestions() {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post("/broker/getllquestions")
    //       .then((res) => {
    //         resolve(res);
    //         console.log("results:::::", res);
    //         this.questions = res.data;
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // getAllAnswers() {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post("/broker/getallanswer")
    //       .then((res) => {
    //         resolve(res);
    //         // console.log("results:::::", res);
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // getAnswerByID(id) {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post(`/broker/getanswer/${id}`)
    //       .then((res) => {
    //         resolve(res);
    //         // console.log("results:::::", res);
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // getQuestionByID(id) {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post(`/broker/getquestion/${id}`)
    //       .then((res) => {
    //         resolve(res);
    //         // console.log("results:::::", res);
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // runCode(data) {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post("/broker/run", data)
    //       .then((res) => {
    //         resolve(res);
    //         console.log("results:::::", res);
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // addSubmission(data) {
    //   return new Promise((resolve, reject) => {
    //     axios
    //       .post(`/broker/submission`, data)
    //       .then((res) => {
    //         resolve(res);
    //         console.log("results:::::", res);
    //       })
    //       .catch((err) => {
    //         reject(err);
    //       });
    //   });
    // },
    // encode(code) {
    //   return new Promise((resolve, reject) => {
    //     let res = btoa(code);
    //     resolve(res);
    //   });
    // },
  },
  persist: {
    storage: sessionStorage, // data in sessionStorage is cleared when the page session ends.
  },
});
