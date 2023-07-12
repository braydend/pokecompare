<script lang="ts">
  export let pokemon;
</script>

<div class="pokemon">
  <span class="name">{pokemon.Name}</span>
  <span>{pokemon.ID}</span>
  <img src={pokemon.Sprite} alt={`${pokemon.Name}'s image'`} />
  <span>{pokemon.Types.join(",")}</span>
 </div>

 <style>
  .pokemon {
    display: flex;
    flex-direction: column;
    border: 1px solid white;
    border-radius: 4px;
  }

  .pokemon > .name {
    font-size: 3rem;
  }
 </style>
