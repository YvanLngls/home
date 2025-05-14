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
  

  // {"particles":
  // {"number":
  // {"value":70,"density":
  // {"enable":true,"value_area":800}},
  // "color":{"value":"#b2ffa8"},
  // "shape":{"type":"circle","stroke":{"width":0,"color":"#000000"},
  // "polygon":{"nb_sides":6},"image":{"src":"img/github.svg","width":100,"height":100}},"opacity":{"value":0.5,"random":false,
  // "anim":{"enable":false,"speed":1,"opacity_min":0.1,"sync":false}},"size":{"value":4,"random":true,"anim":{"enable":false,"speed":4,"size_min":1,"sync":false}},
  // "line_linked":{"enable":true,"distance":250,"color":"#b2ffa8","opacity":0.4,"width":1},
  // "move":{"enable":true,"speed":2,"direction":"top-right","random":true,"straight":false,"out_mode":"out","bounce":false,
  // "attract":{"enable":false,"rotateX":600,"rotateY":1200}}},
  // "interactivity":{"detect_on":"canvas","events":{"onhover":{"enable":true,"mode":"grab"},"onclick":{"enable":false,"mode":"repulse"},"resize":true},
  // "modes":{"grab":{"distance":300,"line_linked":{"opacity":0.1}},
  // "bubble":{"distance":400,"size":40,"duration":2,"opacity":8,"speed":3},"repulse":{"distance":200,"duration":0.4},"push":{"particles_nb":4},"remove":{"particles_nb":2
  // }}},"retina_detect":true});var count_particles, stats, update; stats = new Stats; stats.setMode(0); stats.domElement.style.position = 'absolute'; stats.domElement.style.left = '0px'; stats.domElement.style.top = '0px'; document.body.appendChild(stats.domElement); count_particles = document.querySelector('.js-count-particles'); update = function() { stats.begin(); stats.end(); if (window.pJSDom[0].pJS.particles && window.pJSDom[0].pJS.particles.array) { count_particles.innerText = window.pJSDom[0].pJS.particles.array.length; } requestAnimationFrame(update); }; requestAnimationFrame(update);;
  
  
  const options = ref({
    fullScreen: {
      enable: true,
      zIndex: -1
    },
    background: {
      color: {
        value: '#02010a'
      }
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
  