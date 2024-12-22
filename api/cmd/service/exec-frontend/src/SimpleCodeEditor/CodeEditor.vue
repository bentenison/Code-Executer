<template>
  <div
    :theme="theme"
    class="code-editor"
    :class="{
      'hide-header': !header,
      scroll: scroll,
      'read-only': readOnly,
      wrap: wrap,
    }"
    :style="{
      width: width,
      height: height,
      zIndex: zIndex,
      maxWidth: maxWidth,
      minWidth: minWidth,
      maxHeight: maxHeight,
      minHeight: minHeight,
    }"
  >
    <div class="hljs" :style="{ borderRadius: borderRadius }">
      <div
        class="header"
        :class="{ border: showLineNums }"
        v-if="header"
        :style="{ borderRadius: borderRadius + ' ' + borderRadius + ' 0 0' }"
      >
        <Dropdown
          v-if="displayLanguage"
          :width="langListWidth"
          :title="languageTitle"
          :disabled="languages.length <= 1"
          :defaultDisplay="langListDisplay"
        >
          <ul class="lang-list hljs" :style="{ height: langListHeight }">
            <li
              v-for="(lang, index) in languages"
              :key="index"
              @click="changeLang(lang)"
            >
              {{ lang[1] ? lang[1] : lang[0] }}
            </li>
          </ul>
        </Dropdown>
        <CopyCode @click="copy" v-if="copyCode"></CopyCode>
      </div>
      <div
        class="code-area"
        :style="{
          borderRadius: header
            ? '0 0 ' + borderRadius + ' ' + borderRadius
            : borderRadius,
        }"
      >
        <div
          v-if="showLineNums"
          ref="lineNums"
          class="line-nums hljs"
          :style="{
            fontSize: fontSize,
            paddingTop: header ? '10px' : padding,
            paddingBottom: padding,
            top: top + 'px',
          }"
        >
          <div>1</div>
          <div v-for="num in lineNum" :key="num">{{ num + 1 }}</div>
          <div>&nbsp;</div>
        </div>
        <textarea
          title="textarea"
          :readOnly="readOnly"
          :style="{
            fontSize: fontSize,
            padding: !header
              ? padding
              : lineNums
                ? '10px ' + padding + ' ' + padding
                : '0 ' + padding + ' ' + padding,
            marginLeft: showLineNums ? lineNumsWidth + 'px' : '0',
            width: showLineNums
              ? 'calc(100% - ' + lineNumsWidth + 'px)'
              : '100%',
          }"
          ref="textarea"
          :autofocus="autofocus"
          spellcheck="false"
          @keydown.tab.prevent.stop="tab"
          @scroll="calcScrollDistance"
          :value="modelValue == undefined ? content : modelValue"
          @input="updateValue"
        ></textarea>
        <pre
          :style="{
            paddingRight: scrollBarWidth + 'px',
            paddingBottom: scrollBarHeight + 'px',
            marginLeft: showLineNums ? lineNumsWidth + 'px' : '0',
            width: showLineNums
              ? 'calc(100% - ' + lineNumsWidth + 'px)'
              : '100%',
          }"
        >
        <code
          ref="code"
          v-highlight="contentValue"
          :class="languageClass"
          :style="{
            top: top + 'px',
            left: left + 'px',
            fontSize: fontSize,
            padding: !header ? padding : lineNums ? '10px ' + padding + ' ' + padding : '0 ' + padding + ' ' + padding,
          }">
        </code>
      </pre>
      </div>
    </div>
  </div>
</template>

