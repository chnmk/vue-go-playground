import './assets/main.css'

import { createApp, h } from 'vue'
import App from './App.vue'

import { ApolloClient, InMemoryCache } from '@apollo/client/core'
import { createApolloProvider } from '@vue/apollo-option'
// @ts-ignore
import VueApolloComponents from '@vue/apollo-components'


const cache = new InMemoryCache()
const apolloClient = new ApolloClient({
    uri: 'http://localhost:5173/graphql',
    cache,
})

const apolloProvider = createApolloProvider({
    defaultClient: apolloClient,
})

const app = createApp({
    render: () => h(App),
})

createApp(App).use(apolloProvider).use(VueApolloComponents).mount('#app')