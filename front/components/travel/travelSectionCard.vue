<template>
  <NuxtLink 
    :to="`/travel/${projectId}`"
    class="relative m-2 p-4 pb-8 w-full h-full rounded-xl bg-black/20 backdrop-blur-sm cursor-pointer transition-transform hover:scale-[1.01] hover:shadow-2xl hover:bg-black/30 block no-underline"
  >
    <div class="absolute -top-4 left-6 px-4 py-1 rounded-full">
      <h1 
        class="text-sm font-bold text-white tracking-wider px-3 py-1 rounded-full shadow-lg"
        :class="headerColor" 
      >
        {{ date }} 
      </h1>
    </div>
    
    <div class="mt-4">
      <slot></slot> 
    </div>

  </NuxtLink>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  projectId: { // The ID of the travel project
    type: [String, Number],
    required: true
  },
  date: { // The depearture date
    type: [String],
    required: true
  },
  routeName: { // The name of the route (based on your file structure)
    type: String,
    required: true
  },
  headerColor: { // 💡 NOUVELLE PROP : Classe Tailwind pour la couleur de l'en-tête (e.g., 'bg-red-500')
    type: String,
    default: 'bg-indigo-600'
  }
});
</script>