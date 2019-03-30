<template>
  <b-container fluid id="app">
    <div id="main-view" v-if="isLogedIn">
      <ErrorBadge />
        <div id="sidebar">
          <div id="sidebar-logo">
            <img :src="logo" />
          </div>

          <b-nav vertical class="w-100">
            <b-nav-item class="nav-item" v-for="route in routes" :key=route.to>
              <router-link v-bind:to=route.to class="nav-item" :key=route.to>
                <div class="icon-span">
                  <i class="material-icons">{{route.icon}}</i>
                  <span class="menu-text"> {{route.text}}</span>
                </div>
              </router-link>
            </b-nav-item>
          </b-nav>

          <a href="#" class="log-out-item nav-item" @click="logOut" >
            <span class="icon-span">
              <i class="material-icons">exit_to_app</i>
            </span>
            <span>
              Logga ut
            </span>
          </a>
        </div>
      <router-view/>
    </div>
     <div id="login-view" v-else>
      <Login/>
    </div>
  </b-container>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import TopBar from './TopBar'
import { mapGetters, mapState } from 'vuex'
import ErrorBadge from './ErrorBadge'
import Login from '../views/Login'
export default {
  components: {Login, ErrorBadge, TopBar },
  data () {
    return {
      logo: require('../assets/logo.png')
    }
  },
  created () {
     if (this.isLogedIn) {
      console.log('LOGEDIN' + this.isLogedIn)
      this.$store.dispatch('mainStore/getAppConfig').then((resp) => {
        console.log(resp)
        // this.$store.dispatch('mainStore/getBaseData')
      }).catch((err)=>{
        console.error(err)
      })

      this.$store.dispatch('mainStore/getMyDatasets')
    }
  },
  methods: {
    logOut() {
      console.log('log out')
      this.$store.dispatch('mainStore/logOut')
    },
    set (set) {
      console.log(set)
      if (set == 'datasets') {
        return this.datasets
      }
      return this.documents
    },
    dropdown(e, data) {
        this.set[data] = []
        this.set[data] = this[data]
        console.log(this.set)
      /*
      if (Array.isArray(set)) {
        set.forEach((item) => {
          e.target.parentElement.append
        })
      }*/
      
      console.log(e.target.parentElement)
    }
  },
  computed: {
    ...mapGetters('mainStore', {
      isLogedIn: 'isLogedIn',
      routes: 'appRoutes',
      datasets: 'getDatasets',
      documents: 'getDocuments'
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
  margin: 0;
  padding: 0;
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


#sidebar {
    float: left;
    height: 100vh;
    z-index: 999;
    background: #fff;
    color: #2c3e50;
    transition: all 0.3s;
}
#sidebar-logo {
  display: block;
  background: no-repeat;
  background-size:100%;
  margin:0 0 30px 0;
}

.nav-item {
  display: block;
  text-align: left;
}
.icon-span {
  display: block;
  float: left;
  margin-right: 9px;
}
.menu-text{
  height:24px;
  vertical-align: top;
}
 .log-out-item {
    position: absolute;
    bottom: 0;
 }
</style>
