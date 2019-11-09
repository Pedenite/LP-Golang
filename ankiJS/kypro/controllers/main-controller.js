const fs = require('fs');
const readline = require('readline');

let fetchedBonds;

const getFile = (lineReader, cb) => {
  const bonds = [];
  lineReader.on('line', (line) => {
    const lineSplited = line.split('-');
    bonds.push({ key: lineSplited[0], value: lineSplited[1] });
  }).on('close', () => {
    fetchedBonds = bonds;
    cb();
  });
}

exports.getFetchFile = (req, res) => {
  res.render('index');
}

exports.postFetchFile =(req, res) => {
  const filePath = req.body.filePath;

  const lineReader = readline.createInterface({
    input: fs.createReadStream(filePath)
  });

  getFile(lineReader, () => {
    res.redirect('/bonds');
  });
}

exports.getBondsWorking = (req, res) => {
  const min = 0;
  const max = fetchedBonds.length - 1;
  const random = Math.floor(Math.random() * (+max - +min) + +min);

  res.redirect('/show-value/' + random);
}

exports.getShowValue = (req, res) => {
  const position = req.params.position;
  res.render('bonds-working', {
    key: fetchedBonds[position].key,
    value: fetchedBonds[position].value
  });
}