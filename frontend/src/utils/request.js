import axios from "axios";
import { getToken } from "@/utils/auth";

// create an axios instance
const service = axios.create({
  baseURL: "/api", // url = base url + request url
  withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000, // request timeout
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
    token: "",
  },
});

// request interceptor
service.interceptors.request.use(
  (config) => {
    // token
    config.headers.token = getToken();

    // json 参数转form-urlencoded
    config["transformRequest"] = [
      function(data) {
        if (typeof data === "string") {
          try {
            data = JSON.parse(data);
          } catch (err) {
            throw new Error("参数格式异常，解析失败");
          }
        }
        let ret = "";
        for (const it in data) {
          ret +=
            encodeURIComponent(it) + "=" + encodeURIComponent(data[it]) + "&";
        }
        return ret;
      },
    ];
    return config;
  },
  (error) => {
    console.log(error); // for debug
    return Promise.reject(error);
  }
);

// response interceptor
service.interceptors.response.use(
  (response) => {
    window.console.log(response);
    return response.data;
  },
  (error) => {
    console.log("err" + error); // for debug
    return Promise.reject(error);
  }
);

export default service;
