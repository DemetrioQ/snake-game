<template>
  <div class="scores">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "ScoresComponent",
  props: {
    msg: String,
  },
});
</script>
