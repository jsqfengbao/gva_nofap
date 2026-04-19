const fs = require('fs');
const content = fs.readFileSync('./dist/build/mp-weixin/config/env.js', 'utf8');
console.log(content);
