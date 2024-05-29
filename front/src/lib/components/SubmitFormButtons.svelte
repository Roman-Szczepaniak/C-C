<script lang="ts">
	import { Button, Spinner } from 'flowbite-svelte';
	import ArrowDownOnSquare from 'svelte-heros-v2/ArrowDownOnSquare.svelte';
	import Trash from 'svelte-heros-v2/Trash.svelte';
	import DeletionModal from '$lib/components/DeletionModal.svelte';

	export let loading = false;
	export let canDelete = false;

	let deleteOpen = false;
	export let deleteCallback = () => {};
</script>

<div class="mt-4 flex justify-between items-center">
	{#if canDelete}
		<Button color="red" size="sm" on:click={() => (deleteOpen = true)} disabled={loading}>
			<Trash />
		</Button>
	{:else}
		<div />
	{/if}

	<Button type="submit" size="sm" disabled={loading}>
		{#if loading}<Spinner class="mr-2" size="4" color="white" />{:else}<ArrowDownOnSquare
				class="mr-2"
			/>{/if}
		Save
	</Button>
</div>

<DeletionModal bind:open={deleteOpen} on:click={deleteCallback} />