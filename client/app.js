'use strict'
var React = require('react');

var Hello = React.createClass({
	render: function(){
		return (<div>
			Hell {this.props.name} !!!!!!! 
			</div>
		       )
	}
})

React.render(<Hello name="from React" />, document.getElementById("react-container"))
