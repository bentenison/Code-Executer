<template>
    <div class="fill-in-the-blanks">
      <h2>{{ question }}</h2>
      <pre class="code-block">
        <span v-html="formattedCode"></span>
      </pre>
      <div class="selects">
        <div v-for="(blank, index) in blanks" :key="index" class="select-group">
          <label :for="blank.id">{{ `Select ${blank.id}` }}</label>
          <select
            :id="blank.id"
            v-model="userSelections[blank.id]"
            @change="updateCode"
          >
            <option value="" disabled>Select an option</option>
            <option v-for="option in blank.options" :key="option" :value="option">
              {{ option }}
            </option>
          </select>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      question: {
        type: String,
        required: true,
      },
      codeTemplate: {
        type: String,
        required: true,
      },
      blanks: {
        type: Array,
        required: true,
      },
    },
    data() {
      return {
        userSelections: this.blanks.reduce((acc, blank) => {
          acc[blank.id] = '';
          return acc;
        }, {}),
        formattedCode: '',
      };
    },
    mounted() {
      this.updateCode();
    },
    methods: {
      updateCode() {
        let code = this.codeTemplate;
        // Replace placeholders with user selections
        this.blanks.forEach(blank => {
          const regex = new RegExp(`{{${blank.id}}}`, 'g');
          code = code.replace(regex, this.userSelections[blank.id] || '______');
        });
        // Format the code for HTML display
        this.formattedCode = code.replace(/</g, '&lt;').replace(/>/g, '&gt;');
      }
    },
  };
  </script>
  
  <style scoped>
  .fill-in-the-blanks {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
    background: #f9f9f9;
  }
  .code-block {
    background: #f5f5f5;
    padding: 10px;
    border-radius: 5px;
    overflow-x: auto;
    white-space: pre;
  }
  .select-group {
    margin: 10px 0;
  }
  label {
    display: block;
    margin-bottom: 5px;
  }
  select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  </style>
  