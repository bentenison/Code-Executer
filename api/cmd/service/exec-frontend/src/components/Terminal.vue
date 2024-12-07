<template>
  <exec-term
    class="hljs"
    :theme="currentTheme"
    name="my-terminal"
    ref="execTerm"
    :context="context"
    @exec-cmd="onExecCmd"
    :drag-conf="dragConf"
    :title="title"
    :show-header="false"
    :commands="commands"
  />
  <!-- :init-log="Messages" -->
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
      defaultTheme: { bg: "#0d1117", foreground: "#c9d1d9" },
      context: "🖥️root@tanvirs:~$",
      title: "EPIC terminal",
      currentTheme: "github",
      commands: {
        pwd: {
          key: "pwd",
          title: "Print Working Directory",
          description: "Prints the current working directory.",
          usage: "pwd",
          example: [{ input: "pwd", output: "/home/user" }],
          action: this.pwdCommand,
        },
        ls: {
          key: "ls",
          title: "List Directory Contents",
          description: "Lists files in the current directory.",
          usage: "ls",
          example: [{ input: "ls", output: "Desktop  Documents  Downloads" }],
          action: this.lsCommand,
        },
        cd: {
          key: "cd",
          title: "Change Directory",
          description: "Changes the current directory (simulation).",
          usage: "cd <dir>",
          example: [
            {
              input: "cd Documents",
              output: "Changed directory to /home/user/Documents",
            },
          ],
          action: this.cdCommand,
        },
        echo: {
          key: "echo",
          title: "Echo Command",
          description: "Prints the string to the terminal.",
          usage: "echo <string>",
          example: [{ input: "echo Hello, World!", output: "Hello, World!" }],
          action: this.echoCommand,
        },
        date: {
          key: "date",
          title: "Current Date and Time",
          description: "Displays the current date and time.",
          usage: "date",
          example: [{ input: "date", output: "2024-12-01 14:30:45" }],
          action: this.dateCommand,
        },
        clear: {
          key: "clear",
          title: "Clear Terminal",
          description: "Clears the terminal screen.",
          usage: "clear",
          example: [{ input: "clear", output: "" }],
          action: this.clearCommand,
        },
      },
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
    helpCommand() {
      return `
      Available Commands:
      help - Shows this message
      pwd - Prints current directory
      ls - Lists files in the directory
      cd <dir> - Changes the directory
      echo <string> - Echoes the string
      clear - Clears the terminal screen
      date - Displays the current date and time
      `;
    },

    // Command to print current directory (simulated)
    pwdCommand() {
      return "/home/tanvirs";
    },

    // Simulated ls command
    lsCommand() {
      return `
      Desktop  Documents  Downloads  Music  Pictures  Videos
      `;
    },

    // Simulated cd command (does not change actual state in demo)
    cdCommand(args) {
      if (args.length > 0) {
        return `Changed directory to ${args[0]}`;
      }
      return "Error: No directory specified.";
    },

    // Command to echo a string back
    echoCommand(args) {
      return args.join(" ");
    },
    // Command to show the current date and time
    dateCommand() {
      return new Date().toLocaleString();
    },
    // Optional method to execute commands
    executeCommand(command) {},
    onExecCmd(key, command, success, failed) {
      if (this.commands[command]) {
        success({
          type: "normal",
          class: "success",
          tag: "success",
          content: this.commands[command].action(),
        });
        return;
        // return this.commands[command].action();
      }
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
      // setTimeout(() => {
      let win = document.getElementsByClassName("t-window")[0];
      if (win) {
        win.style.background = e.bg;
        win.style.color = e.foreground;
      }
      // console.log(win)
      // }, 100);
    },
    showMessage(data) {
      console.log("terminal refs ", data);
      let msg = data.toLowerCase();
      if (msg.includes("error") || msg.includes("exception")) {
        this.$refs.execTerm.pushMessage({
          class: "error",
          tag: "ERROR",
          content: data,
        });
      } else {
        this.$refs.execTerm.pushMessage({
          class: "success",
          tag: "SUCCESS",
          content: data,
        });
      }
    },
  },
  mounted() {
    this.changeTheme(this.defaultTheme);
    // updateColor() {
    // Update the CSS variable dynamically
    // document.documentElement.style.setProperty("--t-main-background-color", "red");
    // },
    // console.log(doc2.attributes);
  },
  created() {
    this.emitter.on("changeTerminalTheme", this.changeTheme);
    this.emitter.on("showMessage", this.showMessage);
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
  height: 280px !important;
  border-radius: 12px;
  // width: 100%;
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
