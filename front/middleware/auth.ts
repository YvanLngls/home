export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) {
    const cookieHeader = useRequestHeaders(['cookie']).cookie;
    const token = cookieHeader?.match(/token=([^;]+)/)?.[1];

    if (!token) return navigateTo('/login');

    try {
      const raw = await $fetch<{ valid: boolean }>('http://localhost:8080/verify-token', {
        method: 'POST',
        headers: { cookie: cookieHeader },
        credentials: 'include',
      });
      const res = typeof raw === 'string' ? JSON.parse(raw) : raw;
      if (!res.valid) return navigateTo('/login');
    } catch (err) {
      console.log(err)
      return navigateTo('/login');
    }
  }
});
