var gulp = require('gulp');
var cssnano = require('gulp-cssnano');
var htmlmin = require('gulp-htmlmin');
var jsbeautifier = require('gulp-jsbeautifier');

gulp.task('css-build', function() {
    return gulp.src('./public/main.css')
        .pipe(cssnano())
        .pipe(gulp.dest('./out/public'));
});

gulp.task('css-format', function() {
    return gulp.src(['./public/main.css'])
        .pipe(jsbeautifier())
        .pipe(gulp.dest('./public'));
});

gulp.task('html-build', function() {
    return gulp.src('./public/index.html')
        .pipe(htmlmin({
            collapseWhitespace: true
        }))
        .pipe(gulp.dest('./out/public'));
});

gulp.task('html-format', function() {
    return gulp.src(['./public/index.html'])
        .pipe(jsbeautifier())
        .pipe(gulp.dest('./public'));
});

gulp.task('favicon-build', function() {
    return gulp.src('./public/favicon.ico')
        .pipe(gulp.dest('./out/public'));
});
