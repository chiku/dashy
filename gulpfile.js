// gulpfile.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

var gulp = require('gulp');
var del = require('del');
var zip = require('gulp-zip');

require('./tasks/golang');
require('./tasks/javascript');
require('./tasks/assets');

gulp.task('package', ['build'], function() {
    return gulp.src('./out/**')
        .pipe(zip('dashy.zip'))
        .pipe(gulp.dest('./'));
});

gulp.task('clean', function() {
    return del(['./out/**/*', './dashy.zip']);
});

gulp.task('go', ['go-compile', 'go-test', 'go-format', 'go-lint']);
gulp.task('js', ['js-compile', 'js-lint', 'js-format']);
gulp.task('html', ['html-compile', 'html-format']);
gulp.task('css', ['css-compile', 'css-format']);
gulp.task('favicon', ['favicon-compile']);

gulp.task('prereqs', ['go-prereqs']);
gulp.task('format', ['go-format', 'js-format', 'css-format', 'html-format']);
gulp.task('lint', ['go-lint', 'js-lint']);
gulp.task('compile', ['go-compile', 'js-compile', 'html-compile', 'css-compile', 'favicon-compile']);
gulp.task('test', ['go-test']);

gulp.task('build', ['format', 'lint', 'compile']);

gulp.task('default', ['build', 'test', 'package']);
