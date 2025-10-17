<script setup>

const route = useRoute();

const routeId = route.params.id;
const travelName = "Impressionisme"
const travelCountry = "France"
const travelStartDate = "10/10/2025"
const travelBudget = 2000


</script>

<template>

  <MenusSide/>
  <div class="m-8 p-6 w-full h-full min-h-screen rounded-xl bg-black/20 backdrop-blur-sm">
    <div class="mb-4 flex flex-col">
      <h1 class="text-3xl font-bold text-gray-800 dark:text-gray-200 mb-2 border-b-2 border-secondary pb-2">Détails du voyage : {{ travelName }}, {{travelCountry}} ({{travelStartDate}})</h1>
      <div class="flex">
        <NuxtLink to="/travel" class="text-primary hover:underline hover:text-secondary">
          &larr; Retour à la liste des voyages
        </NuxtLink>
        <NuxtLink to="/travel" class="text-primary hover:text-secondary ml-auto">
          <UIcon name="heroicons-outline:pencil-square" class="w-6 h-6 align-text-bottom" /> 
        </NuxtLink>
      </div>
    </div>

    
    <div class="h-[400px] w-full my-8">
      <TresCanvas
      shadows alpha 
      :clear-alpha="0"
      class="relative z-0">
        <TresPerspectiveCamera
          :position="[-2.2, 2, 0.3]"
          :look-at="[0, 0, 0]"
        />
        <!-- <TresPerspectiveCamera /> -->
        <OrbitControls :enableDamping="true" :dampingFactor="0.05" />
        <TresAmbientLight :intensity="0.1" /> 

        <TresDirectionalLight 
            :position="[5, 5, 5]" 
            :intensity="2"  
            color="#ffffff"
        />

        <TresDirectionalLight 
            :position="[-5, 0, -5]" 
            :intensity="0.8" 
            color="#f555ff"
        />
        <Suspense>
          <TravelGlobe />
        </Suspense>
        <TravelLines
            :start-lat="51.86"
            :start-lon="-90.35"
            :end-lat="45.71"
            :end-lon="-86.01"
        />
      </TresCanvas>
    </div>
    
    <p class="text-lg mb-4">
      Budget : <span class="font-mono bg-gray-800 p-1 rounded">{{ travelBudget }} €</span>
    </p>
    
    <UCard>
    
      <div>
        <p class="text-lg mb-4">
          Ceci est la page de détail pour le projet ayant l'identifiant : 
          <span class="font-mono bg-gray-800 p-1 rounded">{{ routeId }}</span>
        </p>

        <p class="text-gray-600">
          Ici, tu utiliseras l'ID du projet pour charger les données complètes
          depuis ton API ou ton store (titre, dates, photos, journal de bord, etc.).
        </p>
      </div>

      <template #footer>
      </template>
    </UCard>

  </div>
</template>
