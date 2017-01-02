// tasks/golang.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var exec = require('child_process').exec;

var gulp = require('gulp');

var execCommand = function (command, cb) {
    exec(command, function (err, stdout, stderr) {
        console.log(stdout);
        console.error(stderr);
        cb(err);
    });
};

gulp.task('go-prereqs', function (cb) {
    execCommand('glide install', cb);
});

gulp.task('go-compile', ['go-prereqs'], function (cb) {
    execCommand('go build -o out/dashy', cb);
});

gulp.task('go-test', ['go-prereqs'], function (cb) {
    execCommand('go test ./app', cb);
});

gulp.task('go-lint', function (cb) {
    execCommand('go vet . ./app', cb);
});

gulp.task('go-format', function (cb) {
    execCommand('go fmt . ./app', cb);
});
