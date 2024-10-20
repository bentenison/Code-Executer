<template>
  <div>
    <menubar />
    <div
      class="mini-sidebar border-solid border-y-none border-left-none border-right-1 border-white-alpha-20 flex flex-column align-items-center surface-ground"
    >
      <!-- @mouseover="visible = !visible" -->
      <!-- <TieredMenu ref="menu" id="overlay_tmenu" :model="items" popup /> -->
      <div
        class="logo w-100 flex align-items-center justify-content-center p-1 surface-ground"
        style="background-color: #272e38"
      >
        <!-- <Avatar
          image="../../../img/sec-icon.svg"
          class="flex align-items-center justify-content-center mr-2 overflow-hidden"
          size="large"
        /> -->
      </div>
      <TieredMenu ref="menu" id="overlay_tmenu" :model="menuitems" popup />
      <TieredMenu ref="menu1" id="overlay_tmenu1" :model="reportItems" popup />
      <TieredMenu ref="menu2" id="overlay_tmenu2" :model="userItems" popup />
      <TieredMenu ref="menu3" id="overlay_tmenu3" :model="licenceItems" popup />
      <!-- <TieredMenu ref="menu4" id="overlay_tmenu" :model="menuitems" popup /> -->
      <div
        class="menu-icons mt-8 mb-5 flex h-20rem flex-column align-items-center justify-content-between"
      >
        <Button
          @click="toggle"
          icon="pi pi-check"
          text
          rounded
          aria-label="Filter"
          aria-haspopup="true"
          aria-controls="overlay_tmenu"
        />
        <Button
          icon="pi pi-bookmark"
          text
          rounded
          aria-label="Bookmark"
          aria-haspopup="true"
          aria-controls="overlay_tmenu1"
          @click="toggle1"
        />
        <!-- <Button icon="pi pi-search" text rounded aria-label="Search" /> -->
        <Button
          icon="pi pi-user"
          text
          rounded
          aria-label="User"
          aria-haspopup="true"
          aria-controls="overlay_tmenu2"
          @click="toggle2"
        />
        <Button
          icon="pi pi-bell"
          text
          rounded
          aria-label="Notification"
          aria-haspopup="true"
          aria-controls="overlay_tmenu3"
          @click="toggle3"
        />
        <!-- <Button icon="pi pi-heart" text rounded aria-label="Favorite" />
        <Button icon="pi pi-times" text rounded aria-label="Cancel" /> -->
      </div>
    </div>
    <Sidebar v-model:visible="visible" class="" style="width: 250px">
      <!-- @mouseleave="visible = !visible" -->
      <template #container="{ closeCallback }">
        <div class="flex flex-column h-full">
          <div
            class="flex align-items-center justify-content-between px-4 py-3 flex-shrink-0"
            style="background-color: #272e38"
          >
            <span class="inline-flex align-items-center gap-2">
              <!-- <svg
                width="35"
                height="40"
                viewBox="0 0 35 40"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M25.87 18.05L23.16 17.45L25.27 20.46V29.78L32.49 23.76V13.53L29.18 14.73L25.87 18.04V18.05ZM25.27 35.49L29.18 31.58V27.67L25.27 30.98V35.49ZM20.16 17.14H20.03H20.17H20.16ZM30.1 5.19L34.89 4.81L33.08 12.33L24.1 15.67L30.08 5.2L30.1 5.19ZM5.72 14.74L2.41 13.54V23.77L9.63 29.79V20.47L11.74 17.46L9.03 18.06L5.72 14.75V14.74ZM9.63 30.98L5.72 27.67V31.58L9.63 35.49V30.98ZM4.8 5.2L10.78 15.67L1.81 12.33L0 4.81L4.79 5.19L4.8 5.2ZM24.37 21.05V34.59L22.56 37.29L20.46 39.4H14.44L12.34 37.29L10.53 34.59V21.05L12.42 18.23L17.45 26.8L22.48 18.23L24.37 21.05ZM22.85 0L22.57 0.69L17.45 13.08L12.33 0.69L12.05 0H22.85Z"
                  fill="var(--primary-color)"
                />
                <path
                  d="M30.69 4.21L24.37 4.81L22.57 0.69L22.86 0H26.48L30.69 4.21ZM23.75 5.67L22.66 3.08L18.05 14.24V17.14H19.7H20.03H20.16H20.2L24.1 15.7L30.11 5.19L23.75 5.67ZM4.21002 4.21L10.53 4.81L12.33 0.69L12.05 0H8.43002L4.22002 4.21H4.21002ZM21.9 17.4L20.6 18.2H14.3L13 17.4L12.4 18.2L12.42 18.23L17.45 26.8L22.48 18.23L22.5 18.2L21.9 17.4ZM4.79002 5.19L10.8 15.7L14.7 17.14H14.74H15.2H16.85V14.24L12.24 3.09L11.15 5.68L4.79002 5.2V5.19Z"
                  fill="var(--text-color)"
                />
              </svg> -->
              <svg
                version="1.0"
                xmlns="http://www.w3.org/2000/svg"
                width="35"
                height="40"
                viewBox="0 0 35 40"
                fill="none"
              >
                <g id="surface1">
                  <path
                    style="
                      stroke: none;
                      fill-rule: nonzero;
                      fill: var(--primary-color);
                      fill-opacity: 1;
                    "
                    d="M 7.335938 4.136719 C 6.390625 4.730469 4.773438 6.503906 4.476562 7.289062 C 4.328125 7.636719 4.226562 8.25 4.226562 8.644531 L 4.226562 9.359375 L 28.597656 9.359375 L 28.597656 18.351562 C 28.597656 23.300781 28.523438 27.339844 28.425781 27.339844 C 28.300781 27.316406 27.804688 26.722656 27.28125 25.984375 L 26.335938 24.628906 L 19.347656 24.628906 L 19.769531 25.296875 C 20.019531 25.691406 20.492188 26.355469 20.816406 26.773438 L 21.414062 27.585938 L 11.441406 27.585938 L 11.441406 33.496094 L 25.265625 33.496094 L 26.136719 34.410156 C 26.636719 34.925781 27.28125 35.714844 27.578125 36.132812 L 28.128906 36.945312 L 35.339844 36.945312 L 33.746094 34.902344 C 31.660156 32.191406 31.707031 32.339844 32.429688 31.722656 C 32.777344 31.402344 33.300781 30.714844 33.621094 30.171875 C 34.195312 29.1875 34.195312 29.113281 34.269531 18.964844 C 34.320312 11.796875 34.269531 8.449219 34.070312 7.785156 C 33.621094 6.257812 32.652344 5.050781 31.359375 4.335938 L 30.140625 3.695312 L 19.125 3.695312 C 8.851562 3.695312 8.03125 3.71875 7.335938 4.136719 Z M 7.335938 4.136719 "
                  />
                  <path
                    style="
                      stroke: none;
                      fill-rule: nonzero;
                      fill: var(--text-color);
                      fill-opacity: 1;
                    "
                    d="M 4.226562 19.976562 C 4.226562 30 4.203125 29.777344 5.820312 31.476562 C 6.738281 32.464844 7.535156 32.90625 8.929688 33.226562 L 9.949219 33.449219 L 9.949219 24.136719 L 20.890625 24.136719 L 20.890625 18.226562 L 9.949219 18.226562 L 9.949219 16.75 L 25.144531 16.75 L 25.066406 13.84375 L 24.992188 10.960938 L 14.621094 10.910156 L 4.226562 10.835938 Z M 4.226562 19.976562 "
                  />
                </g>
              </svg>
              <span class="font-semibold text-2xl text-primary">FASTQ</span>
            </span>
            <span>
              <Button
                type="button"
                @click="closeCallback"
                icon="pi pi-times"
                rounded
                outlined
                class="h-2rem w-2rem"
              ></Button>
            </span>
          </div>
          <div class="overflow-y-auto">
            <ul class="list-none p-3 m-0">
              <li>
                <div
                  v-ripple
                  v-styleclass="{
                    selector: '@next',
                    enterClass: 'hidden',
                    enterActiveClass: 'slidedown',
                    leaveToClass: 'hidden',
                    leaveActiveClass: 'slideup',
                  }"
                  class="p-3 flex align-items-center justify-content-between text-600 cursor-pointer p-ripple"
                >
                  <span class="font-medium">FAVORITES</span>
                  <i class="pi pi-chevron-down"></i>
                </div>
                <ul class="list-none p-0 m-0 overflow-hidden">
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-home mr-2"></i>
                      <span class="font-medium">Dashboard</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-bookmark mr-2"></i>
                      <span class="font-medium">Bookmarks</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      v-styleclass="{
                        selector: '@next',
                        enterClass: 'hidden',
                        enterActiveClass: 'slidedown',
                        leaveToClass: 'hidden',
                        leaveActiveClass: 'slideup',
                      }"
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-chart-line mr-2"></i>
                      <span class="font-medium">Reports</span>
                      <i class="pi pi-chevron-down ml-auto"></i>
                    </a>
                    <ul
                      class="list-none py-0 pl-3 pr-0 m-0 hidden overflow-y-hidden transition-all transition-duration-400 transition-ease-in-out"
                    >
                      <li>
                        <a
                          v-ripple
                          v-styleclass="{
                            selector: '@next',
                            enterClass: 'hidden',
                            enterActiveClass: 'slidedown',
                            leaveToClass: 'hidden',
                            leaveActiveClass: 'slideup',
                          }"
                          class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                        >
                          <i class="pi pi-chart-line mr-2"></i>
                          <span class="font-medium">Revenue</span>
                          <i class="pi pi-chevron-down ml-auto"></i>
                        </a>
                        <ul
                          class="list-none py-0 pl-3 pr-0 m-0 hidden overflow-y-hidden transition-all transition-duration-400 transition-ease-in-out"
                        >
                          <li>
                            <a
                              v-ripple
                              class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                            >
                              <i class="pi pi-table mr-2"></i>
                              <span class="font-medium">View</span>
                            </a>
                          </li>
                          <li>
                            <a
                              v-ripple
                              class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                            >
                              <i class="pi pi-search mr-2"></i>
                              <span class="font-medium">Search</span>
                            </a>
                          </li>
                        </ul>
                      </li>
                      <li>
                        <a
                          v-ripple
                          class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                        >
                          <i class="pi pi-chart-line mr-2"></i>
                          <span class="font-medium">Expenses</span>
                        </a>
                      </li>
                    </ul>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-users mr-2"></i>
                      <span class="font-medium">Team</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-comments mr-2"></i>
                      <span class="font-medium">Messages</span>
                      <span
                        class="inline-flex align-items-center justify-content-center ml-auto bg-primary border-circle"
                        style="min-width: 1.5rem; height: 1.5rem"
                        >3</span
                      >
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-calendar mr-2"></i>
                      <span class="font-medium">Calendar</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-cog mr-2"></i>
                      <span class="font-medium">Settings</span>
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
            <ul class="list-none p-3 m-0">
              <li>
                <div
                  v-ripple
                  v-styleclass="{
                    selector: '@next',
                    enterClass: 'hidden',
                    enterActiveClass: 'slidedown',
                    leaveToClass: 'hidden',
                    leaveActiveClass: 'slideup',
                  }"
                  class="p-3 flex align-items-center justify-content-between text-600 cursor-pointer p-ripple"
                >
                  <span class="font-medium">APPLICATION</span>
                  <i class="pi pi-chevron-down"></i>
                </div>
                <ul class="list-none p-0 m-0 overflow-hidden">
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-folder mr-2"></i>
                      <span class="font-medium">Projects</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-chart-bar mr-2"></i>
                      <span class="font-medium">Performance</span>
                    </a>
                  </li>
                  <li>
                    <a
                      v-ripple
                      class="flex align-items-center cursor-pointer p-3 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
                    >
                      <i class="pi pi-cog mr-2"></i>
                      <span class="font-medium">Settings</span>
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
          </div>
          <div class="mt-auto overflow-hidden">
            <hr class="mb-3 mx-3 border-top-1 border-none surface-border" />
            <a
              v-ripple
              class="m-3 flex align-items-center cursor-pointer p-3 gap-2 border-round text-700 hover:surface-100 transition-duration-150 transition-colors p-ripple"
            >
              <Avatar
                image="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
                shape="circle"
              />
              <span class="font-bold">Amy Elsner</span>
            </a>
          </div>
        </div>
      </template>
    </Sidebar>
    <!-- <Button icon="pi pi-bars" @click="visible = true" /> -->
    <!-- <Sidebar :visible="true">
        <button v-ripple class="relative overflow-hidden w-full p-link flex align-items-center p-2 pl-3 text-color hover:surface-200 border-noround">
            <Avatar image="/images/avatar/amyelsner.png" class="mr-2" shape="circle" />
            <span class="inline-flex flex-column">
                <span class="font-bold">Amy Elsner</span>
                <span class="text-sm">Admin</span>
            </span>
        </button>
        <PanelMenu :model="items">
          <template #item="{ item }">
            <a v-ripple class="flex align-items-center px-3 py-2 cursor-pointer">
              <span :class="[item.icon, 'text-primary']" />
              <span :class="['ml-2', { 'font-semibold': item.items }]">{{
                item.label
              }}</span>
              <Badge v-if="item.badge" class="ml-auto" :value="item.badge" />
              <span
                v-if="item.shortcut"
                class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1"
                >{{ item.shortcut }}</span
              >
            </a>
          </template>
        </PanelMenu>
    </Sidebar> -->
  </div>
