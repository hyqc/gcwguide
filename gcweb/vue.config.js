module.exports = {
    devServer: {
      proxy: {
        '/api/': {
          target: 'http://127.0.0.1:8083/',
          changeOrigin: true,
          ws: false,
          pathRewrite: {
            '^/api/': '/'
          }
        }
      }
    }
  }