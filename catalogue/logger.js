// logger.js
const pino = require('pino');
const fs = require('fs');

// Create a write stream for logging to a file
const logStream = fs.createWriteStream('./app.log', { flags: 'a' }); // Append mode

// Create a Pino logger instance
const logger = pino({
    level: 'info',
    transport: {
        target: 'pino-pretty', // Optional: for pretty output in development
        options: {
            colorize: true,
        },
    },
    // Stream logs to file as well
    streams: [
        { stream: logStream }, // Log to file
        { stream: process.stdout }, // Log to console
    ],
});

module.exports = logger;