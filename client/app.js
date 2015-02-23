"use strict";
var React = require('react');
var Router= require('react-router');
var { Route, DefaultRoute, NotFoundRoute, Link } = Router;
var Articles = require('./components/articles.js');
var NotFound = require('./components/notfound.js');
var App = require('./components/app.js');
var User = require('./components/user.js');

var routes = (
    <Route handler={App} path="/">
      <DefaultRoute handler={Articles} />
      <Route name="user" handler={User} />
      <NotFoundRoute handler={NotFound} />
    </Route>
    )


Router.run(routes, function(Handler){
  React.render(<Handler/>, document.getElementById("react-container"))
})


