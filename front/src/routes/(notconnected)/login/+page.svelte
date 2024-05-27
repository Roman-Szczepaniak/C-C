<script lang="ts">
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { page } from '$app/stores';

  let email = '';
  let password = '';
  let error = '';

  async function login() {
      try {
          const response = await fetch('http://localhost:8080/users/login', {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify({ email, password })
          });

          if (!response.ok) {
              throw new Error('Login failed');
          }

          const user = await response.json();
          sessionStorage.setItem('access_token', user.token);
          goto('/');
      } catch (err) {
          error = err.message;
      }
  }
</script>

<form on:submit|preventDefault={login}>
  <input type="email" bind:value={email} placeholder="Email" required />
  <input type="password" bind:value={password} placeholder="Password" required />
  <button type="submit">Login</button>
  {#if error}
      <p>{error}</p>
  {/if}
</form>
