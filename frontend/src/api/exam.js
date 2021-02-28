import request from "@/utils/request";

export function fetchExamList(data) {
  return request({
    url: "/exam/list",
    method: "post",
    data,
  });
}

export function fetchExamDetail(data) {
  return request({
    url: "/exam/detail",
    method: "post",
    data,
  });
}

export function addExam(data) {
  return request({
    url: "/exam/add",
    method: "post",
    data,
  });
}

export function editExam(data) {
  return request({
    url: "/exam/edit",
    method: "post",
    data,
  });
}

export function deleteExam(data) {
  return request({
    url: "/exam/delete",
    method: "post",
    data,
  });
}
