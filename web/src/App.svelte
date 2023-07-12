<script lang="ts">
    import Pokemon from "./lib/Pokemon.svelte";

  const randomPokemonPromise = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())
  const randomPokemonPromise2 = fetch("http://localhost:8000/random-pokemon").then(resp => resp.json())
</script>

<main>
  <h1>Pick your favourite pokemon!</h1>

  <div class="comparison">

  {#await randomPokemonPromise}
    <div>Loading...</div>  
  {:then pokemon}
      <Pokemon pokemon={pokemon} />
  {/await}
  {#await randomPokemonPromise2}
  <div>Loading...</div>  
  {:then pokemon}
      <Pokemon pokemon={pokemon} />
  {/await}
</div>

</main>

<style>
  .comparison {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }
</style>
