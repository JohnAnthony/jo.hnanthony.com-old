var express = require('express');

var config;
try {
	config = require('./config.json');
} catch(e) {
	config = {
		port: 80
	};
}

var app = express();
app.use("/", express.static(__dirname + "/build"));
app.listen(config.port || 80);
