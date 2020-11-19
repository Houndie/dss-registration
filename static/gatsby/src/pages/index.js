import React from "react"
import '../styles/style.scss'
import Jumbotron from "react-bootstrap/Jumbotron"
import Container from "react-bootstrap/Container"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"
import Image from "react-bootstrap/Image"
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGraduationCap, faMedal, faMusic, faCircle } from '@fortawesome/free-solid-svg-icons'
import { faFacebook } from '@fortawesome/free-brands-svg-icons'
import {Formik} from 'formik'
import FormText from '../components/FormText.js'
import FormTextArea from '../components/FormTextArea.js'
import Menu from '../components/Menu.js'
import * as Yup from 'yup'

const IconPanel = ({icon, title, children}) => (
	<div className="text-center">
		<FontAwesomeIcon icon={icon} mask={faCircle} size="6x" transform="shrink-7"/>
		<h2>{title}</h2>
		<p>{children}</p>
	</div>
)

const ContactSchema = Yup.object().shape({
	name: Yup.string()
		.min(2, 'Too Short!')
		.required('Required'),
	email: Yup.string()
		.email('Invalid email')
		.required('Required'),
	message: Yup.string()
		.min(2, 'Too Short!')
		.required('Required'),
});

const Home = () => (
	<>
		<Menu/>
		<Jumbotron className="vertical-center horizontal-center" style={{backgroundImage: 'url(ViktorJump.jpg)', backgroundSize: 'cover', height: '450px'}} fluid>
			<h1 className="HeroText">Dayton Swing Smackdown</h1>
		</Jumbotron>
		<Container className="my-5">
			<Row><Col xs="5">
				<Image src="tri_city.jpg" fluid/>
			</Col><Col>
				<h2>We’re doing an intermission year!</h2>
				<p><strong>Spend some time online with us on February 27!</strong></p>
				<p>Due to the covid-19 pandemic, we obviously don’t feel comfortable having an in-person event, but that doesn’t mean we can’t still have a grand time digitally! We’re putting on a one day event with 3 classes, a discussion talk, a DJ party, and THE BATTLE OF THE SWING CITIES TEAM VIDEO COMPETITION! Mark down your calendars now, you don’t want to miss this!</p>
				<p>Because of the reduced nature of this event, we’re offering the intermission completely free so you have no reason not to be there!</p>
			</Col></Row>
		</Container>
		<Container className="my-5">
			<Row><Col>
				<IconPanel icon={faGraduationCap} title="Classes">Three free classes taught digitally!</IconPanel>
			</Col><Col>
				<IconPanel icon={faMedal} title="Competition">Plan with your friends and participate in Smackdown’s first video team competition!!</IconPanel>
			</Col><Col>
				<IconPanel icon={faMusic} title="DJ">Listen and share music with everyone with our jwbx.fm DJ party!</IconPanel>
			</Col></Row>
		</Container>
		<Container className="my-5">
			<Row><Col className="text-center">
				<h2>Have Questions?</h2>
				<i>Contact Us!</i>
				<Formik
					initialValues={{name: '', email: '', message: ''}}
					validationSchema={ContactSchema}
					onSubmit={(values, {setSubmitting}) => {
						var formData = new FormData();

						for (var k in values) {
							formData.append(k, values[k]);
						}
						
						fetch("https://formspree.io/xyyyrenq", {
							method: "POST",
							body: formData,
						}).then(() => {
							setSubmitting(false)
						}).catch((e) => {
							console.log(JSON.stringify(e))
							setSubmitting(false)
						})
					}}
				>
				{({values, handleSubmit, isSubmitting}) => (
					<Form onSubmit={handleSubmit}>
						<Container>
							<Row><Col>
								<FormText label="Name" name="name" type="text"/>
							</Col><Col>
								<FormText label="Email" name="email" type="text"/>
							</Col></Row><Row><Col>
								<FormTextArea label="Message" name="message"/>
							</Col></Row><Row><Col>
								<Button type="submit" disabled={isSubmitting}>Submit</Button>
							</Col></Row>
						</Container>
					</Form>
				)}
				</Formik>
			</Col></Row>
		</Container>
		<Container className="my-5">
			<Row><Col>
				<a href="https://www.facebook.com/Dayton-Swing-Smackdown-120632558063863">
					<FontAwesomeIcon icon={faFacebook} size="6x"/>
				</a>
				<p>Copyright 2021 by Dayton Swing Smackdown</p>
			</Col><Col>
				<h3>Dayton Swing Smackdown</h3>
				<p>We are an event dedicated to creating an exciting and fun atmosphere for new and old dancers alike. See you on the dance floor!</p>
			</Col></Row>
		</Container>
	</>
)

export default Home
