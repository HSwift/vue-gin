import axios from "axios";
import qs from "qs";

function requestLogin(username, password) {
  const url = "/api/login/" + username;
  const postData = { password: password };
  return axios.post(url, qs.stringify(postData));
}

function requestRegister(username, password) {
  const url = "/api/register/" + username;
  const postData = { password: password };
  return axios.post(url, qs.stringify(postData));
}

export { requestLogin, requestRegister };
