<template>
  <div class="home base-view-box">
    <b-container fluid>
      <DropZone />
    </b-container>
  </div>
</template>

<script>
import DropZone from '@/components/DropZone'
export default {
  name: 'home',
  components: {
    DropZone
  }
}
</script>

<style>

</style>