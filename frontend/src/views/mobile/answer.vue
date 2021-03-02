<template>
  <div class="mobile-answer">
    <div class="exam-name">{{ exam.name }}({{ exam.time }}分钟)</div>
    <div>
      <van-field v-model="studentName"
                 label="姓名:"
                 placeholder="姓名" />
    </div>

    <div class="exam-content">
      <div class="question-item"
           v-for="(item,index) in questionList"
           :key="index">
        <div class="question-title">{{index + 1}}. {{ item.title }}</div>
        <div class="question-answer">
          <template v-if="item.questionType === questionConfig.TYPE_CHOOSE">
            <van-field name="radio">
              <template #input>
                <van-radio-group v-model="answerList[item.questionId]"
                                 direction="horizontal">
                  <van-radio :name="chooseIndex + 1"
                             v-for="(chooseItem,chooseIndex) in JSON.parse(item.option)"
                             :key="chooseIndex">{{ chooseItem }}</van-radio>
                </van-radio-group>
              </template>
            </van-field>
          </template>
          <template v-else>
            <van-field v-model="answerList[item.questionId]"
                       rows="2"
                       autosize
                       type="textarea"
                       placeholder="请输入" />
          </template>
        </div>
        <van-divider />
      </div>
    </div>

    <div class="exam-action">
      <van-button type="primary"
                  round
                  @click="handleSubmit()">提交考卷</van-button>
    </div>
  </div>
</template>

<script>
import * as questionConfig from "@/config/question";
import { fetchExamDetail } from "@/api/exam";
import { addAnswer } from "@/api/answer";

export default {
  name: "mobileAnswer",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      loading: false,
      questionConfig: questionConfig,
      examId: "",
      exam: {},
      questionList: [],

      studentName: "",
      answerList: {},
    };
  },
  watch: {},
  created() {
    this.examId = this.$route.query.examId || "";
    if (this.examId) {
      this.handleDetail();
    }
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    handleDetail() {
      this.loading = true;
      fetchExamDetail({
        examId: this.examId,
      })
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.exam = data.exam;
          this.questionList = data.questionList;

          // 答案
          this.questionList.forEach((item) => {
            this.answerList[item.questionId] = "";
          });
        })
        .finally(() => {
          this.loading = false;
        });
    },
    handleSubmit() {
      if (!this.studentName) {
        this.$toast("请填写姓名");
        return;
      }
      this.$toast.loading({
        duration: 0, // 持续展示 toast
        forbidClick: true, // 禁用背景点击
        loadingType: "spinner",
      });
      addAnswer({
        examId: this.exam.examId,
        teacherId: this.exam.userId,
        studentName: this.studentName,
        answerList: JSON.stringify(this.answerList),
      })
        .then((resp) => {
          const { ret, msg } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.$router.replace("/mobile/success");
        })
        .finally(() => {
          this.$toast.clear();
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.mobile-answer {
  min-height: 100vh;
  background: #f2f2f2;
  color: #323232;

  .exam-name {
    padding: 10px;
    text-align: center;
  }

  .exam-content {
    margin-top: 10px;
    background: #ffffff;

    .question-item {
      padding-top: 10px;

      .van-divider {
        margin-top: 10px;
        margin-bottom: 10px;
      }
    }
  }

  .exam-action {
    padding: 16px;

    .van-button {
      width: 100%;
    }
  }
}
</style>
