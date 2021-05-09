const isLeapYear = (v) => {
  if (v % 4 != 0) return false;
  if (v % 100 == 0) return v % 400 == 0;
  return true;
};

const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

readline.question('Please enter the year: ', (v) => {
  if (v == null || v == '' || Number.isNaN(Number(v))) {
    console.log(`'${v}' is not integer`);
    process.exit(1);
  }
  if (isLeapYear(Number(v))) {
    console.log(`'${v}' is leap year`);
  } else {
    console.log(`'${v}' is not leap year`);
  }
  readline.close();
});
