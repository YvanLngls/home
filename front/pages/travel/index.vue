<template>
  <MenusSide/>
  <div class="p-4 m-8 w-full">
    
    <div v-for="section in travelSections" :key="section.title" class="mb-12">
      <h2 class="text-3xl font-bold text-gray-800 dark:text-gray-200 mb-6 border-b-2 border-secondary pb-2">
        {{ section.title }}
      </h2>

      <div class="flex flex-wrap -m-2 h-96 overflow-y-auto border-2 border-primary/50 rounded-lg p-2"> 
        
        <div 
          v-for="trip in section.trips" 
          :key="trip.projectId"
          class="w-full sm:w-1/2 p-2"
        >
          <TravelSectionCard 
            :title="section.title" 
            :projectId="trip.projectId" 
            :date="trip.date"
            :routeName="trip.routeName"
            :headerColor="section.color"
          >
            <p class="text-4xl font-extrabold text-white">{{ trip.city }}, {{ trip.country }}</p>
            <p class="text-gray-300">{{ trip.details }}</p>
          </TravelSectionCard>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

// 💡 NOTE : Dans une application réelle, ces données viendraient d'une API ou d'un Store (Pinia/Vuex).
const travelSections = ref([
  {
    title: "Prochaines Destinations",
    color: "bg-tertiary", // Couleur de l'en-tête de la carte
    trips: [
      {
        city: "Londres",
        country: "Angleterre",
        details: "Countdown details...",
        projectId: "london-trip-2026",
        date:"14/02/2026",
        routeName: "travel" // Nom de la route de destination (à vérifier)
      }
    ]
  },
  {
    title: "Carnets de voyage",
    color: "bg-primary",
    trips: [
      {
        city: "Impressionisme",
        country: "France",
        details: "J'y étais en 2023.",
        projectId: "france-2023",
        date:"08/08/2025",
        routeName: "travel"
      },
      {
        city: "Rome",
        country: "Italie",
        details: "Un souvenir inoubliable.",
        projectId: "rome-2022",
        date:"12/05/2024",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      },
      {
        city: "Tokyo",
        country: "Japon",
        details: "Dans 6 mois !",
        projectId: "tokyo-trip-2026",
        date:"04/06/2026",
        routeName: "travel"
      }
    ]
  }
]);

definePageMeta({
  middleware: 'auth'
});
</script>