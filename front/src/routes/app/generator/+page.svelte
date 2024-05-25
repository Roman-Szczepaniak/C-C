<script>
  let players = [];
  let encounters = [];
  let histories = [];

  const addPlayer = () => {
    players = [...players, { id: players.length + 1, name: `Player ${players.length + 1}` }];
  };

  const addEncounter = () => {
    encounters = [...encounters, { id: encounters.length + 1, name: `Encounter ${encounters.length + 1}` }];
  };

  const addHistory = () => {
    histories = [...histories, { id: histories.length + 1, description: `History ${histories.length + 1}` }];
  };
</script>

<main class="container mx-auto p-4">
  <h1 class="text-3xl font-bold mb-4">Générateur</h1>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <div>
      <h2 class="text-2xl font-bold mb-2">Players</h2>
      <button class="bg-blue-500 text-white px-4 py-2 rounded-md mb-4" on:click={addPlayer}>Add Player</button>
      <ul class="space-y-4">
        {#each players as player}
          <li class="p-4 border rounded-md">{player.name}</li>
        {/each}
      </ul>
    </div>
    <div>
      <h2 class="text-2xl font-bold mb-2">Encounters</h2>
      <button class="bg-blue-500 text-white px-4 py-2 rounded-md mb-4" on:click={addEncounter}>Add Encounter</button>
      <ul class="space-y-4">
        {#each encounters as encounter}
          <li class="p-4 border rounded-md">{encounter.name}</li>
        {/each}
      </ul>
    </div>
    <div>
      <h2 class="text-2xl font-bold mb-2">Histories</h2>
      <button class="bg-blue-500 text-white px-4 py-2 rounded-md mb-4" on:click={addHistory}>Add History</button>
      <ul class="space-y-4">
        {#each histories as history}
          <li class="p-4 border rounded-md">{history.description}</li>
        {/each}
      </ul>
    </div>
  </div>
</main>
