<template>
  <b-container id="app">
    <div id="main-view">
      <ErrorBadge />
      <div class="wrapper">
        <div id="sidebar">
          <div id="sidebar-logo">
            <img :src="logo" />
          </div>
          <router-link v-for="route in routes" v-bind:to=route.to class="nav-item" :key=route.to>
            <span class="icon-span">
              <i class="material-icons">{{route.icon}}</i>
            </span>
            <span>
              {{route.text}}
            </span>
          </router-link>
          <a href="#" class="log-out-item nav-item" @click="logOut" >
            <span class="icon-span">
              <i class="material-icons">exit_to_app</i>
            </span>
            <span>
              Logga ut
            </span>
          </a>
        </div>
      </div>
      <router-view/>
    </div>
  </b-container>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import TopBar from './TopBar'
import { mapGetters } from 'vuex'
import ErrorBadge from './ErrorBadge'
export default {
  components: {ErrorBadge, TopBar },
  data () {
    return {
      logo: require('../assets/logo.png')
    }
  },
  created () {

  },
  methods: {
    logOut() {
      
    }
  },
  computed: {
    ...mapGetters('mainStore', {
      isLogedIn: 'isLogedIn',
      routes: 'appRoutes'
    })
  }
}
</script>

<style>
 @import '../../node_modules/bootstrap-vue/dist/bootstrap-vue.css';
 @import '../../node_modules/bootstrap/dist/css/bootstrap.css';
 @font-face {
    font-family: 'Material Icons';
    font-style: normal;
    font-weight: 400;
    src: url('../assets/fonts/flUhRq6tzZclQEJ-Vdg-IuiaDsNcIhQ8tQ.woff2') format('woff2');
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #f5f5f5;
  min-height: 100vh;
  height: 100%;
  border-right:5px white solid;
}
.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;  /* Preferred icon size */
  display: inline-block;
  line-height: 1;
  text-transform: none;
  letter-spacing: normal;
  word-wrap: normal;
  white-space: nowrap;
  direction: ltr;

  /* Support for all WebKit browsers. */
  -webkit-font-smoothing: antialiased;
  /* Support for Safari and Chrome. */
  text-rendering: optimizeLegibility;

  /* Support for Firefox. */
  -moz-osx-font-smoothing: grayscale;

  /* Support for IE. */
  font-feature-settings: 'liga';
}
.wrapper {
    display: flex;
    width: 100%;
}

#sidebar {
    width: 250px;
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    z-index: 999;
    background: #fff;
    color: #2c3e50;
    transition: all 0.3s;
}
#sidebar-logo {
  display: block;
  width: 300px;
  height:75px;
  background: no-repeat;
  background-size:100%;
  margin-bottom:50px;
  margin-top:5px;
}
div#sidebar-logo > img {
  width: 300px;
}
.nav-item {
  display: block;
  text-align: left;
  margin-left:30px;
  margin-bottom:20px;
}
.icon-span {
  display: block;
  float: left;
  margin-right: 9px;
}
.base-view-box {
 
 }
 .log-out-item {
    position: absolute;
    bottom: 0;
 }
</style>
