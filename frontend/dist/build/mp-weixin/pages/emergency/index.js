"use strict";const e=require("../../common/vendor.js"),S={name:"EmergencyHelp",setup(){const s=e.ref("9:41"),u=e.ref(!0),h=e.ref(23),o=e.ref(!1),a=e.ref(!1),r=e.reactive({title:"",content:"",id:0}),t=e.ref({text:"你比你想象的更坚强，这个时刻会过去的。",author:"戒色社区"}),v=[{text:"你比你想象的更坚强，这个时刻会过去的。",author:"戒色社区"},{text:"每个选择都是新的开始，每一刻都有改变的可能。",author:"戒色助手"},{text:"成长需要时间，对自己要有耐心。",author:"励志语录"}],i=e.ref([{id:1,title:"4-7-8深呼吸练习",type:1,content:`这是一种简单而有效的呼吸技巧，可以帮助您快速平静下来。

步骤：
1. 吸气4秒钟
2. 屏住呼吸7秒钟
3. 呼气8秒钟
4. 重复3-4次

这种呼吸方式可以激活副交感神经系统，帮助身体放松。`,duration:180},{id:2,title:"5分钟正念冥想",type:2,content:`正念冥想可以帮助您专注当下，减少负面情绪的影响。

指导：
1. 找一个安静的地方坐下
2. 闭上眼睛，专注于呼吸
3. 当思绪飘散时，温和地把注意力拉回呼吸
4. 观察身体的感觉，不做判断
5. 保持5分钟

记住：没有'正确'或'错误'的冥想，只要保持观察即可。`,duration:300},{id:5,title:"快速运动指导",type:5,content:`运动可以释放内啡肽，帮助改善心情。这里有一套简单的运动，无需器械：

**热身（1分钟）**
• 原地踏步 30秒
• 手臂绕圈 30秒

**主要运动（3分钟）**
• 俯卧撑 20个（可膝盖着地）
• 深蹲 20个
• 平板支撑 30秒
• 开合跳 20个

**放松（1分钟）**
• 深呼吸 30秒
• 拉伸手臂和腿部 30秒`,duration:300}]),g=()=>{const n=new Date,c=n.getHours().toString().padStart(2,"0"),d=n.getMinutes().toString().padStart(2,"0");s.value=`${c}:${d}`},f=()=>{const n=Math.floor(Math.random()*v.length);t.value=v[n]},m=()=>{e.index.navigateBack()},M=()=>{e.index.showModal({title:"紧急联系",content:"如果遇到生命危险，请立即拨打120或当地急救电话",confirmText:"知道了",showCancel:!1})},w=()=>{const n=i.value.find(c=>c.type===1);n&&l(n)},R=()=>{const n=i.value.find(c=>c.type===5);n&&l(n)},H=n=>{let c;switch(n){case"meditation":c=i.value.find(d=>d.type===2);break;case"music":e.index.showToast({title:"音乐功能开发中",icon:"none"});return;case"reading":e.index.navigateTo({url:"/pages/emergency/articles"});return;case"puzzle":e.index.showToast({title:"游戏功能开发中",icon:"none"});return}c&&l(c)},l=n=>{r.title=n.title,r.content=n.content,r.id=n.id,a.value=!0},x=()=>{a.value=!1},T=()=>{e.index.showToast({title:"开始使用",icon:"success"}),x()},E=()=>{e.index.showToast({title:"感谢你的反馈",icon:"success"})},_=()=>{o.value=!0},y=()=>{o.value=!1},B=n=>{y(),e.index.showModal({title:"求助已发送",content:"你的求助信息已发送给在线志愿者，他们会尽快回复你。",confirmText:"好的",showCancel:!1,success:()=>{e.index.showToast({title:"正在为您匹配志愿者",icon:"loading",duration:2e3})}})},p=()=>{e.index.navigateTo({url:"/pages/community/index?category=4"})};return e.onMounted(()=>{g(),setInterval(g,6e4),setInterval(f,3e4)}),{currentTime:s,isBreathing:u,onlineVolunteers:h,currentQuote:t,showHelpModal:o,showResourceModal:a,selectedResource:r,goBack:m,callEmergency:M,startBreathingExercise:w,startPhysicalExercise:R,openActivity:H,requestHelp:_,closeHelpModal:y,selectHelpType:B,viewStories:p,closeResourceModal:x,useResource:T,rateResource:E}}};function k(s,u,h,o,a,r){return e.e({a:e.t(o.currentTime),b:e.o((...t)=>o.goBack&&o.goBack(...t)),c:e.o((...t)=>o.callEmergency&&o.callEmergency(...t)),d:o.isBreathing?1:"",e:e.o((...t)=>o.startBreathingExercise&&o.startBreathingExercise(...t)),f:e.o((...t)=>o.startPhysicalExercise&&o.startPhysicalExercise(...t)),g:e.o(t=>o.openActivity("puzzle")),h:e.o(t=>o.openActivity("meditation")),i:e.o(t=>o.openActivity("music")),j:e.o(t=>o.openActivity("reading")),k:e.t(o.onlineVolunteers),l:e.o((...t)=>o.requestHelp&&o.requestHelp(...t)),m:e.o((...t)=>o.viewStories&&o.viewStories(...t)),n:e.t(o.currentQuote.text),o:e.t(o.currentQuote.author),p:o.showHelpModal},o.showHelpModal?{q:e.o((...t)=>o.closeHelpModal&&o.closeHelpModal(...t)),r:e.o(t=>o.selectHelpType(1)),s:e.o(t=>o.selectHelpType(2)),t:e.o(t=>o.selectHelpType(3)),v:e.o(t=>o.selectHelpType(4)),w:e.o(()=>{}),x:e.o((...t)=>o.closeHelpModal&&o.closeHelpModal(...t))}:{},{y:o.showResourceModal},o.showResourceModal?{z:e.t(o.selectedResource.title),A:e.o((...t)=>o.closeResourceModal&&o.closeResourceModal(...t)),B:e.t(o.selectedResource.content),C:e.o((...t)=>o.useResource&&o.useResource(...t)),D:e.o((...t)=>o.rateResource&&o.rateResource(...t)),E:e.o(()=>{}),F:e.o((...t)=>o.closeResourceModal&&o.closeResourceModal(...t))}:{})}const q=e._export_sfc(S,[["render",k],["__scopeId","data-v-53254f78"]]);wx.createPage(q);
