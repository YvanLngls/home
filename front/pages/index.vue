<template>

<h1>Bienvenue</h1>

<h3 @click="login()">Login</h3>
</template>

<script setup lang="ts">
const login = () => {
  navigateTo('/login')
};
</script>