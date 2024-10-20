<template>
  <demo></demo>
  <div class=" ">
    <Home></Home>
  </div>
</template>

<script>
import Home from "./views/Home.vue";
import Demo from "./views/Demo.vue";
import { $t, updatePreset, updateSurfacePalette } from "@primevue/themes";
import {useLayout} from "./components/layout"
// import { layoutConfig, setPrimary, setSurface, setPreset, isDarkTheme, setMenuMode } = useLayout();
export default {
  components: {
    Home,
    Demo,
  },
  mounted() {
    // this.$nextTick(function () {
    //   if (localStorage.getItem("simple-code-editor-theme")) {
    //     localStorage.getItem("simple-code-editor-theme") == "light"
    //       ? (document.body.className = "")
    //       : (document.body.className = "dark");
    //   } else {
    //     document.body.className = "dark";
    //   }
    // });
    updateSurfacePalette({
      0: "#ffffff",
      50: "#fafafa",
      100: "#f5f5f5",
      200: "#e5e5e5",
      300: "#d4d4d4",
      400: "#a3a3a3",
      500: "#737373",
      600: "#525252",
      700: "#404040",
      800: "#262626",
      900: "#171717",
      950: "#0a0a0a",
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
body {
  font-family: sans-serif;
  background: var(--surface-ground);
  margin: 0;
}
#app {
  width: 100%;
  height: 100%;
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
