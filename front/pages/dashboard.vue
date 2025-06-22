<template>
  <MenusSide/>
  <div class="p-8">
    <h1 class="text-white">Bienvenue</h1>
  </div>
  <div class="p-8">
    <UButton    @click="playSound">
      Jouer
    </UButton>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
});

const toast = useToast()

async function playSound() {
  try {
    const res = await fetch('http://localhost:8080/music/play', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    credentials: 'include'
    });

    if (!res.ok) {
        const errText = await res.text();
        throw new Error(errText);
    }
    console.log("Created : ",res)
    toast.add({ title: 'Success', description: 'Sound played.', color: 'primary' })
  } catch (err) {
    toast.add({ title: 'Error', description: 'Could not play sound : '+err+'.', color: 'error' })
    console.error('Could not play sound:', err);
  }
}

</script>