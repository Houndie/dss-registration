import React from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Menu from './Menu.js'
import Footer from './Footer.js'
import Hero from './Hero.js'

const Page = ({title, children}) => (
	<>
		<Menu />
		<Hero image='/page_header.png' height='250px' title={title} />
		<Container>
			{children}
		</Container>
		<Footer />
	</>
)

export default Page
