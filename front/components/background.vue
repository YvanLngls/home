<template>
  <div style="position: absolute; top: -20px; left: -300px;">
    <NuxtParticles
      id="tsparticles"
      :options="options"
    ></NuxtParticles>
  </div>
</template>
  
<script setup lang="ts">
  import { ref } from 'vue'
  
  
  const token = useCookie<string>('token')?.value;
  const isAuthenticated = !!token;
  let color = "#8806ce"
  if(isAuthenticated){color = "#00003d"}

  const options = ref({
    fullScreen: {
      enable: true,
      zIndex: -1
    },
    background: {
      color: { value: "transparent" },
      image: "radial-gradient(circle at top left, #010b13, "+color+")",
      position: "center",
      repeat: "no-repeat",
      size: "cover"
    },
    particles: {
      color: {
        value: ["#b2ffa8"],
      },
      opacity:{
        value: 0.3
      },
      density:{
        enable: true,
      },
      links: {
        enable: false,
      },
      move: {
        enable: true,
        speed: 2,
        direction: "top-right",
      },
      shape: {
        type: "circle",
      },
      size: {
        value: { min: 1, max: 5 },
        random: true,
      },
      number: {
        value: 150
      }
    },
    interactivity: {
      events: {
        onHover: {
          enable: true,
          mode: "grab",
        },
      },
      modes: {
        grab: {
          distance: 300,
          links:{
            opacity: 0.1
          }
        },
      },
    },
  })
  </script>
  