<script>
import hljs from "highlight.js";
import Dropdown from "./Dropdown.vue";
import CopyCode from "./CopyCode.vue";
import "./themes/themes-base16.css";
import "./themes/themes.css";
// import { useEditorStore } from "../stores/editor";
export default {
  components: {
    Dropdown,
    CopyCode,
  },
  name: "CodeEditor",
  props: {
    lineNums: {
      type: Boolean,
      default: false,
    },
    modelValue: {
      type: String,
    },
    value: {
      type: String,
    },
    theme: {
      type: String,
      default: "github-dark",
    },
    tabSpaces: {
      type: Number,
      default: 2,
    },
    wrap: {
      type: Boolean,
      default: false,
    },
    readOnly: {
      type: Boolean,
      default: false,
    },
    autofocus: {
      type: Boolean,
      default: false,
    },
    header: {
      type: Boolean,
      default: true,
    },
    width: {
      type: String,
      default: "540px",
    },
    height: {
      type: String,
      default: "auto",
    },
    maxWidth: {
      type: String,
    },
    minWidth: {
      type: String,
    },
    maxHeight: {
      type: String,
    },
    minHeight: {
      type: String,
    },
    borderRadius: {
      type: String,
      default: "12px",
    },
    languages: {
      type: Array,
      required: false,
      default: () => [["javascript", "JavaScript"]],
    },
    langListWidth: {
      type: String,
      default: "110px",
    },
    langListHeight: {
      type: String,
      default: "auto",
    },
    langListDisplay: {
      type: Boolean,
      default: false,
    },
    displayLanguage: {
      type: Boolean,
      default: true,
    },
    copyCode: {
      type: Boolean,
      default: true,
    },
    zIndex: {
      type: String,
      default: "0",
    },
    fontSize: {
      type: String,
      default: "17px",
    },
    padding: {
      type: String,
      default: "20px",
    },
  },
  directives: {
    highlight: {
      mounted(el, binding) {
        el.textContent = binding.value;
        hljs.highlightElement(el);
      },
      updated(el, binding) {
        if (el.scrolling) {
          el.scrolling = false;
        } else {
          el.textContent = binding.value;
          hljs.highlightElement(el);
        }
      },
    },
  },
  data() {
    return {
      // editorStore: useEditorStore(),
      scrollBarWidth: 0,
      scrollBarHeight: 0,
      top: 0,
      left: 0,
      languageClass: null,
      languageTitle: null,
      // languageClass: "hljs language-" + this.languages[0][1],
      // languageTitle: this.languages[0][1]
      //   ? this.languages[0][1]
      //   : this.languages[0][0],
      content: this.value,
      cursorPosition: 0,
      insertTab: false,
      lineNum: 0,
      lineNumsWidth: 0,
      scrolling: false,
      textareaHeight: 0,
      showLineNums: this.wrap ? false : this.lineNums,
    };
  },
  computed: {
    tabWidth() {
      let result = "";
      for (let i = 0; i < this.tabSpaces; i++) {
        result += " ";
      }
      return result;
    },
    contentValue() {
      return this.modelValue == undefined
        ? this.content + "\n"
        : this.modelValue + "\n";
    },
    scroll() {
      return this.height == "auto" ? false : true;
    },
    // languageClass() {
    //   return "language-" + this.languages[0][0];
    // },
    // languageTitle() {
    //   return this.languages[0][1] ? this.languages[0][1] : this.languages[0][0];
    // },
  },
  watch: {
    languages: {
      immediate: true,
      handler(newVal) {
        console.log("watcher called", newVal);
        if (newVal && newVal.length > 0) {
          this.languageClass = "language-" + newVal[0][0];
          this.languageTitle = newVal[0][1] || newVal[0][0];
        }
      },
    },
  },

  methods: {
    updateValue(e) {
      const value = e.target.value;
      const cursorPos = e.target.selectionStart;

      // Check if the last character typed is an opening bracket
      const lastChar = value[cursorPos - 1];
      const brackets = {
        "(": ")",
        "{": "}",
        "[": "]",
      };

      // If the last character is an opening bracket, insert the closing bracket
      if (brackets[lastChar]) {
        const closingBracket = brackets[lastChar];
        // Construct the new content
        this.content =
          value.slice(0, cursorPos) + closingBracket + value.slice(cursorPos);

        // Set the cursor position to just after the closing bracket
        this.$nextTick(() => {
          e.target.value = this.content; // Update the textarea value
          e.target.setSelectionRange(cursorPos + 1, cursorPos + 1); // Move cursor to right after closing bracket
        });
      } else {
        // If no auto-close action, just update the content or emit event
        if (this.modelValue === undefined) {
          this.content = value;
        } else {
          this.$emit("update:modelValue", value);
        }
      }

      // Emit the updated content
      this.$emit("content", this.content);
    },
    changeLang(lang) {
      this.languageTitle = lang[1] ? lang[1] : lang[0];
      this.languageClass = "language-" + lang[0];
      console.log("languages in event emit ", lang);
      this.$emit("lang", lang[0]);
    },
    tab() {
      if (document.execCommand("insertText")) {
        document.execCommand("insertText", false, this.tabWidth);
      } else {
        const cursorPosition = this.$refs.textarea.selectionStart;
        this.content =
          this.content.substring(0, cursorPosition) +
          this.tabWidth +
          this.content.substring(cursorPosition);
        this.cursorPosition = cursorPosition + this.tabWidth.length;
        this.insertTab = true;
      }
    },
    calcScrollDistance(e) {
      // this.$refs.code.scrolling = true;
      if (this.$refs.code) {
        this.scrolling = true;
        this.$refs.code.scrolling = true;
      }
      this.top = -e.target.scrollTop;
      this.left = -e.target.scrollLeft;
    },
    resizer() {
      // textareaResizer
      const textareaResizer = new ResizeObserver((entries) => {
        this.scrollBarWidth =
          entries[0].target.offsetWidth - entries[0].target.clientWidth;
        this.scrollBarHeight =
          entries[0].target.offsetHeight - entries[0].target.clientHeight;
        this.textareaHeight = entries[0].target.offsetHeight;
      });
      textareaResizer.observe(this.$refs.textarea);
      // lineNumsResizer
      const lineNumsResizer = new ResizeObserver((entries) => {
        this.lineNumsWidth = entries[0].target.offsetWidth;
      });
      if (this.$refs.lineNums) {
        lineNumsResizer.observe(this.$refs.lineNums);
      }
    },
    copy() {
      if (document.execCommand("copy")) {
        this.$refs.textarea.select();
        document.execCommand("copy");
        window.getSelection().removeAllRanges();
      } else {
        navigator.clipboard.writeText(this.$refs.textarea.value);
      }
    },
    getLineNum() {
      // lineNum
      const str = this.$refs.textarea.value;
      let lineNum = 0;
      let position = str.indexOf("\n");
      while (position !== -1) {
        lineNum++;
        position = str.indexOf("\n", position + 1);
      }
      // heightNum
      const singleLineHeight = this.$refs.lineNums.firstChild.offsetHeight;
      const heightNum = parseInt(this.textareaHeight / singleLineHeight) - 1;
      // displayed lineNum
      this.lineNum =
        this.height == "auto"
          ? lineNum
          : lineNum > heightNum
            ? lineNum
            : heightNum;
    },
    autoFormat() {
      const formattedCode = this.formatPythonCode(this.content);
      this.content = formattedCode;
      this.$emit("update:modelValue", formattedCode); // Emit the formatted code
    },

    formatPythonCode(code) {
      const lines = code.split("\n");
      let indentLevel = 0;

      const formattedLines = lines.map((line) => {
        const trimmedLine = line.trim();

        // Increase indentation for lines ending with a colon
        if (trimmedLine.endsWith(":")) {
          indentLevel++;
        }
        // Decrease indentation for specific keywords
        if (
          /^(return|pass|break|continue|elif|else|except|finally)\b/.test(
            trimmedLine
          )
        ) {
          indentLevel = Math.max(indentLevel - 1, 0);
        }

        // Indent line
        return " ".repeat(indentLevel * this.tabSpaces) + trimmedLine;
      });

      return formattedLines.join("\n");
    },
    alertThemeChange(e) {
      window.setTimeout(() => {
        const element = document.querySelector(`[theme="${e}"] .hljs`);
        const styles = window.getComputedStyle(element);
        // console.log("mitt", this.emitter);
        // console.log(styles);
        let st = {
          bg: styles.background,
          foreground: styles.color,
        };
        this.emitter.emit("changeTerminalTheme", st);
      }, 100);
    },
  },
  mounted() {
    // console.log("languages are here finally", this.languages);
    this.$emit("content", this.content);
    this.$emit("textarea", this.$refs.textarea);
    this.resizer();
    // this.$emit("lang", this.languages[0][0]);
    // let code = document.getElementsByClassName("hljs")[0];
  },
  beforeUnmount() {
    this.emitter.off("changeLang");
    this.emitter.off("ThemeChanged");
  },
  created() {
    this.emitter.on("ThemeChanged", this.alertThemeChange);
    this.emitter.on("changeLang", this.changeLang);
  },
  updated() {
    if (this.insertTab) {
      this.$refs.textarea.setSelectionRange(
        this.cursorPosition,
        this.cursorPosition
      );
      this.insertTab = false;
    }
    if (this.lineNums) {
      if (this.scrolling) {
        this.scrolling = false;
      } else {
        this.getLineNum();
      }
    }
  },
};
</script>

