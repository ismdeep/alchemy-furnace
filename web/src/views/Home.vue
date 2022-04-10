<template>
  <el-main>
    <el-dialog v-model="dialogData.visible"
               :title="dialogData.title"
               :before-close="dialogDestroy"
               destroy-on-close
               width="80%">
      <pre>test</pre>
    </el-dialog>

    <el-tabs v-if="data.tasks.length > 0"
             v-model="data.selectedTaskName"
             type="border-card"
             class="demo-tabs"
    >
      <el-tab-pane v-for="task in data.tasks" :label="task.name" :name="task.name">
        <div>
          <el-card v-for="trigger in task.triggers" style="margin-bottom: 8px">
            <template #header>
              <div class="card-header">
                <span>{{ trigger.name }}</span>
                <div>
                  <el-popconfirm title="Are you sure to run?" @confirm="runTrigger(task,trigger)">
                    <template #reference>
                      <el-button class="button" type="text">
                        <el-icon>
                          <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ba633cb8="">
                            <path fill="currentColor" d="M384 192v640l384-320.064z"></path>
                          </svg>
                        </el-icon>
                      </el-button>
                    </template>
                  </el-popconfirm>
                  <el-button class="button" type="text">
                    <el-icon>
                      <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ba633cb8="">
                        <path fill="currentColor"
                              d="M832 512a32 32 0 1 1 64 0v352a32 32 0 0 1-32 32H160a32 32 0 0 1-32-32V160a32 32 0 0 1 32-32h352a32 32 0 0 1 0 64H192v640h640V512z"></path>
                        <path fill="currentColor"
                              d="m469.952 554.24 52.8-7.552L847.104 222.4a32 32 0 1 0-45.248-45.248L477.44 501.44l-7.552 52.8zm422.4-422.4a96 96 0 0 1 0 135.808l-331.84 331.84a32 32 0 0 1-18.112 9.088L436.8 623.68a32 32 0 0 1-36.224-36.224l15.104-105.6a32 32 0 0 1 9.024-18.112l331.904-331.84a96 96 0 0 1 135.744 0z"></path>
                      </svg>
                    </el-icon>
                  </el-button>
                </div>
              </div>
            </template>
            <span v-for="run in trigger.recent_runs">
              <el-tag :type="getTagType(run.status,run.exit_code)" @click="showDialog(run)"
                      style="margin: 4px 4px 4px 4px;">
                <el-icon v-if="run.status===0 || run.status===1">
                  <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ba633cb8=""><path
                      fill="currentColor"
                      d="M512 64a32 32 0 0 1 32 32v192a32 32 0 0 1-64 0V96a32 32 0 0 1 32-32zm0 640a32 32 0 0 1 32 32v192a32 32 0 1 1-64 0V736a32 32 0 0 1 32-32zm448-192a32 32 0 0 1-32 32H736a32 32 0 1 1 0-64h192a32 32 0 0 1 32 32zm-640 0a32 32 0 0 1-32 32H96a32 32 0 0 1 0-64h192a32 32 0 0 1 32 32zM195.2 195.2a32 32 0 0 1 45.248 0L376.32 331.008a32 32 0 0 1-45.248 45.248L195.2 240.448a32 32 0 0 1 0-45.248zm452.544 452.544a32 32 0 0 1 45.248 0L828.8 783.552a32 32 0 0 1-45.248 45.248L647.744 692.992a32 32 0 0 1 0-45.248zM828.8 195.264a32 32 0 0 1 0 45.184L692.992 376.32a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0zm-452.544 452.48a32 32 0 0 1 0 45.248L240.448 828.8a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0z"></path></svg>
                </el-icon>
                {{ getFromNow(run.created_at) }}
              </el-tag>
            </span>
          </el-card>
        </div>
        <div style="text-align: right; margin-top: 8px">
          <el-button class="button" type="text">
            <el-icon>
              <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ba633cb8="">
                <path fill="currentColor"
                      d="M832 512a32 32 0 1 1 64 0v352a32 32 0 0 1-32 32H160a32 32 0 0 1-32-32V160a32 32 0 0 1 32-32h352a32 32 0 0 1 0 64H192v640h640V512z"></path>
                <path fill="currentColor"
                      d="m469.952 554.24 52.8-7.552L847.104 222.4a32 32 0 1 0-45.248-45.248L477.44 501.44l-7.552 52.8zm422.4-422.4a96 96 0 0 1 0 135.808l-331.84 331.84a32 32 0 0 1-18.112 9.088L436.8 623.68a32 32 0 0 1-36.224-36.224l15.104-105.6a32 32 0 0 1 9.024-18.112l331.904-331.84a96 96 0 0 1 135.744 0z"></path>
              </svg>
            </el-icon>
          </el-button>
        </div>
      </el-tab-pane>
    </el-tabs>
  </el-main>
</template>

<script lang="ts" setup>

import {reactive} from "vue";
import axios from "axios";
import {ElMessage} from "element-plus";
import moment from "moment/moment";

const data = reactive({
  tasks: [],
  selectedTaskName: ''
})

const dialogData = reactive({
  visible: false,
  title: 'Test',
})

const showDialog = (e: any) => {
  dialogData.visible = true
  console.log(e)

}

const dialogDestroy = (done: () => void) => {
  done()
}

const runTrigger = (taskInfo: any, e: any) => {
  console.log(e)
  axios.post(`/api/v1/tasks/${taskInfo.id}/triggers/${e.id}/runs`, null, {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('token')
    }
  }).then(() => {
    ElMessage.success("Success")
  })
}

function loadData() {
  axios.get(`/api/v1/tasks`, {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('token')
    }
  }).then((res) => {
    if (res.data.code !== 0) {
      ElMessage.error(res.data.msg)
      return
    }
    if (res.data.data.length > 0 && data.selectedTaskName === '') {
      data.selectedTaskName = res.data.data[0].name
    }
    data.tasks = res.data.data
  })
}

function getFromNow(t: moment.MomentInput | undefined) {
  return moment(t).fromNow()
}

function getTagType(status: any, exit_code: any) {
  if (status === 0 || status === 1) {
    return "info"
  }
  if (status === 2 && exit_code === 0) {
    return "success"
  }
  if (status === 2 && exit_code !== 0) {
    return "danger"
  }
  if (status === 3) {
    return "warning"
  }
  return "info"
}

// main
loadData()
setInterval(() => {
  loadData()
}, 1000)

</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>

<style lang="scss">
.masonry {
  width: 100%;
  margin: 10px auto;
  column-count: 2;
  column-gap: 20px;

  .item {
    width: 100%;
    break-inside: avoid;
    margin-bottom: 10px;
  }
}

@media screen and (max-width: 960px) {
  .masonry {
    columns: 1;
  }
}

pre {
  display: block;
  padding: 9.5px;
  margin: 0 0 0 0;
  font-size: 13px;
  line-height: 1.2;
  word-break: break-all;
  word-wrap: break-word;
  background-color: #f5f5f5;
  border: 1px solid #ccc;
  border-radius: 4px;
}
</style>
