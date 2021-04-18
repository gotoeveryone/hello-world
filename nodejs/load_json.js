const fs = require('fs');
const path = require('path');

const jsonDir = path.dirname(__dirname);
const jsonData = JSON.parse(fs.readFileSync(`${jsonDir}/test.json`));
jsonData.forEach((v) => console.log(`id: ${v.id}, name: ${v.name}`));
