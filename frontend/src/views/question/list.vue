<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>题目列表</span>
      </div>
    </template>
    <div class="question-list">
      <div>
        <el-form :inline="true"
                 :model="formSearch">
          <el-form-item label="题目类别:">
            <el-select v-model="formSearch.questionType"
                       placeholder="请选择"
                       clearable>
              <el-option :label="item"
                         v-for="(item,index) in questionConfig.typeLabels"
                         :key="index"
                         :value="index"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       @click="getList">查询</el-button>
            <el-button type="success"
                       @click="handleAdd">新增</el-button>
          </el-form-item>
        </el-form>
      </div>

      <el-table :data="list"
                border>
        <el-table-column prop="questionId"
                         label="ID">
        </el-table-column>
        <el-table-column label="类别">
          <template #default="scope">
            {{ questionConfig.typeLabels[scope.row.questionType] }}
          </template>
        </el-table-column>
        <el-table-column prop="title"
                         label="题目">
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
import * as questionConfig from "@/config/question";
import { fetchQuestionList, deleteQuestion } from "@/api/question";

export default {
  name: "questionList",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      listLoading: false,
      questionConfig: questionConfig,

      formSearch: {
        questionType: "",
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
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    getList() {
      this.listLoading = true;
      fetchQuestionList(this.formSearch)
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
    // 新增
    handleAdd() {
      this.$router.push({
        path: "/question/edit",
      });
    },
    // 编辑
    handleEdit(item) {
      this.$router.push({
        path: "/question/edit",
        query: {
          questionId: item.questionId,
        },
      });
    },
    // 删除
    handleDelete(item) {
      this.$confirm(`确认删除"${item.title}"?`).then((_) => {
        this.listLoading = true;
        deleteQuestion({ questionId: item.questionId })
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
