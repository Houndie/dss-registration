import React, {useState} from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Menu from './Menu.js'
import Footer from './Footer.js'
import Hero from './Hero.js'

const Page = ({title, children}) => {
	const [gAuth, setGAuth] = useState(null)
	return (
		<>
			<Menu gAuth={gAuth} setGAuth={setGAuth} />
			<Hero image='/page_header.png' height='250px' title={title} />
			<Container>
				{children}
			</Container>
			<Footer />
		</>
	)
}

export default Page
