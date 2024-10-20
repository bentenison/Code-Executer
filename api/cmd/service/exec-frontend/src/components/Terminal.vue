<template>
  <exec-term
    class="hljs"
    :theme="currentTheme"
    name="my-terminal"
    :context="context"
    @exec-cmd="onExecCmd"
    :drag-conf="dragConf"
    :title="title"
    :init-log="Messages"
    :show-header="false"
  />
</template>
<script>
import execTerm from "vue-web-terminal";
//  Light theme: vue-web-terminal/lib/theme/light.css
import "vue-web-terminal/lib/theme/dark.css";


export default {
  name: "Terminal",
  components: {
    execTerm,
  },
  data() {
    return {
      defaultTheme:{bg:"#0d1117",foreground:"#c9d1d9"},
      context: "🖥️root@tanvirs:~$",
      title: "EPIC terminal",
      currentTheme: "github",
      Messages: [
        {
          type: "normal",
          content: "Initializing terminal...",
          class: "info",
        },
        {
          type: "normal",
          content: "Loading configuration...",
          class: "info",
        },
        {
          type: "normal",
          content: "Configuration loaded successfully.",
          class: "success",
        },
        {
          type: "normal",
          content: "Starting code execution...",
          class: "info",
        },
        {
          type: "code",
          content: "const result = calculate();",
          class: "info",
        },
        // {
        //   type: "json",
        //   content: {
        //     status: "success",
        //     data: {
        //       value: 42,
        //     },
        //   },
        //   class: "info",
        // },
        {
          type: "normal",
          content: "Execution completed with no errors.",
          class: "success",
        },
        {
          type: "table",
          content: [
            ["Step", "Status"],
            ["Initialization", "Success"],
            ["Loading Config", "Success"],
            ["Execution", "Success"],
          ],
          class: "info",
        },
        {
          type: "normal",
          content: "Cleaning up resources...",
          class: "info",
        },
        {
          type: "normal",
          content: "Terminating terminal session.",
          class: "system",
        },
      ],
      dragConf: {
        width: "50%",
        height: "70%",
        zIndex: -0,
        init: {
          x: 200,
          y: 200,
        },
        pinned: false,
      },
    };
  },
  methods: {
    onExecCmd(key, command, success, failed) {
      if (key === "fail") {
        failed("Something wrong!!!");
      } else {
        let allClass = ["success", "error", "system", "info", "warning"];

        let clazz = allClass[Math.floor(Math.random() * allClass.length)];
        success({
          type: "normal",
          class: clazz,
          tag: clazz,
          content: `Your command is '${command}'`,
        });
      }
    },
    changeTheme(e) {
      console.log("event catched!!!", e);
      let win = document.getElementsByClassName("t-window")[0];
      // console.log(win)
      win.style.background = e.bg;
      win.style.color = e.foreground;
    },
  },
  mounted() {
    this.changeTheme(this.defaultTheme)
    // updateColor() {
    // Update the CSS variable dynamically
    // document.documentElement.style.setProperty("--t-main-background-color", "red");
    // },
    // console.log(doc2.attributes);
  },
  created() {
    this.emitter.on("changeTerminalTheme", this.changeTheme);
  },
};
</script>
<style lang="scss">
.t-container {
  all: unset !important;
  // display: inline-block;
  position: relative !important;
  width: auto !important;
  height: auto !important;
  left: auto !important;
  top: auto !important;
  z-index: auto !important;
}
.terminal .t-window {
  // all: unset !important;
  position: static;
  // position: relative;
  margin-top: 10px !important;
  height: 300px !important;
  border-radius: 12px;
  width: 100%;
  transition: background 0.5s ease-in-out;
  // background: #0d1117;
}
.t-cmd-help {
  // all: unset !important;
  position: absolute;
  // position: static;
  // right: 20px;
  // top: 0px;
  // top: 0px;
  // bottom: 0px;
  // z-index: 100;
  display: none;
}
.t-cmd-tips {
  // all: unset !important;
  // position: static;
  // right: 20px;
  top: 0px;
}
pre {
  color: grey;
}
// .hljs{
//   all: inherit;
//   background: transparent;
// }
</style>
