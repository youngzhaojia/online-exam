<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <el-page-header @back="$router.back()"
                        content="答题详情">
        </el-page-header>
      </template>
      <div class="answer-edit"
           v-loading="loading">
        <el-row>
          <el-col :span="20">
            <el-form ref="form"
                     label-width="120px">
              <el-form-item label="学生名:">
                {{ answerDetail.studentName }}
              </el-form-item>

              <el-form-item label="试卷名:">
                {{ answerDetail.examName }}
              </el-form-item>

              <el-form-item label="交卷时间:">
                {{ formatDateTime(answerDetail.createTime) }}
              </el-form-item>

              <el-form-item label="答题内容:">

              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script>
import * as questionConfig from "@/config/question";
import { fetchAnswerDetail } from "@/api/answer";
import { formatDateTime } from "@/utils/datetime";

export default {
  name: "answerDetail",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      loading: false,
      questionConfig: questionConfig,
      answerId: "",
      answerDetail: {},
      questionList: [],
    };
  },
  watch: {},
  created() {
    this.answerId = this.$route.query.answerId || "";
    if (this.answerId) {
      this.handleDetail();
    }
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    formatDateTime,
    handleDetail() {
      this.loading = true;
      fetchAnswerDetail({
        answerId: this.answerId,
      })
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.answerDetail = data.answer;
          this.questionList = data.questionList.reduce(
            (accumulator, current) => {
              accumulator[current.questionId] = current;
              return accumulator;
            },
            {}
          );
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
