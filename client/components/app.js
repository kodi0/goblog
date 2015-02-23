var React = require('react');
var Router = require('react-router');
var {RouteHandler, Link} = Router;

var App = React.createClass({
	render: function(){
		return (
      <div>
		   <header>
        <ul>
           <li><Link to="user">User</Link></li>
        </ul>
       </header>
       <RouteHandler/>
      </div>
		  )
	}
})

module.exports= App;
