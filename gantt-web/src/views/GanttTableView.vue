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
        <el-table-column prop="20240910" label="2024-09-10" width="180" />
        <el-table-column prop="20240911" label="2024-09-11" width="180" />
      </el-table>
    </main>
  </div>
</template>

<script>
export default {
  name: "GanttTableView",
  data() {
    return {
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