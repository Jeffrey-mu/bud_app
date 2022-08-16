<script>
  import Header from '../component/Header.svelte'
  import Content from '../component/Content.svelte'
  import postData from '../db/postData.json'
  const postList = postData
</script>

<Header />
<Content {postList} />

<br />

<style>
</style>
