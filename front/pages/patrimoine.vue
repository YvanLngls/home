<template>
  <MenusSide/>
  <div class="p-8">
    <h1 class="text-white">Bienvenue</h1>
  </div>
</template>

<script lang="ts">
definePageMeta({
  middleware: 'auth'
});

</script>