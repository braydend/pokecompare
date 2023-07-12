<script lang="ts">
    export let pokemon: {Name: string, votes: number}[]

    $: sortedLeaderboard = pokemon.sort(({ votes: firstVotes }, {votes: secondVotes }) => (secondVotes - firstVotes))
</script>

<div class="leaderboard">
    {#each sortedLeaderboard as selectedPokemon }
        <div>{selectedPokemon.votes} - {selectedPokemon.Name}</div>
    {/each}
</div>

<style>
    .leaderboard {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }
</style>