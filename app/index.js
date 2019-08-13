'use strict';

var util       = require('util'),
    path       = require('path'),
    generators = require('yeoman-generator'),
    _          = require('lodash'),
    _s         = require('underscore.string'),
    pluralize  = require('pluralize'),
    asciify    = require('asciify');


module.exports = generators.Base.extend({
  constructor: function () {
    generators.Base.apply(this, arguments);

    this.option('flat', {
      type: Boolean,
      required: false,
      defaults: false,
      desc: 'When specified, generators will be created at the top level of the project.'
    });
  },

  initializing: function () {
    this.config.set('structure', this.options.flat ? 'flat' : 'nested');
    this.generatorsPrefix = this.options.flat ? '' : 'generators/';
    this.appGeneratorDir = this.options.flat ? 'app' : 'generators';
  },

  prompting: {
    askFor: function () {
      var done = this.async();
      this.log('\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '|g|o| |e|c|h|o| |g|e|n|e|r|a|t|o|r|\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '\n'
      );

      var prompts = [{
        type: 'input',
        name: 'baseName',
        message: 'What is the name of your application?',
        store   : true,
        default: 'myapp'
      },
      {
        type: 'input',
        name: 'httpLib',
        message: 'What HTTP library would you like to use, http, httprouter or echo? [http]',
        store   : true,
        default: 'http'
      }];

      this.prompt(prompts, function (props) {
        this.baseName = props.baseName;
        this.httpLib = props.httpLib;

        done();
      }.bind(this));
    }
  },

  writing: {
    app: function () {
      var httpLib = this.httpLib;
      this.copy('gitignore', '.gitignore');

      var modelsDir       = 'src/models/',
          controllersDir  = 'src/controllers/',
      this.mkdir(modelsDir);
      this.mkdir(controllersDir);

      this.fs.copyTpl(
        this.templatePath('_server_' + httpLib + '.go'),
        this.destinationPath('server.go'),
        {
          baseName: this.baseName
        }
      );

      this.fs.copyTpl(
        this.templatePath('src/controllers/_hello_' + httpLib + '.go'),
        this.destinationPath(controllersDir + 'hello.go'),
        {
          baseName: this.baseName
        }
      );

      var packages = "";
      if (httpLib == "echo") {
        packages = "github.com/labstack/echo"
      } else if (httpLib == "httprouter") {
        packages = "github.com/julienschmidt/httprouter"
      }

      this.fs.copyTpl(
        this.templatePath('_Goopfile'),
        this.destinationPath('Goopfile'),
        {
          packages: packages
        }
      );


      this.template(controllersDir + '_hello_' + httpLib + '_test.go', controllersDir + 'hello_test.go');
      this.template(modelsDir + '_hello.go', modelsDir + 'hello.go');
      this.template(modelsDir + '_hello_test.go', modelsDir + 'hello_test.go');
      this.template('README.md.erb', 'README.md');
      this.spawnCommand('go', ['mod', 'init'], {'cwd': this.destinationPath('src')});
    }
  }
});
