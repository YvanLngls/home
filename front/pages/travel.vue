<template>
  <MenusSide/>
  <div class="p-4 my-8 w-full">
    
    <TravelSectionCard 
      title="Prochaines Destinations"
      projectId="tokyo-trip-2026"
      routeName="">
        <p class="text-4xl font-extrabold text-white">Londres, Angleterre</p>
        <p class="text-gray-300">Countdown details...</p>
    </TravelSectionCard>

    <TravelSectionCard title="Histoires de Voyages">
      <p class="text-4xl font-extrabold text-white">Impressionisme, France</p>
      <p class="text-gray-300">Countdown details...</p>
    </TravelSectionCard>

  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
});
</script>
