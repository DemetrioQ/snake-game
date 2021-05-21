<template>
  <div class="scores">
    <ScoresComponent msg="Scores" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ScoresComponent from "@/components/ScoresComponent.vue"; // @ is an alias to /src

export default Vue.extend({
  name: "Scores",
  components: {
    ScoresComponent,
  },
});
</script>
