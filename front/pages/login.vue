<template>
  <div class="min-h-screen flex flex-col items-center justify-center">
    <div class="p-8 rounded-lg backdrop-filter backdrop-blur-xs bg-white/10 shadow-lg">
    <UForm :state="state" :validate="validate" @submit="onSubmit"
    class="space-y-4 opacity-100">
    <h1 class="text-2xl font-bold mb-4">Login</h1>
      <UFormField label="Username">
        <UInput v-model="state.username"/>
      </UFormField>

      <UFormField label="Password">
        <UInput v-model="state.password" type="password"/>
      </UFormField>
      
      <div class="flex justify-center mt-8">
        <UButton type="submit">
          Submit
        </UButton>
      </div>
    </UForm>
  </div>
  </div>
</template>
  
  <script setup lang="ts">
  import { useRouter } from 'vue-router';
  
  const state = reactive({
    username: undefined,
    password: undefined
  })
  
  const validate = (state: any): FormsError[] => {
    const errors = []
    if (!state.username) errors.push({ name: 'username', message: 'Required' })
    if (!state.password) errors.push({ name: 'password', message: 'Required' })
    return errors
  }
  
  const toast = useToast()
  const router = useRouter();
  
  async function onSubmit(event: FormSubmitEvent<typeof state>) {
    try {
      const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: state.username,
        password: state.password
      }),
      credentials: 'include' // important: allows cookies
    });

    if (!res.ok) {
      const errText = await res.text();
      throw new Error(errText);
    }

    const data = await res.json();
    console.log('Logged in:', data);
    toast.add({ title: 'Success', description: 'The form has been submitted.', color: 'primary' })

    router.push("/dashboard")

    // Optional: redirect or show success
    } catch (err) {
      toast.add({ title: 'Error', description: 'Could not login : '+err+'.', color: 'error' })
      console.error('Could not login:', err);
    }
  }
  
</script>