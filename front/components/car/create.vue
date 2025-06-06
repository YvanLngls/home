<template>
    <div class="space-y-4 rounded-lg backdrop-filter backdrop-blur-xs bg-black/30 shadow-lg drop-shadow-xl p-6 flex flex-col items-center">

        <UForm :validate="validate" :state="car" @submit="onSubmit" class="flex flex-col space-y-4">
            <UFormField label="Marque">
                <UInput v-model="car.brand"/>
            </UFormField>
            <UFormField label="Modèle">
                <UInput v-model="car.model"/>
            </UFormField>
            <div class="flex flex-col items-center mt-4 hover:scale-105 transition">
                <UButton type="submit" class="bg-primary text-white px-6 py-2 rounded hover:bg-secondary transition">
                    Créer
                </UButton>
            </div>
        </UForm>
    </div>
</template>

<script setup lang="ts">
    const car = reactive({
        brand: '',
        model: ''
    })
  
    const validate = (car: any): FormsError[] => {
        const errors = []
        if (!car.brand) errors.push({ name: 'marque', message: 'Required' })
        if (!car.model) errors.push({ name: 'modèle', message: 'Required' })
        return errors
    }

    const toast = useToast()
    async function onSubmit(event: FormSubmitEvent<typeof car>) {
        try {
            const res = await fetch('http://localhost:8080/car/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                brand: car.brand,
                model: car.model
            }),
            credentials: 'include'
            });
        

            if (!res.ok) {
                const errText = await res.text();
                throw new Error(errText);
            }
            console.log("Created : ",res)
            toast.add({ title: 'Success', description: 'The car has been created.', color: 'primary' })
        } catch (err) {
        toast.add({ title: 'Error', description: 'Could not create the car : '+err+'.', color: 'error' })
        console.error('Could not create the car:', err);
        }
    }


</script>