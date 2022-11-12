import React, {useEffect, useState} from "react"
import '../styles/style.scss'
import Container from "react-bootstrap/Container"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"
import Image from "react-bootstrap/Image"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { IconProp } from '@fortawesome/fontawesome-svg-core'
import { faGraduationCap, faMedal, faMusic, faCircle } from '@fortawesome/free-solid-svg-icons'
import Menu from '../components/Menu'
import Footer from '../components/Footer'
import Hero from '../components/Hero'

interface IconPanelProps {
	link: string
	icon: IconProp
	title: string
	children: React.ReactNode
}
const IconPanel = ({link, icon, title, children}:IconPanelProps) => (
	<div className="text-center">
		<a href={link} className="no-special-color">
			<FontAwesomeIcon icon={icon} mask={faCircle} size="6x" transform="shrink-7"/>
		</a>
		<h2>{title}</h2>
		<p>{children}</p>
	</div>
)

const Home = () => (
	<>
		<Menu/>
		<Hero image='ViktorJump.jpg' height='450px' title='Dayton Swing Smackdown' />
		<Container className="my-5">
			<h2>February 24 - February 26, 2023</h2>
		</Container>
		<Container className="my-5">
			<Row><Col xs="5">
				<Image src="tri_city.jpg" fluid/>
			</Col><Col className="align-self-center">
				<p>Dayton Swing Smackdown is a swing dancing event held every year on the last full weekend of February. It features over 9 hours of dancing, 13 hours of instruction, Solo Jazz Competition, Mix n Match Competition, and The Battle of the Swing Cities Team Routine Competition.  Smackdown is now on it’s 15th  year, and getting better with age.  In addition to a dedication to providing a quality weekend, it is one of Smackdown’s core goals to be accessible to everyone, from the experienced dance community, to brand new dancers.  Come and join us in February!</p>
			</Col></Row>
		</Container>
		<Container className="my-5">
			<Row><Col>
				<IconPanel link="/classes" icon={faGraduationCap} title="Classes">Smackdown offers three levels of classes to provide the best instruction for you!</IconPanel>
			</Col><Col>
				<IconPanel link="/competitions" icon={faMedal} title="Competition">Get your team together and compete in the Battle of the Swing Cities Team Competition!</IconPanel>
			</Col><Col>
				<IconPanel link="/music" icon={faMusic} title="Music">Dance to our live band on Friday, or to our collection of amazing DJs!</IconPanel>
			</Col></Row>
		</Container>
		<Container className="my-5">
			<Row><Col className="text-center">
				<h2>Have Questions?</h2>
				<i>Contact Us at info@daytonswingsmackdown.com</i>
			</Col></Row>
		</Container>
		<Footer />
	</>
)


export default Home
