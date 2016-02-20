// tasks/javascript.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

var gulp = require('gulp');
var uglify = require('gulp-uglify');
var jshint = require('gulp-jshint');
var prettify = require('gulp-jsbeautifier');
var sourcemaps = require('gulp-sourcemaps');
var jasmine = require('gulp-jasmine');

var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var source = require('vinyl-source-stream');

gulp.task('js-compile', function() {
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

gulp.task('js-test', function() {
    return gulp.src('javascript/test/**/*.test.js')
        .pipe(jasmine());
});

gulp.task('js-lint', function() {
    return gulp.src(['./javascript/**/*.js', './gulpfile.js', './tasks/*.js'])
        .pipe(jshint())
        .pipe(jshint.reporter('default'));
});

gulp.task('js-format-main', function() {
    return gulp.src(['./javascript/**/*.js'])
        .pipe(prettify())
        .pipe(gulp.dest('./javascript'));
});

gulp.task('js-format-gulpfile', function() {
    return gulp.src(['./gulpfile.js'])
        .pipe(prettify())
        .pipe(gulp.dest('./'));
});

gulp.task('js-format-tasks', function() {
    return gulp.src(['./tasks/*.js'])
        .pipe(prettify())
        .pipe(gulp.dest('./tasks'));
});

gulp.task('js-format', ['js-format-main', 'js-format-gulpfile', 'js-format-tasks']);
