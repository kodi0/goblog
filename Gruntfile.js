module.exports = function(grunt){
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    browserify: {
      dist: {
        files: {  
          './server/public/dist/js/bundle.js': ['client/app.js']
	      },
	      options: {
	        transform: ['reactify','es6ify']
	      }
      }
    },
    watch: {
      client:{
        files: ['client/**/*.js'],
        tasks: ['browserify'],
      },			  
      livereload:{
        files: ['./public/dist/js/bundle.js','server/app.go'],
        options: {
          livereload: true,
        }
      }
    }  
  });
  grunt.loadNpmTasks('grunt-browserify')
  grunt.loadNpmTasks('grunt-contrib-watch')

  grunt.registerTask('default', ['watch'])
  grunt.registerTask('build',['browserify'])
}
