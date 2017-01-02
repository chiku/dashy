// gulpfile.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var gulp = require('gulp');
var del = require('del');
var zip = require('gulp-zip');

require('./tasks/javascript');
require('./tasks/assets');

gulp.task('package', ['build'], function () {
    return gulp.src('./out/**')
        .pipe(zip('dashy.zip'))
        .pipe(gulp.dest('./'));
});

gulp.task('clean', function () {
    return del(['./out/**/*', './dashy.zip']);
});

gulp.task('js', ['js-compile', 'js-test', 'js-format']);
gulp.task('html', ['html-compile', 'html-format']);
gulp.task('css', ['css-compile', 'css-format']);
gulp.task('favicon', ['favicon-compile']);

gulp.task('format', ['css-format', 'html-format']);
gulp.task('compile', ['js-compile', 'html-compile', 'css-compile', 'favicon-compile']);
gulp.task('test', ['js-test']);

gulp.task('build', ['format', 'compile']);

gulp.task('default', ['build', 'test', 'package']);
