export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) {
    const cookieHeader = useRequestHeaders(['cookie']).cookie;
    const token = cookieHeader?.match(/token=([^;]+)/)?.[1];

    if (!token) return navigateTo('/login');

    const { data, error } = await $fetch('http://localhost:8080/verify-token', {
      method: 'POST',
      headers: { cookie: cookieHeader },
      credentials: 'include',
    });

    if (error || !data?.valid) return navigateTo('/login');
  }
});
