<template>
  <b-container fluid class=login>
    <b-row>
      <b-col cols="5"></b-col>
      <b-col cols="2"><h2>Login</h2></b-col>
      <b-col cols="5"></b-col>
      </b-row>
    <b-row>
      <b-col cols="5"></b-col>
      <b-col cols="2">
        <LoginBox/>
      </b-col>
      <b-col cols="5">
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
// @ is an alias to /src
import LoginBox from '@/components/LoginBox.vue'

export default {
  name: 'login',
  components: {
    LoginBox
  }
}
</script>

<style>

.login {
  padding-top:100px;
}
label {
  float:left;
}


.bgc-white, .bgcH-white:hover {
    background-color: #fff!important;
}
</style>
