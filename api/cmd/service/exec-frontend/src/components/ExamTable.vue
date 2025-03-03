<template>
  <div class="flex align-items-center justify-content-center">
    <DataTable
      v-model:expandedRows="expandedRows"
      :value="settings"
      dataKey="id"
      @rowExpand="onRowExpand"
      @rowCollapse="onRowCollapse"
      tableStyle="min-width: 90rem"
    >
      <template #header>
        <div class="flex flex-wrap justify-content-end gap-2">
          <Button text icon="pi pi-plus" label="Expand All" @click="addRow" />
        </div>
      </template>
      <Column expander style="width: 5rem" />
      <Column field="tags" header="Tags">
        <template #body="">
          <tagging class="w-100" @addedTags="getTags" /> </template
      ></Column>
      <!-- <Column field="difficulty" header="Difficulty">
        <template #body="">
          <Dropdown
            v-model="selectedDiff"
            :options="difficulties"
            optionLabel="label"{
        tags: this.selectedTags,
        lang: this.selectedLanguage.value,
        is_qc: this.showQCDoneOnly,
        page: 1,
        row: 8,
      }
            placeholder="Select Difficulty"
            class=""
          />
        </template>
      </Column> -->
      <Column field="easyQuestion" header="Easy Questions">
        <!-- <template #body="">
          <InputNumber v-model="easy" inputId="integeronly" fluid />
        </template> -->
      </Column>
      <Column field="mediumQuestion" header="Medium Questions">
        <!-- <template #body="">
          <InputNumber v-model="medium" inputId="integeronly" fluid />
        </template> -->
      </Column>
      <Column field="highQuestion" header="High Questions">
        <!-- <template #body="">
          <InputNumber v-model="high" inputId="integeronly" fluid />
        </template> -->
      </Column>
      <Column field="number" header="Total Questions"></Column>
      <Column field="" header="Activities">
        <template #body="slotProps">
          <div class="flex gap-3">
            <Button
              icon="pi pi-search"
              outlined
              @click="addQuestions(slotProps.data.id)"
            ></Button>
            <Button icon="pi pi-trash" severity="danger" outlined></Button>
          </div>
        </template>
      </Column>
      <!-- <Column header="Status">
        <template #body="slotProps">
          <Tag
            :value="slotProps.data.inventoryStatus"
            :severity="getSeverity(slotProps.data)"
          />
        </template>
      </Column> -->
      <template #expansion="slotProps">
        <div class="p-4">
          <!-- <h5>Orders for {{ rowSetting[slotProps.data.id] }}</h5> -->
          <DataTable :value="rowSetting[slotProps.data.id].questions">
            <Column field="id" header="Id" sortable></Column>
            <Column field="title" header="Title" sortable></Column>
            <Column field="created_at" header="Date" sortable>
              <template #body="slotProps">
                {{ formatDate(slotProps.data.answer.createdat) }}
              </template>
            </Column>
            <!-- <Column field="tags" header="Tags">
              <template #body="slotProps">
                <div
                  class="flex flex-row"
                  v-for="val in slotProps.data.tags"
                  :key="val"
                >
                  {{ val }}
                </div>
              </template>
            </Column> -->
            <Column field="tags" header="Tags" sortable></Column>
            <!-- <Column field="status" header="Status" sortable>
              <template #body="slotProps">
                <Tag
                  :value="slotProps.data.status.toLowerCase()"
                  :severity="getOrderSeverity(slotProps.data)"
                />
              </template>
            </Column>
            <Column headerStyle="width:4rem">
              <template #body>
                <Button icon="pi pi-search" />
              </template>
            </Column> -->
          </DataTable>
        </div>
      </template>
    </DataTable>
    <Toast />
  </div>
</template>

<script>
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputNumber from "primevue/inputnumber";
import Tagging from "./Tagging.vue";
import { useCreatorStore } from "../stores/creator";
import { format, parseISO } from "date-fns";
export default {
  components: {
    DataTable,
    Column,
    Tagging,
    InputNumber,
  },
  data() {
    return {
      easy: null,
      medium: null,
      high: null,
      settings: [],
      expandedRows: {},
      difficulties: [
        { label: "Low", value: "low" },
        { label: "Medium", value: "medium" },
        { label: "High", value: "high" },
      ],
      selectedDiff: null,
      rowSetting: new Map(),
      tags: null,
      creatorStore: useCreatorStore(),
      newRow: {
        id: "",
        tags: "",
        easyQuestion: 0,
        mediumQuestion: 0,
        highQuestion: 0,
        number: 0,
        questions: [],
      },
      filteredQuestion: null,
      TotalRecords: 0,
    };
  },
  computed: {},
  mounted() {
    this.settings = [];
    this.rowSetting = new Map();
  },
  methods: {
    formatDate(isoDate) {
      const date = parseISO(isoDate);
      // Get the humanized distance from now
      return format(date, "MMMM do, yyyy, h:mm a");
    },
    onRowExpand(event) {
      this.$toast.add({
        severity: "info",
        summary: "Questions Expanded",
        detail: event.data.name,
        life: 3000,
      });
    },
    onRowCollapse(event) {
      this.$toast.add({
        severity: "success",
        summary: "Questions Collapsed",
        detail: event.data.name,
        life: 3000,
      });
    },
    addRow() {
      this.newRow.id = `id-${Date.now()}`;
      this.settings.push(this.newRow);
      this.rowSetting[`id-${Date.now()}`] = this.newRow;
      this.newRow = {};
      // console.log("Settings Map", this.rowSetting);
    },
    getTags(val) {
      // console.log("tags",val)
      this.tags = val;
    },
    addQuestions(id) {
      // console.log("selected tags>>>>>>>>>>", this.tags);
      console.log("selected tags>>>>>>>>>>", id);
      let filters = {
        tags: this.tags,
        lang: "python",
        is_qc: true,
        page: 1,
        row: 100,
      };
      this.query(filters, id);
    },
    query(filters, id) {
      this.creatorStore
        .queryWithFilters(filters)
        .then((res) => {
          // console.log("results>>>>>>>>>>>>>>", res);
          // this.filteredQuestion = res.Documents;
          const index = this.settings.findIndex((obj) => obj.id === id);
          this.rowSetting[id].questions = res.Documents;
          console.log("rowSetting>>>>>>>>", this.rowSetting);
          if (index == -1) {
            return;
          }
          for (let i = 0; i < res.Documents.length; i++) {
            // console.log("difficulty",res.Documents[i].difficulty)
            if(res.Documents[i].difficulty==="easy"){
              this.rowSetting[id].easyQuestion += 1
            }else if(res.Documents[i].difficulty==="medium"){
              this.rowSetting[id].mediumQuestion += 1
            }else{
              this.rowSetting[id].highQuestion += 1
            }
          }
          this.settings[index].questions = res.Documents;
          this.TotalRecords = res.Count;
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
  },
};
</script>
