// gulpfile.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var gulp = require('gulp');
var del = require('del');

require('./tasks/javascript');
require('./tasks/assets');

gulp.task('clean', function () {
    return del(['./out/public/**/*', './dashy.zip']);
});

gulp.task('js', ['js-compile', 'js-test', 'js-format']);
gulp.task('html', ['html-compile', 'html-format']);
gulp.task('css', ['css-compile', 'css-format']);
gulp.task('favicon', ['favicon-compile']);

gulp.task('format', ['css-format', 'html-format']);
gulp.task('compile', ['js-compile', 'html-compile', 'css-compile', 'favicon-compile']);
gulp.task('test', ['js-test']);

gulp.task('build', ['format', 'compile']);

gulp.task('default', ['build', 'test']);
