<script>
	import BoxTitle from './BoxTitle.svelte';
	import classNames from 'classnames';
</script>

<div
	class={classNames(
		'p-4 flex rounded-lg border border-gray-200 bg-white shadow-md dark:border-gray-700 dark:bg-gray-800 flex-col',
		$$props.class
	)}
>
	{#if $$slots.title}
		<BoxTitle><slot name="title" /></BoxTitle>
	{/if}
	<slot />
</div>