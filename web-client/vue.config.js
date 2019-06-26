const Config = require('webpack-chain')

// Instantiate the configuration with a new API
const config = new Config()

module.exports = {
  chainWebpack: config => {
    config.module
      .rule('scss')
      .test(/sebgroup.*\.scss$/)
      .use('style-loader', 'css-loader', 'sass-loader')
      .end()
  }
}

module.exports = config.toConfig();