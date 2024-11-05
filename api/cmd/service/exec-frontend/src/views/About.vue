<template>
  <div class="">
    <div class="card flex justify-content-center">
      <Stepper value="1" linear class="basis-100rem mt-3">
        <StepList>
          <Step value="1">Select Questions</Step>
          <Step value="2">Create Questions</Step>
          <Step value="3">QC</Step>
        </StepList>
        <StepPanels>
          <StepPanel v-slot="{ activateCallback }" value="1">
            <div class="flex flex-column p-5">
              <stepper-form />
            </div>
            <div class="flex p-4 justify-content-end">
              <Button
                label="Next"
                icon="pi pi-arrow-right"
                @click="activateCallback('2')"
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
              <div class="flex justify-content-between gap-5">
                <Textarea
                  variant="filled"
                  rows="5"
                  cols="30"
                  class=""
                  @input="changeValue"
                  v-model="rawJson"
                />
                <JsonViewer :value="val" copyable boxed sort />
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
            <div class="flex flex-column p-5">
              <Message
                severity="error"
                variant="outlined"
                v-if="error"
                class="m-2"
                >{{ error }}</Message
              >
              <div class="flex justify-content-between gap-5">
                <Textarea
                  variant="filled"
                  rows="5"
                  cols="30"
                  class=""
                  @input="changeValue"
                  v-model="rawJson"
                />
                <JsonViewer :value="val" copyable boxed sort />
              </div>
            </div>
            <div class="p-4">
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
  </div>
</template>

<script>
import Stepper from "primevue/stepper";
import StepList from "primevue/steplist";
import StepPanels from "primevue/steppanels";
import StepItem from "primevue/stepitem";
import Step from "primevue/step";
import StepPanel from "primevue/steppanel";
import StepperForm from "../components/StepperForm.vue";
export default {
  components: {
    Stepper,
    StepList,
    StepPanel,
    StepPanels,
    Step,
    StepItem,
    StepperForm,
  },
  data() {
    return {
      rawJson: "", // Holds the raw JSON input
      formattedJson: "", // Holds the formatted JSON output
      error: "", // Error message if the JSON is invalid
      val: null,
    };
  },
  methods: {
    changeValue(e) {
      // console.log("Event value chnaged", e);
      try {
        this.val = JSON.parse(this.rawJson);
        this.error = null;
      } catch (e) {
        if (this.rawJson.length > 0) {
          this.error = "Invalid JSON: " + e.message; // Show the error message
          console.log("errorororo", this.error);
          this.formattedJson = ""; // Clear the formatted JSON output
        } else {
          this.error = null;
        }
      }
    },
  },
};
</script>

<style lang="scss">
.jv-container.jv-light {
  min-width: 50rem;
}
.p-textarea {
  width: 50rem;
  padding: 13px !important;
}
</style>
