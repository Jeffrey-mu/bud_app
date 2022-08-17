<script lang="typescript">
  export let postList
  import { fade } from 'svelte/transition'
  const states = postList.map((item) => ({ id: item.id, state: false }))
  function findItem(id) {
    return states.findIndex((item) => item.id === id)
  }
</script>

<div class="content">
  <ul class="content-ul">
    {#each postList as { content, description, main_title, id, picture, time } (id)}
      <li>
        <img class="picture" src={'http://47.104.212.164:3000/' + picture} alt="img" width="80%" />
        <h2>{main_title}</h2>
        <p class="time">{time}</p>
        <p class="description">{description}</p>
        {#if states.find((item) => item.id === id && item.state)}
          <img
            class="show-content"
            src={'/up.png'}
            alt="img"
            width="32"
            on:click={() => {
              let index = findItem(id)
              states[index].state = false
            }}
          />
        {:else}
          <img
            class="show-content"
            src={'/down.png'}
            alt="img"
            width="32"
            on:click={() => {
              let index = findItem(id)
              states[index].state = true
            }}
          />
          <div bind:innerHTML={content} contenteditable transition:fade={{ delay: 250, duration: 300 }} />
        {/if}
        <hr />
      </li>
    {/each}
  </ul>
</div>

<style>
  .content {
    font-family: Helvetica, Arial, sans-serif;
    margin: 0 auto;
    padding: 0;
    max-width: 100%;
    list-style: none;
    height: 91.2vh;
    background: conic-gradient(from 192deg at 46.5% 42.58%, #fa8792 -54.94deg, #561bbe 28.29deg, #70e6fb 157.82deg, #fbf8b3 220.83deg, #fa8792 305.06deg, #561bbe 388.29deg);
    overflow: auto;
  }
  .content-ul {
    max-width: 900px;
    margin: 0 auto;
  }
  .time {
    color: #999;
    font-size: 14px;
  }
  .show-content {
    cursor: pointer;
  }
  .content li {
    text-align: center;
    margin-top: 20px;
  }
  .picture {
    border-radius: 20px;
    transition: all 0.2s;
    margin: 0 auto;
  }
  .picture:hover {
    transition: all 0.2s;
    transform: scale(1.05);
  }
  .description:hover {
    transition: all 0.2s;
    color: rgb(240, 126, 27);
  }
  p:hover {
    transition: all 0.2s;
    color: rgb(240, 126, 27);
  }
  .p-content {
    overflow: hidden;
  }
</style>
