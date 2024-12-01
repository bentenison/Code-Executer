<template>
  <div
    class="container flex flex-column align-items-center justify-content-center"
    style="overflow: hidden;min-height: 80vh;"
  >
    <div class="text-center mb-5">
      <h2 class="section-title">
        Our <span class="section-subtitle">interactive</span> features
      </h2>
    </div>
    <div class="grid justify-content-center">
      <Card class="col-3 feature-card feature-card-1">
        <template #content>
          <div class="feature-icon">
            <i class="pi pi-question-circle" style="font-size: 3rem"></i>
          </div>
          <h3>Fun <span class="highlight">Quiz</span></h3>
          <p>Test your understanding with a short but fun quizzes!</p>
          <div class="card-footer flex justify-content-end">
            <Button label="Start" severity="contrast" variant="text"></Button>
          </div>
        </template>
      </Card>
      <Card class="col-3 feature-card feature-card-2">
        <template #content>
          <div class="feature-icon">
            <i class="pi pi-chart-line" style="font-size: 3rem"></i>
          </div>
          <h3>Creative <span class="highlight">Practice Activities</span></h3>
          <p>
            Discover enjoyable activities such as coloring, crafting, and
            science.
          </p>
          <div class="card-footer flex justify-content-end">
            <Button label="Start" severity="contrast" variant="text"></Button>
          </div>
        </template>
      </Card>
      <Card class="col-3 feature-card feature-card-3">
        <template #content>
          <div class="feature-icon">
            <i class="pi pi-bullseye" style="font-size: 3rem"></i>
          </div>
          <h3>Learn with <span class="highlight">Challanges</span></h3>
          <p>Learn something new while your kids playing games!</p>
          <div class="card-footer flex justify-content-end">
            <Button label="Start" severity="contrast" variant="text"></Button>
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.section-title {
  font-family: "Pacifico", cursive;
  font-size: 4rem;
  color: #6a0dad;
}
.section-subtitle {
  font-size: 4rem;
  color: var(--p-text-color);
}
.feature-card {
  border-radius: 20px;
  padding: 20px;
  text-align: center;
  color: #fff;
  margin: 10px;
}
.feature-card h3 {
  font-size: 1.5rem;
}
.feature-card p {
  font-size: 1rem;
}
.feature-card .highlight {
  font-family: "Pacifico", cursive;
  font-size: 2rem;
}
.feature-card-1 {
  background-color: #adb5bd;
}
.feature-card-2 {
  background-color: #7f00ff;
}
.feature-card-3 {
  background-color: #ffcc00;
}
.feature-icon {
  //   font-size: 3rem;
  margin-bottom: 10px;
}
</style>
