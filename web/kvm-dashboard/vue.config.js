const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    open: true, // auto open browser
    host: '127.0.0.1',
    port: 8080,
  },
  transpileDependencies: true,
  lintOnSave: false // close eslint detect
})
