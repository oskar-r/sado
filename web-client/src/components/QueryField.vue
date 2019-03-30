<template>
     <b-form class="query-form">
        <b-form-input 
                      type="text"
                      v-model="query"
                      placeholder="query">
        </b-form-input>
      <b-button v-on:click="execQuery" variant="primary">Query</b-button>
    </b-form>
</template>
<script>

export default {
  data () {
    return {
      query: ''
    }
  },
  methods: {
    execQuery () {
    }
  }
}
</script>

<style scoped>
  div#login-box {
    margin-top: 3em;
  }
  a {
    margin-top:1em;
    display: block;
  }
  form{
    display: block;
    height: 400px;
    width: 400px;
    margin: auto;
    text-align: center;
    line-height: normal;
  }
</style>
