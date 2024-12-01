<template>
  <div class="flex flex-column">
    <h2 class="text-center text-muted">QC Questions</h2>
    <filters />
    <!-- <Paginator :rows="5" @page="pageChanged"  :rowsPerPageOptions="[5,10, 20, 30]"></Paginator> -->
    <div
      class="min-w-full min-h-full flex justify-content-center align-items-center flex-wrap gap-3"
      v-if="creatorStore.filteredQuestion.length > 0"
    >
      <div
        class="flex align-items-center flex-wrap"
        style="width: 30rem; height: 17rem"
        v-for="q in creatorStore.filteredQuestion"
        :key="q.id"
      >
        <Card class="h-full shadow-3">
          <template #title>
            <div class="flex flex-1 align-items-center gap-3">
              <div class="">
                <i class="pi pi-question-circle" style="font-size: 2rem"></i>
              </div>
              <div class="text flex flex-column gap-2">
                <p class="p-0 m-0 text-lg align-self-start">{{ q.title }}</p>
                <div class="flex gap-3">
                  <Tag
                    icon="pi pi-language"
                    severity="success"
                    :value="q.language"
                  ></Tag>
                  <Tag
                    icon="pi pi-thumbtack"
                    severity="warn"
                    :value="q.difficulty"
                  ></Tag>
                </div>
              </div>
            </div>
          </template>
          <template #content>
            <p class="m-0 mt-2">
              {{ q.description }}
            </p>
          </template>
          <template #footer>
            <div class="flex gap-4 mt-1 mb-2">
              <div class="flex flex-column gap-2 w-full">
                <div class="flex gap-4">
                  <span class="p-0 m-0" v-for="t in q.tags" :key="t">{{
                    t
                  }}</span>
                </div>
                <p class="p-0 m-0">{{ q.answer.created_at }}</p>
              </div>
              <Button label="QC" @click="handleQCQuestion(q)" v-if="!q.is_qc" />
              <Button
                icon="pi pi-check"
                severity="success"
                v-if="q.is_qc"
                v-tooltip="'QC Success'"
                rounded
              />
            </div>
          </template>
        </Card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Filters from "../components/Filters.vue";
import { useCreatorStore } from "../stores/creator";
import { useEditorStore } from "../stores/editor";
import { useMainStore } from "../stores/main";
export default {
  components: { Filters },
  data() {
    return {
      creatorStore: useCreatorStore(),
      editorStore: useEditorStore(),
      mainStore: useMainStore(),
      langId: null,
      executedBy: null,
    };
  },
  methods: {
    getAllQuestions() {
      this.creatorStore
        .getAllQuestions()
        .then((res) => {
          this.$toast.add({
            severity: "success",
            summary: "fetched all question",
            detail: "",
            life: 3000,
          });
        })
        .catch((err) => {
          this.$toast.add({
            severity: "error",
            summary: "broker service is down! contact administrator.",
            detail: "",
            life: 3000,
          });
        });
    },
    prepareData(currQuestion) {},
    handleQCQuestion(currQuestion) {
      // this.editorStore.encode();
      console.log("currQuestion", currQuestion);
      this.mainStore.togglePageBlock();
      axios.defaults.baseURL = "/server";
      this.editorStore
        .getLanguageID(currQuestion.language.toLowerCase())
        .then((res) => {
          console.log("results>>>>>", res);
          this.langId = res;

          currQuestion.file_extension = this.langId.file_extension;
          // var payload = {};
          //   let payload = {
          //     language_code: this.langId.id,
          //     language: this.currQuestion.language,
          //     // code_snippet: this.content,
          //     question_id: this.currQuestion.id,
          //     user_id: "51fc3552-45e0-4982-9adb-50d8cc46c46d",
          //     file_extension: this.langId.file_extension,
          //   };
          this.editorStore
            .qcQuestion(currQuestion)
            .then((res) => {
              this.executedBy = res.data.containerID;
              //   this.emitter.emit("showMessage", res.data.output);
              this.getAllQuestions();
              this.mainStore.togglePageBlock();
              //   this.mainStore.togglePageBlock();
              this.$toast.add({
                severity: "success",
                summary: "QC done.Code executed successfully.",
                detail: res.data.output,
                life: 3000,
              });
            })
            .catch((err) => {
              this.mainStore.togglePageBlock();
              this.$toast.add({
                severity: "error",
                summary: "service is down! contact administrator.",
                detail: err,
                life: 3000,
              });
            });
        })
        .catch((err) => {
          this.mainStore.togglePageBlock();
          this.$toast.add({
            severity: "error",
            summary: "error in qc! contact admin",
            detail: err,
            life: 3000,
          });
        });
    },
  },
  mounted() {
    this.getAllQuestions();
  },
};
</script>

<style lang="scss" scoped></style>
