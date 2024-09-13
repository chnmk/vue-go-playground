<template>
  <div class="wrapper">
    <h1 class="message">
      vue-go-playground
    </h1>
    <DataReader />
    <!--
      <HelloGraphQL />
    -->
  </div>
</template>

<script setup lang="ts">
import DataReader from './DataReader.vue'
// @ts-ignore
// import HelloGraphQL from './HelloGraphQL.vue';
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