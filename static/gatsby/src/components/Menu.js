import React from 'react'
import Navbar from "react-bootstrap/Navbar"
import NavDropdown from "react-bootstrap/NavDropdown"
import Nav from "react-bootstrap/Nav"
import Container from "react-bootstrap/Container"
import { GoogleLogin, GoogleLogout } from "react-google-login"

const Menu = ({gAuth, setGAuth}) => (
	<Container>
		<Navbar expand="lg">
			<Navbar.Brand href="/">Dayton Swing Smackdown</Navbar.Brand>
			<Navbar.Toggle aria-controls="basic-navbar-nav"/>
			<Navbar.Collapse className="justify-content-end">
				<Nav>
					<NavDropdown title="Registration">
						<NavDropdown.Item href="/registration">Register Now</NavDropdown.Item>
						<NavDropdown.Item href="/pricing-and-tshirts">Pricing & T-Shirts</NavDropdown.Item>
						<NavDropdown.Item href="/housing">Housing</NavDropdown.Item>
						<NavDropdown.Item href="/volunteer">Volunteer</NavDropdown.Item>
					</NavDropdown>
					<Nav.Link href="/classes">Classes</Nav.Link>
					<Nav.Link href="/instructors">Instructors</Nav.Link>
					<NavDropdown title="Competitions">
						<NavDropdown.Item href="/competition-info">Competition Info</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2011">2011 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2012">2012 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2013">2013 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2014">2014 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2015">2015 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2016">2016 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2017">2017 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2018">2018 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2019">2019 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2020">2020 Competition Results</NavDropdown.Item>
						<NavDropdown.Item href="/competition-results/2021">2021 Competition Results</NavDropdown.Item>
					</NavDropdown>
					<NavDropdown title="Venues">
						<NavDropdown.Item href="/event-venues">Event Venues</NavDropdown.Item>
						<NavDropdown.Item href="/local-faire">Local Faire</NavDropdown.Item>
					</NavDropdown>
					<Nav.Link href="/music">Music</Nav.Link>
					<Nav.Link href="/schedule">Schedule</Nav.Link>
					<Nav.Link href="/safety">Safety</Nav.Link>
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

export default Menu
