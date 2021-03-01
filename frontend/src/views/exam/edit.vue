<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <el-page-header @back="$router.back()"
                        :content="`试卷${formData.examId ? '编辑' : '新增'}`">
        </el-page-header>
      </template>
      <div class="quesion-edit"
           v-loading="loading">
        <el-row>
          <el-col :span="20">
            <el-form ref="form"
                     :model="formData"
                     :rules="rules"
                     label-width="120px">

              <el-form-item label="试卷名称:"
                            placeholder="请填写试卷名称"
                            prop="name">
                <el-input v-model.trim="formData.name"></el-input>
              </el-form-item>

              <el-form-item label="考试时间:"
                            placeholder="请填写考试时间"
                            prop="time">
                <el-input v-model.trim="formData.time">
                  <template #append>分钟</template>
                </el-input>
              </el-form-item>

              <el-form-item label="试卷试题:"
                            prop="questionList">
                <el-button type="success"
                           @click="handleOpenDialog()"
                           size="mini">+ 添加试题</el-button>

                <el-table :data="formData.questionList"
                          border>
                  <el-table-column prop="questionId"
                                   label="ID"
                                   width="100">
                  </el-table-column>
                  <el-table-column label="类别"
                                   width="120">
                    <template #default="scope">
                      {{ questionConfig.typeLabels[scope.row.questionType] }}
                    </template>
                  </el-table-column>
                  <el-table-column prop="title"
                                   label="题目">
                  </el-table-column>

                  <el-table-column label="操作"
                                   width="100">
                    <template #default="scope">
                      <el-button type="danger"
                                 size="mini"
                                 @click="handleDeleteQuestion(scope.$index)">删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-button type="default"
                     @click="$router.back()">返回</el-button>
          <el-button type="primary"
                     @click="handleSave()">保存</el-button>
        </el-row>
      </div>
    </el-card>

    <el-dialog title="添加试题"
               v-model="dialogVisible">

      <el-select placeholder="请选择试题"
                 v-model="questionId"
                 filterable
                 style="width: 100%;">
        <el-option v-for="(item,index) in questionList"
                   :key="index"
                   :value="item.questionId"
                   :label="`${item.questionId}: ${item.title}`"></el-option>
      </el-select>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary"
                     @click="handleAddQuestion()"
                     :disabled="questionId <= 0">添加</el-button>
        </span>
      </template>
    </el-dialog>

  </div>

</template>

<script>
import * as questionConfig from "@/config/question";
import { fetchExamDetail, addExam, editExam } from "@/api/exam";
import { fetchQuestionList } from "@/api/question";

export default {
  name: "quesionEdit",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      loading: false,
      questionConfig: questionConfig,
      formData: {
        examId: "",
        name: "",
        time: "",
        questionList: [],
      },
      rules: {
        name: [{ required: true, message: "请输入试卷名称", trigger: "blur" }],
        time: [{ required: true, message: "请输入考试名称", trigger: "blur" }],
        questionList: [
          { required: true, message: "请添加试卷试题", trigger: "blur" },
        ],
      },
      dialogVisible: false,
      questionId: "",
      questionList: {},
    };
  },
  watch: {},
  created() {
    this.formData.examId = this.$route.query.examId || "";
    if (this.formData.examId) {
      this.handleDetail();
    }
    this.getQuestionList();
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    handleDetail() {
      this.loading = true;
      fetchExamDetail({
        examId: this.formData.examId,
      })
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.formData = Object.assign(data.exam);
          this.formData.questionList = data.questionList;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    getQuestionList() {
      this.listLoading = true;
      return fetchQuestionList({
        pageSize: 10000,
      })
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.questionList = data.list;
        })
        .finally(() => {
          this.listLoading = false;
        });
    },
    handleSave() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.loading = true;
          const actionApi = this.formData.examId ? editExam : addExam;
          // 参数
          const params = Object.assign({}, this.formData);
          params.questionList = params.questionList.map(
            (item) => item.questionId
          );
          params.questionList = [...params.questionList].join(",");
          actionApi(params)
            .then((resp) => {
              const { ret, msg } = resp;
              if (ret) {
                this.$message.error(msg);
                return;
              }
              this.$message.success("保存成功");
              this.$router.back();
            })
            .finally(() => {
              this.loading = false;
            });
        }
      });
    },
    // 添加试题
    handleOpenDialog() {
      this.dialogVisible = true;
      this.questionId = "";
    },
    // 添加试题
    handleAddQuestion() {
      const index = this.questionList.findIndex(
        (item) => item.questionId === this.questionId
      );
      if (index >= 0) {
        this.formData.questionList.push(this.questionList[index]);
      }
      this.dialogVisible = false;
    },
    // 添加试题
    handleDeleteQuestion(index) {
      this.formData.questionList.splice(index, 1);
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
