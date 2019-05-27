import axios from "axios";

export default class ApiClient {
  config = {
    //baseURL: 'https://some-domain.com/api/',
    timeout: 2000,
    headers: { "X-Custom-Header": "foobar" }
  };

  instance = null;

  constructor(config = null) {
    if (config !== null) {
      this.config = config;
    }

    this.instance = axios.create(this.config);
  }

  all(url) {
    return this.instance.get(url);
  }

  find(id) {
    return id;
  }

  create(url, payload) {
      return this.instance.post(url, payload);
  }

  update(url, payload) {
      return this.instance.put(url, payload);
  }

  delete(url) {
    return this.instance.delete(url);
  }
}
