'use strict'

const express = require('express');
const app = express();
const logger = require('./libs/logger');

app.get('/', function (req, res) {
	logger.info('hello in the handler..')
	res.send("hello.. i'm pusheen the cat");
});

app.listen(15120, function () {
	  logger.info('Example app listening on port 15120!');
});
