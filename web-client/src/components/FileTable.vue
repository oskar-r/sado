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
      <template slot="show_details" slot-scope="row">
        <!-- As `row.showDetails` is one-way, we call the toggleDetails function on @change -->
        <b-form-checkbox @change="test(row)" v-model="row.detailsShowing">
        </b-form-checkbox>
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
import { mapGetters } from 'vuex'

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
            key: 'show_details',
            lable: 'Details',
            formatter: 'button'
          }
        ]
      }
    },

    methods: {
      test(row) {
        console.log(row)
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
        var icon =''
        switch (item.category) {
          case "document":
            icon ='<i class="material-icons">description</i>'
            break;
          case "dataset": 
            icon = '<i class="material-icons">view_column</i>'
            break;
        }
        return icon + ' '+ item.name
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
  .wrap-text {
    word-break: break-all;
  }
</style>
