const data = require('./data.json');

for (let d of data) {
  console.log(`* [${d.id}. ${d.title}](problems/${d.id}/explanation.md)`);
}