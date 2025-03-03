<template>
  <div class="exam-details-form px-4">
    <!-- <h2>{{ isEditMode ? "Edit Exam Details" : "Create Exam Details" }}</h2> -->
    <div class="grid flex flex-wrap">
      <div class="col-6">
        <div class="field flex flex-column">
          <label for="examCode">Exam Code</label>
          <InputText id="examCode" v-model="examDetails.examCode" required />
        </div>
      </div>

      <div class="col-6">
        <div class="field flex flex-column">
          <label for="examName">Exam Name</label>
          <InputText id="examName" v-model="examDetails.examName" required />
        </div>
      </div>
      <div class="col-6">
        <div class="field flex flex-column">
          <label for="defaultLanguageId">Default Language ID</label>
          <InputText
            id="defaultLanguageId"
            v-model="examDetails.defaultLanguageId"
            required
          />
        </div>
      </div>

      <div class="col-6">
        <div class="field flex flex-column">
          <label for="allowedLanguages">Allowed Languages</label>
          <Dropdown
            id="allowedLanguages"
            v-model="examDetails.allowedLanguages"
            :options="languageOptions"
            multiple
          />
        </div>
      </div>
      <div class="col-4">
          <div class="field flex flex-column">
            <label for="durationInSec">Duration (seconds)</label>
            <InputNumber
              id="durationInSec"
              v-model="examDetails.durationInSec"
              required
            />
          </div>
        </div>
        <div class="col-4">
          <div class="field flex flex-column">
            <label for="totalMarks">Total Marks</label>
            <InputNumber
              id="totalMarks"
              v-model="examDetails.totalMarks"
              required
            />
          </div>
        </div>
        <div class="col-4">
          <div class="field flex flex-column">
            <label for="minimumPassingMarks">Minimum Passing Marks</label>
            <InputNumber
              id="minimumPassingMarks"
              v-model="examDetails.minimumPassingMarks"
              required
            />
          </div>
        </div>
        <!-- <div class="col-12">
          <div class="field">
              <label for="instructions">Instructions</label>
              <Textarea id="instructions" v-model="examDetails.instructions.text" rows="3" />
            </div>
        </div> -->
    </div>
    <div>
      <h3>Sections</h3>
      <Button
        label="Add Section"
        @click="addSection"
        class="p-button-outlined"
      />
      <!-- v-for="(section, index) in examDetails.sections" -->
      <div
        :key="index"
        class="card section border-round p-2 my-2"
      >
        <div class="grid flex gap-3 justify-content-center grid-nogutter">
          <div class="col-5">
            <div class="field flex flex-column">
              <label :for="'sectionName' + index">Section Name</label>
              <InputText
                :id="'sectionName' + index"
                v-model="section.sectionName"
              />
            </div>
          </div>
          <div class="col-5">
            <div class="field flex flex-column">
              <label :for="'totalItems' + index">Total Items</label>
              <InputNumber
                :id="'totalItems' + index"
                v-model="section.totalItems"
              />
            </div>
          </div>
          <div class="col-5">
            <div class="field flex flex-column">
              <label :for="'totalMarks' + index">Total Marks</label>
              <InputNumber
                :id="'totalMarks' + index"
                v-model="section.totalMarks"
              />
            </div>
          </div>
          <div class="col-5">
            <div class="field flex flex-column">
              <label :for="'durationInSec' + index">Duration (seconds)</label>
              <InputNumber
                :id="'durationInSec' + index"
                v-model="section.durationInSec"
              />
            </div>
          </div>
        </div>
        <Button
          label="Remove Section"
          @click="removeSection(index)"
          class="p-button-danger"
        />
      </div>

      <!-- <Button type="submit" label="Submit" class="p-button" /> -->
    </div>
  </div>
</template>

<script>
import InputText from "primevue/inputtext";
import InputNumber from "primevue/inputnumber";
import Dropdown from "primevue/dropdown";
import Textarea from "primevue/textarea";
import Button from "primevue/button";
import { Form } from "@primevue/forms";
export default {
  components: {
    InputText,
    InputNumber,
    Dropdown,
    Textarea,
    Button,
    Form,
  },
  props: {
    isEditMode: {
      type: Boolean,
      default: false,
    },
    initialExamDetails: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
    return {
      section: {},
      examDetails: { ...this.initialExamDetails },
      languageOptions: [
        { label: "Java", value: "java" },
        { label: "Python", value: "python" },
        { label: "C++", value: "c++" },
        // Add more languages as needed
      ],
    };
  },
  methods: {
    addSection() {
      this.examDetails.sections.push({
        sectionId: "",
        sectionName: "",
        totalItems: 0,
        totalMarks: 0,
        durationInSec: 0,
        questionItemList: [],
        sectionConfigs: [],
      });
    },
    removeSection(index) {
      this.examDetails.sections.splice(index, 1);
    },
    submitForm() {
      console.log("Exam Details Submitted:", this.examDetails);
    },
  },
};
</script>

<style scoped>
.exam-details-form {
  min-width: 100rem;
  /* margin: auto; */
  /* padding: 20px; */
}

.section {
  /* margin-bottom: 1em; */
}

/* .border-1 {
  border: 1px solid #ccc;
} */

.border-round {
  border-radius: 10px;
}
</style>
