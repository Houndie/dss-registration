import React from 'react'
import Navbar from "react-bootstrap/Navbar"
import NavDropdown from "react-bootstrap/NavDropdown"
import Nav from "react-bootstrap/Nav"
import Container from "react-bootstrap/Container"
import LoginLogout from './LoginLogout'

export default () => (
	<Container>
		<Navbar expand="lg">
			<Navbar.Brand href="/">Dayton Swing Smackdown</Navbar.Brand>
			<Navbar.Toggle aria-controls="basic-navbar-nav"/>
			<Navbar.Collapse className="justify-content-end">
				<Nav>
					<NavDropdown title="Registrations" id="menu-discounts">
						<NavDropdown.Item href="/admin/registrations">List</NavDropdown.Item>
					</NavDropdown>
					<LoginLogout/>
				</Nav>
			</Navbar.Collapse>
		</Navbar>
	</Container>
)
