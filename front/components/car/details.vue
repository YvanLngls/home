<template>
    <div class="space-y-4 rounded-lg backdrop-filter backdrop-blur-xs bg-black/30 shadow-lg drop-shadow-xl p-8 pr-64">

        <div v-if="loading">
            <p>Loading car information...</p>
        </div>
        <div v-else-if="error">
            <p>Error: {{ error }}</p>
        </div>
        <div v-else-if="car">
            <div>
                <h2 class="text-xl font-bold">Informations générales</h2>
                <h3>Marque de la voiture : {{car.brand}}</h3>
                <h3>Modèle de la voiture : {{car.model}}</h3>
            </div>
        
            <div>
                <h2 class="text-xl font-bold">Entretien et maintenance</h2>
                <h3>todo</h3>
            </div>
            <div>
                <h2 class="text-xl font-bold">Assurance et documents</h2>
                <h3>todo</h3>
            </div>
            <div>
                <h2 class="text-xl font-bold">Carburant et consommation</h2>
                <h3>todo</h3>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

// Define the type for the car object
interface Car {
  id: string;
  brand: string;
  model: string;
}

const car = ref<Car | null>(null);
const loading = ref<boolean>(true);
const error = ref<string | null>(null);

onMounted(async () => {
    try {
        const response = await fetch(`http://localhost:8080/car/get?id=1`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include'
            });
        console.log(response)
        if (!response.ok) {
        throw new Error('Network response was not ok');
        }
        car.value = await response.json();   
    
    } catch (err) {
        console.error('Error fetching car information:', err);
        error.value = err instanceof Error ? err.message : 'An unknown error occurred';
    } finally {
        loading.value = false;
    }
});
</script>