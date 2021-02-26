import request from "@/utils/request";

export function fetchAnswerList(data) {
  return request({
    url: "/answer/list",
    method: "post",
    data,
  });
}
