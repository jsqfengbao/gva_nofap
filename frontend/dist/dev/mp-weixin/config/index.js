"use strict";
const config_env = require("./env.js");
function getApiUrl(path = "") {
  return config_env.buildApiUrl(path);
}
exports.getApiUrl = getApiUrl;
