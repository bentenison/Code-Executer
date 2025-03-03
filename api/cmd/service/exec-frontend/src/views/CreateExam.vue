<template>
  <div class="flex flex-column align-items-center justify-content-center">
    <div class="title">
      <h2 class="text-center">Exam Creator</h2>
    </div>
    <Stepper value="1" linear class="flex-1" style="width: 100rem">
      <StepList>
        <Step value="1">Exam Settings</Step>
        <Step value="2">Create Sections</Step>
        <Step value="3">Preview</Step>
      </StepList>
      <StepPanels>
        <StepPanel v-slot="{ activateCallback }" value="1">
          <div class="flex align-items-center justify-content-center p-5">
            <exam-form />
          </div>
          <div class="flex p-4 justify-content-end">
            <Button
              label="Next"
              icon="pi pi-arrow-right"
              @click="
                activateCallback('2');
              "
            />
          </div>
        </StepPanel>
        <StepPanel v-slot="{ activateCallback }" value="2">
          <div class="flex flex-column p-5">
            <Message
              severity="error"
              variant="outlined"
              v-if="error"
              class="m-2"
              >{{ error }}</Message
            >
            <div class="flex justify-content-center gap-5">
              <ExamTable />
            </div>
          </div>
          <div class="flex p-4 justify-content-between">
            <Button
              label="Back"
              severity="secondary"
              icon="pi pi-arrow-left"
              @click="activateCallback('1')"
            />
            <Button
              label="Next"
              icon="pi pi-arrow-right"
              iconPos="right"
              @click="activateCallback('3')"
            />
          </div>
        </StepPanel>
        <StepPanel v-slot="{ activateCallback }" value="3">
          <div class="flex flex-column p-5" style="min-width: 100rem">
            <q-cards />
          </div>
          <div class="flex p-4 justify-content-between">
            <Button
              label="Back"
              severity="secondary"
              icon="pi pi-arrow-left"
              @click="activateCallback('2')"
            />
          </div>
        </StepPanel>
      </StepPanels>
    </Stepper>
  </div>
</template>

<script>
import ExamTable from "../components/ExamTable.vue";
import Stepper from "primevue/stepper";
import StepList from "primevue/steplist";
import StepPanels from "primevue/steppanels";
import StepItem from "primevue/stepitem";
import Step from "primevue/step";
import StepPanel from "primevue/steppanel";
import ExamForm from "../components/ExamForm.vue";
export default {
  components: {
    ExamTable,
    Stepper,
    StepList,
    StepPanels,
    StepItem,
    Step,
    StepPanel,
    ExamForm,
  },
  data() {
    return {};
  },
};
</script>

<style lang="scss" scoped></style>
