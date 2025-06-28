"use strict";
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};
require("../common/vendor.js");
function installInterceptors() {
  return __async(this, null, function* () {
    console.log("🚀 开始安装拦截器...");
    try {
      prototypeInterceptor.install();
      console.log("✅ 原型拦截器安装完成");
      requestInterceptor.install();
      console.log("✅ 请求拦截器安装完成");
      yield routeInterceptor.install();
      console.log("✅ 路由拦截器安装完成");
      console.log("🎉 所有拦截器安装完成");
    } catch (error) {
      console.error("❌ 拦截器安装失败:", error);
      throw error;
    }
  });
}
exports.installInterceptors = installInterceptors;
