<template>
  <div
    class="min-h-screen min-w-100vw overflow-hidden flex flex-column align-items-center justify-content-center"
  >
    <h3>Question Types</h3>
    <code-display
      :question="questions[0].question"
      :codeTemplate="questions[0].codeTemplate"
      :blanks="questions[0].blanks"
    />
    <div v-html="formattedCodeTemplate"></div>
    <div id="options" class="options"></div>
    <div id="result" class="result"></div>
    <CodeEditor
      id="editor"
      ref="editor"
      :line-nums="true"
      :key="3"
      :v-html="formattedCodeTemplate"
      :theme="theme"
      width="120rem"
      height="500px"
      lang-list-height="300px"
      :fontSize="fontSize"
      @content="getContent"
      @lang="getLanguage"
      :languages="[['java', 'JAVA']]"
    ></CodeEditor>
  </div>
</template>

<script>
import CodeEditor from "../SimpleCodeEditor/CodeEditor.vue";
import CodeDisplay from "../components/qt/CodeDisplay.vue";
export default {
  components: { CodeEditor, CodeDisplay },
  data() {
    return {
      questions: [
        {
          question: "Fill in the blanks in the following code.",
          codeTemplate: `public  {{BLANK_1}} Example {
      public static void {{BLANK_2}}(String[] args) {
          System.out.println("Hello, and !");
      }
  }`,
          blanks: [
            {
              id: "BLANK_1",
              options: ["World", "Java", "Everyone"],
              correct: "World",
            },
            {
              id: "BLANK_2",
              options: ["Students", "Developers", "Coders"],
              correct: "Students",
            },
          ],
        },
      ],
      codeWithBlanks: null,
    };
  },
  computed: {
    formattedCodeTemplate() {
      const question = this.questions[0];
      let formattedTemplate = question.codeTemplate;

      question.blanks.forEach((blank) => {
        const userInput = "______";
        formattedTemplate = formattedTemplate.replace(
          `{{${blank.id}}}`,
          `<span class="blank" data-id="${blank.id}">${userInput}</span>`
        );
      });

      return formattedTemplate;
    },
  },
  methods: {
    loadProgress() {
      const progress = JSON.parse(localStorage.getItem("progress")) || {};
      return progress[0] || {};
    },

    saveProgress(blankId, value) {
      const progress = JSON.parse(localStorage.getItem("progress")) || {};
      if (!progress[0]) {
        progress[0] = {};
      }
      progress[0][blankId] = value;
      localStorage.setItem("progress", JSON.stringify(progress));
    },

    showOptions(event) {
      const blankId = event.target.dataset.id; // Get the blank ID from the data attribute
      const rect = event.target.getBoundingClientRect();
      const optionsContainer = document.getElementById("options");
      optionsContainer.style.display = "block";
      optionsContainer.style.top = `${rect.bottom + window.scrollY}px`;
      optionsContainer.style.left = `${rect.left}px`;

      const options = this.questions[0].blanks.find(
        (b) => b.id === blankId
      ).options;
      optionsContainer.innerHTML = ""; // Clear previous options
      options.forEach((option) => {
        const optionDiv = document.createElement("div");
        optionDiv.innerText = option;
        optionDiv.onclick = () => this.fillBlank(blankId, option);
        optionsContainer.appendChild(optionDiv);
      });
    },

    fillBlank(blankId, value) {
      const blank = document.querySelector(`.blank[data-id="${blankId}"]`);
      if (blank) {
        blank.innerText = value; // Fill in the blank
        this.saveProgress(blankId, value); // Save progress
      }
      document.getElementById("options").style.display = "none"; // Hide options
      this.validateAnswers();
    },

    validateAnswers() {
      const question = this.questions[0];
      const userAnswers = this.loadProgress();
      const resultDiv = document.getElementById("result");
      const allCorrect = question.blanks.every(
        (blank) => userAnswers[blank.id] === blank.correct
      );

      resultDiv.innerText = allCorrect
        ? "All answers are correct!"
        : "Some answers are incorrect.";
    },
  },
  mounted() {
    //   this.renderCode();
    // Attach click event to the blanks after rendering
    this.$nextTick(() => {
      const blanks = document.querySelectorAll(".blank");
      blanks.forEach((blank) => {
        blank.addEventListener("click", this.showOptions);
      });
    });
  },
};
</script>

<style lang="scss" scoped>
.blank {
  border-bottom: 1px dashed #000;
  cursor: pointer;
  display: inline-block;
}
.options {
  display: none;
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  z-index: 10;
}
.options div {
  padding: 5px;
  cursor: pointer;
}
.options div:hover {
  background-color: #f0f0f0;
}
.result {
  margin-top: 20px;
  font-weight: bold;
}
</style>
