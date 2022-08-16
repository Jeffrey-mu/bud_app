<script>
  import Header from '../component/Header.svelte'
  import Content from '../component/Content.svelte'
  import homeData from '../db/homeData.json'
  const postList = homeData
</script>

<Header />
<Content {postList} />

<br />

<style>
</style>
