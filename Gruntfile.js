module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    browserify: {
      dist: {
        files: {
            'server/public/dist/scripts/bundle.js': [
			  'client/scripts/app.js'
			]
	      },
	      options: {
	        transform: ['reactify']
	      }
      }
    },
    watch: {
      scripts:{
        files: ['client/scripts/**/*.js'],
        tasks: ['browserify'],
      },
	  styles: {
	    files: ['client/styles/**/*.css'],
		tasks: ['copy:main']
	  },
      livereload:{
        files: [
		  'server/public/dist/scripts/bundle.js',
		  'server/app.go'
		]
      }
    },
	copy: {
	  main: {
	    files: [
		  {
			src: 'client/bower_components/normalize-css/normalize.css',
			dest: 'server/public/dist/styles/normalize.css'
	      },
		  {
			src: 'client/styles/main.css',
			dest: 'server/public/dist/styles/main.css'
		  }
		]
	  }
	}
  });
  
  grunt.loadNpmTasks('grunt-browserify');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');

  grunt.registerTask('default', ['watch']);
  grunt.registerTask('build', ['copy:main', 'browserify']);
}
