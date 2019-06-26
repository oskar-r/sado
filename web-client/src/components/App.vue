<template>
  <b-container fluid id="app">
    <div id="main-view" v-if="isLogedIn">
      <nav class="navbar navbar-dark bg-dark">
        <span class="navbar-brand" href="#">
          Archive - Demo
        </span>
        <form class="form-row">
          <div class="role-label">Role:</div>
          <div class="col" v-if="myRoles.length > 1">
            <b-dropdown id="role-dd" right :text="activeRole" split variant="bg-primary-element" class="m-md-2" :placeholder="activeRole" aria-label="Role">
              <b-dropdown-item v-for="role in myRoles" :key=role @click="changeRole(role)" :active="activeRole == role ? true : false">{{role}}</b-dropdown-item>
            </b-dropdown>
          </div>
          <div v-else class="role-label">
            {{activeRole}}
          </div>
        </form>
      </nav>
      <ErrorBadge />
        <nav id="sidebar" class="bg-dark">
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
        </nav>
      <router-view/>
    </div>
     <div id="login-view" v-else>
      <Login/>
    </div>
  </b-container>
</template>

<script>
// import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
//import '@sebgroup/bootstrap'
import '../assets/css/seb-bootstrap.min.css'
import '../assets/fontawesome.min.css'
import TopBar from './TopBar'
import { mapGetters, mapState } from 'vuex'
import ErrorBadge from './ErrorBadge'
import Login from '../views/Login'
import * as ws from '../utility/websockets.js'

export default {
  components: {Login, ErrorBadge, TopBar },
  data () {
    return {
      logo: require('../assets/logo.png'),
      showDD:false
    }
  },
  created () {
     if (this.isLogedIn) {
      console.log('LOGEDIN' + this.isLogedIn)
      this.$store.dispatch('mainStore/getAppConfig').then((resp) => {
      }).catch((err)=>{
        console.error(err)
      })
      ws.connect(process.env.VUE_APP_WS_SERVER, this.$store, this.$toasted)
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
    numberOfRoles() {
      console.log(this.myRoles)
      // return this.myRoles.length()
      return this.myRoles.length() > 1 ? true : false
    },
    changeRole(role) {
      if (role != this.activeRole) { //Change the role 
        this.$store.dispatch('mainStore/changeRole', role).then(() => {
          this.$store.dispatch('mainStore/getAppConfig')
        }).catch((error) =>{
          console.log(error)
        })
      }
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
      myRoles: 'roles',
      activeRole: 'getActiveRole',
    })
  }
}
</script>

<style>
 @import '../../node_modules/bootstrap-vue/dist/bootstrap-vue.css';
 @import '../assets/css/seb-bootstrap.min.css';
 /*@import '../../node_modules/bootstrap/dist/css/bootstrap.css';*/

 @font-face {
    font-family: 'Material Icons';
    font-style: normal;
    font-weight: 400;
    src: url('../assets/fonts/flUhRq6tzZclQEJ-Vdg-IuiaDsNcIhQ8tQ.woff2') format('woff2');
  }
  @font-face {
    font-family: 'Font Awesome 5 Free';
    font-style: normal;
    font-weight: 900;
    font-display: auto;
    src: url("../assets/fonts/fa-solid-900.eot");
    src: url("../assets/fonts/fa-solid-900.eot?#iefix") format("embedded-opentype"),url("../assets/fonts/fa-solid-900.woff2") format("woff2"),url("../assets/fonts/fa-solid-900.woff") format("woff"),url("../assets/fonts/fa-solid-900.ttf") format("truetype"),url("../assets/fonts/fa-solid-900.svg#fontawesome") format("svg");
  }
  @font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-Light.eot");
  src: url("../assets/fonts/SEBSansSerif-Light.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-Light.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-Light.woff") format("woff");
  font-weight: 300;
  font-style: normal; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-Regular.eot");
  src: url("../assets/fonts/SEBSansSerif-Regular.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-Regular.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-Regular.woff") format("woff");
  font-weight: 400;
  font-style: normal; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-Medium.eot");
  src: url("../assets/fonts/SEBSansSerif-Medium.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-Medium.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-Medium.woff") format("woff");
  font-weight: 500;
  font-style: normal; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-Bold.eot");
  src: url("../assets/fonts/SEBSansSerif-Bold.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-Bold.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-Bold.woff") format("woff");
  font-weight: 700;
  font-style: normal; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-LightItalic.eot");
  src: url("../assets/fonts/SEBSansSerif-LightItalic.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-LightItalic.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-LightItalic.woff") format("woff");
  font-weight: 300;
  font-style: italic; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-RegularItalic.eot");
  src: url("../assets/fonts/SEBSansSerif-RegularItalic.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-RegularItalic.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-RegularItalic.woff") format("woff");
  font-weight: 400;
  font-style: italic; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-MediumItalic.eot");
  src: url("../assets/fonts/SEBSansSerif-MediumItalic.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-MediumItalic.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-MediumItalic.woff") format("woff");
  font-weight: 500;
  font-style: italic; }

@font-face {
  font-family: "SEBSansSerif";
  src: url("../assets/fonts/SEBSansSerif-BoldItalic.eot");
  src: url("../assets/fonts/SEBSansSerif-BoldItalic.eot?#iefix") format("embedded-opentype"), url("../assets/fonts/SEBSansSerif-BoldItalic.woff2") format("woff2"), url("../assets/fonts/SEBSansSerif-BoldItalic.woff") format("woff");
  font-weight: 700;
  font-style: italic; }

  .fa,
  .fas {
    font-family: 'Font Awesome 5 Free';
    font-weight: 900; 
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
  /*-webkit-transform: translate3d(-150px, 0, 0);
  transform: translate3d(-150px, 0, 0);*/
  transition: -webkit-transform 400ms ease-in-out;
  transition: transform 400ms ease-in-out;
  transition: transform 400ms ease-in-out, -webkit-transform 400ms ease-in-out;
  width: 200px;
  padding-top:2rem;
}
.toast-info {
    color: #fff !important;
    background-color: #673AB6 !important;
    border-color: #673AB6 !important;
}
.toast-error {
    color: #fff !important;
    background-color:#F03529 !important;
    border-color: #F03529 !important;
}
.toasted {
  display:block!important;
}
.toasted-container.bottom-right {
    right: 1% !important;
    bottom: 1% !important;
        max-width: 15%;
}
.toasted  i.fa {
  padding-top:5px!important;
  float:left!important;
}
.toast-title{
  float: left;
  font-weight: 600;
  padding-top:6px;
}
.toast-message {
    float: left;
    display: block;
    width: 100%;
    padding-top: 6px;
    padding-bottom:3px;
    word-break: break-all;
}
.nav-item {
  display: block;
  text-align: left;
  color: #fff;
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
 #role-dd{
   margin-right:5%;
 }
 .role-label{
   color:#fff;
   font-size: 1.2em;
   font-weight: bold;
   padding-top:18px;
 }
</style>