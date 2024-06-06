<script lang="ts">
    import { Input, Label, Textarea, Card, Heading } from 'flowbite-svelte';
    import { goto } from '$app/navigation';
    import { del, formDataToJson, post, put } from '$lib/api';
    import type { Action } from '$lib/interfaces';
    import * as notification from '$lib/notifications';
    import SubmitFormButtons from '$lib/components/SubmitFormButtons.svelte';

    export let action: Action = {} as Action;

    async function deleteAction() {
        try {
            const response = await del(`/actions/${action.id}`);
            if (response?.detail) {
                notification.error('Impossible to delete');
                return;
            }
            notification.success('Action deleted');
            await goto(`/monsters`);
        } catch (error) {
            notification.error('An error occurred while deleting the action');
        }
    }

    let loading = false;
    let savePromise = Promise.resolve();
    $: savePromise, savePromise.then(() => (loading = false));

    async function handleSubmit(event: SubmitEvent) {
        loading = true;
        const formData = new FormData(event.target as HTMLFormElement);
        const data = formDataToJson(formData);
        
        if (action.id) {
            data.id = action.id;
            const updateAction = await put(`/actions/${action.id}`, data);
            if (updateAction?.detail) {
                notification.error('Something wrong!');
                return;
            }
            action = updateAction;
            notification.success('Action updated successfully');
        } else {
            delete data.id;
            const createAction = await post(`/actions`, data);
            if (createAction?.detail) {
                notification.error('Something wrong!');
                return;
            }
            action = createAction;
            notification.success('Action created successfully');
            await goto(`/monsters`);
        }
    }
</script>

<Card class="mb-6">
    <Heading tag="h2" class="mb-4">
        {#if action.id}
            Edit Action
        {:else}
            Create Action
        {/if}
    </Heading>

    <form on:submit|preventDefault={(event) => (savePromise = handleSubmit(event))}>
        <div class="grid gap-4 mb-6">
            <Label>
                <span class="text-white">Name</span>
                <Input id="name" name="name" type="text" bind:value={action.name} required class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Description</span>
                <Textarea id="description" name="description" rows="5" bind:value={action.description} required class="mt-1 block w-full" />
            </Label>
        </div>
        <SubmitFormButtons bind:loading canDelete={action.id} deleteCallback={deleteAction} />
    </form>
</Card>
