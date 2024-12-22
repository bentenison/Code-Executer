<template>
  <div class="">
    <h2 class="text-center text-muted">QC Questions</h2>
    <filters @filtersChanged="changeFilter" />
    <div v-if="filteredQuestion.length > 0"></div>
    <div v-else class="mt-3 text-center">
      <p>No questions found with the selected filters.</p>
    </div>
    <!-- <Paginator :rows="5" @page="pageChanged"  :rowsPerPageOptions="[5,10, 20, 30]"></Paginator> -->
    <div class="flex flex-column align-items-center">
      <div
        class="min-w-full min-h-full flex justify-content-center align-items-center flex-wrap gap-3"
        v-if="filteredQuestion.length > 0"
        :key="page"
      >
        <div
          class="flex align-items-center flex-wrap"
          style="width: 30rem; height: 17rem"
          v-for="q in filteredQuestion"
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
                <Button
                  label="QC"
                  @click="handleQCQuestion(q)"
                  v-if="!q.is_qc"
                />
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
      <div class="mt-5" style="width: 100vw" v-if="filteredQuestion.length > 0">
        <Paginator
          class="flex align-items-center justify-content-center"
          :template="{
            default:
              'FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink JumpToPageDropdown JumpToPageInput',
          }"
          :rows="10"
          :totalRecords="TotalRecords"
          @page="fetchNext"
          @update:rows="rowChanged"
        >
        </Paginator>
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
import Paginator from "primevue/paginator";
export default {
  components: { Filters, Paginator },
  data() {
    return {
      creatorStore: useCreatorStore(),
      editorStore: useEditorStore(),
      mainStore: useMainStore(),
      langId: null,
      executedBy: null,
      filters: null,
      filteredQuestion: [],
      page: 0,
      TotalRecords:0
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
    fetchNext(e) {
      // console.log("pages changed", e);
      // first: 10
      // page: 1
      // pageCount: 12
      // rows: 10
      this.filters.page = e.page + 1;
      //       {
      //   "tags": "",
      //   "lang": "python",
      //   "is_qc": false,
      //   "page": 1,
      //   "row": 10
      // }
      this.query();
    },
    rowChanged(e) {
      console.log("row chnaged", e);
    },
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
              // this.getAllQuestions();
              this.query()
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
    query() {
      this.creatorStore
        .queryWithFilters(this.filters)
        .then((res) => {
          // console.log("results>>>>>>>>>>>>>>", res);
          this.filteredQuestion = res.Documents;
          this.TotalRecords = res.Count
          this.$toast.add({
            severity: "success",
            summary: "questions fetched successfully.",
            detail: "",
            life: 3000,
          });
        })
        .catch((err) => {
          this.$toast.add({
            severity: "error",
            summary: "",
            detail: err,
            life: 3000,
          });
        });
    },
    changeFilter(e) {
      // console.log("filters updated", e);
      this.filters = e;
      console.log("calling");
      this.query();
    },
  },
  mounted() {
    // this.getAllQuestions();
  },
};
</script>

<style lang="scss" scoped></style>
