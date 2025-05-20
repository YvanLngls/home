<template>
  <div class="p-4 m-4 rounded-lg backdrop-filter backdrop-blur-xs bg-black/30 shadow-lg drop-shadow-xl">
    <UNavigationMenu orientation="vertical" :items="items" class="data-[orientation=vertical]:w-48" />
  </div>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const items = ref<NavigationMenuItem[][]>([
  [
    {
      label: 'Dashboard',
      icon: 'heroicons:squares-2x2',
      to: '/dashboard'
    },
    {
      label: 'Paramètres',
      icon: 'heroicons:cog-6-tooth',
      to: '/settings',
    }
  ],
  [
    {
      label: 'Domotique',
      icon: 'heroicons:home',
      to: '/home-automation'
    },
    {
      label: 'Patrimoine',
      icon: 'heroicons:presentation-chart-line',
      to: '/patrimoine',
    },
    {
      label: 'Voiture',
      icon: 'hugeicons:car-03',
      to: '/car',
    }
  ]
])
</script>