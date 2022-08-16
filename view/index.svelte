<script>
  import Header from './component/Header.svelte'
  import Content from './component/Content.svelte'
  export let postList
</script>

<Header />
<Content {postList} />

<br />

<style>
</style>
