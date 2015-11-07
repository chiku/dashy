var exec = require('child_process').exec;

var gulp = require('gulp');

var execCommand = function(command, cb) {
  exec(command, function (err, stdout, stderr) {
    console.log(stdout);
    console.error(stderr);
    cb(err);
  });
};

gulp.task('go-restore', function(cb) {
  execCommand('godep restore', cb);
});

gulp.task('go-build', ['go-restore'], function(cb) {
  execCommand('go build -o out/dashy.exe', cb);
});

gulp.task('go-test', ['go-restore'], function(cb) {
  execCommand('go test ./app', cb);
});

gulp.task('go-format', function(cb) {
  execCommand('go fmt . ./app', cb);
});
