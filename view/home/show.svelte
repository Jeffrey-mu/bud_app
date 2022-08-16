<script>
  export let home = {}
</script>

<h1>Show Home</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(home) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(home) as key}
      <td>{home[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/home`}>Back</a>
<span> | </span>
<a href={`/home/${home.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
