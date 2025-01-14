<template>
  <div id="gantt_table_view">
    <header>
      <h1>任务列表</h1>
      <div>
        <el-button @click="dialogAdd=true">新增</el-button>
      </div>
    </header>
    <main>
      <gantt-table-list ref="ganttTable" />
    </main>
    <el-dialog v-model="dialogAdd">
      <gantt-table-add @submitted="handleAddSubmitted" />
    </el-dialog>
  </div>
</template>

<script>
import GanttTableList from "@/components/GanttTableList.vue";
import GanttTableAdd from "@/components/GanttTableAdd.vue";
export default {
  name: "GanttTableView",
  components: {GanttTableAdd, GanttTableList},
  data() {
    return {
      dialogAdd: false
    }
  },
  created() {
  },
  methods: {
    handleAddSubmitted() {
      this.$refs.ganttTable.refresh()
      this.dialogAdd = false
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