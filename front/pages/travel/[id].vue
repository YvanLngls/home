<template>
  <MenusSide/>
  <div class="p-8">
    
    <UCard>
      <template #header>
        <h1 class="text-3xl font-bold">Détails du voyage : {{ routeId }}</h1>
      </template>

      <div>
        <p class="text-lg mb-4">
          Ceci est la page de détail pour le projet ayant l'identifiant : 
          <span class="font-mono bg-gray-100 p-1 rounded">{{ routeId }}</span>
        </p>

        <p class="text-gray-600">
          Ici, tu utiliseras l'ID du projet pour charger les données complètes
          depuis ton API ou ton store (titre, dates, photos, journal de bord, etc.).
        </p>
      </div>

      <template #footer>
        <NuxtLink to="/travel" class="text-primary hover:underline">
          &larr; Retour à la liste des voyages
        </NuxtLink>
      </template>
    </UCard>

  </div>
</template>

<script setup>
// 1. Importer le hook de routage de Nuxt pour accéder aux paramètres de l'URL
const route = useRoute();

// 2. Récupérer le paramètre 'id' qui correspond au nom du fichier ([id].vue)
const routeId = route.params.id;

// Tu peux ajouter des métadonnées spécifiques à la page si tu le souhaites
useHead({
  title: `Voyage - ${routeId}`
})

// 💡 ÉTAPE SUIVANTE : 
// Ici, tu mettrais ta logique pour charger les détails du voyage :
// const { data: tripDetails } = await useAsyncData(
//   `trip-${routeId}`,
//   () => $fetch(`/api/trips/${routeId}`) 
// );
</script>