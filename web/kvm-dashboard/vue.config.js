const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    open: true, // auto open browser
    host: '127.0.0.1',
    port: 8080,
    proxy: {
      '/golang_web': {
        target: 'http://127.0.0.1:5001/',
        changeOrigin: true,
        pathRewrite: {
          '^/golang_web': '/' // delete /api path
        }
      },
      // '/golang_ws': {
      //   target: 'http://127.0.0.1:5001/',
      //   ws: true, // proxy websockets
      //   changeOrigin: true,
        // pathRewrite: {
        //   '^/golang_ws': '/' // delete /api path
        // }
      }
    }
  },
  transpileDependencies: true,
  lintOnSave: false // close eslint detect
})
