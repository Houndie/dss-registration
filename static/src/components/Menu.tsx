import React from 'react'
import Navbar from "react-bootstrap/Navbar"
import NavDropdown from "react-bootstrap/NavDropdown"
import Nav from "react-bootstrap/Nav"
import Container from "react-bootstrap/Container"

export default () => (
	<Container>
		<Navbar expand="lg">
			<Navbar.Brand href="/">Dayton Swing Smackdown</Navbar.Brand>
			<Navbar.Toggle aria-controls="basic-navbar-nav" />
			<Navbar.Collapse className="justify-content-end">
				<Nav>
					<NavDropdown title="Registration" id="menu-registration">
						<NavDropdown.Item href="https://dayton-swing-smackdown-2024.dancecamps.org/">Register Now</NavDropdown.Item>
						<NavDropdown.Item href="/pricing-and-tshirts">Pricing</NavDropdown.Item>
						<NavDropdown.Item href="/housing">Housing</NavDropdown.Item>
					</NavDropdown>
					<Nav.Link href="/classes">Classes</Nav.Link>
					<Nav.Link href="/instructors">Instructors</Nav.Link>
					<NavDropdown title="Competitions" id="menu-competitions">
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
					</NavDropdown>
					<NavDropdown title="Venues" id="menu-venues">
						<NavDropdown.Item href="/event-venues">Event Venues</NavDropdown.Item>
						<NavDropdown.Item href="/local-faire">Local Faire</NavDropdown.Item>
					</NavDropdown>
					<Nav.Link href="/music">Music</Nav.Link>
					<Nav.Link href="/schedule">Schedule</Nav.Link>
					<NavDropdown title="Safety" id="menu-safety">
						<NavDropdown.Item href="/covid">Covid</NavDropdown.Item>
						<NavDropdown.Item href="/safety">Safety</NavDropdown.Item>
					</NavDropdown>
				</Nav>
			</Navbar.Collapse>
		</Navbar>
	</Container>
)
