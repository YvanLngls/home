<template>
  <MenusSide/>
  <div class="p-8">
    <h1 class="text-white">Bienvenue</h1>

    <h3>Nom : {{name}}</h3>
    <h3>Date de création : {{creationDate}}</h3>
    <h3>Dernière connexion: {{lastConnectionDate}}</h3>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
});

const name = ref("")
const creation = ref("")
const lastConnection = ref("")

const { data: user } = await useFetch('http://localhost:8080/getUserInfos', {
  method: 'GET',credentials: 'include'
})
let parsedUser = JSON.parse(user.value)
name.value = parsedUser?.username || ''
creation.value = parsedUser?.createdAt || ''
lastConnection.value = parsedUser?.lastConnection || ''

function formatDate(dateString) {
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' };
  const date = new Date(dateString);
  return date.toLocaleDateString('en-UK', options);
}

const creationDate = formatDate(creation.value);
const lastConnectionDate = formatDate(lastConnection.value);
</script>
