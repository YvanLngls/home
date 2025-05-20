<template>
  <MenusSide/>
  <div class="p-8 space-y-16 w-full">
    <h1 class="text-2xl font-bold">Voiture</h1>
    <div class="flex flex-col w-full justify-center items-center p-8 space-y-16">
      <div class="hover:scale-105 transition">
        <NuxtLink to="/car/create" class="bg-primary text-white px-6 py-2 rounded hover:bg-secondary transition">
          Ajouter un véhicule
        </NuxtLink>
      </div>
  
      <div>
        Liste des véhicules...
        <CarDetails/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
definePageMeta({
  middleware: 'auth'
});
</script>