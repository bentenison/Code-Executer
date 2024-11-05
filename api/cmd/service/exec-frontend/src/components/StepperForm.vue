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
      <Textarea v-model="finalPrompt" rows="10" cols="30"></Textarea>
    </div>
  </div>
</template>

<script>
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
      console.log("FormDATA>>>>>>>>>>>>>", this.formData);
      let part = `You are tasked with generating ${this.numberInput} programming question and its corresponding answer in JSON format for ${this.selectedLanguage} languages. The topic to generate the questions on is ${this.selectedTopic}. The question and answer should contain the following fields:`;
      let sampleQuest = `{
  "title": "Check Leap Year",
  "description": "Write a function that determines whether a given year is a leap year.",
  "input": {
    "description": "You will receive a single integer year as input.",
    "expected": "The function should return a boolean value, True if the year is a leap year, False otherwise."
  },
  "output": {
    "description": "The output will be a boolean indicating whether the year is a leap year."
  },
  "template_code": "def main(year):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (2020,),\n        (1900,),\n        (2000,),\n        (2024,)\n    ]\n    expected_outputs = [\n        True,\n        False,\n        True,\n        True\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "leap year",
    "beginner"
  ],
  "user_logic_template": {
    "description": "Insert your logic below:",
    "code": "def main(year):\n    # Your code here\n\nif __name__ == '__main__':\n    main(2020)"
  },
  "testcase_template": {
    "description": "You can use the following test case structure to validate your function:",
    "code": "if __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (2020,),\n        (1900,),\n        (2000,),\n        (2024,)\n    ]\n    expected_outputs = [\n        True,\n        False,\n        True,\n        True\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)"
  },
  "testcases": [
    {
      "input": 2020,
      "expectedOutput": true
    },
    {
      "input": 1900,
      "expectedOutput": false
    },
    {
      "input": 2000,
      "expectedOutput": true
    },
    {
      "input": 2024,
      "expectedOutput": true
    }
  ],
  "answer": {
    "id": "2",
    "logic": "def main(year):\n    if (year % 4 == 0 and year % 100 != 0) or (year % 400 == 0):\n        return True\n    return False",
    "created_at": "2023-10-20T00:00:00Z",
    "updated_at": "2023-10-20T00:00:00Z",
    "testcases": [
      {
        "input": 2020,
        "expectedOutput": true
      },
      {
        "input": 1900,
        "expectedOutput": false
      },
      {
        "input": 2000,
        "expectedOutput": true
      },
      {
        "input": 2024,
        "expectedOutput": true
      }
    ]
  },
  "id": "2"
}
`;
      this.finalPrompt =
        part +
        " " +
        sampleQuest +
        " " +
        "follow the format correctly and generate questions accurately by analyzing the given sample format.user_logic_template.code should contain if __name__ == '__main__': to run the user function for debugging before submitting. There should not be the solution written already in the main function. The output should be single file.";
    },
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
