<template>
  <div id="game-app">
    <div id="game">
      <GameComponent />
    </div>
    <div id="border" />
  </div>
</template>

<script>
import Vue from "vue";
import GameComponent from "@/components/GameComponent.vue";

export default Vue.extend({
  name: "Game",
  components: {
    GameComponent,
  },
});
</script>

<style>
#game-app {
  font-family: "Trebuchet MS", "Lucida Sans Unicode", "Lucida Grande", "Lucida Sans", Arial, sans-serif;
  text-align: left;
  backdrop-filter: black;
  color: cyan;
  display: flex;
}
#game {
  width: 50ww;
  height: 100vh;
}
#border {
  border-right: 2px solid cyan;
}
</style>