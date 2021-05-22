<template>
  <div class="scores">
    <Scores msg="Scores" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Scores from "@/components/ScoresComponent.vue"; // @ is an alias to /src

export default Vue.extend({
  name: "Scores",
  components: {
    Scores,
  },
});
</script>
