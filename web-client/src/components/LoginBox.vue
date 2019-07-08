<template>
  <div id="login-box">
     <b-form>
      <b-form-group id="usernameLbl"
                    label="Username"
                    label-for="username">
        <b-form-input id="username"
                      type="email"
                      v-model="username"
                      required
                      placeholder="Username">
        </b-form-input>
      </b-form-group>
      <b-form-group id="passwordLbl"
                    label="Password"
                    label-for="password">
        <b-form-input id="password"
                      type="password"
                      v-model="password"
                      required
                      placeholder="Password">
        </b-form-input>
      </b-form-group>
      <b-button v-on:click="logIn" variant="primary">Log in</b-button>
    </b-form>
  </div>
</template>
<script>
import * as ws from '../utility/websockets.js'
export default {
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    logIn () {
      var self = this
      if (this.username.length > 0 && this.password.length > 0) {
        this.$store.dispatch('mainStore/logIn', { username: this.username, password: this.password }).then((resp) => {
          self.$store.dispatch('mainStore/getAppConfig')
          ws.connect(process.env.VUE_APP_WS_SERVER, self.$store, self.$toasted, self)
        }).catch((err) => {
          console.error(err)
        })
      }
    },
    register () {
    },
    passwordForgotten () {
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
</style>
