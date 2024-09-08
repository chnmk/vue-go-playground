<template>
  <ApolloMutation
    :mutation="gql => gql`
      mutation SayHello ($hello: String!) {
        hello (name: $name)
      }
    `"
    :variables="{ hello }"
  >
    <template v-slot="{ mutate, loading, error }">
      <div class="button-wrapper">
        <button :disabled="loading" @click="mutate()" class="hello">Click me</button>
        <p v-if="error">An error occurred: {{ error }}</p>
      </div>
    </template>
  </ApolloMutation>
</template>

<script>
const data = "Hello from Vue through GraphQL!"
</script>

<style scoped>
.button-wrapper {
  margin-top: 1rem;
  display: flex;
  justify-content: center;
}

.hello {
  background-color: aqua;
  border-radius: 20px;
  border: 1px solid gray;
  padding: 0.8rem 1rem;
  transition: .3s;
}

.hello:hover {
  cursor: pointer;
  background-color: aquamarine;
}
</style>