<template>
  <div class="reader-wrapper">
    <h2 class="title">
      Select database:
    </h2>
    <div class="select-box">
      <div class="select-item" @click="selected='REST'">
        REST
      </div>
      <div class="select-item" @click="selected='gRPC'">
        gRPC
      </div>
      <div class="select-item" @click="selected='GraphQL'">
        GraphQL
      </div>
      <div class="select-item" @click="selected='Kafka'">
        Kafka
      </div>
      <div class="select-item" @click="selected='RabbitMQ'">
        RabbitMQ
      </div>
      <div class="select-item" @click="selected='SQLite'">
        SQLite
      </div>
      <div class="select-item" @click="selected='PostgreSQL'">
        PostgreSQL
      </div>
      <div class="select-item" @click="selected='Redis'">
        Redis
      </div>
      <div class="select-item" @click="selected='Elastic'">
        Elastic
      </div>
      <div class="select-item" @click="selected='Clickhouse'">
        Clickhouse
      </div>
      <div class="select-item" @click="selected='R'">
        R
      </div>
    </div>
    <div class="display">
      Current DB: {{ selected }}
    </div>
    <div class="button-wrapper">
      
      <button @click="connect" class="button">
        Connect
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

let selected = ref("Go (REST API)")

function connect() {
  if (selected.value === "GraphQL") {
    alert("GraphQL connection placeholder.")
    return
  }
  
  let url = "/api/" + selected.value

  axios.post(url, {
    message: 'Hello from Vue'
  })
  .then(function (response) {
    console.log(response);
    alert("Success! Check console for details.")
  })
  .catch((error) => {
    window.alert(`The API returned an error: ${error}`);
  })
}
</script>

<style scoped>
.reader-wrapper {
  display: grid;
  justify-content: center;
}

.title {
  text-decoration: underline;
  display: flex;
  justify-content: center;
}

.select-box {
  background-color: #dff8d5;
  border: 2px solid darkgreen;
  padding: 10px 5px 20px 5px;
  min-width: 250px;
  max-height: 350px;
  overflow-y: scroll;
}

.select-item {
  background-color: darkgreen;
  border: 2px solid black;
  padding: 5px 10px 5px 10px;
  border-radius: 10px;
  color: aliceblue;
  margin: 10px;
  transition: .2s;
}

.select-item:hover {
  cursor: pointer;
  background-color: darkolivegreen;
}

.display {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

.button-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}

.button {
  width: 200px;
  background-color: darkgreen;
  color: aliceblue;
  font-size: 1.1rem;
  border-radius: 20px;
  border: 2px solid black;
  padding: 0.8rem 1rem;
  transition: .3s;
}

.button:hover {
  cursor: pointer;
  background-color: darkolivegreen;
}
</style>
