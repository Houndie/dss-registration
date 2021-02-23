import React from 'react'
import Navbar from "react-bootstrap/Navbar"
import NavDropdown from "react-bootstrap/NavDropdown"
import Nav from "react-bootstrap/Nav"
import Container from "react-bootstrap/Container"
import { GoogleLogin, GoogleLogout } from "react-google-login"

export default ({gAuth, setGAuth}) => (
	<Container>
		<Navbar expand="lg">
			<Navbar.Brand href="/">Dayton Swing Smackdown</Navbar.Brand>
			<Navbar.Toggle aria-controls="basic-navbar-nav"/>
			<Navbar.Collapse className="justify-content-end">
				<Nav>
					<NavDropdown title="Discounts">
						<NavDropdown.Item href="/admin/discounts">List</NavDropdown.Item>
						<NavDropdown.Item href="/admin/discounts/add">Add</NavDropdown.Item>
					</NavDropdown>
					{ gAuth ? (
						<GoogleLogout
							clientId={`${process.env.GATSBY_CLIENT_ID}`}
							buttonText="Logout"
							onLogoutSuccess={() => {setGAuth(null)}}
						/>
					) : (
						<GoogleLogin
							clientId={`${process.env.GATSBY_CLIENT_ID}`}
							buttonText="Login"
							onSuccess={(newAuth) => setGAuth(newAuth)}
							onFailure={() => console.log(`NO ${process.env.GATSBY_CLIENT_ID}`)}
							cookiePolicy={'single_host_origin'}
							isSignedIn={true}
						/>
					)}
				</Nav>
			</Navbar.Collapse>
		</Navbar>
	</Container>
)
