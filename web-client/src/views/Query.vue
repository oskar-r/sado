<template>
    <div class="query-page">
      <QueryField />
    </div>
</template>

<script>
import QueryField from '@/components/QueryField'
export default {
  name: 'query',
  components: {
    QueryField
  }
}
</script>

<style>
.query-page{
  width:100%;
  padding-top:3em;
}
</style>