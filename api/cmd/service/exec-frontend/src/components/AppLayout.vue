<!-- <template> -->
<template>
  <div>
    <div class="layout-wrapper" :class="'layout-static'">
      <menubar />
      <!-- <app-sidebar></app-sidebar> -->
      <!-- <AppSideBar /> -->
      <demo/>
      <div class="layout-main-container">
        <div class="layout-main flex-auto">
          <router-view></router-view>
        </div>
        <app-footer></app-footer>
      </div>
      <div class="layout-mask animate-fadein"></div>
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
export default {
  components: { Menubar, Demo, AppFooter, AppSideBar },
};
</script>

<style lang="scss" scoped></style>
