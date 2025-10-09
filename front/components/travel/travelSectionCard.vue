<template>
  <NuxtLink 
    :to="{ name: routeName, params: { id: projectId } }"
    class="relative my-8 m-2 p-4 pb-8 w-1/2 h-full border-2 border-primary rounded-xl bg-black/10 backdrop-blur-sm cursor-pointer transition-transform hover:scale-[1.01] hover:shadow-2xl block no-underline"
  >
  <!-- <div class="relative my-8 m-2 p-4 pb-8 w-1/2 h-full border-2 border-primary rounded-xl bg-black/10 backdrop-blur-sm"
  @click="goToProjectDetail"> -->
    <div class="absolute -top-4 left-6 px-4 py-1 rounded-full">
      <h1 class="text-sm font-bold text-white tracking-wider bg-indigo-600 px-3 py-1 rounded-full shadow-lg">
        {{ title }} 
      </h1>
    </div>
    
    <div class="mt-4">
      <slot></slot> 
    </div>
  <!-- </div> -->

  </NuxtLink>
</template>

<script setup>
import { defineProps } from 'vue';
// No need for useRouter when using NuxtLink!

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  projectId: { // The ID of the travel project
    type: [String, Number],
    required: true
  },
  routeName: { // The name of the route (based on your file structure)
    type: String,
    // Example: The file is pages/voyages/[id].vue
    // The route name is derived automatically (e.g., 'voyages-id')
    required: true
  }
});
</script>