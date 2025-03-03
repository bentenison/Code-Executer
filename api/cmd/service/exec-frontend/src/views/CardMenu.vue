<template>
  <div
    class="flex flex-column align-items-center justify-content-center"
    style="min-height: 100vh"
  >
    <h2>Community Cards</h2>
    <div class="card-container">
      <Card
        class="card"
        v-for="language in editorStore.languages"
        :key="language.id"
      >
        <template #header>
          <img
            alt="SuperRare image"
            class="card-img-top"
            height="50"
            :src="getImagePath(language.name.toUpperCase() + '.svg')"
            width="300"
          />
        </template>
        <template #content>
          <h5 class="card-title">{{ language.name }}</h5>
          <p class="card-text">
            {{ language.description }}
          </p>
        </template>
        <template #footer>
          <div class="flex justify-content-end">
            <div class="flex align-items-center gap-3">
              <span>
                <i class="pi pi-heart" style="color: var(--p-red-500)"></i>
                299.5K
              </span>
              <span>
                <i class="pi pi-users" style="color: var(--p-slate-500)"></i>
                71.4K
              </span>
            </div>
            <!-- <div class="flex gap-2">
              <Tag severity="secondary" value="Secondary">ART</Tag>
              <Tag severity="secondary" value="Secondary">Crypto</Tag>
              <Tag severity="secondary" value="Secondary">NFT</Tag>
            </div> -->
          </div>
          <!-- {{ language }} -->
          <div class="flex mt-2 justify-content-end flex-wrap gap-3">
            <Button
              label="Documentation"
              severity="info"
              @click="openNewWindow(language.documentation_url)"
              text
            />
            <Button
              label="Challenges"
              variant="text"
              text
              icon="pi pi-check"
              iconPos="right"
              @click="createChallenge(language)"
            />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import { useChallengeStore } from "../stores/challenges";
import { useEditorStore } from "../stores/editor";

export default {
  data() {
    return {
      editorStore: useEditorStore(),
      challengeStore: useChallengeStore(),
    };
  },
  methods: {
    createChallenge(language) {
      this.challengeStore.selectedLanguage = language;
      let payload = {
        user_id: "51fc3552-45e0-4982-9adb-50d8cc46c46d",
        username: "tanvirs",
        selected_language: language.name.toLowerCase(),
      };
      this.challengeStore
        .prepareChallenges(payload)
        .then((res) => {
          this.challengeStore
            .createChallenges(payload)
            .then((res) => {
              this.$router.push("/challengeViewer")
            })
            .catch((err) => {
              this.$toast.add({
                severity: "error",
                summary: "unable to create challenges",
                detail: err,
                life: 3000,
              });
            });
        })
        .catch((err) => {
          this.$toast.add({
            severity: "error",
            summary: "unable to prepare challenge",
            detail: err,
            life: 3000,
          });
        });
    },
    openNewWindow(url) {
      // Open the URL in a new window/tab
      window.open(url, "_blank");
    },
    getImagePath(fileName) {
      //   console.log("filemane", fileName);
      if (fileName.includes("#")) {
        fileName = "CSHARP.svg";
        return `/img/${fileName}`;
      }
      return `/img/${fileName}`;
    },
  },
};
</script>

<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
  align-items: center;
}
.card {
  border: none;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 350px;
}
.card img {
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  width: 100%;
  height: 120px;
  object-fit: cover;
}

.tags span {
  background-color: #e9ecef;
  border-radius: 5px;
  padding: 2px 5px;
  font-size: 12px;
}
</style>
