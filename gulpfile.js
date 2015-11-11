var gulp = require('gulp');

require('./tasks/golang');
require('./tasks/javascript');
require('./tasks/assets');

gulp.task('go', ['go-build', 'go-test', 'go-format', 'go-lint']);
gulp.task('js', ['js-build', 'js-lint']);
gulp.task('assets', ['html-build', 'css-build', 'favicon-build']);

gulp.task('build', ['go-build', 'js-build', 'html-build', 'css-build', 'favicon-build']);
gulp.task('test', ['go-test']);
gulp.task('format', ['go-format']);
gulp.task('lint', ['go-lint', 'js-lint']);

gulp.task('default', ['go', 'js', 'assets']);
