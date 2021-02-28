import request from "@/utils/request";

export function fetchQuestionList(data) {
  return request({
    url: "/question/list",
    method: "post",
    data,
  });
}

export function fetchQuestionDetail(data) {
  return request({
    url: "/question/detail",
    method: "post",
    data,
  });
}

export function addQuestion(data) {
  return request({
    url: "/question/add",
    method: "post",
    data,
  });
}

export function editQuestion(data) {
  return request({
    url: "/question/edit",
    method: "post",
    data,
  });
}

export function deleteQuestion(data) {
  return request({
    url: "/question/delete",
    method: "post",
    data,
  });
}
