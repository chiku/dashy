var del = require('del');

var gulp = require('gulp');
var zip = require('gulp-zip');

gulp.task('package', ['build'], function() {
    return gulp.src('./out/**')
		.pipe(zip('dashy.zip'))
        .pipe(gulp.dest('./'));
});

gulp.task('clean', function () {
  return del(['./out/**/*', './dashy.zip' ]);
});
