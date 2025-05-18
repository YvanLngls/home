<template>
  <div class="p-4 rounded-lg backdrop-filter backdrop-blur-xs bg-black/10 shadow-lg">
    <UNavigationMenu orientation="vertical" :items="items" class="data-[orientation=vertical]:w-48" />
  </div>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const items = ref<NavigationMenuItem[][]>([
  [
    {
      label: 'Dashboard',
      icon: 'i-lucide-book-open',
      to: '/dashboard'
    },
    {
      label: 'Paramètres',
      icon: 'i-simple-icons-github',
      to: '/settings',
    }
  ],
  [
    {
      label: 'Domotique',
      icon: 'i-lucide-database',
      to: '/domotique'
    },
    {
      label: 'Patrimoine',
      icon: 'i-lucide-box',
      to: '/patrimoine',
    },
    {
      label: 'Voiture',
      icon: 'i-lucide-box',
      to: '/car',
    }
  ]
])
</script>