<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>答题列表</span>
      </div>
    </template>
    <div class="answer-list">
      <div>
        <el-form :inline="true"
                 :model="formSearch">
          <el-form-item label="试卷:">
            <el-select v-model="formSearch.examId"
                       placeholder="请选择"
                       clearable>
              <el-option v-for="(item,index) in allExamList"
                         :key="index"
                         :label="item.name"
                         :value="item.examId"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="getList">查询</el-button>
          </el-form-item>
        </el-form>
      </div>

      <el-table :data="list"
                border>
        <el-table-column prop="answerId"
                         label="ID">
        </el-table-column>

        <el-table-column label="试卷名"
                         prop="examName">
        </el-table-column>

        <el-table-column label="学生名"
                         prop="userName">
        </el-table-column>

        <el-table-column label="交卷时间"
                         width="200">
          <template #default="scope">
            {{ formatDateTime(scope.row.createTime) }}
          </template>
        </el-table-column>

        <el-table-column label="操作"
                         width="100">
          <template #default="scope">
            <el-button type="success"
                       size="mini"
                       @click="handleEdit(scope.row)">详情</el-button>
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
import { formatDateTime } from "@/utils/datetime";
import { fetchAnswerList } from "@/api/answer";
import { fetchExamList } from "@/api/exam";

export default {
  name: "answerList",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      listLoading: false,
      allExamList: [],
      formSearch: {
        examId: "",
        page: 1,
        pageSize: 10,
      },
      list: [],
      total: 0,
    };
  },
  watch: {},
  created() {
    this.getList();
    this.getExamList();
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    formatDateTime,
    getList() {
      this.listLoading = true;
      fetchAnswerList(this.formSearch)
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.list = data.list;
          this.total = data.total;
        })
        .finally(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(val) {
      this.formSearch.page = val;
      this.getList();
    },
    getExamList() {
      fetchExamList({
        pageSize: 10000,
      }).then((resp) => {
        const { ret, msg, data } = resp;
        if (ret) {
          this.$message.error(msg);
          return;
        }
        this.allExamList = data.list;
      });
    },
    // 详情
    handleDetail() {
      this.$router.push({
        path: "/answer/detail",
        query: {
          answerId: item.answerId,
        },
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
