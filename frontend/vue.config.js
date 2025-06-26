const CopyWebpackPlugin = require('copy-webpack-plugin')
const path = require('path')

module.exports = {
  configureWebpack: {
    plugins: [
      new CopyWebpackPlugin({
        patterns: [
          // 复制静态资源到微信小程序目录
          {
            from: path.resolve(__dirname, 'src/static'),
            to: path.resolve(__dirname, 'dist/dev/mp-weixin/static'),
            globOptions: {
              ignore: ['.*']
            },
            noErrorOnMissing: true
          },
          // 复制项目配置文件
          {
            from: path.resolve(__dirname, 'project.config.json'),
            to: path.resolve(__dirname, 'dist/dev/mp-weixin/project.config.json'),
            noErrorOnMissing: true
          },
          {
            from: path.resolve(__dirname, 'project.private.config.json'),
            to: path.resolve(__dirname, 'dist/dev/mp-weixin/project.private.config.json'),
            noErrorOnMissing: true
          },
          // 复制sitemap.json
          {
            from: path.resolve(__dirname, 'sitemap.json'),
            to: path.resolve(__dirname, 'dist/dev/mp-weixin/sitemap.json'),
            noErrorOnMissing: true
          }
        ]
      })
    ]
  },
  
  // 开发服务器配置
  devServer: {
    port: 3000,
    host: '0.0.0.0',
    hot: true,
    open: false
  },
  
  // 针对不同平台的配置
  chainWebpack: config => {
    // 微信小程序特殊配置
    if (process.env.UNI_PLATFORM === 'mp-weixin') {
      // 优化构建
      config.optimization.minimize(process.env.NODE_ENV === 'production')
      
      // 处理静态资源
      config.module
        .rule('images')
        .test(/\.(png|jpe?g|gif|svg)(\?.*)?$/)
        .use('url-loader')
        .loader('url-loader')
        .options({
          limit: 4096,
          fallback: {
            loader: 'file-loader',
            options: {
              name: 'static/images/[name].[hash:8].[ext]'
            }
          }
        })
    }
  }
} 