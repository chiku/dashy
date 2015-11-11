var gulp = require('gulp');
var uglify = require('gulp-uglify');
var jshint = require('gulp-jshint');

var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var source = require('vinyl-source-stream');

gulp.task('js-build', function() {
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
