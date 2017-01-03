// tasks/javascript.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var gulp = require('gulp');
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');

var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var source = require('vinyl-source-stream');

gulp.task('js-compile', function () {
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
