<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>试卷列表</span>
      </div>
    </template>
    <div class="exam-list">
      <div>
        <el-form :inline="true"
                 :model="formSearch">
          <el-form-item>
            <el-button type="primary"
                       @click="getList">刷新</el-button>
            <el-button type="success"
                       @click="handleAdd">新增</el-button>
          </el-form-item>
        </el-form>
      </div>

      <el-table :data="list"
                border>
        <el-table-column prop="examId"
                         label="ID">
        </el-table-column>
        <el-table-column label="试卷名称"
                         prop="name">
        </el-table-column>
        <el-table-column prop="time"
                         label="考试时间">
          <template #default="scope">
            {{ scope.row.time }}分钟
          </template>
        </el-table-column>

        <el-table-column prop="time"
                         label="考试链接">
          <template #default="scope">
            <el-popover placement="top-start"
                        trigger="hover">
              <img :src="qrList[scope.$index]"
                   alt="">
              <template #reference>
                <el-button size="mini">二维码</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>

        <el-table-column label="操作"
                         width="150">
          <template #default="scope">
            <el-button type="primary"
                       size="mini"
                       @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="danger"
                       size="mini"
                       @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination @current-change="handleCurrentChange"
                     background
                     layout="total, prev, pager, next"
                     :total="total">
      </el-pagination>

    </div>
  </el-card>
</template>

<script>
import QRCode from "qrcode";
import { fetchExamList, deleteExam } from "@/api/exam";

export default {
  name: "examList",
  props: {},
  computed: {},
  data() {
    return {
      listLoading: false,

      formSearch: {
        examType: "",
        page: 1,
        pageSize: 10,
      },
      list: [],
      total: 0,
      qrList: [],
    };
  },
  watch: {},
  created() {
    this.getList();
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    getList() {
      this.listLoading = true;
      fetchExamList(this.formSearch)
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.list = data.list;
          this.total = data.total;
          this.createQr();
        })
        .finally(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(val) {
      this.formSearch.page = val;
      this.getList();
    },
    // 二维码地址
    createQr() {
      this.list.forEach((item, key) => {
        const path = `${location.protocol}//${location.hostname}:${location.port}/mobile/answer?examId=${item.examId}`;
        QRCode.toDataURL(path).then((url) => {
          this.qrList[key] = url;
        });
      });
    },
    // 新增
    handleAdd() {
      this.$router.push({
        path: "/exam/edit",
      });
    },
    // 编辑
    handleEdit(item) {
      this.$router.push({
        path: "/exam/edit",
        query: {
          examId: item.examId,
        },
      });
    },
    // 删除
    handleDelete(item) {
      this.$confirm(`确认删除"${item.name}"?`).then((_) => {
        this.listLoading = true;
        deleteExam({ examId: item.examId })
          .then((resp) => {
            const { ret, msg } = resp;
            if (ret) {
              this.$message.error(msg);
              return;
            }
            this.$message.success("删除成功");
            this.getList();
          })
          .finally(() => {
            this.listLoading = false;
          })
          .catch(() => {});
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.el-pagination {
  margin-top: 20px;
}
</style>
