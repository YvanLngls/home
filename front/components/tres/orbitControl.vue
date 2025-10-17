<script setup lang="ts">
import { useTres } from '@tresjs/core'

const { camera, renderer } = useTres()
</script>

<template>
  <TresOrbitControls
    v-if="renderer"
    :args="[camera, renderer?.domElement]"
  />
</template>
