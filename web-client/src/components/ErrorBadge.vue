<template>
    <div v-if="error !== undefined && (error.code !== 0 && error.code !== '')">
        <div id="error-badge"> Code: {{error.code}} <br> Message: {{error.message}}</div>
        <div class="close-x" @click="close">X</div>
    </div>
</template>

<script>

import { mapGetters } from 'vuex'

export default {
    data () {
        return {
            // isLogedIn: this.$store.logedIn
        }
    },
    created () {
        console.log(this.$store)
        console.log(process.env)
    },
    computed: {
        ...mapGetters('mainStore', {
            error: 'errorMessage'
        })
    },
    methods: {
        close () {
            this.$store.dispatch('mainStore/resetErrorMessage')
        }
    }
}
</script>

<style>
    div#error-badge {
        z-index:99 !important;
        position: absolute !important;
        top:0 !important;
        height:60px !important;
        width: 100% !important;
        background-color:#ff7961 !important;
        padding: 3px !important;
        text-align: center !important;
        color: white;
        font-size:1.4em;
    }
    .close-x{
        position: absolute;
        top: 0px;
        z-index: 100;
        color: white;
        right: 8px;
        font-size: 1.1em;
    }
</style>