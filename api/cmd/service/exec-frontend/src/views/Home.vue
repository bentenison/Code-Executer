<template>
  <div :class="{ 'menu-on': isMenuDisplayed }">
    <!-- <transition name="fade">
      <div class="menu" v-show="isMenuDisplayed">
        <div class="menu-title">
          <div>Theme</div>
          <svg
            @click="toggleMenu"
            xmlns="http://www.w3.org/2000/svg"
            width="22"
            height="22"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="currentColor"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
            <path d="M18 6l-12 12"></path>
            <path d="M6 6l12 12"></path>
          </svg>
        </div>
        <div class="menu-body">
         
        </div>
      </div>
    </transition> -->
    <!-- demo -->
    <div class="flex justify-content-between">
      <!-- <CodeEditor
          :autofocus="true"
          :theme="theme"
          v-model="animationCode"
          width="100%"
          :languages="[
            ['html', 'HTML'],
            ['javascript', 'JS'],
            ['scss', 'SCSS'],
          ]"
        ></CodeEditor> -->
      <div class="flex flex-column mt-2">
        <CodeEditor
          v-if="currQuestion"
          :line-nums="true"
          :key="currQuestion.id"
          :theme="theme"
          :value="currQuestion.user_logic_template.code"
          width="80rem"
          height="500px"
          lang-list-height="300px"
          :fontSize="fontSize"
          @content="getContent"
          @lang="getLanguage"
          :languages="editorStore.langArr"
        ></CodeEditor>
        <terminal />
        <div class="flex align-items-center justify-content-end">
          <p class="w-full p-0 m-0">
            code executed by <strong>{{ executedBy }}</strong>
          </p>
          <!-- <div class="flex gap-5">
            <Button
              type="button"
              label="Previous"
              icon="pi pi-arrow-circle-left text-lg"
              severity="secondary"
              raised
              class="text-lg px-2"
              :loading="loading"
              @click="previous"
            />
            <Button
              type="button"
              label="Next "
              icon="pi pi-arrow-circle-right text-lg"
              severity="secondary"
              raised
              class="text-lg px-2"
              :loading="loading"
              @click="next"
            />
          </div> -->
        </div>
      </div>
      <div class="w-60rem pl-4">
        <!-- <div class="w-30rem flex"> -->
        <div
          class="w-60rem gap-5 mb-2 flex align-items-center justify-content-end"
        >
          <Button
            v-for="(dt, index) in challengeStore.challengeQuestions"
            :key="index"
            raised
            rounded
            outlined
            @click="gotoQuest(index)"
            >{{ index + 1 }}</Button
          >
          <Button
            type="button"
            label="Run"
            class="text-lg font-medium"
            icon="pi pi-play"
            :loading="runBtnLoading"
            @click="handlecodeRun"
            severity="success"
            outlined
          />
          <Button
            type="button"
            class="text-lg font-medium"
            label="Submit"
            icon="pi pi-arrow-circle-right"
            :loading="loading"
            @click="load"
            severity="info"
            outlined
          />
        </div>
        <!-- </div> -->
        <Tabs
          value="0"
          class="w-full"
          style="margin-top: 1rem"
          v-if="currQuestion"
        >
          <TabList>
            <Tab value="0" as="div" class="flex align-items-center gap-2">
              <i class="pi pi-info-circle" style="font-size: 1.5rem"></i>
              <span class="font-bold whitespace-nowrap">Instructions</span>
            </Tab>
            <Tab value="1" as="div" class="flex align-items-center gap-2">
              <i class="pi pi-clipboard" style="font-size: 1.5rem"></i>
              <span class="font-bold whitespace-nowrap">Notes</span>
            </Tab>
            <Tab v-slot="slotProps" value="2" asChild>
              <div
                :class="['flex align-items-center gap-2', slotProps.class]"
                @click="slotProps.onClick"
                v-bind="slotProps.a11yAttrs"
              >
                <i class="pi pi-question-circle" style="font-size: 1.5rem"></i>
                <span class="font-bold whitespace-nowrap">IDE help</span>
                <Badge value="2" />
              </div>
            </Tab>
          </TabList>
          <TabPanels>
            <TabPanel value="0" as="div" class="m-0 tab-height">
              <div class="flex flex-column gap-2">
                <div
                  class="logotext flex gap-2 align-items-center justify-content-center"
                >
                  <svg
                    width="40px"
                    height="40px"
                    viewBox="0 0 40 40"
                    version="1.1"
                  >
                    <g id="surface1">
                      <path
                        style="
                          stroke: none;
                          fill-rule: nonzero;
                          fill: var(--p-primary-color);
                          fill-opacity: 1;
                        "
                        d="M 6.824219 -0.109375 C 7.785156 -0.121094 7.785156 -0.121094 8.765625 -0.136719 C 9.460938 -0.140625 10.15625 -0.144531 10.871094 -0.148438 C 11.582031 -0.152344 12.296875 -0.160156 13.03125 -0.164062 C 14.539062 -0.171875 16.046875 -0.179688 17.558594 -0.183594 C 19.867188 -0.195312 22.175781 -0.222656 24.484375 -0.25 C 25.949219 -0.257812 27.414062 -0.261719 28.878906 -0.265625 C 29.570312 -0.277344 30.261719 -0.289062 30.972656 -0.300781 C 35.773438 -0.285156 35.773438 -0.285156 38.234375 1.1875 C 39.949219 3.160156 40.082031 4.261719 40.109375 6.824219 C 40.121094 7.785156 40.121094 7.785156 40.136719 8.765625 C 40.140625 9.460938 40.144531 10.15625 40.148438 10.871094 C 40.152344 11.582031 40.160156 12.296875 40.164062 13.03125 C 40.171875 14.539062 40.179688 16.046875 40.183594 17.558594 C 40.195312 19.867188 40.222656 22.175781 40.25 24.484375 C 40.257812 25.949219 40.261719 27.414062 40.265625 28.878906 C 40.277344 29.570312 40.289062 30.261719 40.300781 30.972656 C 40.285156 35.773438 40.285156 35.773438 38.8125 38.234375 C 36.839844 39.949219 35.738281 40.082031 33.175781 40.109375 C 32.214844 40.121094 32.214844 40.121094 31.234375 40.136719 C 30.191406 40.140625 30.191406 40.140625 29.128906 40.148438 C 28.417969 40.152344 27.703125 40.160156 26.96875 40.164062 C 25.460938 40.171875 23.953125 40.179688 22.441406 40.183594 C 20.132812 40.195312 17.824219 40.222656 15.515625 40.25 C 14.050781 40.257812 12.585938 40.261719 11.121094 40.265625 C 10.085938 40.28125 10.085938 40.28125 9.027344 40.300781 C 4.226562 40.285156 4.226562 40.285156 1.765625 38.8125 C 0.0507812 36.839844 -0.0820312 35.738281 -0.109375 33.175781 C -0.117188 32.535156 -0.125 31.894531 -0.136719 31.234375 C -0.140625 30.539062 -0.144531 29.84375 -0.148438 29.128906 C -0.152344 28.417969 -0.160156 27.703125 -0.164062 26.96875 C -0.171875 25.460938 -0.179688 23.953125 -0.183594 22.441406 C -0.195312 20.132812 -0.222656 17.824219 -0.25 15.515625 C -0.257812 14.050781 -0.261719 12.585938 -0.265625 11.121094 C -0.28125 10.085938 -0.28125 10.085938 -0.300781 9.027344 C -0.285156 4.226562 -0.285156 4.226562 1.1875 1.765625 C 3.160156 0.0507812 4.261719 -0.0820312 6.824219 -0.109375 Z M 15.625 5 C 16.207031 5.691406 16.792969 6.378906 17.375 7.066406 C 17.699219 7.449219 18.023438 7.832031 18.359375 8.226562 C 19.359375 9.382812 19.359375 9.382812 20.449219 10.425781 C 21.554688 11.515625 22.066406 12.25 22.5 13.75 C 21.964844 16.664062 20.660156 19.230469 19.375 21.875 C 18.882812 22.914062 18.394531 23.957031 17.902344 24.996094 C 17.140625 26.613281 16.375 28.226562 15.59375 29.828125 C 15.355469 30.320312 15.117188 30.8125 14.871094 31.316406 C 14.546875 31.976562 14.546875 31.976562 14.214844 32.648438 C 13.667969 33.742188 13.667969 33.742188 13.75 35 C 16.675781 35.015625 19.601562 35.027344 22.527344 35.035156 C 23.523438 35.039062 24.519531 35.042969 25.515625 35.046875 C 26.945312 35.054688 28.375 35.058594 29.800781 35.0625 C 30.25 35.0625 30.699219 35.066406 31.160156 35.070312 C 32.230469 35.070312 33.304688 35.039062 34.375 35 C 35.359375 34.015625 35.078125 33.039062 35.082031 31.675781 C 35.082031 30.816406 35.082031 30.816406 35.085938 29.9375 C 35.085938 29.3125 35.085938 28.691406 35.082031 28.046875 C 35.082031 27.410156 35.085938 26.773438 35.085938 26.121094 C 35.085938 24.773438 35.085938 23.425781 35.082031 22.078125 C 35.078125 20.007812 35.082031 17.9375 35.085938 15.867188 C 35.085938 14.5625 35.085938 13.257812 35.082031 11.953125 C 35.085938 11.328125 35.085938 10.707031 35.085938 10.0625 C 35.085938 9.488281 35.082031 8.917969 35.082031 8.324219 C 35.082031 7.816406 35.078125 7.308594 35.078125 6.789062 C 35.128906 5.617188 35.128906 5.617188 34.375 5 C 32.988281 4.945312 31.601562 4.933594 30.210938 4.9375 C 29.796875 4.941406 29.382812 4.941406 28.953125 4.941406 C 27.621094 4.945312 26.292969 4.953125 24.960938 4.960938 C 24.058594 4.964844 23.160156 4.96875 22.257812 4.96875 C 20.046875 4.976562 17.835938 4.988281 15.625 5 Z M 6.875 8.125 C 6.460938 8.539062 6.050781 8.949219 5.625 9.375 C 5.824219 10.691406 5.824219 10.691406 6.25 11.875 C 7.75 12.300781 7.75 12.300781 9.375 12.5 C 9.789062 12.085938 10.199219 11.675781 10.625 11.25 C 10.425781 9.933594 10.425781 9.933594 10 8.75 C 8.5 8.324219 8.5 8.324219 6.875 8.125 Z M 6.875 8.125 "
                      />
                      <path
                        style="
                          stroke: none;
                          fill-rule: nonzero;
                          fill: var(--p-primary-color);
                          fill-opacity: 1;
                        "
                        d="M 25.898438 26.757812 C 26.390625 26.75 26.886719 26.742188 27.394531 26.734375 C 28.679688 26.730469 29.964844 26.796875 31.25 26.875 C 31.664062 27.289062 32.074219 27.699219 32.5 28.125 C 32.34375 29.726562 32.34375 29.726562 31.875 31.25 C 29.980469 32.199219 28.015625 31.992188 25.9375 31.992188 C 25.507812 32 25.074219 32.007812 24.632812 32.015625 C 23.503906 32.019531 22.375 31.949219 21.25 31.875 C 20 30.625 20 30.625 19.882812 29.023438 C 20.140625 25.691406 23.246094 26.761719 25.898438 26.757812 Z M 25.898438 26.757812 "
                      />
                    </g>
                  </svg>
                  <div class="logoline flex flex-column">
                    <h2 class="p-0 m-0 text-xxl">EPIC</h2>
                    <h5 class="p-0 m-0 text-color">Assessment</h5>
                  </div>
                </div>
                <div class="headline mt-2">
                  <h2 class="text-lg p-0 m-0">
                    {{ currQuestion.title }}
                  </h2>

                  <p class="line-height-2">
                    {{ currQuestion.description }}
                  </p>
                </div>
                <div class="instruction">
                  <p class="text-color font-bold">
                    Simplified Challenge Instructions
                  </p>
                  <p class="line-height-2 flex gap-3">
                    <span class="font-semibold">Input: </span>
                    {{ currQuestion.input.description }} <br />
                  </p>
                  <p class="line-height-2 flex gap-3">
                    <span class="font-semibold">Expected: </span>
                    {{ currQuestion.input.expected }}
                  </p>
                  <p class="line-height-2 flex gap-3">
                    <span class="font-semibold">Output: </span>
                    {{ currQuestion.output.description }} <br />
                  </p>
                </div>
                <div class="lists mt-0">
                  <ul class="flex flex-column gap-3">
                    <li class="line-height-3">
                      Run the code once to see how the error messages work.
                      Click the RUN TESTS button to see the output.
                    </li>
                    <li class="line-height-3">
                      <p>Following are the testcases for the assessment</p>
                      <span
                        v-for="test in currQuestion.testcases"
                        :key="test.id"
                        class="flex flex-column"
                        >{{ test.input }} : {{ test.expectedOutput }}</span
                      >
                    </li>
                    <li class="line-height-3">
                      After your developer is confident in their solution, they
                      use the SUBMIT SOLUTION button above to run the code
                      against all the tests and produce a score. Now you can
                      click REVIEW ASSESSMENT to complete the assessment and
                      leave any comments.
                    </li>
                  </ul>
                </div>
                <div class="note"></div>
              </div>
            </TabPanel>
            <TabPanel value="1" as="div" class="m-0 tab-height">
              <notes />
            </TabPanel>
            <TabPanel v-slot="slotProps" value="2" class="" asChild>
              <div
                v-show="slotProps.active"
                :class="slotProps.class"
                v-bind="slotProps.a11yAttrs"
                class="h-full"
              >
                <h2 class="text-center m-0 mb-2">Challange IDE</h2>
                <p class="m-0">
                  The code challenge IDE provides an in-browser IDE for editing,
                  testing, and running assessments.
                </p>
                <h2 class="text-center m-0 my-3">Quick Intro</h2>
                <hr />
                <ul class="flex flex-column gap-2">
                  <li class="text-lg">
                    Read through the provided Instructions.
                  </li>
                  <li class="text-lg">
                    Edit the provided files, and optionally add your own, to
                    solve the given problem.
                  </li>
                  <li class="text-lg">
                    Run your solution against the provided tests using
                    <strong>RUN TESTS</strong>, unless this is a web-only
                    challenge without tests.
                  </li>
                  <li class="text-lg">
                    Results will be shown in the Run Output pane.
                  </li>
                </ul>
              </div>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </div>
      <!-- <CodeEditor
          :read-only="true"
          v-model="themeDemo"
          theme="atom-one-dark"
          width="100%"
          :languages="[['html', 'HTML']]"
        ></CodeEditor> -->
    </div>
  </div>
