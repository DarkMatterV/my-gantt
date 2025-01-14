<template>
  <el-form v-model="info">
    <el-form-item label="No">
      <el-input v-model="info.no"></el-input>
    </el-form-item>
    <el-form-item label="任务">
      <el-input v-model="info.title"></el-input>
    </el-form-item>
    <el-form-item label="状态">
      <el-input v-model="info.status"></el-input>
    </el-form-item>
    <el-form-item label="进展">
      <el-input v-model="info.record_str"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button @click="handleSubmit">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {addTask} from "@/apis/task";
export default {
  name: "GanttTableAdd",
  data() {
    return {
      info: {
        no: "",
        title: "",
        status: "",
        record_str: "",
        pid: 0,
      }
    }
  },
  methods: {
    handleSubmit() {
      const data = {
        no: this.info.no,
        title: this.info.title,
        status: this.info.status,
        record: {},
        pid: this.info.pid
      }
      this.$message({message: '提交成功', type: 'success'})
      this.$emit('submitted')
      // addTask(data).then(resp => {
      //   this.$message({message: '提交成功', type: 'success'})
      //   this.$emit('close')
      // })
    }
  }
}
</script>

<style scoped>

</style>