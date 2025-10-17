<script setup lang="ts">
import { Vector3, MathUtils, CatmullRomCurve3, TubeGeometry, BufferGeometry, LineBasicMaterial, Line } from 'three';
import { ref } from 'vue';

// ⚠️ Rayon de votre globe (doit correspondre à TresSphereGeometry: [2, 16, 64])
const GLOBE_RADIUS = 2; 
const lineRef = ref();

// 1. Définir les "voyages" (Lat/Lon en degrés)
// Exemple : Paris (48.86° N, 2.35° E) à New York (40.71° N, -74.01° O)
const props = defineProps<{
  startLat: number;
  startLon: number;
  endLat: number;
  endLon: number;
}>();
// const travels = [
//   { start: { lat: 51.86, lon: -90.35 }, end: { lat: 45.71, lon: -86.01 } },
//   // Ajoutez d'autres trajets ici
// ];

/**
 * Convertit des coordonnées Lat/Lon en coordonnées Three.js (Vector3).
 * @param lat Latitude en degrés
 * @param lon Longitude en degrés
 */
const latLonToVector3 = (lat: number, lon: number): Vector3 => {
  const latRad = MathUtils.degToRad(lat);
  const lonRad = MathUtils.degToRad(lon);
  
  // Three.js utilise généralement l'axe Y comme axe vertical (polaire)
  const x = GLOBE_RADIUS * Math.cos(latRad) * Math.sin(lonRad);
  const y = GLOBE_RADIUS * Math.sin(latRad);
  const z = GLOBE_RADIUS * Math.cos(latRad) * Math.cos(lonRad);
  
  return new Vector3(x, y, z);
};

// 2. Générer les géométries de lignes
const generateArc = (startLat: number, startLon: number, endLat: number, endLon: number) => {
  const startV = latLonToVector3(startLat, startLon);
  const endV = latLonToVector3(endLat, endLon);

  // Le point de contrôle pour élever l'arc
  const midV = new Vector3()
    .addVectors(startV, endV)
    .divideScalar(2)
    .normalize()
    .multiplyScalar(GLOBE_RADIUS * 1.1); // 1.5 pour élever la courbe au-dessus de la surface
  
  // Utilisation d'une CatmullRomCurve pour un tracé lisse
  const curve = new CatmullRomCurve3([startV, midV, endV]);
  
  // Obtenir les points de la courbe pour la géométrie de la ligne
  const points = curve.getPoints(50); // 50 segments pour la ligne
  
  return points;
};

// 3. Créer une seule géométrie pour toutes les lignes (optimisation)
const linePoints: Vector3[] = [];

const points = generateArc(props.startLat, props.startLon, props.endLat, props.endLon);
points.forEach(p => linePoints.push(p));
// travels.forEach(t => {
    
    // Pour une TresLine, il suffit d'ajouter tous les points
    // Vous devez cependant gérer les sauts si vous voulez une seule TresLine
// });

// Créez une instance de géométrie Three.js
const geometry = new BufferGeometry().setFromPoints(linePoints);

</script>

<template>
  <TresLine :geometry="geometry" ref="lineRef">
    <TresLineBasicMaterial 
        color="#5000ff" 
        :linewidth="5"
        :linecap="'round'" 
        :linejoin="'round'" 
    />
  </TresLine>
  <!-- <TresLine :geometry="geometry" ref="lineRef">
    <TresLineBasicMaterial color="#5000ff" :linewidth="2" />
  </TresLine> -->
</template>