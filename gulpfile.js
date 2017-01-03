// gulpfile.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var gulp = require('gulp');
var cssnano = require('gulp-cssnano');
var htmlmin = require('gulp-htmlmin');
var jsbeautifier = require('gulp-jsbeautifier');
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');

var del = require('del');
var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var source = require('vinyl-source-stream');

gulp.task('js-compile', function () {
    return browserify('./javascript/src/main.js')
        .bundle()
        .pipe(source('app.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({
            loadMaps: true
        }))
        .pipe(uglify())
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('./out/public'));
});

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
