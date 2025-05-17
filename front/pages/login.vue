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
  
  <script setup lang="ts">
  import { ref } from 'vue';
  
  const username = ref('');
  const password = ref('');
    
  const handleSubmit = async () => {
  try {
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value
      }),
      credentials: 'include' // important: allows cookies
    });

    if (!res.ok) {
      const errText = await res.text();
      throw new Error(errText);
    }

    const data = await res.json();
    console.log('Logged in:', data);

    // Optional: redirect or show success
  } catch (err) {
    console.error('Login error:', err);
  }
};

  </script>
  
  <style scoped>
  /* Aucun fond gris ajouté */
  </style>
  