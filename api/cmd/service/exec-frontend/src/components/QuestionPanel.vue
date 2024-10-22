<template>
  <div>
    <div class="flex align-items-center justify-content-start gap-2 mb-2">
      <Avatar
        image="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
        shape="circle"
      />
      <span class="font-bold">Amy Elsner</span>
    </div>
    <div class="flex items-center gap-2">
      <Button
        label="Account"
        icon="pi pi-user"
        class="flex-auto"
        severity="success"
        text
      ></Button>
      <Button
        label="Logout"
        icon="pi pi-sign-out"
        severity="danger"
        text
      ></Button>
    </div>
    <hr class="bg-surface-100" />
    <div class="question-pallete flex flex-column">
      <div
        class="flex align-items-center justify-content-between max-h-30rem pb-2 overflow-y-scroll gap-2 flex-wrap w-full"
      >
        <Button
          class="w-3rem h-3rem shadow-4"
          v-for="index in 50"
          :key="index"
          severity="success"
        >
          <span class="text-muted">{{ index }}</span>
        </Button>
      </div>
      <div class="pallete-footer flex align-items-center mt-5">
        <!-- <Button label="Account" icon="pi pi-user" class="flex-auto" severity="success" text></Button> -->
        <Button
          type="button"
          label="End Exam"
          icon="pi pi-sign-out"
          class="flex-auto"
          severity="danger"
          raised
        ></Button>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped></style>
