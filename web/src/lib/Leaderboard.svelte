<script lang="ts">
    export let pokemon: {Name: string, votes: number}[]

    $: showAllRankings = false

    $: sortedLeaderboard = pokemon.sort(({ votes: firstVotes }, {votes: secondVotes }) => (secondVotes - firstVotes)).filter((_, i) => showAllRankings ? true : i<10)

    function toggleShowAllRankings() {
        showAllRankings = !showAllRankings
    }
</script>
<div>
    <h2>Leaderboard</h2>
    <div class="leaderboard">
        {#each sortedLeaderboard as selectedPokemon }
            <div>{selectedPokemon.votes} - {selectedPokemon.Name}</div>
        {/each}
    </div>
    <button on:click={toggleShowAllRankings}>{showAllRankings ? "Show Top 10" : "Show all"}</button>
</div>

<style>
    .leaderboard {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        max-height: 20rem;
        overflow-y: scroll;
    }

    button {
        margin: 1rem;
    }
</style>