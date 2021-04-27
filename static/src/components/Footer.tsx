import React from 'react'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faFacebook } from '@fortawesome/free-brands-svg-icons'

const Footer = () => (
	<Container className="my-5">
		<Row><Col>
			<a href="https://www.facebook.com/Dayton-Swing-Smackdown-120632558063863" className="no-special-color">
				<FontAwesomeIcon icon={faFacebook} size="6x"/>
			</a>
			<p>Copyright 2021 by Dayton Swing Smackdown</p>
		</Col><Col>
			<h3>Dayton Swing Smackdown</h3>
			<p>We are an event dedicated to creating an exciting and fun atmosphere for new and old dancers alike. See you on the dance floor!</p>
		</Col></Row>
	</Container>
)

export default Footer
