var gulp = require('gulp');
var minifyCss = require('gulp-minify-css');
var minifyHTML = require('gulp-minify-html');

gulp.task('css-build', function() {
  return gulp.src('public/main.css')
    .pipe(minifyCss())
    .pipe(gulp.dest('./out/public'));
});

gulp.task('html-build', function() {
  return gulp.src('public/index.html')
    .pipe(minifyHTML())
    .pipe(gulp.dest('./out/public'));
});
