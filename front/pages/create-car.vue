<template>
  <MenusSide/>
  <div class="p-8 space-y-16 w-full">
    <h1 class="text-2xl font-bold">Ajouter un véhicule</h1>
    <div class="flex flex-col w-full justify-center items-center p-8 space-y-16">
      <CarCreate/>
    </div>
  </div>
</template>

<script lang="ts">
definePageMeta({
  middleware: 'auth'
});

</script>