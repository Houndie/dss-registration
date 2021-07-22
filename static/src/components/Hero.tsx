import React from 'react'

interface HeroProps {
	image: string
	height: string
	title: string
}

export default ({image, height, title}: HeroProps) => (
	<div className="jumbotron vertical-center horizontal-center" style={{backgroundImage: 'url('+image+')', backgroundSize: 'cover', backgroundPosition: 'center top', height: height}}>
		<h1 className="HeroText">{title}</h1>
	</div>
)
