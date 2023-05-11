const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    open: true, // auto open browser
    host: '127.0.0.1',
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8089/api2',
        pathRewrite: { '^/api': '' },
        changeOrigin: true,
        pathRewrite: {
          '^/api': '' // delete /api path
        }
      }
    }
  },
  transpileDependencies: true,
  lintOnSave: false // close eslint detect
})
