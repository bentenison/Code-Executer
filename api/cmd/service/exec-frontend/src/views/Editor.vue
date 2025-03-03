<template>
  <div
    class="flex flex-column align-items-center justify-content-center mt-5"
    style="min-height: 100vh"
  >
    <!-- demo -->
    <div class="flex justify-content-between">
      <div class="flex flex-column mt-2">
        <CodeEditor
          v-if="currQuestion && editorStore.langArr"
          :line-nums="true"
          :key="currQuestion.id"
          :theme="theme"
          :value="formattedCode"
          width="120rem"
          height="500px"
          lang-list-height="300px"
          :fontSize="fontSize"
          @content="getContent"
          @lang="getLanguage"
          :languages="editorStore.langArr"
        ></CodeEditor>
        <terminal :key="$route.name" :isSingle="true" />
        <div class="flex align-items-center justify-content-end p-0 m-0">
          <p class="w-full">
            code executed by <strong>{{ executedBy }}</strong>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Terminal from "../components/Terminal.vue";
import CodeEditor from "../SimpleCodeEditor/CodeEditor.vue";
import { useEditorStore } from "../stores/editor";
import { mapState } from "pinia";
import { useMainStore } from "../stores/main";
export default {
  name: "Home",
  components: {
    CodeEditor,
    Terminal,
    // Button
  },
  data() {
    return {
      fontSize: "17px",
      currQuestion: null,
      currIndex: 0,
      content: null,
      lang: [],
      langArr: [],
      formattedCode: "",
      executedBy: null,
      langId: null,
      runBtnLoading: false,
      theme: "github-dark",
      demo: "// Please edit it",
      demoLanguages1: `<CodeEditor :languages="[['cpp', 'C++']]" />`,
      demoLanguages2: `<CodeEditor :languages="[['cpp', 'C++'],['python', 'Python'],['php', 'PHP']]" />`,
      isMenuDisplayed: false,
      code: '<CodeEditor v-model="value"></CodeEditor>',
      animationCode: "",
      prog: `from typing import List, Dict, Any
  from datetime import datetime
  
  class TestCase:
      def __init__(self, input: Any, expected_output: Any):
          self.input = input
          self.expected_output = expected_output`,
      editorStore: useEditorStore(),
      mainStore: useMainStore(),
    };
  },
  computed: {
    ...mapState(useEditorStore, ["questions"]),
    themeDemo() {
      return `<CodeEditor theme="${this.theme}"></CodeEditor>`;
    },
    // codeTemplate() {
    //   console.log("currQuest", this.currQuestion.user_logic_template.code);
    //   return this.currQuestion.user_logic_template.code;
    // },
  },
  methods: {
    switchTheme(theme) {
      this.theme = theme;
      this.emitter.emit("ThemeChanged", theme);
    },
    getLanguage(lang) {
      console.log("The current language is: " + lang);
    },
    getContent(content) {
      // console.log("The content is: " + content);
      this.editorStore.encode(content).then((res) => {
        this.content = res;
      });
    },
    getTextarea(node) {
      console.log("The textarea is: " + node);
    },
    toggleMenu() {
      this.isMenuDisplayed = this.isMenuDisplayed ? false : true;
    },
    changeFontSize(e) {
      this.fontSize = e + "px";
    },
    handleCodeFormat() {
      this.mainStore.togglePageBlock();
      let payload = {
        language: this.currQuestion.language,
        code: this.content,
      };
      // console.log("langIDDDDDDDDDDD",this.langId);
      this.editorStore
        .formatCode(payload)
        .then((res) => {
          //   this.executedBy = res.data.containerID;
          //   this.emitter.emit("showMessage", res.data.output);
          //   this.runBtnLoading = false;
          this.formattedCode = res.data.formatted_code;
          this.mainStore.togglePageBlock();
        })
        .catch((err) => {
          this.mainStore.togglePageBlock();
          //   this.runBtnLoading = false;
          this.$toast.add({
            severity: "error",
            summary: "service is down! contact administrator.",
            detail: err,
            life: 3000,
          });
        });
    },
    handlecodeRun() {
      // this.editorStore.encode();
      this.mainStore.togglePageBlock();
      this.runBtnLoading = true;
      // var payload = {};
      let payload = {
        language_code: this.langId.id,
        language: this.currQuestion.language,
        code_snippet: this.content,
        question_id: this.currQuestion.id,
        user_id: "51fc3552-45e0-4982-9adb-50d8cc46c46d",
        file_extension: this.langId.file_extension,
      };
      // console.log("langIDDDDDDDDDDD",this.langId);
      this.editorStore
        .runCode(payload)
        .then((res) => {
          this.executedBy = res.data.containerID;
          this.emitter.emit("showMessage", res.data.output);
          this.runBtnLoading = false;
          this.mainStore.togglePageBlock();
        })
        .catch((err) => {
          this.mainStore.togglePageBlock();
          this.runBtnLoading = false;
          this.$toast.add({
            severity: "error",
            summary: "service is down! contact administrator.",
            detail: err,
            life: 3000,
          });
        });
    },
    setupEditor() {
      this.editorStore
        .getAllQuestions()
        .then((res) => {
          if (this.questions.length > 0) {
            this.currQuestion = this.questions[this.currIndex];
            this.editorStore
              .getLanguageID(this.currQuestion.language.toLowerCase())
              .then((res) => {
                console.log("results>>>>>", res);
                this.langId = res;
                this.lang = [];
                this.lang.push(this.currQuestion.language);
                this.lang.push(this.currQuestion.language_code.toUpperCase());
                this.editorStore.changeLanguage(this.lang);
                // this.emitter.emit("changeLang", this.lang);
                // this.emitter.emit("changeLang", this.editorStore.langArr);
                console.log(";angArray>>>>>", this.editorStore.langArr);
              });

            this.$toast.add({
              severity: "success",
              summary: "fetched all question",
              detail: "",
              life: 3000,
            });
          }
        })
        .catch((err) => {
          this.$toast.add({
            severity: "error",
            summary: "broker service is down! contact administrator.",
            detail: "",
            life: 3000,
          });
        });
    },
    toggleLineNums(e) {
      console.log("LineNums", e);
    },
  },
  mounted() {
    let code = "<CodeEdit";
    let index = 9;
    const timer = setInterval(() => {
      code += this.code[index];
      this.animationCode = code;
      index++;
      if (code == this.code) {
        clearInterval(timer);
      }
    }, 100);
    this.emitter.on("ThemeChangeSettings", this.switchTheme);
    this.emitter.on("increaseFont", this.changeFontSize);
    this.emitter.on("decreaseFont", this.changeFontSize);
    this.emitter.on("lineNums", this.toggleLineNums);
    this.emitter.on("formatCode", this.handleCodeFormat);
    this.setupEditor();
  },
};
</script>

<style lang="scss">
@font-face {
  font-family: "Quicksand";
  src:
    url("../assets/font/Quicksand-Regular.woff2") format("woff2"),
    url("../assets/font/Quicksand-Regular.woff") format("woff");
  font-weight: normal;
  font-style: normal;
  font-display: swap;
}
.p-tabpanel.p-tabpanel-active {
  min-height: 44.5rem !important;
  // border-radius: 10px !important;
}
// .tab-height{
//   min-height: 100% !important;
// }
</style>
