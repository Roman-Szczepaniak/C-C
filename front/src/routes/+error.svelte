<script>
	import '../app.postcss';
	import { page } from '$app/stores';
	import { GradientButton } from 'flowbite-svelte';
</script>

<div
	class="flex flex-col justify-center items-center px-6 mx-auto h-screen xl:px-0 dark:bg-gray-900"
>
	<div class="text-center xl:max-w-4xl">
		{#if $page.status === 404}
			<h1
				class="mb-3 text-2xl font-bold leading-tight text-gray-900 sm:text-4xl lg:text-5xl dark:text-white"
			>
				Page not found
			</h1>
			<p class="mb-5 text-base font-normal text-gray-500 md:text-lg dark:text-gray-400">
				Oops! Looks like you followed a bad link. If you think this is a problem with us, please
				tell us.
			</p>
		{:else if $page.status === 500}
			<h1
				class="mb-3 text-2xl font-bold leading-tight text-gray-900 sm:text-4xl lg:text-5xl dark:text-white"
			>
				Something has gone seriously wrong
			</h1>
			<p class="mb-5 text-base font-normal text-gray-500 md:text-lg dark:text-gray-400">
				It's always time for a coffee break. We should be back by the time you finish your coffee.
			</p>
		{:else}
			<h1>{$page.status}: {$page.error?.message}</h1>
		{/if}

		<GradientButton color="purpleToBlue" href="/">
			<svg
				class="mr-2 -ml-1 w-5 h-5"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					fill-rule="evenodd"
					d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
					clip-rule="evenodd"
				/></svg
			>
			Go back home
		</GradientButton>
	</div>
</div>