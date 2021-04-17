Array.from(new Array(30).keys(), (n) => {
  const v = n + 1;
  const m = [];
  if (v % 3 === 0) {
    m.push('fizz');
  }
  if (v % 5 === 0) {
    m.push('buzz');
  }
  console.log(m.length ? m.join('') : v.toString());
});
