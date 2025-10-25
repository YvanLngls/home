<template>
  <div class="border border-gray-200 rounded-xl shadow-md p-5 mb-6">
    
    <header class="border-b border-secondary pb-3 mb-4 flex justify-between items-center">
      <h2 class="text-2xl font-semibold">
        Jour {{ day }} 
        <span v-if="date" class="text-sm font-normal italic ml-2">({{ date }})</span>
      </h2>
    </header>

    <div v-if="activities.length > 0">

      <UTable 
        :data="activities" 
        :columns="columns"
      >
      </UTable>

      
    </div>
    
    <p v-else class="text-center py-3">
      ❓ Informations non renseignées
    </p>
    <div class="w-full py-8 px-16" v-if="nuggetList.length>0">
      <UCarousel
        v-slot="{ item }"
        :items="nuggetList"
        :ui="{ item: 'basis-1/2' }"
        arrows
        :prev-icon="'heroicons:chevron-left-16-solid'"
        :next-icon="'heroicons:chevron-right-16-solid'"
        :autoplay="{ delay: 5000 }"
        loop
        dots
        class="w-full"
      >
        <div class="p-4 border rounded-lg bg-gray-50 dark:bg-gray-800 h-full">
          <h3 class="text-lg font-semibold mb-2">{{ item.name }}</h3>
          <p class="text-gray-700 dark:text-gray-300">{{ item.description }}</p>
        </div>
      </UCarousel>
    </div>

  </div>
</template>

<script setup lang="ts">

interface Nugget {
  name: string;
  description: string;
}

interface Activity {
  time: string;
  description: string;
  location?: string;
  link?: string;
}

interface DayProps {
  day: number;
  date?: string; // Optionnel
  activities: Activity[];
  nuggetList: Nugget[];
}

// 2. Utilisation de defineProps avec l'interface
const props = withDefaults(
  defineProps<DayProps>(), 
  {
    // Définition des valeurs par défaut
    activities: () => [],
    nuggetList: () => [],
    date: ''
  }
);
const columns: TableColumn<Activity>[] = [
  {
    // L'identifiant pour lier la donnée à la colonne
    accessorKey: 'time', 
    header: 'Heure', 
    // Utiliser la 'class' ici ne fonctionne que pour l'en-tête.
    // Pour la cellule, nous utilisons le slot ou la propriété 'cell'.
  },
  {
    accessorKey: 'description',
    header: 'Activité',
  },
  {
    accessorKey: 'location',
    header: 'Lieu / Lien',
    cell: ({ row }) => {
      // console.log(row.original)
      return h('a', { href: row.original.link, target: '_blank', class: 'text-blue-500 hover:underline' }, row.original.location || 'Lien');
      // return `${row.original.link}`
      }

  },
];
// 3. Destructuration pour un accès plus propre dans le template (optionnel mais courant)
const { day, date, activities, nuggetList} = props;
</script>