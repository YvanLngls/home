export default defineNuxtRouteMiddleware((to, from) => {
  const token = useCookie<string>('token')?.value;

  if (!token) {
    return navigateTo('/login');
  }

  // Optional: decode token, check expiry etc.
});