</template>

<script>
import PanelMenu from "primevue/panelmenu";
import TieredMenu from "primevue/tieredmenu";
import Menubar from "../components/Menubar.vue";

export default {
  data() {
    return {
      visible: false,
      reportItems: [
        {
          label: "Reports",
          icon: "pi pi-chart-bar",
          shortcut: "⌘+R",
          items: [
            {
              label: "Sales",
              icon: "pi pi-chart-line",
              badge: 3,
            },
            {
              label: "Products",
              icon: "pi pi-list",
              badge: 6,
            },
          ],
        },
      ],
      licenceItems: [
        {
          label: "Mail",
          icon: "pi pi-envelope",
          badge: 5,
          items: [
            {
              label: "Compose",
              icon: "pi pi-file-edit",
              shortcut: "⌘+N",
            },
            {
              label: "Inbox",
              icon: "pi pi-inbox",
              badge: 5,
            },
            {
              label: "Sent",
              icon: "pi pi-send",
              shortcut: "⌘+S",
            },
            {
              label: "Trash",
              icon: "pi pi-trash",
              shortcut: "⌘+T",
            },
          ],
        },
      ],
      userItems: [
        {
          label: "Profile",
          icon: "pi pi-user",
          shortcut: "⌘+W",
          items: [
            {
              label: "Settings",
              icon: "pi pi-cog",
              shortcut: "⌘+O",
            },
            {
              label: "Privacy",
              icon: "pi pi-shield",
              shortcut: "⌘+P",
            },
          ],
        },
      ],
      menuitems: [
        {
          label: "File",
          icon: "pi pi-file",
          items: [
            {
              label: "New",
              icon: "pi pi-plus",
              items: [
                {
                  label: "Document",
                  icon: "pi pi-file",
                  shortcut: "⌘+N",
                },
                {
                  label: "Image",
                  icon: "pi pi-image",
                  shortcut: "⌘+I",
                },
                {
                  label: "Video",
                  icon: "pi pi-video",
                  shortcut: "⌘+L",
                },
              ],
            },
            {
              label: "Open",
              icon: "pi pi-folder-open",
              shortcut: "⌘+O",
            },
            {
              label: "Print",
              icon: "pi pi-print",
              shortcut: "⌘+P",
            },
          ],
        },
        {
          label: "Edit",
          icon: "pi pi-file-edit",
          items: [
            {
              label: "Copy",
              icon: "pi pi-copy",
              shortcut: "⌘+C",
            },
            {
              label: "Delete",
              icon: "pi pi-times",
              shortcut: "⌘+D",
            },
          ],
        },
        {
          label: "Search",
          icon: "pi pi-search",
          shortcut: "⌘+S",
        },
        {
          separator: true,
        },
        {
          label: "Share",
          icon: "pi pi-share-alt",
          items: [
            {
              label: "Slack",
              icon: "pi pi-slack",
              badge: 2,
            },
            {
              label: "Whatsapp",
              icon: "pi pi-whatsapp",
              badge: 3,
            },
          ],
        },
      ],
    };
  },
  components: { PanelMenu, TieredMenu, Menubar },
  methods: {
    toggle(event) {
      this.$refs.menu.toggle(event);
    },
    toggle1(event) {
      this.$refs.menu1.toggle(event);
    },
    toggle2(event) {
      this.$refs.menu2.toggle(event);
    },
    toggle3(event) {
      this.$refs.menu3.toggle(event);
    },
  },
  mounted() {
    console.log(this.$refs);
  },
};
</script>

<style scoped>
.mini-sidebar {
  position: fixed;
  height: 100%;
  top: 0;
  left: 0;
  width: 3.8rem;
  z-index: 999;
  /* background: white; */
  /* border-right: 1px solid var(--surface-500); */
  /* box-shadow: 0px 3px 5px rgba(0, 0, 0, 0.02), 0px 0px 5px rgba(0, 0, 0, 0.05),
    0px 3px 7px rgba(0, 0, 0, 0.8); */
  display: flex;
  flex-direction: column;
}
</style>
