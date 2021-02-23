import React, {useState} from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Menu from './Menu.js'
import AdminMenu from './AdminMenu.js'
import Footer from './Footer.js'
import Hero from './Hero.js'
import AuthGuard from './AuthGuard.js'

const GenericPage = ({title, children, menu, requireAuth}) => {
	const [gAuth, setGAuth] = useState(null)
	return (
		<>
			{menu(gAuth, setGAuth)}
			<Hero image='/page_header.png' height='250px' title={title} />
			<Container>
				{ requireAuth ? (
					<AuthGuard gAuth={gAuth}>
						{() => children(gAuth)}
					</AuthGuard>
				) : children(gAuth)}
			</Container>
			<Footer />
		</>
	)
}

export default ({title, children}) => (
	<GenericPage 
		title={title} 
		children={children} 
		menu={(gAuth, setGAuth) => <Menu gAuth={gAuth} setGAuth={setGAuth} />} 
	/>
)

const AdminPage = ({title, children}) => (
	<GenericPage 
		title={title} 
		children={children} 
		menu={(gAuth, setGAuth) => <AdminMenu gAuth={gAuth} setGAuth={setGAuth} />} 
		requireAuth={true}
	/>
)

export {AdminPage}
