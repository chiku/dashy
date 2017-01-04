// gulpfile.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var gulp = require('gulp');
var cssnano = require('gulp-cssnano');
var htmlmin = require('gulp-htmlmin');
var jsbeautifier = require('gulp-jsbeautifier');

gulp.task('css-compile', function () {
    return gulp.src('./public/main.css')
        .pipe(cssnano())
        .pipe(gulp.dest('./out/public'));
});

gulp.task('css-format', function () {
    return gulp.src(['./public/main.css'])
        .pipe(jsbeautifier())
        .pipe(gulp.dest('./public'));
});

gulp.task('html-compile', function () {
    return gulp.src('./public/index.html')
        .pipe(htmlmin({
            collapseWhitespace: true
        }))
        .pipe(gulp.dest('./out/public'));
});

gulp.task('html-format', function () {
    return gulp.src(['./public/index.html'])
        .pipe(jsbeautifier())
        .pipe(gulp.dest('./public'));
});

gulp.task('favicon-compile', function () {
    return gulp.src('./public/favicon.ico')
        .pipe(gulp.dest('./out/public'));
});

gulp.task('format', ['css-format', 'html-format']);
gulp.task('compile', ['html-compile', 'css-compile', 'favicon-compile']);

gulp.task('build', ['format', 'compile']);

gulp.task('default', ['build']);
