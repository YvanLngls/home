<template>
    <div class="min-h-screen flex items-center justify-center">
      <div class="bg-white p-8 rounded-lg shadow-md w-80">
        <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Login</h2>
  
        <form @submit.prevent="handleSubmit">
          <div class="mb-4">
            <label for="username" class="block text-sm font-medium text-gray-600">Username</label>
            <input 
              type="text" 
              id="username" 
              v-model="username" 
              class="w-full p-2 mt-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" 
              required 
            />
          </div>
  
          <div class="mb-6">
            <label for="password" class="block text-sm font-medium text-gray-600">Password</label>
            <input 
              type="password" 
              id="password" 
              v-model="password" 
              class="w-full p-2 mt-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" 
              required 
            />
          </div>
  
          <button 
            type="submit" 
            class="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const username = ref('');
  const password = ref('');
  
  const handleSubmit = () => {
    // Logique de connexion ici
    console.log('Se connecter avec:', username.value, password.value);
  };
  </script>
  
  <style scoped>
  /* Aucun fond gris ajouté */
  </style>
  