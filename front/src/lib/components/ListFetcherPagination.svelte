<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { Button } from 'flowbite-svelte';

	export let showPrevious: boolean;
	export let hasMore: boolean;
	export let disabled: boolean = false;

	const dispatch = createEventDispatcher();
</script>

{#if showPrevious || hasMore}
	<div class="flex space-x-3 my-2">
		{#if showPrevious}
			<Button
				color="light"
				class="dark:text-gray-400"
				{disabled}
				on:click={() => dispatch('previous')}
			>
				<svg
					class="mr-2 w-5 h-5"
					fill="currentColor"
					viewBox="0 0 20 20"
					xmlns="http://www.w3.org/2000/svg"
					><path
						fill-rule="evenodd"
						d="M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z"
						clip-rule="evenodd"
					/></svg
				>
				Previous
			</Button>
		{/if}
		{#if hasMore}
			<Button color="light" class="dark:text-gray-400" {disabled} on:click={() => dispatch('next')}>
				Next
				<svg
					class="ml-2 w-5 h-5"
					fill="currentColor"
					viewBox="0 0 20 20"
					xmlns="http://www.w3.org/2000/svg"
					><path
						fill-rule="evenodd"
						d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z"
						clip-rule="evenodd"
					/></svg
				>
			</Button>
		{/if}
	</div>
{/if}