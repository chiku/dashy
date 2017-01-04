var webpack = require('webpack');

module.exports = {
    entry: './javascript/src/main.js',

    output: {
        path: __dirname,
        filename: 'out/public/app.js'
    },

    devtool: 'source-map',

    plugins: [
        new webpack.optimize.UglifyJsPlugin(),
        new webpack.optimize.OccurrenceOrderPlugin(),
        new webpack.optimize.DedupePlugin()
    ]
};
