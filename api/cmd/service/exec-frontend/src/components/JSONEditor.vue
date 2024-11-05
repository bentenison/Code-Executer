<template>
  <div>
    <JsonViewer :value="val" copyable boxed sort theme="jv-light" />
  </div>
</template>

<script>
import JsonViewer from "vue3-json-viewer";
// if you used v1.0.5 or latster ,you should add import "vue3-json-viewer/dist/index.css"
import "vue3-json-viewer/dist/index.css";
export default {
  components: {
    JsonViewer,
  },
  data() {
    return {
      val: {
        glossary: {
          title: "example glossary",
          GlossDiv: {
            title: "S",
            GlossList: {
              GlossEntry: {
                ID: "SGML",
                SortAs: "SGML",
                GlossTerm: "Standard Generalized Markup Language",
                Acronym: "SGML",
                Abbrev: "ISO 8879:1986",
                GlossDef: {
                  para: "A meta-markup language, used to create markup languages such as DocBook.",
                  GlossSeeAlso: ["GML", "XML"],
                },
                GlossSee: "markup",
              },
            },
          },
        },
      },
    };
  },
};
</script>

<style lang="scss" scoped></style>
