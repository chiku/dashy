var gulp = require('gulp');
var uglify = require('gulp-uglify');
var jshint = require('gulp-jshint');
var prettify = require('gulp-jsbeautifier');

var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var source = require('vinyl-source-stream');

gulp.task('js-compile', function() {
    return browserify('./public/main.js')
        .bundle()
        .pipe(source('app.js'))
        .pipe(buffer())
        .pipe(uglify())
        .pipe(gulp.dest('./out/public'));
});

gulp.task('js-lint', function() {
    return gulp.src(['./public/main.js', './gulpfile.js', './tasks/*.js'])
        .pipe(jshint())
        .pipe(jshint.reporter('default'));
});

gulp.task('js-format-main', function() {
    return gulp.src(['./public/main.js'])
        .pipe(prettify())
        .pipe(gulp.dest('./public'));
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
