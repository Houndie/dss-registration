import React, {useEffect, useState} from "react"
import '../styles/style.scss'
import Container from "react-bootstrap/Container"
import Alert from "react-bootstrap/Alert"
import Spinner from "react-bootstrap/Spinner"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"
import Image from "react-bootstrap/Image"
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGraduationCap, faMedal, faMusic, faCircle } from '@fortawesome/free-solid-svg-icons'
import {Formik} from 'formik'
import FormField from '../components/FormField.js'
import Menu from '../components/Menu.js'
import Footer from '../components/Footer.js'
import Hero from '../components/Hero.js'
import * as Yup from 'yup'
import Recaptcha from 'react-recaptcha';
import formsTwirp from "../rpc/forms_pb_twirp.js"
import {ImportRecaptchaEffect} from '../components/recaptcha.js'

const IconPanel = ({link, icon, title, children}) => (
	<div className="text-center">
		<a href={link} className="no-special-color">
			<FontAwesomeIcon icon={icon} mask={faCircle} size="6x" transform="shrink-7"/>
		</a>
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

const Home = () => {
	return (
		<>
			<Menu/>
			<Hero image='ViktorJump.jpg' height='450px' title='Dayton Swing Smackdown' />
			<Container className="my-5">
				<Row><Col xs="5">
					<Image src="tri_city.jpg" fluid/>
				</Col><Col className="align-self-center">
					<p>Dayton Swing Smackdown is a swing dancing event held every year on the last full weekend of February. It features over 9 hours of dancing, 13 hours of instruction, Solo Jazz Competition, Mix n Match Competition, and The Battle of the Swing Cities Team Routine Competition.  Smackdown is now on it’s 12th  year, and getting better with age.  In addition to a dedication to providing a quality weekend, it is one of Smackdown’s core goals to be accessible to everyone, from the experienced dance community, to brand new dancers.  Come and join us in February!</p>
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
					<i>Contact Us!</i>
					<ContactUs />
				</Col></Row>
			</Container>
			<Footer />
		</>
	)
}

const ContactUs = () => {
	const formsClient = formsTwirp.createFormsClient(`${process.env.GATSBY_BACKEND}`)
	const [recaptchaLoaded, setRecaptchaLoaded] = useState(false)
	const [formSubmitted, setFormSubmitted] = useState(false)
	let recaptchaInstance
	useEffect(ImportRecaptchaEffect)
	if(formSubmitted){
		return (
			<Alert variant="success">
				Thank you for submitting your message!  We will get back to you as soon as we can.
			</Alert>
		)
	}
	return (
		<Formik
			initialValues={{name: '', email: '', message: '', recaptchaResponse: ''}}
			validationSchema={ContactSchema}
			onSubmit={(values, {setSubmitting}) => {
				if (values.recaptchaResponse === '') {
					recaptchaInstance.execute()
					return
				}

				const req = new formsTwirp.ContactUsReq()
				req.setName(values.name)
				req.setEmail(values.email)
				req.setMsg(values.message)
				req.setRecaptchaResponse(values.recaptchaResponse)

				formsClient.contactUs(req).then(() => {
					setFormSubmitted(true)
					setSubmitting(false)
				}).catch((e) => {
					console.log(JSON.stringify(e))
					setSubmitting(false)
				})
			}}
		>
		{({handleSubmit, submitForm, isSubmitting, setFieldValue}) => (
			<Form onSubmit={handleSubmit}>
				<Container>
					<Form.Row>
						<FormField label="Name" name="name" type="text"/>
						<FormField label="Email" name="email" type="text"/>
					</Form.Row><Form.Row>
						<FormField label="Message" name="message" as="textarea"/>
					</Form.Row><Form.Row>
					</Form.Row>
					<Form.Row><Col>
					{ recaptchaLoaded ? (
							<Button type="submit" disabled={isSubmitting}>Submit</Button>
					) : (
						<>
							<Spinner animation="border" />
							<p>Please wait while form loads</p>
						</>
					)}
					</Col></Form.Row>
					<Form.Row><Col>
						<br/>
						<center>
							<Recaptcha
								badge="inline"
								ref={e => recaptchaInstance = e}
								render="explicit"
								sitekey="6LcZg0AaAAAAANPdf6dXHEmQLf0fD2DR-Es7ztVH"
								size="invisible"
								verifyCallback={(response) => {
									setFieldValue('recaptchaResponse', response)
									submitForm()
								}}
								onloadCallback={() => setRecaptchaLoaded(true) }
								expiredCallback={() => setFieldValue('recaptchaResponse', '')}
							/>
						</center>
					</Col></Form.Row>

				</Container>
			</Form>
		)}
		</Formik>
	)
}

export default Home
