<template>
  <div class="card flex justify-content-center">
    <Select
      v-model="selectedCountry"
      :options="Themes"
      filter
      optionLabel="name"
      placeholder="Select a Theme"
      class="w-full md:w-56"
      @change="changeTheme"
    >
      <template #value="slotProps">
        <div v-if="slotProps.value" class="flex align-items-center">
          <img
            :alt="slotProps.value.label"
            src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
            :class="`mr-2 flag flag-${slotProps.value.code.toLowerCase()}`"
            style="width: 18px"
          />
          <div>{{ slotProps.value.name }}</div>
        </div>
        <span v-else>
          {{ slotProps.placeholder }}
        </span>
      </template>
      <template #option="slotProps">
        <div class="flex align-items-center">
          <img
            :alt="slotProps.option.label"
            src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
            :class="`mr-2 flag flag-${slotProps.option.code.toLowerCase()}`"
            style="width: 18px"
          />
          <div>{{ slotProps.option.name }}</div>
        </div>
      </template>
    </Select>
  </div>
</template>

<script>
import Select from "primevue/select";
export default {
  components: {
    Select,
  },
  props: {
    Themes: {
      type: Array,
      default: [],
    },
  },
  data() {
    return {
      selectedCountry: null,
    };
  },
  methods: {
    changeTheme(e) {
      console.log("chnage theme",this.selectedCountry.name );
      this.emitter.emit("ThemeChangeSettings",this.selectedCountry.name)
    },
  },
};
</script>
