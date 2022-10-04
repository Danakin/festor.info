const path = require('path')
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
    entry: './ui/assets/js/app.js',
    watch: true,
    output: {
      filename: 'app.js',
      path: path.resolve(__dirname, 'ui', 'static', 'js'),
    },
    plugins: [new MiniCssExtractPlugin({
        filename: "../css/app.css"
    })],
    module: {
      rules: [
        {
          test: /\.s[ac]ss$/i,
          use: [
            MiniCssExtractPlugin.loader,
            // Translates CSS into CommonJS
            "css-loader",
            // Compiles Sass to CSS
            "sass-loader",
          ],
        },
      ],
    },
  };