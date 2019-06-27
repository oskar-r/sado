<template>
     <b-form class="query-form">
        <b-form-textarea
          id="textarea"
          v-model="query"
          placeholder="query"
          rows="3"
        ></b-form-textarea>
      <b-button v-on:click="execQuery" variant="primary" class="q-button">Query</b-button>
    </b-form>
</template>
<script>
import { mapGetters, mapState } from 'vuex'
export default {
  data () {
    return {
      query: ''
    }
  },
  created () {
     if (this.selectedDataset !== undefined && this.selectedDataset !== '') {
       this.query = 'SELECT * FROM '+ this.selectedDataset + ' LIMIT 10'
      }
  },
  methods: {
    async execQuery () {
      try {
        var result = await this.$store.dispatch('mainStore/runQuery', this.query)
        console.log(result)
      } catch (e) {
        console.log(e.message)
      }
    },
    defaultQuery() {
      console.log(this.selectedDataset)
      if (this.selectedDataset != '' || this.selectedDataset !== undefined) {
        return 'SELECT * FROM '+ this.selectedDataset + ' LIMIT 10;'
      }
      return ''
    }
  },
  computed: {
    ...mapGetters('mainStore', {
      selectedDataset: 'getSelectedDataset',
      activeRole: 'getActiveRole'
    })
  }
}
</script>

<style scoped>
  form{
    display: block;
    height: 400px;
    width: 50%;
    margin: auto;
    text-align: center;
    line-height: normal;
  }
  .q-button{
    margin-top:3px;
    float:left;
  }
</style>
