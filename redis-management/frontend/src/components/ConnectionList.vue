<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" name="item.id" @click="getInfo(item.id)">
          <template #title>
            <div class="item">
              <div>
                {{item.name}}
              </div>
              <div style="display: flex">
                <DbInfo @click.stop  :data="item"></DbInfo>
                <ConnectionManage @click.stop title="编辑" btn-type="text" :data="item" @emit-connection-list="connectionList" />
                <el-popconfirm title="确认删除?" @confirm="connectionDelete(item.id)">
                  <template #reference>
                    <el-button link type="danger" @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>
          </template>
          <div v-for="db in infoDbList" @click="selectDB(db.key,item.id)">
            <div v-if="db.key!==selectDbKey" class="my-item">{{db.key}}({{db.number}})</div>
            <div v-else class="my-select-item">{{db.key}}({{db.number}})</div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </main>
</template>

<script setup>
import {defineProps, ref, watch} from "vue";
import {ConnectionDelete, ConnectionList, DbList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
import ConnectionManage from "./ConnectionManage.vue";

let list=ref()
let props=defineProps(['update'])
let infoDbList=ref()
let selectDbKey=ref()
let emits=defineEmits(['emit-select-db'])
//监听update的值，发生变动更新列表
watch(props,(newUpdate)=>{
  connectionList()
})
function connectionList(){
  ConnectionList().then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    list.value=res.data
  })
}
//程序刚进入也需要调用一次
connectionList()
function connectionDelete(id){
  ConnectionDelete(id).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type:"sucess",
    })
    connectionList()
  })
}

function getInfo(id){
  DbList(id).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    infoDbList.value=res.data
  })
}

function selectDB(db,connection_id){
  selectDbKey.value=db
  emits('emit-select-db',Number(db.substring(2)),connection_id)
}
</script>

<style scoped>

</style>