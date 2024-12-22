<!-- <template> -->
<!-- <demo></demo>
  <div class="">
    <Home></Home>
  </div> -->
<template>
  <div class="flex flex-column justify-content-center">
    <menubar />
    <div id="app-loader" v-if="mainStore.isBlocked">
      <ProgressSpinner></ProgressSpinner>
    </div>
    <Toast />
    <router-view />
  </div>
</template>
<!-- </template> -->

<script>
import { $t, updatePreset, updateSurfacePalette } from "@primevue/themes";
import { useLayout } from "./components/layout";
import ProgressSpinner from "primevue/progressspinner";
import { useMainStore } from "./stores/main";
import { useEditorStore } from "./stores/editor";
import Aura from "@primevue/themes/aura";
import Lara from "@primevue/themes/lara";
import Menubar from "./components/Menubar.vue";
export default {
  components: { ProgressSpinner, Menubar },
  data() {
    return {
      mainStore: useMainStore(),
      editorStore: useEditorStore(),
    };
  },
  mounted() {
    const { toggleDarkMode, setSurface } = useLayout();
    updatePreset(Lara);
    toggleDarkMode();
    // updateSurfacePalette({
    //   surfaceBackground: 'red'
    // });
    // this.$nextTick(function () {
    //   if (localStorage.getItem("simple-code-editor-theme")) {
    //     localStorage.getItem("simple-code-editor-theme") == "light"
    //       ? (document.body.className = "")
    //       : (document.body.className = "dark");
    //   } else {
    //     document.body.className = "dark";
    //   }
    // });
  },
  created() {
    this.editorStore
      .getAllLanguages()
      .then((res) => {})
      .catch((err) => {
        this.$toast.add({
          severity: "error",
          summary: "broker service is down! contact administrator.",
          detail: err,
          life: 3000,
        });
      });
    this.editorStore
      .getQuestTemplates()
      .then((res) => {})
      .catch((err) => {
        this.$toast.add({
          severity: "error",
          summary: "broker service is down! contact administrator.",
          detail: err,
          life: 3000,
        });
      });
  },
  methods: {
    switchTheme() {
      if (document.body.className == "") {
        document.body.className = "dark";
        localStorage.setItem("simple-code-editor-theme", "dark");
      } else {
        document.body.className = "";
        localStorage.setItem("simple-code-editor-theme", "light");
      }
    },
  },
};
</script>

<style lang="scss">
// @import "./assets/theme/lisa.css";
// @import "./assets/theme/lisa-dark.css";
#app-loader {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background: rgba(255, 255, 255, 0.8);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
}
body {
  font-family: sans-serif;
  background: var(--surface-ground);
  margin: 0;
}
#app {
  width: 100%;
  height: 100%;
}
.container {
}
// .float-button {
//   cursor: pointer;
//   color: var(--white);
//   position: fixed;
//   z-index: 5;
//   bottom: 30px;
//   right: 20px;
//   width: 44px;
//   height: 44px;
//   border-radius: 50%;
//   background: var(--grey-9);
//   text-align: center;
//   line-height: 44px;
//   opacity: 0.9;
//   > div {
//     width: 100%;
//     height: 100%;
//     display: flex;
//     align-items: center;
//     justify-content: center;
//   }
//   .sun {
//     display: none;
//   }
//   .moon {
//     display: flex;
//   }
// }
.dark {
  .sun {
    display: flex;
  }
  .moon {
    display: none;
  }
}
</style>
