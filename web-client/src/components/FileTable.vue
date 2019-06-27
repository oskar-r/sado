<template>
  <div class="file-table">
    <b-table striped 
        responsive 
        borderless 
        text-left
        head-variant='dark'
        :fields="fields" 
        :items="files">

      <span slot="iconName" slot-scope="data" v-html="data.value" />
      <template slot="context_menu" slot-scope="row">
        
        <b-dropdown size="sm" dropright variant="link" toggle-class="text-decoration-none" no-caret class="context-menu">
          <template slot="button-content"><i class="fas fa-ellipsis-h"></i></template>
          <b-dropdown-item  @click="query(row)">Query</b-dropdown-item>
          <b-dropdown-item  @click="preview(row)">Preview</b-dropdown-item>
          <b-dropdown-item  @click="deleteDataset(row)">Delete</b-dropdown-item>
      </b-dropdown>
      </template>
      <template slot="row-details" slot-scope="row">
        <b-card>
          <b-row class="mb-2">
            <b-col sm="3" class="text-sm-right"><b>Columns</b></b-col>
            <b-col>
              <span v-for="(c,i) in row.item.preview" :key="i" class="wrap-text">{{ c }}<br></span>
              </b-col>
          </b-row>
          <b-button size="sm" @click="row.toggleDetails">Hide Details</b-button>
        </b-card>
      </template>
    </b-table>
  </div>
</template>
<script>
import { mapGetters, mapState } from 'vuex'
import * as fileIcon from '../utility/documentTypes.js'

export default {
   data() {
      return {
        fields: [
          {
            key: 'iconName',
            label: 'Name',
            formatter: 'iconName'
          },
          {
            key: 'fileSize',
            label: 'Size',
            formatter: (value, key, item) => {
              var str = '' 
              if (item.size > 1024*1000*1000) {
                str =  (item.size/(1024*1000*1000)).toFixed(2) + ' Gb'
              }else if (item.size > 1024*1000) {
                str =  (item.size/(1024*1000)).toFixed(2) + ' Mb'
              } else {
                str =  (item.size/1024).toFixed(2) + ' Kb'
              }
              return str
            }
          },
          {
            // A virtual column with custom formatter
            key: 'modtime',
            label: 'Last Modified',
            formatter: (value, key, item) => {
              return new Date(item.last_modified).toLocaleString()
            }
          },
          {
            key: 'context_menu',
            label: ''
          }
        ]
      }
    },

    methods: {
      preview(row) {
       if (row.item.preview === undefined || row.item.preview.length == 0) {
          this.$store.dispatch('mainStore/previewFile', row.item.name).then((item) => {
            if (item.includes(',',0))
              row.item.preview = item.split(",")
            else if(item.includes(';',0))
              row.item.preview = item.split(";")
            else {
              row.item.preview = item
            }

            row.toggleDetails()
          }).catch((error) => {
            console.error(error)
          })
       } else {
         row.toggleDetails()
       }
       
      },
      fileSize(value,key,item) {
        var str = '' 
        if (item.size > 1024*1000*1000) {
          str =  (item.size/(1024*1000*1000)).toFixed(2) + ' Gb'
        }else if (item.size > 1024*1000) {
          str =  (item.size/(1024*1000)).toFixed(2) + ' Mb'
        } else {
          str =  (item.size/1024).toFixed(2) + ' Kb'
        }
        return str
      },
      iconName(value,key,item) {
        return '<i class="fa fa-'+ fileIcon.byType(item.content_type) +'"></i><span class="file-icon-text">'+ item.name +'</span>'
      },
      query(row) {
        this.$store.dispatch('mainStore/selectDataset', row.item.name)
        this.$router.push({ path: 'query' })
      },
      showMenu() {
        console.log("hej")
      }
    },
   computed: {
      ...mapGetters('mainStore', {
            datasets: 'getDatasets'
        })
   },
   computed: {
    files () {
      console.log(this.$store.state.mainStore.datasets)
      
      return this.$store.state.mainStore.datasets.map((item) => {
        return item
      })
    }
  }
}

</script>

<style scoped>
  .file-table{
    text-align:left;
  }
  .file-table >>> .file-icon-text {
    margin-left:1rem;
  }
  .wrap-text {
    word-break: break-all;
  }
  .table-responsive {
    overflow-x:visible;
  }
  /*.three-dots:after {
    content: '\2807';
    font-size: 20px;
  }*/
  .context-menu >>> button {
    border:0!important;
    background-color: transparent!important;
  }
  .context-menu >>> .dropdown-menu {
    border: 1px solid #474747;
    transform: translate3d(10px, 15px, 0px);
  }
  .context-menu >>> .dropdown-item {
    padding: .05rem .5rem!important;
  }
</style>
