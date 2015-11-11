var gulp = require('gulp');
var minifyCss = require('gulp-minify-css');
var minifyHTML = require('gulp-minify-html');
var prettify = require('gulp-jsbeautifier');

gulp.task('css-build', function() {
    return gulp.src('./public/main.css')
        .pipe(minifyCss())
        .pipe(gulp.dest('./out/public'));
});

gulp.task('css-format', function() {
    return gulp.src(['./public/main.css'])
        .pipe(prettify())
        .pipe(gulp.dest('./public'));
});

gulp.task('html-build', function() {
    return gulp.src('./public/index.html')
        .pipe(minifyHTML())
        .pipe(gulp.dest('./out/public'));
});

gulp.task('html-format', function() {
    return gulp.src(['./public/index.html'])
        .pipe(prettify())
        .pipe(gulp.dest('./public'));
});

gulp.task('favicon-build', function() {
    return gulp.src('./public/favicon.ico')
        .pipe(gulp.dest('./out/public'));
});
