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
                <div class="question-item"
                     v-for="(questionItem, questionIndex) in questionList"
                     :key="questionIndex">
                  <h3>{{ questionIndex + 1 }}. {{questionItem.title}}</h3>
                  <div class="question-answer">
                    <!-- 1.选择题 -->
                    <template v-if="questionItem.questionType == questionConfig.TYPE_CHOOSE">
                      <template v-if="isAnswerEmpty(questionItem.questionId)">
                        <div>未填写</div>
                      </template>

                      <div v-for="(chooseItem, chooseIndex) in JSON.parse(questionItem.option)"
                           :key="chooseIndex"
                           class="question-choose-item"
                           :class="{
                             'question-choose-active': getChooseActive(questionItem.questionId, chooseIndex),
                           }">
                        {{ chooseIndex + 1 }}: {{chooseItem }}
                      </div>
                    </template>

                    <!-- 2.填空题 -->
                    <template v-else>
                      <template v-if="isAnswerEmpty(questionItem.questionId)">
                        未填写内容
                      </template>
                      <template v-else>
                        {{ getInputContent(questionItem.questionId) }}
                      </template>
                    </template>
                  </div>
                </div>
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
      answerList: {},
      questionList: [],
      activeNames: [],
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
          this.answerList = JSON.parse(data.answer.answerList);
          this.questionList = data.questionList;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    isAnswerEmpty(questionId) {
      if (!this.answerList.hasOwnProperty(questionId)) {
        return true;
      }
      return false;
    },
    // 填空题答案
    getInputContent(questionId) {
      if (this.answerList[questionId] == "") {
        return "未填写内容";
      }
      return this.answerList[questionId];
    },
    // 选择题选的哪个
    getChooseActive(questionId, chooseIndex) {
      if (this.isAnswerEmpty(questionId)) {
        return false;
      }
      return this.answerList[questionId] === chooseIndex + 1;
    },
  },
};
</script>

<style lang="scss" scoped>
.question-item {
  h3 {
    margin-top: 0;
    margin-bottom: 0;
  }

  .question-answer {
    border: 1px solid #e4e7ed;
    padding: 5px;

    .question-choose-item {
      color: #909399;

      &.question-choose-active {
        color: #409eff;
      }
    }
  }
}
</style>
