<template>
  <el-card class="box-card">
    <template #header>
      <el-page-header @back="$router.back()"
                      :content="`题目${formData.questionId ? '编辑' : '新增'}`">
      </el-page-header>
    </template>
    <div class="quesion-edit"
         v-loading="loading">
      <el-row>
        <el-col :span="16">
          <el-form ref="form"
                   :model="formData"
                   :rules="rules"
                   label-width="120px">

            <el-form-item label="试题类别:"
                          prop="questionType">
              <el-select v-model="formData.questionType"
                         placeholder="请选择试题类别"
                         style="width: 100%">
                <el-option :label="item"
                           v-for="(item,index) in questionConfig.typeLabels"
                           :key="index"
                           :value="index * 1"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="试题题目:"
                          placeholder="请填写试题题目"
                          prop="title">
              <el-input v-model.trim="formData.title"></el-input>
            </el-form-item>

            <el-form-item v-if="formData.questionType == questionConfig.TYPE_CHOOSE"
                          label="试题选项:"
                          prop="option">
              <div :key="index"
                   v-for="(tag,index) in formData.option">
                <el-tag closable
                        :disable-transitions="false"
                        @close="handleClose(index)">
                  {{tag}}
                </el-tag>
              </div>

              <el-input class="input-new-tag"
                        v-if="optionInputVisible"
                        v-model="optionInputValue"
                        placeholder="请输入选项内容，如:（A: 以上都不对）"
                        ref="saveTagInput"
                        size="small"
                        @keyup.enter="handleInputConfirm"
                        @blur="handleInputConfirm">
              </el-input>
              <el-button v-else
                         type="success"
                         size="small"
                         @click="handleShowInput">+ 添加选项</el-button>
            </el-form-item>

            <el-form-item label="试题备注:"
                          placeholder="请填写试题备注"
                          prop="remark">
              <el-input type="textarea"
                        rows="3"
                        v-model.trim="formData.remark"></el-input>
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
</template>

<script>
import * as questionConfig from "@/config/question";
import { fetchQuestionDetail, addQuestion, editQuestion } from "@/api/question";

export default {
  name: "quesionEdit",
  components: {},
  props: {},
  computed: {},
  data() {
    return {
      questionConfig: questionConfig,
      loading: false,
      formData: {
        questionId: "",
        questionType: "",
        title: "",
        remark: "",
        option: [],
      },
      rules: {
        questionType: [
          { required: true, message: "请选择试题类别", trigger: "blur" },
        ],
        title: [{ required: true, message: "请输入试题题目", trigger: "blur" }],
        option: [
          { required: true, message: "请输入试题选项", trigger: "blur" },
        ],
      },

      optionInputValue: "",
      optionInputVisible: false,
    };
  },
  watch: {},
  created() {
    this.formData.questionId = this.$route.query.questionId || "";
    if (this.formData.questionId) {
      this.handleDetail();
    }
  },
  mounted() {},
  beforeUnmount() {},
  methods: {
    handleDetail() {
      this.loading = true;
      fetchQuestionDetail({
        questionId: this.formData.questionId,
      })
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            this.$message.error(msg);
            return;
          }
          this.formData = Object.assign({}, data);
          if (this.formData.questionType == questionConfig.TYPE_CHOOSE) {
            this.formData.option = JSON.parse(this.formData.option);
          } else {
            this.formData.option = [];
          }
        })
        .finally(() => {
          this.loading = false;
        });
    },
    handleSave() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.loading = true;
          const actionApi = this.formData.questionId
            ? editQuestion
            : addQuestion;
          // 参数
          const params = Object.assign({}, this.formData);

          if (params.questionType == questionConfig.TYPE_CHOOSE) {
            params.option = JSON.stringify([...params.option]);
          } else {
            params.option = "";
          }
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
    handleClose(index) {
      this.formData.option.splice(index, 1);
    },
    handleShowInput() {
      this.optionInputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let optionInputValue = this.optionInputValue;
      if (optionInputValue) {
        this.formData.option.push(optionInputValue);
      }
      this.optionInputVisible = false;
      this.optionInputValue = "";
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
