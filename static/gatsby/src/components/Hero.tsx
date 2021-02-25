import React from 'react'
import Jumbotron from 'react-bootstrap/Jumbotron'

interface HeroProps {
	image: string
	height: string
	title: string
}

export default ({image, height, title}: HeroProps) => (
	<Jumbotron className="vertical-center horizontal-center" style={{backgroundImage: 'url('+image+')', backgroundSize: 'cover', backgroundPosition: 'center top', height: height}} fluid>
		<h1 className="HeroText">{title}</h1>
	</Jumbotron>
)
