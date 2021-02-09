import React from 'react'
import Jumbotron from 'react-bootstrap/Jumbotron'

const Hero = ({image, height, title}) => (
	<Jumbotron className="vertical-center horizontal-center" style={{backgroundImage: 'url('+image+')', backgroundSize: 'cover', backgroundPosition: 'center top', height: height}} fluid>
		<h1 className="HeroText">{title}</h1>
	</Jumbotron>
)

export default Hero

