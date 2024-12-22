<template>
  <div
    class="grid container flex align-items-center justify-content-center w-full"
  >
    <div
      class="col-4 flex align-items-center flex-wrap"
    style="width: 30rem; height: 20rem"
      v-for="q in creatorStore.questions"
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
                <Chip
                  :label="q.language"
                  class="text-sm"
                  icon="pi pi-language"
                  style="
                    background-color: var(--p-green-100);
                    color: var(--p-green-600);
                    border: 1px solid var(--p-green-600);
                  "
                />
                <Chip
                  :label="q.difficulty"
                  class="text-sm"
                  icon="pi pi-thumbtack"
                  style="
                    background-color: var(--p-orange-100);
                    color: var(--p-orange-600);
                    border: 1px solid var(--p-orange-600);
                  "
                />
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
            <Button label="QC" />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import { useCreatorStore } from "../stores/creator";
export default {
  components: {},
  data() {
    return {
      creatorStore: useCreatorStore(),
    };
  },
};
</script>

<style lang="scss" scoped></style>
