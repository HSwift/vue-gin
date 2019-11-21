import axios from "axios";
import qs from "qs";

function requestLogin(username, password, code, captchaid) {
  const url = "/api/login/" + username;
  const postData = { password: password, code: code, captchaid: captchaid };
  return axios.post(url, qs.stringify(postData));
}

function requestRegister(username, password, code, captchaid) {
  const url = "/api/register/" + username;
  const postData = { password: password, code: code, captchaid: captchaid };
  return axios.post(url, qs.stringify(postData));
}

function requestCaptcha() {
  const url = "/captcha";
  return axios.get(url);
}

export { requestLogin, requestRegister, requestCaptcha };
