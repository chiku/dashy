var webpack = require('webpack');

module.exports = {
    entry: './javascript/src/main.js',

    output: {
        path: __dirname,
        filename: 'out/public/app.js',
        publicPath: '/public/'
    },

    stats: {
        colors: true,
        reasons: true
    },

    devtool: 'source-map',

    plugins: [
        new webpack.optimize.UglifyJsPlugin(),
        new webpack.optimize.OccurrenceOrderPlugin(),
        new webpack.optimize.DedupePlugin()
    ]
};
