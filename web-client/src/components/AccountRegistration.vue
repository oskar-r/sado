<template>
  <div id="account-registration">
      <div v-if="activeRole =='admin'">
      <h2>Create account</h2>
      <b-form @submit.prevent="create" v-if="show">
        <b-form-group
            id="input-group-1"
            label="Username:"
            label-for="input-1"
        >
            <b-form-input
            id="input-1"
            v-model="form.username"
            type="text"
            required
            placeholder="Username"
            ></b-form-input>
        </b-form-group>

        <b-form-group
            id="input-group-2"
            label="Password:"
            label-for="input-2"
        >
            <b-form-input
            id="input-2"
            v-model="form.password"
            type="password"
            required
            placeholder="Password"
            ></b-form-input>
        </b-form-group>
        <b-form-group
            id="input-group-3"
            label="Bucketname:"
            label-for="input-3"
        >
            <b-form-input
            id="input-3"
            v-model="form.bucket"
            type="text"
            required
            placeholder="Bucket name"
            ></b-form-input>
        </b-form-group>
        <b-button type="submit" variant="primary">Create</b-button>
      </b-form>
      <div v-if="created">Account {{form.username}} created</div>
      </div>
      <div v-else>No access to page</div>
  </div>
</template>
<script>

import Accounts from '../api/Accounts'
import { mapGetters } from 'vuex'
export default {
   data() {
      return {
          form: {
              username: '',
              password: '',
              bucket: ''
          },
          show:true,
          created:false
      }
    },
    methods: {  
        create() {
            var that = this
            console.log(Accounts)
            Accounts.create(this.form).then((resp) => {
                console.log(resp)
                that.show = false
            }).catch((err)=> {
                console.error(err)
            })
        } 
    },
   computed: {
    ...mapGetters('mainStore', {
      activeRole: 'getActiveRole'
    })
   }
}

</script>

<style scoped>
    #account-registration {
        padding-top: 40px;
        width: 420px;
        margin: auto;
    }
    form{
        display: block;
        height: 400px;
        width: 400px;
        margin: auto;

        text-align: center;
        line-height: normal;
        border-radius: 10px;
    }
</style>
