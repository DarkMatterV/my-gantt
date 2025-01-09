<template>
  <div id="gantt_table_view">
    <header>
      <h1>任务列表</h1>
    </header>
    <main>
      <el-table :data="tableData" border row-key="id" default-expand-all>
        <el-table-column prop="no" label="No"/>
        <el-table-column prop="title" label="Title"/>
        <el-table-column prop="status" label="状态"/>
        <el-table-column v-for="item in dateList" :key="item.id" :label="`${item.date} ${item.week}`">
          <template #default="scope">
            <div>{{ displayRecord(scope.row, item.date) }}</div>
          </template>
        </el-table-column>
      </el-table>
    </main>
  </div>
</template>

<script>
import { getDateList } from "../apis/date";
import { getTaskList } from "../apis/task";
export default {
  name: "GanttTableView",
  data() {
    return {
      dateList: [],
      tableData: [
        {
          id: 1,
          no: '1',
          title: '任务1',
          status: 'Done',
          20240910: '',
          20240911: ''
        },
        {
          id: 2,
          no: '2',
          title: '任务2',
          status: 'Done',
          20240910: '',
          20240911: ''
        },
        {
          id: 3,
          no: '3',
          title: '任务3',
          status: 'Todo',
          20240910: '',
          20240911: '',
          children: [
            {
              id: 5,
              no: '3-1',
              title: '任务3-1',
              status: 'Todo',
              20240910: '',
              20240911: ''
            }
          ]
        },
        {
          id: 4,
          no: '4',
          title: '任务4',
          status: 'Done',
          20240910: '',
          20240911: ''
        }
      ]
    }
  },
  created() {
    this.handleGetDateList()
    this.handleGetTaskList()
  },
  methods: {
    handleGetDateList() {
      getDateList().then(resp => {
        this.dateList = resp.data || []
      })
    },
    handleGetTaskList() {
      getTaskList().then(resp => {
        this.tableData = resp.data || []
      })
    },
    displayRecord(row, date) {
      if (row && 'record' in row && date in row.record) {
        return row.record[date].status + " " + row.record[date].desc
      }
      return "-"
    }
  }
}
</script>

<style scoped>
#gantt_table_view {
}
#gantt_table_view header {
  text-align: center;
  margin-bottom: 10px;
}
#gantt_table_view main {

}
</style>