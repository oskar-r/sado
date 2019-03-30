<template>
  <div id="file-drag-drop">
    <form ref="fileform">
      <div class="drop-files">
        <span class="big-icon">
          <i class="material-icons">cloud_upload</i>
        </span>
        <span class="drop-files">Drop file here</span>
      </div>
    </form>
  </div>
</template>

<script>
  export default {
    data(){
      return {
        dragAndDropCapable: false,
        files: [],
        showModal:false,
        showUserModal:false,
        showWarningModal:false,
        error:{},
        parsedUsers: [],
        uploadPercentage:0
      }
    },
    mounted(){
      this.dragAndDropCapable = this.determineDragAndDropCapable();
      var self = this
      if( this.dragAndDropCapable ){
        ['drag', 'dragstart', 'dragend', 'dragover', 'dragenter', 'dragleave', 'drop'].forEach( function( evt ) {
         
          this.$refs.fileform.addEventListener(evt, function(e){
            e.preventDefault()
            e.stopPropagation()
          }.bind(this), false)
        }.bind(this))

        this.$refs.fileform.addEventListener('drop', function(e){
          //self.upload(e.dataTransfer.files[0])
          for(let i = 0; i < e.dataTransfer.files.length; i++ )
            this.upload(e.dataTransfer.files[i])
        }.bind(this))
      }
    },
    methods: {
      determineDragAndDropCapable(){
        var div = document.createElement('div')
        return ( ( 'draggable' in div )
                || ( 'ondragstart' in div && 'ondrop' in div ) )
                && 'FormData' in window
                && 'FileReader' in window
      },
      upload(file) {

        this.$store.dispatch('mainStore/upload',file).then((resp) => {
          console.log(resp)
        }).catch((err)=> {
          console.error(err)
        })

        /*
        let formData = new FormData();
        console.log(this.files)
       for( var i = 0; i < this.files.length; i++ ){
          let file = this.files[i];
          formData.append('files[' + i + ']', file);
        }

        axios.post( 'http://localhost:8091/upload',
          formData,
          {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
          }
        ).then(function(){
          console.log('SUCCESS!!');
        })
        .catch(function(){
          console.log('FAILURE!!');
        });
        this.files = []
        */
        /*this.$store.dispatch('mainStore/upload',formData).then((resp) => {
          console.log(resp)
        }).catch((err)=> {
          console.error(err)
        })*/
      }
    }
  }
</script>

<style scoped>
#file-drag-drop {
     padding-top:40px;
}
  form{
    display: block;
    height: 400px;
    width: 400px;
    background: #ccc;
    margin: auto;

    text-align: center;
    line-height: normal;
    border-radius: 10px;
    border: dashed black 1px;
  }
  .drop-files {
    margin-top:200px;
  }
  form > span {
    line-height: normal!important;
  }
  .big-icon {
    display: block;
  }
  .big-icon > i {
    font-size: 70px;
  }
</style>