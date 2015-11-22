var gulp = require('gulp');

require('./tasks/golang');
require('./tasks/javascript');
require('./tasks/assets');
require('./tasks/package');

gulp.task('go', ['go-build', 'go-test', 'go-format', 'go-lint']);
gulp.task('js', ['js-build', 'js-lint']);
gulp.task('html', ['html-build', 'html-format']);
gulp.task('css', ['css-build', 'css-format']);
gulp.task('favicon', ['favicon-build']);

gulp.task('format', ['go-format', 'js-format', 'css-format', 'html-format']);
gulp.task('lint', ['go-lint', 'js-lint']);
gulp.task('build', ['go-build', 'js-build', 'html-build', 'css-build', 'favicon-build']);
gulp.task('test', ['go-test']);

gulp.task('default', ['go', 'js', 'css', 'html', 'favicon', 'package']);
