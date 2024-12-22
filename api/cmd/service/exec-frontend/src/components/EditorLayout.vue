<!-- <template> -->
<template>
  <div>
    <div class="layout-wrapper">
      <menubar />
      <!-- <app-sidebar></app-sidebar> -->
      <!-- <AppSideBar /> -->
      <!-- <demo/> -->
      <div class="layout-main-container">
        <div class="layout-main">
          <home />
        </div>

        <app-footer></app-footer>
      </div>
      <!-- <div class="layout-mask animate-fadein"></div> -->
    </div>
    <Toast />
  </div>
</template>
<!-- </template> -->

<script>
import Demo from "../views/Demo.vue";
import Menubar from "./Menubar.vue";
import AppFooter from "./AppFooter.vue";
import AppSideBar from "./AppSideBar.vue";
import Home from "../views/Home.vue";
export default {
  components: { Menubar, Demo, AppFooter, AppSideBar, Home },
};
</script>

<style lang="scss" scoped>
.main-content {
  flex: 1; /* Take the remaining space */
  background-color: #ffffff; /* Background color */
  padding: 20px; /* Padding for content */
}

.sidebar {
  width: 3rem; /* Width of the sidebar */
  background-color: red; /* Background color */
  padding: 20px; /* Padding for content */
}
</style>
