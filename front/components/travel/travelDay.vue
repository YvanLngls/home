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
      <!-- <div 
        v-for="(activity, index) in activities" 
        :key="index" 
        class="flex items-start mb-3"
      >
        <div class="font-bold w-16">{{ activity.time }}</div>
        <div class="">
          <p>{{ activity.description }}</p>
        </div>
        <div class="mx-4">
          <a :href="activity.link" target="_blank" class="text-primary hover:text-secondary font-medium transition duration-150 ease-in-out">{{activity.location}}</a>
        </div>
      </div> -->
    </div>
    
    <p v-else class="text-center py-3">
      Rien de prévu pour le moment. 🌴
    </p>

  </div>
</template>

<script setup lang="ts">

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
}

// 2. Utilisation de defineProps avec l'interface
const props = withDefaults(
  defineProps<DayProps>(), 
  {
    // Définition des valeurs par défaut
    activities: () => [],
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
const { day, date, activities } = props;
</script>