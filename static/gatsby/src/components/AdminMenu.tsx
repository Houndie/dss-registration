import React from 'react'
import Navbar from "react-bootstrap/Navbar"
import NavDropdown from "react-bootstrap/NavDropdown"
import Nav from "react-bootstrap/Nav"
import Container from "react-bootstrap/Container"
import { GoogleLogin, GoogleLogout, GoogleLoginResponse } from "react-google-login"
import {AuthResult, AuthStatus} from './auth'

const isGoogleLoginResponse = (res: any): res is GoogleLoginResponse => {
	return (res as GoogleLoginResponse).accessToken !== undefined
}

interface AdminMenuProps {
	auth: AuthResult
}

export default ({auth}: AdminMenuProps) => (
	<Container>
		<Navbar expand="lg">
			<Navbar.Brand href="/">Dayton Swing Smackdown</Navbar.Brand>
			<Navbar.Toggle aria-controls="basic-navbar-nav"/>
			<Navbar.Collapse className="justify-content-end">
				<Nav>
					<NavDropdown title="Discounts" id="menu-discounts">
						<NavDropdown.Item href="/admin/discounts">List</NavDropdown.Item>
						<NavDropdown.Item href="/admin/discounts/add">Add</NavDropdown.Item>
					</NavDropdown>
					{ auth.clients.status== AuthStatus.SignedIn ? (
						<Nav.Link href="#" onClick={auth.signOut}>Log Out</Nav.Link>
					) : (
						<Nav.Link href="#" onClick={auth.signIn}>Log In</Nav.Link>
					)}
				</Nav>
			</Navbar.Collapse>
		</Navbar>
	</Container>
)
