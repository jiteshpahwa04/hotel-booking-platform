require("ts-node/register"); // This line enables typescript support
const config = require("./db.config");

// Access the 'default' property from the TypeScript module
module.exports = config.default;