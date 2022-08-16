<script lang="ts">
  let routes = [
    {
      name: 'index',
      link: '/',
    },
    {
      name: 'home',
      link: '/home',
    },
    {
      name: 'post',
      link: '/post',
    },
  ]
</script>

<ul class="header">
  {#each routes as { link, name } (link)}
    <li>
      <a href={link}><span>{name}</span></a>
    </li>
  {/each}
</ul>

<style>
  .header {
    font-family: Helvetica, Arial, sans-serif;
    margin: 0;
    width: 100vw;
    height: 50px;
    background-color: black;
    list-style: none;
    display: flex;
  }
  .header li a {
    display: inline-block;
    line-height: 50px;
    color: #fff;
    margin: 0 20px;
    text-decoration: none;
  }
</style>