</template>

<script>
import Terminal from "../components/Terminal.vue";
import CodeEditor from "../SimpleCodeEditor/CodeEditor.vue";
// import Button from "primevue/button"
import Tabs from "primevue/tabs";
import Tab from "primevue/tab";
import TabList from "primevue/tablist";
import TabPanel from "primevue/tabpanel";
import TabPanels from "primevue/tabpanels";
import Notes from "../components/Notes.vue";
import IDESettings from "../components/IDESettings.vue";
import IDEInstructions from "../components/IDEInstructions.vue";
import { useEditorStore } from "../stores/editor";
import { mapState } from "pinia";
import { useMainStore } from "../stores/main";
import { useChallengeStore } from "../stores/challenges";

export default {
  name: "Home",
  components: {
    CodeEditor,
    Terminal,
    Tab,
    Tabs,
    TabList,
    TabPanel,
    TabPanels,
    Notes,
    IDESettings,
    IDEInstructions,
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
      executedBy: null,
      langId: null,
      runBtnLoading: false,
      items: [
        { route: "/dashboard", label: "Dashboard", icon: "pi pi-home" },
        {
          route: "/transactions",
          label: "Transactions",
          icon: "pi pi-chart-line",
        },
        { route: "/products", label: "Products", icon: "pi pi-list" },
        { route: "/messages", label: "Messages", icon: "pi pi-inbox" },
      ],
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
      challengeStore: useChallengeStore(),
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
    handlecodeRun() {
      // this.editorStore.encode();
      this.mainStore.togglePageBlock();
      this.runBtnLoading = true;
      console.log("payload", this.currQuestion);
      // var payload = {};
      let payload = {
        language_code: this.challengeStore.selectedLanguage.id,
        language: this.currQuestion.language,
        code_snippet: this.content,
        question_id: this.currQuestion.id,
        user_id: "51fc3552-45e0-4982-9adb-50d8cc46c46d",
        file_extension: this.challengeStore.selectedLanguage.file_extension,
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
      let questIds = [];
      this.challengeStore.challenges.questions.forEach((element) => {
        questIds.push(element.question_id);
      });
      console.log("challeges Questiopns", questIds);
      this.challengeStore
        .fetchChallengeQuests(questIds)
        .then((res) => {
          // console.log("all questions", res);
          this.challengeStore.challenges.questions.forEach((element, index) => {
            // questIds.push(element.question_id);
            if (!element.is_completed) {
              this.currIndex = index;
            }
          });
          // console.log(res[this.currIndex])
          this.currQuestion = res[this.currIndex];
          // console.log("currentQuest",this.currQuestion)
          this.lang = [];
          this.lang.push(this.currQuestion.language);
          this.lang.push(this.currQuestion.language_code.toUpperCase());
          this.editorStore.changeLanguage(this.lang);
        })
        .catch((err) => {
          console.log("error", err);
          this.$toast.add({
            severity: "error",
            summary: "broker service is down! contact administrator.",
            detail: "",
            life: 3000,
          });
        });
      // this.editorStore
      //   .getAllQuestions()
      //   .then((res) => {
      //     this.currQuestion = this.questions[this.currIndex];
      //     this.editorStore
      //       .getLanguageID(this.currQuestion.language.toLowerCase())
      //       .then((res) => {
      //         // console.log("results>>>>>", res);
      //         this.langId = res;
      //         this.lang = [];
      //         this.lang.push(this.currQuestion.language);
      //         this.lang.push(this.currQuestion.language_code.toUpperCase());
      //         this.editorStore.changeLanguage(this.lang);
      //         // this.emitter.emit("changeLang", this.lang);
      //         // this.emitter.emit("changeLang", this.editorStore.langArr);
      //       });

      //     this.$toast.add({
      //       severity: "success",
      //       summary: "fetched all question",
      //       detail: "",
      //       life: 3000,
      //     });
      //   })
    },
    next() {
      this.currIndex++;
      if (
        this.currIndex ===
        this.challengeStore.challengeQuestions.length - 1
      ) {
        this.currIndex = 0;
      }
      this.currQuestion =
        this.challengeStore.challengeQuestions[this.currIndex];
      this.lang = [];
      this.lang.push(this.currQuestion.language);
      this.lang.push(this.currQuestion.language_code.toUpperCase());
      this.editorStore.changeLanguage(this.lang);
      // this.emitter.emit("changeLang", this.editorStore.langArr);
      // console.log("next", this.currIndex);
    },
    previous() {
      if (this.currIndex === 0) {
        this.currIndex = this.challengeStore.challengeQuestions.length;
      }
      this.currIndex--;
      this.currQuestion =
        this.challengeStore.challengeQuestions[this.currIndex];
      this.lang = [];
      this.lang.push(this.currQuestion.language);
      this.lang.push(this.currQuestion.language_code.toUpperCase());
      this.editorStore.changeLanguage(this.lang);
      // this.emitter.emit("changeLang", this.lang);
    },
    gotoQuest(idx) {
      this.currQuestion = this.challengeStore.challengeQuestions[idx];
      this.lang = [];
      this.lang.push(this.currQuestion.language);
      this.lang.push(this.currQuestion.language_code.toUpperCase());
      this.editorStore.changeLanguage(this.lang);
    },
    toggleLineNums(e) {
      console.log("LineNums", e);
    },
    fetchQuests() {
      console.log("challeges Questiopns", this.challengeStore.challenges);
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
    this.setupEditor();
    // this.
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
  min-height: 47.5rem !important;
  // border-radius: 10px !important;
}
// .tab-height{
//   min-height: 100% !important;
// }
</style>
