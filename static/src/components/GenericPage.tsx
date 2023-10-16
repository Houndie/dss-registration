import React from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Footer from './Footer'
import Hero from './Hero'

interface GenericPageProps {
	title: string;
	menu: () => React.ReactNode
}

const GenericMenu : React.Fc<GenericPageProps> = ({title, menu, children}) => (
	<>
		{menu()}
		<Hero image='/page_header.png' height='250px' title={title} />
		<Container>
			{ children }
		</Container>
		<Footer />
	</>
)

export default GenericMenu
