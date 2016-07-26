'use strict';

const winston = require('winston');
winston.emitErrs = true;

const logger = new (winston.Logger)({
  transports: [
    new (winston.transports.Console)({
      timestamp: function() {
        return new Date().toISOString().replace(/T/, ' ').replace(/\..+/, '');
      },
      formatter: function(options) {
        return options.timestamp() +' '+ options.level.toUpperCase() +' '+ (undefined !== options.message ? options.message : '') +
          (options.meta && Object.keys(options.meta).length ? '\n\t'+ JSON.stringify(options.meta) : '' );
      }
    })
  ]
});

module.exports = logger;
module.exports.stream = {
  write(message) {
    logger.info(message);
  }
};