<style lang="scss" scoped>
.code-editor {
  position: relative;
}
.code-editor > div {
  width: 100%;
  height: 100%;
}
/* header */
.code-editor .header {
  box-sizing: border-box;
  position: relative;
  z-index: 1;
  height: 34px;
}
.code-editor .header > .dropdown {
  position: absolute;
  top: 12px;
  left: 18px;
}
.code-editor .header > .copy-code {
  position: absolute;
  top: 10px;
  right: 12px;
}
/* code-area */
.code-editor .code-area {
  position: relative;
  z-index: 0;
  text-align: left;
  overflow: hidden;
}
/* font style */
.code-editor .code-area > textarea,
.code-editor .code-area > pre > code,
.code-editor .line-nums > div {
  font-family: Consolas, Monaco, monospace;
  line-height: 1.5;
}
.code-editor .code-area > textarea:hover,
.code-editor .code-area > textarea:focus-visible {
  outline: none;
}
.code-editor .code-area > textarea {
  position: absolute;
  z-index: 1;
  top: 0;
  left: 0;
  overflow-y: hidden;
  box-sizing: border-box;
  caret-color: rgb(127, 127, 127);
  color: transparent;
  white-space: pre;
  word-wrap: normal;
  border: 0;
  width: 100%;
  height: 100%;
  background: none;
  resize: none;
}
.code-editor .code-area > pre {
  box-sizing: border-box;
  position: relative;
  z-index: 0;
  overflow: hidden;
  font-size: 0;
  margin: 0;
}
.code-editor .code-area > pre > code {
  background: none;
  display: block;
  position: relative;
  overflow-x: visible !important;
  border-radius: 0;
  box-sizing: border-box;
  margin: 0;
}
/* wrap code */
.code-editor.wrap .code-area > textarea,
.code-editor.wrap .code-area > pre > code {
  white-space: pre-wrap;
  word-wrap: break-word;
}
/* hide-header */
.code-editor.hide-header.scroll .code-area {
  height: 100%;
}
/* scroll */
.code-editor.scroll .code-area {
  height: calc(100% - 34px);
}
.code-editor.scroll .code-area > textarea {
  overflow: auto;
}
.code-editor.scroll .code-area > pre {
  width: 100%;
  height: 100%;
  overflow: hidden;
}
/* dropdown */
.code-editor .list {
  -webkit-user-select: none;
  user-select: none;
  height: 100%;
  font-family: sans-serif;
}
.code-editor .list > .lang-list {
  border-radius: 5px;
  box-sizing: border-box;
  overflow: auto;
  font-size: 13px;
  padding: 0;
  margin: 0;
  list-style: none;
  text-align: left;
}
.code-editor .list > .lang-list > li {
  font-size: 13px;
  transition:
    background 0.16s ease,
    color 0.16s ease;
  box-sizing: border-box;
  padding: 0 12px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 30px;
}
.code-editor .list > .lang-list > li:first-child {
  padding-top: 5px;
}
.code-editor .list > .lang-list > li:last-child {
  padding-bottom: 5px;
}
.code-editor .list > .lang-list > li:hover {
  background: rgba(160, 160, 160, 0.4);
}
/* line-nums */
.code-editor .line-nums {
  min-width: 36px;
  text-align: right;
  box-sizing: border-box;
  position: absolute;
  left: 0;
  padding-right: 8px;
  padding-left: 8px;
  opacity: 0.3;
}
.code-editor .line-nums::after {
  content: "";
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  border-right: 1px solid currentColor;
  opacity: 0.5;
}
.code-editor .header.border::after {
  content: "";
  position: absolute;
  width: 100%;
  height: 1px;
  bottom: 0;
  left: 0;
  background: currentColor;
  opacity: 0.15;
}
</style>
