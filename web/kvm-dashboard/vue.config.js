const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    open: true, // auto open browser
    host: '127.0.0.1',
    port: 8080,
    proxy: {
      '/web': {
        target: 'http://127.0.0.1:28080/',
        changeOrigin: true,
        pathRewrite: {
          '^/web': '/' // delete /api path
        }
      }
    }
  },
  transpileDependencies: true,
  lintOnSave: false // close eslint detect
})
