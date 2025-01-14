<template>
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
</template>

<script>
import { getDateList } from "@/apis/date";
import { getTaskList } from "@/apis/task";
export default {
  name: "GanttTableList",
  data() {
    return {
      dateList: [],
      tableData: []
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
    refresh() {
      this.handleGetDateList()
      this.handleGetTaskList()
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
</style>