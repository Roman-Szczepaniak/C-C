<script lang="ts">
	import { Card, Heading, Input, Label, Textarea } from 'flowbite-svelte';
	import { invalidate } from '$app/navigation';
	import { del, formDataToJson, post, put } from '$lib/api';
	import type { ICard } from '$lib/interfaces';
	import * as notification from '$lib/notifications';
	import SubmitFormButtons from '$lib/components/SubmitFormButtons.svelte';

	export let card: ICard = {} as ICard;
    export let monster_id: number;

	async function deleteCard() {
		try {
			const response = await del(`/cards/${card.id}`);
			if (response?.detail) {
				notification.error('Impossible to delete');
				return;
			}
			notification.success('Card deleted');
			await invalidate('app:monster')
		} catch (error) {
			notification.error('An error occurred while deleting the card');
		}
	}

	let loading = false;
	let savePromise = Promise.resolve();
	$: savePromise, savePromise.then(() => (loading = false));
	async function handleSubmit(event: SubmitEvent) {
		loading = true;
		const formData = new FormData(event.target as HTMLFormElement);
		const data = formDataToJson(formData);
        data.ca = Number(data.ca);
		data.pv = Number(data.pv);
		
		if (card.id) {
			data.id = card.id;
			const updateCard = await put(`/cards/${card.id}`, data);
			if (updateCard?.detail) {
				notification.error('Something wrong!');
				return;
			}
			card = updateCard;
			notification.success('Card updated successfully');
		} else {
			delete data.id;
			const createCard = await post(`/cards`, data);
			if (createCard?.detail) {
				notification.error('Something wrong!');
				return;
			}
			card = createCard;
            const updateMonster = await put(`/monsters/${monster_id}`, { card_id: card.id });
			if (updateMonster?.detail) {
				notification.error('Something wrong when updating monster!');
				return;
			}
			notification.success('Card created and associated with monster successfully');
            await invalidate('app:monster')
		}
	}
</script>

<Card size="xl">
    <Heading tag="h2" class="mb-4">
        {#if card.id}
            Edit Card
        {:else}
            Create Card
        {/if}
    </Heading>
    <form on:submit|preventDefault={(event) => (savePromise = handleSubmit(event))}>
        <div class="grid gap-4 mb-6">
            <Label>
                <span class="text-white">Description</span>
                <Textarea id="description" name="description" rows="3" bind:value={card.description} required class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">CA</span>
                <Input id="ca" name="ca" type="number" bind:value={card.ca} required class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">PV</span>
                <Input id="pv" name="pv" type="number" bind:value={card.pv} required class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Speed</span>
                <Input id="speed" name="speed" type="text" bind:value={card.speed} required class="mt-1 block w-full" />
            </Label>
            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
                <Label>
                    <span class="text-white">Strength</span>
                    <Input id="strength" name="strength" type="text" bind:value={card.strength} required class="mt-1 block w-full" />
                </Label>
                <Label>
                    <span class="text-white">Dexterity</span>
                    <Input id="dexterity" name="dexterity" type="text" bind:value={card.dexterity} required class="mt-1 block w-full" />
                </Label>
                <Label>
                    <span class="text-white">Constitution</span>
                    <Input id="constitution" name="constitution" type="text" bind:value={card.constitution} required class="mt-1 block w-full" />
                </Label>
                <Label>
                    <span class="text-white">Intelligence</span>
                    <Input id="intelligence" name="intelligence" type="text" bind:value={card.intelligence} required class="mt-1 block w-full" />
                </Label>
                <Label>
                    <span class="text-white">Wisdom</span>
                    <Input id="wisdom" name="wisdom" type="text" bind:value={card.wisdom} required class="mt-1 block w-full" />
                </Label>
                <Label>
                    <span class="text-white">Charisma</span>
                    <Input id="charisma" name="charisma" type="text" bind:value={card.charisma} required class="mt-1 block w-full" />
                </Label>
            </div>
            <Label>
                <span class="text-white">Save Throws</span>
                <Input id="save_throws" name="save_throws" type="text" bind:value={card.save_throws} class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Damage Resist</span>
                <Input id="damage_resist" name="damage_resist" type="text" bind:value={card.damage_resist} class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Damage Immune</span>
                <Input id="damage_immune" name="damage_immune" type="text" bind:value={card.damage_immune} class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Condition Immune</span>
                <Input id="condition_immune" name="condition_immune" type="text" bind:value={card.condition_immune} class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Senses</span>
                <Input id="senses" name="senses" type="text" bind:value={card.senses} class="mt-1 block w-full" />
            </Label>
            <Label>
                <span class="text-white">Languages</span>
                <Input id="languages" name="languages" type="text" bind:value={card.languages} class="mt-1 block w-full" />
            </Label>
        </div>
        <SubmitFormButtons bind:loading canDelete={card.id} deleteCallback={deleteCard} />
    </form>
</Card>
