<template>
  <div class="card">
    <Editor v-model="value" editorStyle="height: 320px" />
  </div>
</template>
<script>
import Editor from "primevue/editor";
export default {
  components: {
    Editor,
  },
  data() {
    return {};
  },
};
</script>

<style lang="scss" scoped></style>
