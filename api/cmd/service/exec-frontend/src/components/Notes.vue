<template>
  <div class="card notes-panel">
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

<style lang="scss" scoped>
.notes-panel{
  min-height: 44rem;
  width: 49rem;
  border-radius: 10px;
}
</style>
