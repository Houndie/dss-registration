const React = require('react')

exports.onPreRenderHTML = ({ getHeadComponents, replaceHeadComponents}) => {
	const headComponents = getHeadComponents()
	headComponents.push(<script key="googlerecaptcha" src="https://www.google.com/recaptcha/api.js" async defer/>)
	replaceHeadComponents(headComponents)

}
