<template>
  <div class="card fixed z-5 top-0 w-full">
    <Menubar :model="items" class="text-muted border-none">
      <template #start>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
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
                fill: rgb(0%, 77.254903%, 35.294119%);
                fill-opacity: 1;
              "
              d="M 6.824219 -0.109375 C 7.785156 -0.121094 7.785156 -0.121094 8.765625 -0.136719 C 9.460938 -0.140625 10.15625 -0.144531 10.871094 -0.148438 C 11.582031 -0.152344 12.296875 -0.160156 13.03125 -0.164062 C 14.539062 -0.171875 16.046875 -0.179688 17.558594 -0.183594 C 19.867188 -0.195312 22.175781 -0.222656 24.484375 -0.25 C 25.949219 -0.257812 27.414062 -0.261719 28.878906 -0.265625 C 29.570312 -0.277344 30.261719 -0.289062 30.972656 -0.300781 C 35.773438 -0.285156 35.773438 -0.285156 38.234375 1.1875 C 39.949219 3.160156 40.082031 4.261719 40.109375 6.824219 C 40.121094 7.785156 40.121094 7.785156 40.136719 8.765625 C 40.140625 9.460938 40.144531 10.15625 40.148438 10.871094 C 40.152344 11.582031 40.160156 12.296875 40.164062 13.03125 C 40.171875 14.539062 40.179688 16.046875 40.183594 17.558594 C 40.195312 19.867188 40.222656 22.175781 40.25 24.484375 C 40.257812 25.949219 40.261719 27.414062 40.265625 28.878906 C 40.277344 29.570312 40.289062 30.261719 40.300781 30.972656 C 40.285156 35.773438 40.285156 35.773438 38.8125 38.234375 C 36.839844 39.949219 35.738281 40.082031 33.175781 40.109375 C 32.214844 40.121094 32.214844 40.121094 31.234375 40.136719 C 30.191406 40.140625 30.191406 40.140625 29.128906 40.148438 C 28.417969 40.152344 27.703125 40.160156 26.96875 40.164062 C 25.460938 40.171875 23.953125 40.179688 22.441406 40.183594 C 20.132812 40.195312 17.824219 40.222656 15.515625 40.25 C 14.050781 40.257812 12.585938 40.261719 11.121094 40.265625 C 10.085938 40.28125 10.085938 40.28125 9.027344 40.300781 C 4.226562 40.285156 4.226562 40.285156 1.765625 38.8125 C 0.0507812 36.839844 -0.0820312 35.738281 -0.109375 33.175781 C -0.117188 32.535156 -0.125 31.894531 -0.136719 31.234375 C -0.140625 30.539062 -0.144531 29.84375 -0.148438 29.128906 C -0.152344 28.417969 -0.160156 27.703125 -0.164062 26.96875 C -0.171875 25.460938 -0.179688 23.953125 -0.183594 22.441406 C -0.195312 20.132812 -0.222656 17.824219 -0.25 15.515625 C -0.257812 14.050781 -0.261719 12.585938 -0.265625 11.121094 C -0.28125 10.085938 -0.28125 10.085938 -0.300781 9.027344 C -0.285156 4.226562 -0.285156 4.226562 1.1875 1.765625 C 3.160156 0.0507812 4.261719 -0.0820312 6.824219 -0.109375 Z M 15.625 5 C 16.207031 5.691406 16.792969 6.378906 17.375 7.066406 C 17.699219 7.449219 18.023438 7.832031 18.359375 8.226562 C 19.359375 9.382812 19.359375 9.382812 20.449219 10.425781 C 21.554688 11.515625 22.066406 12.25 22.5 13.75 C 21.964844 16.664062 20.660156 19.230469 19.375 21.875 C 18.882812 22.914062 18.394531 23.957031 17.902344 24.996094 C 17.140625 26.613281 16.375 28.226562 15.59375 29.828125 C 15.355469 30.320312 15.117188 30.8125 14.871094 31.316406 C 14.546875 31.976562 14.546875 31.976562 14.214844 32.648438 C 13.667969 33.742188 13.667969 33.742188 13.75 35 C 16.675781 35.015625 19.601562 35.027344 22.527344 35.035156 C 23.523438 35.039062 24.519531 35.042969 25.515625 35.046875 C 26.945312 35.054688 28.375 35.058594 29.800781 35.0625 C 30.25 35.0625 30.699219 35.066406 31.160156 35.070312 C 32.230469 35.070312 33.304688 35.039062 34.375 35 C 35.359375 34.015625 35.078125 33.039062 35.082031 31.675781 C 35.082031 30.816406 35.082031 30.816406 35.085938 29.9375 C 35.085938 29.3125 35.085938 28.691406 35.082031 28.046875 C 35.082031 27.410156 35.085938 26.773438 35.085938 26.121094 C 35.085938 24.773438 35.085938 23.425781 35.082031 22.078125 C 35.078125 20.007812 35.082031 17.9375 35.085938 15.867188 C 35.085938 14.5625 35.085938 13.257812 35.082031 11.953125 C 35.085938 11.328125 35.085938 10.707031 35.085938 10.0625 C 35.085938 9.488281 35.082031 8.917969 35.082031 8.324219 C 35.082031 7.816406 35.078125 7.308594 35.078125 6.789062 C 35.128906 5.617188 35.128906 5.617188 34.375 5 C 32.988281 4.945312 31.601562 4.933594 30.210938 4.9375 C 29.796875 4.941406 29.382812 4.941406 28.953125 4.941406 C 27.621094 4.945312 26.292969 4.953125 24.960938 4.960938 C 24.058594 4.964844 23.160156 4.96875 22.257812 4.96875 C 20.046875 4.976562 17.835938 4.988281 15.625 5 Z M 6.875 8.125 C 6.460938 8.539062 6.050781 8.949219 5.625 9.375 C 5.824219 10.691406 5.824219 10.691406 6.25 11.875 C 7.75 12.300781 7.75 12.300781 9.375 12.5 C 9.789062 12.085938 10.199219 11.675781 10.625 11.25 C 10.425781 9.933594 10.425781 9.933594 10 8.75 C 8.5 8.324219 8.5 8.324219 6.875 8.125 Z M 6.875 8.125 "
            />
            <path
              style="
                stroke: none;
                fill-rule: nonzero;
                fill: rgb(0%, 77.254903%, 35.294119%);
                fill-opacity: 1;
              "
              d="M 25.898438 26.757812 C 26.390625 26.75 26.886719 26.742188 27.394531 26.734375 C 28.679688 26.730469 29.964844 26.796875 31.25 26.875 C 31.664062 27.289062 32.074219 27.699219 32.5 28.125 C 32.34375 29.726562 32.34375 29.726562 31.875 31.25 C 29.980469 32.199219 28.015625 31.992188 25.9375 31.992188 C 25.507812 32 25.074219 32.007812 24.632812 32.015625 C 23.503906 32.019531 22.375 31.949219 21.25 31.875 C 20 30.625 20 30.625 19.882812 29.023438 C 20.140625 25.691406 23.246094 26.761719 25.898438 26.757812 Z M 25.898438 26.757812 "
            />
          </g>
        </svg>
      </template>
      <template>
        <SplitButton
          label="Save"
          class="text-primary"
          :model="items"
          @click="save"
        ></SplitButton>
      </template>
      <template #end>
        <div class="flex align-items-center gap-2">
          <Button type="button" @click="toggleDarkMode">
            <i
              :class="[
                'pi',
                { 'pi-moon': isDarkTheme, 'pi-sun': !isDarkTheme },
              ]"
            ></i>
          </Button>
          <Button @click="emitter.emit('PopOpened', $event)" class="">
            <i class="pi pi-palette"></i>
          </Button>
          <Appconfigurator />
          <InputText placeholder="Search" type="text" class="w-32 sm:w-auto" />
          <Avatar
            image="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
            shape="circle"
          />
        </div>
      </template>
    </Menubar>

    <Dialog
      v-model:visible="visible"
      modal
      header="IDE Settings"
      :style="{ width: '45rem', height: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
    >
      <div class="flex flex-column gap-3">
        <div class="flex">
          <p class="text-grey font-semibold">Input Control</p>
        </div>
        <div class="flex w-20rem gap-3">
          <p>Mode :</p>
          <SelectButton
            v-model="value"
            :options="ThemeOptions"
            @change="toggleDarkMode"
            aria-labelledby="basic"
          />
        </div>

        <div class="flex gap-3">
          <p class="text-muted line-height-1">Theme :</p>
          <theme-selector :Themes="themes" />
        </div>
        <div class="flex">
          <p class="text-grey font-semibold">Editor Control</p>
        </div>
        <div class="flex gap-4">
          <p>Line Number :</p>
          <SelectButton
            v-model="value"
            :options="LineOptions"
            aria-labelledby="basic"
          />
        </div>
        <div class="flex gap-5">
          <p>Auto format :</p>
          <SelectButton
            v-model="value"
            :options="options"
            aria-labelledby="basic"
          />
        </div>
        <div class="flex gap-6">
          <p>Font Size :</p>
          <div class="flex gap-3 align-items-center">
            <Button
              icon="pi pi-minus"
              @click="decreaseFontSize"
              :disabled="fontSize <= 14"
            />
            <span>{{ fontSize }}</span>
            <Button
              icon="pi pi-plus"
              @click="increaseFontSize"
              :disabled="fontSize >= 20"
            />
          </div>
        </div>
      </div>
    </Dialog>
  </div>
