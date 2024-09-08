<template>
  <div class="wrapper">
    <h1 class="message">
      Hello, Vue!
    </h1>
    <HelloButton />
    <!--
      <HelloGraphQL />
    -->
  </div>
</template>

<script setup lang="ts">
import HelloButton from './HelloButton.vue'
// @ts-ignore
import HelloGraphQL from './HelloGraphQL.vue';
// import DataDisplay from './DataDisplay.vue'
</script>

<style scoped>
.wrapper {
  background-color: aliceblue;
  width: 100vw;
  height: 100vh;
}
.message {
  text-align: center;
  color: darkgreen;
  margin-top: 0;
  padding-top: 1rem;
  font-weight: bold;
}
</style>