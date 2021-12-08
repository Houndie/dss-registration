import React from "react"
import Spinner from "react-bootstrap/Spinner"

export default () => (
	<>
		<Spinner animation="border" role="status">
			<span className="visually-hidden">Loading Page...</span>
		</Spinner>
		<span>  Loading Page...</span>
	</>
)