</template>

<script>
import { useLayout } from "../components/layout";
import Appconfigurator from "./Appconfigurator.vue";
import SelectButton from "primevue/selectbutton";
import ThemeSelector from "./ThemeSelector.vue";
export default {
  components: {
    Appconfigurator,
    SelectButton,
    ThemeSelector,
  },
  data() {
    return {
      isDarkTheme: false,
      visible: false,
      fontSize: 17,
      themes: [
        { name: "a11y-dark", code: "a11y-dark" },
        { name: "a11y-light", code: "a11y-light" },
        { name: "agate", code: "agate" },
        { name: "an-old-hope", code: "an-old-hope" },
        { name: "androidstudio", code: "androidstudio" },
        { name: "arduino-light", code: "arduino-light" },
        { name: "arta", code: "arta" },
        { name: "ascetic", code: "ascetic" },
        { name: "atom-one-dark-reasonable", code: "atom-one-dark-reasonable" },
        { name: "atom-one-dark", code: "atom-one-dark" },
        { name: "atom-one-light", code: "atom-one-light" },
        { name: "brown-paper", code: "brown-paper" },
        { name: "codepen-embed", code: "codepen-embed" },
        { name: "color-brewer", code: "color-brewer" },
        { name: "dark", code: "dark" },
        { name: "default", code: "default" },
        { name: "devibeans", code: "devibeans" },
        { name: "docco", code: "docco" },
        { name: "far", code: "far" },
        { name: "felipec", code: "felipec" },
        { name: "foundation", code: "foundation" },
        { name: "github-dark-dimmed", code: "github-dark-dimmed" },
        { name: "github-dark", code: "github-dark" },
        { name: "github", code: "github" },
        { name: "gml", code: "gml" },
        { name: "googlecode", code: "googlecode" },
        { name: "gradient-dark", code: "gradient-dark" },
        { name: "gradient-light", code: "gradient-light" },
        { name: "grayscale", code: "grayscale" },
        { name: "hybrid", code: "hybrid" },
        { name: "idea", code: "idea" },
        { name: "intellij-light", code: "intellij-light" },
        { name: "ir-black", code: "ir-black" },
        { name: "isbl-editor-dark", code: "isbl-editor-dark" },
        { name: "isbl-editor-light", code: "isbl-editor-light" },
        { name: "kimbie-dark", code: "kimbie-dark" },
        { name: "kimbie-light", code: "kimbie-light" },
        { name: "lightfair", code: "lightfair" },
        { name: "lioshi", code: "lioshi" },
        { name: "magula", code: "magula" },
        { name: "mono-blue", code: "mono-blue" },
        { name: "monokai-sublime", code: "monokai-sublime" },
        { name: "monokai", code: "monokai" },
        { name: "night-owl", code: "night-owl" },
        { name: "nnfx-dark", code: "nnfx-dark" },
        { name: "nnfx-light", code: "nnfx-light" },
        { name: "nord", code: "nord" },
        { name: "obsidian", code: "obsidian" },
        { name: "panda-syntax-dark", code: "panda-syntax-dark" },
        { name: "panda-syntax-light", code: "panda-syntax-light" },
        { name: "paraiso-dark", code: "paraiso-dark" },
        { name: "paraiso-light", code: "paraiso-light" },
        { name: "pojoaque", code: "pojoaque" },
        { name: "purebasic", code: "purebasic" },
        { name: "qtcreator-dark", code: "qtcreator-dark" },
        { name: "qtcreator-light", code: "qtcreator-light" },
        { name: "rainbow", code: "rainbow" },
        { name: "routeros", code: "routeros" },
        { name: "school-book", code: "school-book" },
        { name: "shades-of-purple", code: "shades-of-purple" },
        { name: "srcery", code: "srcery" },
        { name: "stackoverflow-dark", code: "stackoverflow-dark" },
        { name: "stackoverflow-light", code: "stackoverflow-light" },
        { name: "sunburst", code: "sunburst" },
        { name: "tokyo-night-dark", code: "tokyo-night-dark" },
        { name: "tokyo-night-light", code: "tokyo-night-light" },
        { name: "tomorrow-night-blue", code: "tomorrow-night-blue" },
        { name: "tomorrow-night-bright", code: "tomorrow-night-bright" },
        { name: "vs", code: "vs" },
        { name: "vs2015", code: "vs2015" },
        { name: "xcode", code: "xcode" },
        { name: "xt256", code: "xt256" },
        { name: "base16-3024", code: "base16-3024" },
        { name: "base16-apathy", code: "base16-apathy" },
        { name: "base16-apprentice", code: "base16-apprentice" },
        { name: "base16-ashes", code: "base16-ashes" },
        {
          name: "base16-atelier-cave-light",
          code: "base16-atelier-cave-light",
        },
        { name: "base16-atelier-cave", code: "base16-atelier-cave" },
        {
          name: "base16-atelier-dune-light",
          code: "base16-atelier-dune-light",
        },
        { name: "base16-atelier-dune", code: "base16-atelier-dune" },
        {
          name: "base16-atelier-estuary-light",
          code: "base16-atelier-estuary-light",
        },
        { name: "base16-atelier-estuary", code: "base16-atelier-estuary" },
        {
          name: "base16-atelier-forest-light",
          code: "base16-atelier-forest-light",
        },
        { name: "base16-atelier-forest", code: "base16-atelier-forest" },
        {
          name: "base16-atelier-heath-light",
          code: "base16-atelier-heath-light",
        },
        { name: "base16-atelier-heath", code: "base16-atelier-heath" },
        {
          name: "base16-atelier-lakeside-light",
          code: "base16-atelier-lakeside-light",
        },
        { name: "base16-atelier-lakeside", code: "base16-atelier-lakeside" },
        {
          name: "base16-atelier-plateau-light",
          code: "base16-atelier-plateau-light",
        },
        { name: "base16-atelier-plateau", code: "base16-atelier-plateau" },
        {
          name: "base16-atelier-savanna-light",
          code: "base16-atelier-savanna-light",
        },
        { name: "base16-atelier-savanna", code: "base16-atelier-savanna" },
        {
          name: "base16-atelier-seaside-light",
          code: "base16-atelier-seaside-light",
        },
        { name: "base16-atelier-seaside", code: "base16-atelier-seaside" },
        {
          name: "base16-atelier-sulphurpool-light",
          code: "base16-atelier-sulphurpool-light",
        },
        {
          name: "base16-atelier-sulphurpool",
          code: "base16-atelier-sulphurpool",
        },
        { name: "base16-atlas", code: "base16-atlas" },
        { name: "base16-bespin", code: "base16-bespin" },
        {
          name: "base16-black-metal-bathory",
          code: "base16-black-metal-bathory",
        },
        {
          name: "base16-black-metal-burzum",
          code: "base16-black-metal-burzum",
        },
        {
          name: "base16-black-metal-dark-funeral",
          code: "base16-black-metal-dark-funeral",
        },
        {
          name: "base16-black-metal-gorgoroth",
          code: "base16-black-metal-gorgoroth",
        },
        {
          name: "base16-black-metal-immortal",
          code: "base16-black-metal-immortal",
        },
        { name: "base16-black-metal-khold", code: "base16-black-metal-khold" },
        {
          name: "base16-black-metal-marduk",
          code: "base16-black-metal-marduk",
        },
        {
          name: "base16-black-metal-mayhem",
          code: "base16-black-metal-mayhem",
        },
        { name: "base16-black-metal-nile", code: "base16-black-metal-nile" },
        { name: "base16-black-metal-venom", code: "base16-black-metal-venom" },
        { name: "base16-black-metal", code: "base16-black-metal" },
        { name: "base16-brewer", code: "base16-brewer" },
        { name: "base16-bright", code: "base16-bright" },
        { name: "base16-brogrammer", code: "base16-brogrammer" },
        { name: "base16-brush-trees-dark", code: "base16-brush-trees-dark" },
        { name: "base16-brush-trees", code: "base16-brush-trees" },
        { name: "base16-chalk", code: "base16-chalk" },
        { name: "base16-circus", code: "base16-circus" },
        { name: "base16-classic-dark", code: "base16-classic-dark" },
        { name: "base16-classic-light", code: "base16-classic-light" },
        { name: "base16-codeschool", code: "base16-codeschool" },
        { name: "base16-colors", code: "base16-colors" },
        { name: "base16-cupcake", code: "base16-cupcake" },
        { name: "base16-cupertino", code: "base16-cupertino" },
        { name: "base16-danqing", code: "base16-danqing" },
        { name: "base16-darcula", code: "base16-darcula" },
        { name: "base16-dark-violet", code: "base16-dark-violet" },
        { name: "base16-darkmoss", code: "base16-darkmoss" },
        { name: "base16-darktooth", code: "base16-darktooth" },
        { name: "base16-decaf", code: "base16-decaf" },
        { name: "base16-default-dark", code: "base16-default-dark" },
        { name: "base16-default-light", code: "base16-default-light" },
        { name: "base16-dirtysea", code: "base16-dirtysea" },
        { name: "base16-dracula", code: "base16-dracula" },
        { name: "base16-edge-dark", code: "base16-edge-dark" },
        { name: "base16-edge-light", code: "base16-edge-light" },
        { name: "base16-eighties", code: "base16-eighties" },
        { name: "base16-embers", code: "base16-embers" },
        { name: "base16-equilibrium-dark", code: "base16-equilibrium-dark" },
        {
          name: "base16-equilibrium-gray-dark",
          code: "base16-equilibrium-gray-dark",
        },
        {
          name: "base16-equilibrium-gray-light",
          code: "base16-equilibrium-gray-light",
        },
        { name: "base16-equilibrium-light", code: "base16-equilibrium-light" },
        { name: "base16-espresso", code: "base16-espresso" },
        { name: "base16-eva-dim", code: "base16-eva-dim" },
        { name: "base16-eva", code: "base16-eva" },
        { name: "base16-flat", code: "base16-flat" },
        { name: "base16-framer", code: "base16-framer" },
        { name: "base16-fruit-soda", code: "base16-fruit-soda" },
        { name: "base16-gigavolt", code: "base16-gigavolt" },
        { name: "base16-github", code: "base16-github" },
        { name: "base16-google-dark", code: "base16-google-dark" },
        { name: "base16-google-light", code: "base16-google-light" },
        { name: "base16-grayscale-dark", code: "base16-grayscale-dark" },
        { name: "base16-grayscale-light", code: "base16-grayscale-light" },
        { name: "base16-green-screen", code: "base16-green-screen" },
        { name: "base16-gruvbox-dark-hard", code: "base16-gruvbox-dark-hard" },
        {
          name: "base16-gruvbox-dark-medium",
          code: "base16-gruvbox-dark-medium",
        },
        { name: "base16-gruvbox-dark-pale", code: "base16-gruvbox-dark-pale" },
        { name: "base16-gruvbox-dark-soft", code: "base16-gruvbox-dark-soft" },
        {
          name: "base16-gruvbox-light-hard",
          code: "base16-gruvbox-light-hard",
        },
        {
          name: "base16-gruvbox-light-medium",
          code: "base16-gruvbox-light-medium",
        },
        {
          name: "base16-gruvbox-light-soft",
          code: "base16-gruvbox-light-soft",
        },
        { name: "base16-hardcore", code: "base16-hardcore" },
        { name: "base16-harmonic16-dark", code: "base16-harmonic16-dark" },
        { name: "base16-harmonic16-light", code: "base16-harmonic16-light" },
        { name: "base16-heetch-dark", code: "base16-heetch-dark" },
        { name: "base16-heetch-light", code: "base16-heetch-light" },
        { name: "base16-helios", code: "base16-helios" },
        { name: "base16-hopscotch", code: "base16-hopscotch" },
        { name: "base16-horizon-dark", code: "base16-horizon-dark" },
        { name: "base16-horizon-light", code: "base16-horizon-light" },
        { name: "base16-humanoid-dark", code: "base16-humanoid-dark" },
        { name: "base16-humanoid-light", code: "base16-humanoid-light" },
        { name: "base16-ia-dark", code: "base16-ia-dark" },
        { name: "base16-ia-light", code: "base16-ia-light" },
        { name: "base16-icy-dark", code: "base16-icy-dark" },
      ],
      value: "One-Way",
      options: ["On", "Off"],
      ThemeOptions: ["Light", "Dark"],
      LineOptions: ["Show", "Hide"],
      items: [
        {
          label: "Settings",
          icon: "pi pi-cog",
          command: () => {
            this.visible = true;
            this.$toast.add({
              severity: "success",
              summary: "Updated",
              detail: "Data Updated",
              life: 3000,
            });
          },
        },
        {
          label: "Projects",
          icon: "pi pi-search",
          badge: 3,
        },
      ],
    };
  },
  methods: {
    openSettings() {
      this.emitter.emit("OpenDialog", true);
    },
    toggleDarkMode() {
      const { toggleDarkMode } = useLayout();
      toggleDarkMode();
      this.isDarkTheme = !this.isDarkTheme; // Sync with local state
    },
    increaseFontSize() {
      this.fontSize++;
      this.emitter.emit("increaseFont", this.fontSize);
    },
    decreaseFontSize() {
      this.fontSize--;
      this.emitter.emit("decreaseFont", this.fontSize);
    },
  },
  created() {
    // this.emitter.on('toggleMode',this.toggleDarkMode)
  },
};
</script>

<style lang="scss">
// .p-menubar{
//   background: var(--p-surface-600) !important;
// }
</style>
