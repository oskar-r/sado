<template>
    <b-container fluid class="file-table">
      <b-row>
        <b-col cols="1"></b-col>
        <b-col cols="10"><FileTable /></b-col>
        <b-col cols="1"></b-col>
      </b-row>
    </b-container>
</template>

<script>
import FileTable from '@/components/FileTable'
export default {
  name: 'datasets',
  components: {
    FileTable
  }
}
</script>

<style>
 .file-table {
   margin-top:1rem;
 }
</style>