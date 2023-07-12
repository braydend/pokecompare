<script lang="ts">
  import Leaderboard from "./lib/Leaderboard.svelte";
  import Pokemon from "./lib/Pokemon.svelte";

  let randomPokemonPromise = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())
  let randomPokemonPromise2 = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())

  let leaderboard = [];

  function shufflePokemon() {
    randomPokemonPromise = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())
    randomPokemonPromise2 = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())
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

  <div class="comparison">
    {#await randomPokemonPromise}
      <div>Loading...</div>  
    {:then pokemon}
      <div on:click={() => addPokemonToLeaderboard(pokemon)}>
        <Pokemon pokemon={pokemon} />
      </div>
    {/await}
    {#await randomPokemonPromise2}
    <div>Loading...</div>  
    {:then pokemon}
      <div on:click={() => addPokemonToLeaderboard(pokemon)}>
        <Pokemon pokemon={pokemon} />
      </div>
    {/await}
  </div>
  <Leaderboard pokemon={leaderboard} />

</main>

<style>
  .comparison {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }
</style>
