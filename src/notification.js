const notifier = require("node-notifier");

const loadNotify = (title, comment, ...args) => {
  //TODO args
  //console.log(args[0]);
  notifier.notify({
    title: title,
    message: comment,
  });
};

module.exports = loadNotify;
