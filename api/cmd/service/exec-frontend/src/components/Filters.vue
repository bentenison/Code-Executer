<template>
  <div>
    <!-- Filter Form -->
    <div class="flex justify-content-center gap-5 align-items-center mb-3">
      <!-- Tags Filter -->
      <div class="p-field flex align-items-center">
        <label class="w-8rem" for="tagFilter">Filter by Tag</label>
        <div class="tag-input-container">
          <div class="chips-input-wrapper">
            <!-- Chips inside the input field -->
            <div class="chips-container">
              <span v-for="(tag, index) in tags" :key="index" class="chip">
                {{ tag }}
                <span class="remove-btn" @click="removeTag(index)">×</span>
              </span>
            </div>

            <!-- Input Text to type a new tag -->
            <InputText
              v-model="currentTag"
              @keydown.enter="addTag"
              :disabled="tags.length >= 5"
              placeholder="Enter a tag and press Enter"
              class="input-field"
            />
          </div>
        </div>
      </div>

      <!-- Language Filter -->
      <div class="p-field flex gap-2 align-items-center">
        <label for="languageFilter">Filter by Language</label>
        <Dropdown
          v-model="selectedLanguage"
          :options="languages"
          optionLabel="label"
          placeholder="Select Language"
          class="mr-2"
        />
      </div>

      <!-- QC Done Filter -->
      <div class="p-field flex gap-2 align-items-center">
        <label for="qcDoneFilter">QC Done Only</label>
        <Checkbox
          v-model="showQCDoneOnly"
          binary
          label="Show QC Done"
          class="mr-2"
        />
      </div>

      <!-- Apply Filters Button -->
      <Button label="Apply Filters" icon="pi pi-search" @click="applyFilters" />
    </div>

    <!-- Display Questions -->
  </div>
</template>

<script>
import { ref, computed } from "vue";
import Checkbox from "primevue/checkbox";
import Button from "primevue/button";
// import { useCreatorStore } from "../stores/creator";
import { useEditorStore } from "../stores/editor";
export default {
  name: "App",
  components: {
    Checkbox,
    Button,
  },
  data() {
    return {
      currentTag: "", // Current tag being typed
      tags: [],
      selectedTags: "",
      // creatorStore: useCreatorStore(),
      editorStore: useEditorStore(),
      // Example questions data
      questions: [],
      languages: [
        // { label: "Python", value: "python" },
        // { label: "JavaScript", value: "javascript" },
      ],
      selectedLanguage: null,
      showQCDoneOnly: false,
      // filteredQuestions: [],
    };
  },
  computed: {},
  mounted() {
    this.languages = [];
    for (let index = 0; index < this.editorStore.languages.length; index++) {
      const element = this.editorStore.languages[index];
      let langObj = {
        value: element.name.toLowerCase(),
        label: element.name,
      };
      this.languages.push(langObj);
    }
  },
  methods: {
    addTag() {
      const trimmedTag = this.currentTag.trim();
      if (trimmedTag && !this.tags.includes(trimmedTag)) {
        this.tags.push(trimmedTag);
        if (this.selectedTags.length > 0) {
          this.selectedTags = this.selectedTags + "," + trimmedTag;
        } else {
          this.selectedTags = trimmedTag;
        }
        this.currentTag = ""; // Clear input after adding the tag
      }
    },
    removeTag(index) {
      this.tags.splice(index, 1); // Remove tag from the array
    },
    applyFilters() {
      this.$emit("filtersChanged", {
        tags: this.selectedTags,
        lang: this.selectedLanguage.value,
        is_qc: this.showQCDoneOnly,
        page: 1,
        row: 8,
      });
    },
  },
};
</script>

<style>
/* Optional styles */
.p-d-flex {
  display: flex;
}

.p-jc-between {
  justify-content: space-between;
}

.p-ai-center {
  align-items: center;
}

.p-mr-2 {
  margin-right: 1rem;
}

.p-mt-3 {
  margin-top: 1.5rem;
}

.p-card {
  background-color: #fff;
  padding: 1rem;
  margin-bottom: 1rem;
  border: 1px solid #ddd;
}
.tag-input-container {
  width: 100%;
  display: flex;
  flex-direction: column;
}

.chips-input-wrapper {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  position: relative;
  /* border: 1px solid var(--p-form-field-border-color); */
  padding: 8px;
  border-radius: 4px;
}

.chips-container {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  max-width: 100%;
}

.chip {
  background-color: var(--p-primary-color);
  color: white;
  padding: 5px 10px;
  border-radius: 20px;
  font-size: 14px;
  margin-right: 5px;
  display: flex;
  align-items: center;
}

.remove-btn {
  margin-left: 5px;
  cursor: pointer;
  font-weight: bold;
  font-size: 14px;
}

.input-field {
  flex: 1;
  border: none;
  outline: none;
  padding: 5px 10px;
  font-size: 14px;
  background-color: transparent;
}

.input-field:focus {
  border-color: transparent;
  outline-color: transparent;
}
</style>
