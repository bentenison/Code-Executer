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
      <div class="flex flex-column">
        <CodeEditor
          :line-nums="true"
          :theme="theme"
          :value="prog"
          width="80rem"
          height="450px"
          @content="getContent"
          @lang="getLanguage"
        ></CodeEditor>
        <terminal />
      </div>
      <div class="w-60rem pl-4">
        <!-- <div class="w-30rem flex"> -->
        <div class="w-40rem gap-2 flex align-items-center justify-content-end">
          <Button
            type="button"
            label="Run"
            class="text-lg font-medium"
            icon="pi pi-play"
            :loading="loading"
            @click="load"
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
        <Tabs value="0">
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
            <TabPanel value="0" as="p" class="m-0">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut
              enim ad minim veniam, quis nostrud exercitation ullamco laboris
              nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
              reprehenderit in voluptate velit esse cillum dolore eu fugiat
              nulla pariatur. Excepteur sint occaecat cupidatat non proident,
              sunt in culpa qui officia deserunt mollit anim id est laborum.
            </TabPanel>
            <TabPanel value="1" as="p" class="m-0">
              <notes/>
            </TabPanel>
            <TabPanel v-slot="slotProps" value="2" asChild>
              <div
                v-show="slotProps.active"
                :class="slotProps.class"
                v-bind="slotProps.a11yAttrs"
              >
                <p class="m-0">
                  At vero eos et accusamus et iusto odio dignissimos ducimus qui
                  blanditiis praesentium voluptatum deleniti atque corrupti quos
                  dolores et quas molestias excepturi sint occaecati cupiditate
                  non provident, similique sunt in culpa qui officia deserunt
                  mollitia animi, id est laborum et dolorum fuga. Et harum
                  quidem rerum facilis est et expedita distinctio. Nam libero
                  tempore, cum soluta nobis est eligendi optio cumque nihil
                  impedit quo minus.
                </p>
              </div>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </div>
      <ide-settings :visible="true"/>
      <!-- <CodeEditor
          :read-only="true"
          v-model="themeDemo"
          theme="atom-one-dark"
          width="100%"
          :languages="[['html', 'HTML']]"
        ></CodeEditor> -->
      <!-- <div class="button-group">
          <button
            :class="{ selected: theme == 'github-dark' }"
            @click="switchTheme('github-dark')"
          >
            github-dark
          </button>
          <button
            :class="{ selected: theme == 'github' }"
            @click="switchTheme('github')"
          >
            github
          </button>
          <button
            :class="{ selected: theme == 'gradient-dark' }"
            @click="switchTheme('gradient-dark')"
          >
            gradient-dark
          </button>
          <button
            :class="{ selected: theme == 'hybrid' }"
            @click="switchTheme('hybrid')"
          >
            hybrid
          </button>
          <button
            :class="{ selected: theme == 'isbl-editor-dark' }"
            @click="switchTheme('isbl-editor-dark')"
          >
            isbl-editor-dark
          </button>
          <button
            :class="{ selected: theme == 'vs2015' }"
            @click="switchTheme('vs2015')"
          >
            vs2015
          </button>
          <button
            :class="{ selected: theme == 'atom-one-dark' }"
            @click="switchTheme('atom-one-dark')"
          >
            atom-one-dark
          </button>
          <button style="color: var(--main-5)" @click="toggleMenu">
            All themes
          </button>
        </div> -->
    </div>
    <!-- <terminal /> -->
    <!-- <div class="footer surface-card">
      The playground system powered by
      <a target="_blank" href="https://github.com/bentenison"> Bentenison</a>
    </div> -->
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
import Notes from '../components/Notes.vue';
import IDESettings from '../components/IDESettings.vue';
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
    // Button
  },
  data() {
    return {
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
    };
  },
  computed: {
    themeDemo() {
      return `<CodeEditor theme="${this.theme}"></CodeEditor>`;
    },
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
      console.log("The content is: " + content);
    },
    getTextarea(node) {
      console.log("The textarea is: " + node);
    },
    toggleMenu() {
      this.isMenuDisplayed = this.isMenuDisplayed ? false : true;
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
    this.emitter.on("ThemeChangeSettings",this.switchTheme)
  },
};
</script>

<style lang="scss" scoped>
@font-face {
  font-family: "Quicksand";
  src: url("../assets/font/Quicksand-Regular.woff2") format("woff2"),
    url("../assets/font/Quicksand-Regular.woff") format("woff");
  font-weight: normal;
  font-style: normal;
  font-display: swap;
}
</style>
