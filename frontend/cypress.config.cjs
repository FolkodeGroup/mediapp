const { defineConfig } = require('cypress');

module.exports = defineConfig({
  e2e: {
    baseUrl: 'http://localhost:3000', // Correcto para desarrollo local
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});