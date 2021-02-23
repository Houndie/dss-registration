import React from 'react'

export default({gAuth, children}) => {
	if(!gAuth) {
		return <p>You must be logged in to view this page!</p>
	}
	return children()
}
