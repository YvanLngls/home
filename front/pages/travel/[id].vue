<!-- <template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 to-blue-900">
    <UContainer class="py-8">
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-white mb-2">
          My Travel Journey
        </h1>
        <p class="text-blue-200">
          Exploring the world, one destination at a time
        </p>
      </div>

      <div class="grid lg:grid-cols-3 gap-8">

        Travel List
        <div class="space-y-6">
          <UCard class="bg-white/5 backdrop-blur-sm border-white/10">
            <template #header>
              <h2 class="text-xl font-semibold text-white">Travel Stats</h2>
            </template>

            <div class="grid grid-cols-2 gap-4 text-center">
              <div class="p-4 bg-blue-500/10 rounded-lg">
                <div class="text-2xl font-bold text-blue-300">{{ travelLocations.length }}</div>
                <div class="text-blue-200 text-sm">Countries Visited</div>
              </div>
              <div class="p-4 bg-green-500/10 rounded-lg">
                <div class="text-2xl font-bold text-green-300">{{ totalCities }}</div>
                <div class="text-green-200 text-sm">Cities Explored</div>
              </div>
            </div>
          </UCard>

          <UCard class="bg-white/5 backdrop-blur-sm border-white/10">
            <template #header>
              <h2 class="text-xl font-semibold text-white">Recent Travels</h2>
            </template>

            <div class="space-y-3">
              <div
                v-for="location in recentLocations"
                :key="location.id"
                class="flex items-center space-x-3 p-3 rounded-lg hover:bg-white/5 transition-colors"
              >
                <div class="w-3 h-3 rounded-full bg-red-400"></div>
                <div class="flex-1">
                  <div class="text-white font-medium">{{ location.city }}</div>
                  <div class="text-blue-200 text-sm">{{ location.country }}</div>
                </div>
                <div class="text-white/60 text-sm">
                  {{ location.date }}
                </div>
              </div>
            </div>
          </UCard>
        </div>
      </div>

      Add Travel Form
      <UCard class="mt-8 bg-white/5 backdrop-blur-sm border-white/10">
        <template #header>
          <h2 class="text-xl font-semibold text-white">Add New Travel</h2>
        </template>

        <form @submit.prevent="addNewLocation" class="space-y-4">
          <div class="grid md:grid-cols-4 gap-4">
            <UFormGroup label="City" class="text-white">
              <UInput v-model="newLocation.city" placeholder="Enter city name" />
            </UFormGroup>

            <UFormGroup label="Country" class="text-white">
              <UInput v-model="newLocation.country" placeholder="Enter country" />
            </UFormGroup>

            <UFormGroup label="Latitude" class="text-white">
              <UInput v-model.number="newLocation.lat" type="number" step="any" placeholder="e.g., 40.7128" />
            </UFormGroup>

            <UFormGroup label="Longitude" class="text-white">
              <UInput v-model.number="newLocation.lng" type="number" step="any" placeholder="e.g., -74.0060" />
            </UFormGroup>
          </div>

          <div class="flex justify-end">
            <UButton type="submit" icon="i-heroicons-plus" class="bg-blue-500 hover:bg-blue-600">
              Add Location
            </UButton>
          </div>
        </form>
      </UCard>
    </UContainer>
  </div>
</template> -->

<script setup>

const route = useRoute();

const routeId = route.params.id;
const travelName = "Impressionisme"
const travelCountry = "France"
const travelStartDate = "10/10/2025"
const travelBudget = 2000

// Sample travel data
const travelLocations = ref([
  { id: 1, city: 'New York', country: 'USA', lat: 40.7128, lng: -74.0060, date: '2024-01-15' },
  { id: 2, city: 'London', country: 'UK', lat: 51.5074, lng: -0.1278, date: '2024-02-20' },
  { id: 3, city: 'Tokyo', country: 'Japan', lat: 35.6762, lng: 139.6503, date: '2024-03-10' },
  { id: 4, city: 'Paris', country: 'France', lat: 48.8566, lng: 2.3522, date: '2024-04-05' },
  { id: 5, city: 'Sydney', country: 'Australia', lat: -33.8688, lng: 151.2093, date: '2024-05-12' }
])

const newLocation = ref({
  city: '',
  country: '',
  lat: null,
  lng: null
})

// Computed properties
const totalCities = computed(() => travelLocations.value.length)

const recentLocations = computed(() => 
  [...travelLocations.value]
    .sort((a, b) => new Date(b.date) - new Date(a.date))
    .slice(0, 5)
)

// Methods
function addNewLocation() {
  if (!newLocation.value.city || !newLocation.value.country || !newLocation.value.lat || !newLocation.value.lng) {
    return
  }

  travelLocations.value.push({
    id: Date.now(),
    city: newLocation.value.city,
    country: newLocation.value.country,
    lat: newLocation.value.lat,
    lng: newLocation.value.lng,
    date: new Date().toISOString().split('T')[0]
  })

  // Reset form
  newLocation.value = {
    city: '',
    country: '',
    lat: null,
    lng: null
  }
}
</script>

<template>

  <MenusSide/>
  <div class="m-8 p-6 w-full h-full rounded-xl bg-black/20 backdrop-blur-sm">
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

    <p class="text-lg mb-4">
      Budget : <span class="font-mono bg-gray-800 p-1 rounded">{{ travelBudget }} €</span>
    </p>
    
    <!-- Globe Section -->
    <div class="lg:col-span-2">
      <div class="bg-black/20 rounded-xl p-4 backdrop-blur-sm border border-white/10">
        <div class="h-96 lg:h-[500px] rounded-lg overflow-hidden">
          <TravelGlobe :locations="travelLocations" />
        </div>
      </div>
    </div>

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
