<script setup lang="ts">
import { useLoop } from '@tresjs/core'
import { useTexture } from '@tresjs/cientos' // 1. Importer useTexture
import { ref } from 'vue'

// 2. Charger la texture (le chemin part de /public/)
// 'await' est nécessaire car le chargement est asynchrone
const { state: texture, isLoading, error } = useTexture('/textures/earth.jpg')

// const { map } = await useTexture({
//   map: '/textures/earth.jpg',
// })
console.log('Texture chargée :', texture)

const sphereRef = ref()
const { onBeforeRender } = useLoop()

onBeforeRender(({ elapsed }) => {
  if (sphereRef.value) {
    // Faisons tourner le globe seulement sur l'axe Y pour un effet plus réaliste
    // sphereRef.value.rotation.y = elapsed * 0.1 
  }
})
</script>

<template>
  <TresMesh ref="sphereRef" :position="[0, 0, 0]"> 
    <!-- <TresSphereGeometry :args="[2, 64, 128]" /> -->
    <TresIcosahedronGeometry :args="[2, 5]" />
    <TresMeshStandardMaterial 
        v-if="texture" 
        :map="texture"
        
        :metalness="0.4"       :roughness="0.7"       :emissive="0x222222"   :emissive-map="lightsMap" :emissive-intensity="1" /> 
    <!-- <TresMeshStandardMaterial :map="texture"/> -->
     <!-- color="#888888"/> -->
  </TresMesh>

  </template>

<!-- <script setup lang="ts">
import { useTexture } from '@tresjs/cientos'

const { state: texture, isLoading, error } = useTexture(path)
</script>

<template>
  <TresMesh>
    <TresSphereGeometry />
    <TresMeshStandardMaterial :map="texture" />
  </TresMesh>
</template> -->