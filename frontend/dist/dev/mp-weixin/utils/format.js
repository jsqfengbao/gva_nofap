"use strict";
function formatNumber(num) {
  if (!num) return "0";
  if (num >= 1e6) {
    return (num / 1e6).toFixed(1) + "M";
  }
  if (num >= 1e3) {
    return (num / 1e3).toFixed(1) + "k";
  }
  return num.toString();
}
function getCurrentTime() {
  const now = /* @__PURE__ */ new Date();
  return `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
}
exports.formatNumber = formatNumber;
exports.getCurrentTime = getCurrentTime;
