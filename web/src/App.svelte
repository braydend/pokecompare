<script lang="ts">
  import Leaderboard from "./lib/Leaderboard.svelte";
  import Pokemon from "./lib/Pokemon.svelte";

  let generateMatchupPromise = fetch("http://localhost:8000/generate-matchup", { 
    method: "POST",  
    body: JSON.stringify({previousIds: []})
  }).then(resp => resp.json())

  let leaderboard = [];

  function shufflePokemon() {
    const payload = leaderboard.length >= 20 ? {previousIds: leaderboard.map(({ ID}) => ID)} : {previousIds: []}
    generateMatchupPromise = fetch("http://localhost:8000/generate-matchup", {
      method: "POST", 
      body: JSON.stringify(payload)
    }).then(resp => resp.json())
  }

  function addPokemonToLeaderboard(pokemon) {
    const existingPokemon = leaderboard.find(({ ID}) => pokemon.ID === ID);

    if (existingPokemon) {
      leaderboard = [...leaderboard.filter(({ ID}) => ID !== pokemon.ID), {...existingPokemon, votes: existingPokemon.votes + 1}]
      shufflePokemon();
      return;
    }

    leaderboard = [...leaderboard, {...pokemon, votes: 1} ];
    shufflePokemon();
  }
</script>

<main>
  <h1>Pick your favourite pokemon!</h1>
  <div class="container">
  <div class="comparison">
    {#await generateMatchupPromise}
      <div>Loading...</div>  
    {:then data}
      <div on:click={() => addPokemonToLeaderboard(data.PokemonOne)}>
        <Pokemon pokemon={data.PokemonOne} />
      </div>
      <div on:click={() => addPokemonToLeaderboard(data.PokemonTwo)}>
        <Pokemon pokemon={data.PokemonTwo} />
      </div>
    {/await}
  </div>
  <Leaderboard pokemon={leaderboard} />
</div>

</main>

<style>
  .container {
    display: grid;
    grid-template-columns: 2fr 1fr;
  }
  
  .comparison {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }
</style>
