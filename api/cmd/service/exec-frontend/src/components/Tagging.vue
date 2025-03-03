<template>
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
</template>

<script>
//   import { InputText } from 'primevue/inputtext';

export default {
  components: {
    //   InputText
  },
  // props: {
  //   value: {
  //     type: String,
  //     required: true,
  //   },
  // },
  data() {
    return {
      currentTag: "", // Current tag being typed
      tags: [], // List of tags
    };
  },
  methods: {
    addTag() {
      const trimmedTag = this.currentTag.trim();
      if (trimmedTag && !this.tags.includes(trimmedTag)) {
        this.tags.push(trimmedTag);
        this.currentTag = ""; // Clear input after adding the tag
        this.$emit("addedTags", this.tags);
      }
    },
    removeTag(index) {
      this.tags.splice(index, 1); // Remove tag from the array
      this.$emit("addedTags", this.tags);
    },
  },
};
</script>

<style scoped>
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
  border: 1px solid var(--p-form-field-border-color);
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
  background-color: #007bff;
  color: white;
  padding: 5px 10px;
  border-radius: 20px;
  font-size: 14px;
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
