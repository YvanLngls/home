// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['nuxt-particles', '@nuxt/ui'],
  css: ['~/assets/css/main.css', '~/assets/css/fonts.css'],
  ui: {
    theme: {
      colors: ['primary', 'secondary', 'tertiary', 'accent', 'success', 'error']
    }
  }
})