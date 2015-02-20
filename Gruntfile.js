module.exports = function(grunt){
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    browserify: {
      dist: {
        files: {  
          './client/components.js': ['client/app.js']
	      },
	      options: {
	        transform: ['reactify']
	      }
      }
    },
    concat: {
      options: {
        separator: ';',
      },
      dist:{
        src: ['./client/bower_components/dist/jquery.js',
              './client/bower_components/materialize/dist/js/materialize.js',
              './client/components.js'],
        dest: './client/bundle.js'
      }
    },
    uglify: {
      my_target: {
        options: {
          mangle: false,
        },
        files: {
          './server/public/dist/js/bundle.min.js': ['./client/bundle.js']
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
  })
  grunt.loadNpmTasks('grunt-browserify')
  grunt.loadNpmTasks('grunt-contrib-concat')
  grunt.loadNpmTasks('grunt-contrib-uglify')
  grunt.loadNpmTasks('grunt-contrib-watch')

  grunt.registerTask('default', ['watch'])
  grunt.registerTask('build',['browserify','concat','uglify'])
}
