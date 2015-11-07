var gulp = require('gulp');

require('./tasks/golang')
require('./tasks/javascript')
require('./tasks/assets')

gulp.task('go', ['go-build', 'go-test', 'go-format']);
gulp.task('js', ['js-build']);
gulp.task('assets', ['html-build', 'css-build']);

gulp.task('build', ['go-build', 'js-build', 'html-build', 'css-build']);
gulp.task('test', ['go-test']);
gulp.task('format', ['go-format']);

gulp.task('default', ['go', 'js', 'assets']);
