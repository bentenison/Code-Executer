<template>
  <div class="form-container flex flex-column">
    <h3 class="text-center">Programming Language Selection Form</h3>

    <form
      @submit.prevent="handleSubmit"
      class="px-5 py-2 flex flex-column min-w-full"
    >
      <!-- Topic Selection Dropdown -->
      <div class="flex flex-auto gap-5 justify-content-center">
        <div class="form-group flex flex-column gap-2">
          <label for="language">Programming Language</label>
          <Dropdown
            v-model="selectedLanguage"
            :options="languages"
            option-label="label"
            option-value="value"
            placeholder="Select a language"
            id="language"
          />
        </div>
        <div class="flex flex-column gap-2">
          <label for="topic">Topic</label>
          <Dropdown
            v-model="selectedTopic"
            :options="topics"
            option-label="label"
            option-value="value"
            placeholder="Select a topic"
            id="topic"
          />
        </div>
        <div class="flex flex-column gap-2">
          <label for="numberInput">Number</label>
          <InputNumber
            v-model="numberInput"
            id="numberInput"
            :min="0"
            :max="100"
            showButtons
          />
        </div>
      </div>
      <div class="flex mt-3 align-items-center justify-content-center">
        <Button
          class="w-20rem"
          icon="pi pi-search"
          label="Submit"
          @click="handleSubmit"
        ></Button>
      </div>
    </form>

    <div class="flex m-2 w-full justify-content-center">
      <Textarea v-model="finalPrompt" rows="10" cols="30" disabled></Textarea>
    </div>
    <p class="text-center" style="color: var(--p-primary-600)">
      <strong>NOTE: </strong>Copy the above prompt and paste it into chatGPT
    </p>
  </div>
</template>

<script>
import { useCreatorStore } from "../stores/creator";
import { useEditorStore } from "../stores/editor";

export default {
  data() {
    return {
      selectedTopic: null, // Store selected topic
      numberInput: null, // Store number input
      selectedLanguage: null, // Store selected language
      topics: [
        { label: "Strings", value: "strings" },
        { label: "Conditional Statements", value: "conditionals" },
        { label: "Loops", value: "loops" },
        { label: "Arrays", value: "arrays" },
        { label: "Functions", value: "functions" },
        { label: "Object-Oriented Programming", value: "oop" },
        { label: "Error Handling", value: "errorHandling" },
        { label: "Recursion", value: "recursion" },
        { label: "Data Structures", value: "dataStructures" },
        { label: "Algorithms", value: "algorithms" },
        { label: "Databases", value: "databases" },
        { label: "Web Development", value: "webDevelopment" },
        { label: "Mobile Development", value: "mobileDevelopment" },
        { label: "Game Development", value: "gameDevelopment" },
      ],
      languages: [
        { label: "JavaScript", value: "javascript" },
        { label: "Python", value: "python" },
        { label: "Java", value: "java" },
        { label: "C++", value: "cpp" },
        { label: "Ruby", value: "ruby" },
        { label: "C#", value: "csharp" },
        { label: "C", value: "c" },
        { label: "PHP", value: "php" },
        { label: "Swift", value: "swift" },
        { label: "Go", value: "go" },
        { label: "Rust", value: "rust" },
        { label: "Kotlin", value: "kotlin" },
        { label: "TypeScript", value: "typescript" },
        { label: "R", value: "r" },
        { label: "Lua", value: "lua" },
        { label: "Perl", value: "perl" },
        { label: "Scala", value: "scala" },
        { label: "Dart", value: "dart" },
        { label: "SQL", value: "sql" },
        { label: "Haskell", value: "haskell" },
        { label: "MATLAB", value: "matlab" },
        { label: "Shell Script", value: "shell" },
        { label: "VHDL", value: "vhdl" },
        { label: "Assembly", value: "assembly" },
      ],
      formData: null, // Store the submitted form data
      finalPrompt: null,
      editorStore: useEditorStore(),
      creatorStore: useCreatorStore(),
    };
  },
  methods: {
    handleSubmit() {
      // Collect form data when the form is submitted
      this.formData = {
        topic: this.selectedTopic,
        number: this.numberInput,
        language: this.selectedLanguage,
      };
      let sampleQuest = ``;
      for (
        let index = 0;
        index < this.editorStore.questionTemplates.length;
        index++
      ) {
        const element = this.editorStore.questionTemplates[index];
        // console.log("element----------------",this.selectedLanguage)
        if (element.language === this.selectedLanguage) {
          console.log("element----------------", element);
          sampleQuest = JSON.stringify(element, null, 2);
        }
      }
      //   console.log("FormDATA>>>>>>>>>>>>>", this.formData);
      let part = `You are tasked with generating ${this.numberInput} programming question and its corresponding answer in JSON format for ${this.selectedLanguage} language. The topic to generate the questions on is ${this.selectedTopic}. The question and answer should contain the following fields:`;

      this.finalPrompt =
        part +
        " " +
        sampleQuest +
        " " +
        `You are tasked with generating programming questions in JSON format based on the following structure:

    Title: A descriptive title of the question.

    Description: A brief explanation of what the user needs to implement.

    Input:
        Description: Description of the input format.
        Expected: The type of output that the function should return.

    Output:
        Description: A description of what the function should output.

    Template Code: A code template for the function with placeholders for user logic.

    Language: The programming language used for the question.

    Language Code: The language code (e.g., py for Python, java for Java, etc.).

    Difficulty: The difficulty level of the question (e.g., easy, medium, hard).

    Tags: Relevant tags for the question (e.g., leap year, beginner, conditionals).

    User Logic Template:
        Description: Provide a description of where the user should place their code logic.
        Code: A code snippet where the user can add their implementation. The main function should not already contain any logic. Instead, use placeholders to mark where the user should implement their solution. If the language is Python, for example, use if __name__ == '__main__': to run the user's function.
        CodeRunTemplate : A skeleton with a signle test case and placeholder for user function

    Test Case Template:
        Description: Description of how the user can validate the function with test cases.
        Code: A skeleton code that allows the user to run multiple test cases to validate the logic.
    
    Testcases: A list of test cases with inputs and expected outputs.
    ExecTemplate: A template with placeholders for UserLogic and TestCases
    Answer:
        ID: A unique identifier for the solution.
        Logic: The correct solution for the problem.
        Created At: The timestamp of when the answer was created.
        Updated At: The timestamp of when the answer was last updated.
        Testcases: A list of test cases that were used to validate the solution
the testcase_template.code should be a function where define one variable loop trough the testcases and pass input to the user defined function and compare the function output to extectedOutput if all testcases matches the expectedOutput then return true else return false. 
The code in answer.logic and testcase_template.code make sure for every question it should be correct.
`;
    },
  },
  mounted() {
    this.creatorStore
      .getAllConcepts()
      .then((res) => {})
      .catch((err) => {
        this.$toast.add({
          severity: "error",
          summary: "error in getting programming concepts",
          detail: err,
          life: 3000,
        });
      });
  },
};
</script>

<style scoped>
.form-container {
  width: 80rem;
  /* margin: 50px auto;
    padding: 20px; */
  /* border: 1px solid #ddd; */
  border-radius: 8px;
  background-color: var(--p-surface-card);
}

/* .form-group {
  margin-bottom: 20px;
} */

h2 {
  text-align: center;
}

button {
  width: 100%;
}
</style>
