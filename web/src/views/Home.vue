<template>
  <el-collapse>
    <el-collapse-item v-for="task in data.tasks" :title="task.name" :name="task.name">
      <el-card>
        <template #header>
          <div class="card-header">
            <span>{{ task.name }}</span>
            <el-button class="button" type="text">Edit</el-button>
          </div>
        </template>
        <el-row :gutter="12">
          <el-col :span="12" v-for="trigger in task.triggers">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>{{ trigger.name }}</span>
                  <el-popconfirm title="Are you sure to run?" @confirm="runTrigger(trigger)">
                    <template #reference>
                      <el-button class="button" type="text">Run</el-button>
                    </template>
                  </el-popconfirm>
                  <el-button class="button" type="text">Edit</el-button>
                </div>
              </template>
              <span v-for="run in trigger.recent_runs">
              <el-tag v-if="run.status===2 && run.exit_code===0"
                      style="margin: 4px 4px 4px 4px;" type="success">{{ getFromNow(run.created_at) }}</el-tag>
              <el-tag v-if="run.status===2 && run.exit_code===1"
                      style="margin: 4px 4px 4px 4px;" type="danger">{{ getFromNow(run.created_at) }}</el-tag>
            </span>
            </el-card>
          </el-col>
        </el-row>
      </el-card>
    </el-collapse-item>
  </el-collapse>
</template>

<script lang="ts" setup>

import {reactive} from "vue";
import axios from "axios";
import {ElMessage} from "element-plus";
import moment from "moment/moment";

const data = reactive({
  tasks: []
})

const runTrigger = (e: any) => {
  console.log(e)
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

    data.tasks = res.data.data
  })
}

function getFromNow(t: moment.MomentInput | undefined) {
  return moment(t).fromNow()
}

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

.text {
  font-size: 14px;
}

.box-card {
  width: 480px;
  margin-top: 8px;
  margin-bottom: 8px;
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
