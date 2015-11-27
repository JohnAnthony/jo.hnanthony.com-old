var gulp = require('gulp');
var sass = require('gulp-sass');
var concat = require('gulp-concat');
var jade = require('gulp-jade');

gulp.task('build', ['sass', 'jade']);

gulp.task('sass', function() {
    gulp.src('./src/sass/*.scss')
        .pipe(concat('bundle.css'))
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('./build'));
});

gulp.task('jade', function() {
    gulp.src('./src/index.jade')
	.pipe(jade())
        .pipe(gulp.dest('./build'));
});

gulp.task('watch', ['build'], function() {
    gulp.watch('./src/index.html', ['html'])
    gulp.watch('./src/sass/*.scss', ['sass'])
});
