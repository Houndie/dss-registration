import React, {useEffect, useState} from 'react'
import Page from '../components/Page'
import Form from "react-bootstrap/Form"
import Col from "react-bootstrap/Col"
import Container from "react-bootstrap/Container"
import Button from "react-bootstrap/Button"
import Spinner from "react-bootstrap/Spinner"
import FormField from '../components/FormField'
import FormCheck from '../components/FormCheck'
import SlidingScale from '../components/SlidingScale'
import RadioGroup from '../components/RadioGroup'
import {Formik} from 'formik'
import * as Yup from 'yup'
import Recaptcha from 'react-recaptcha'
import google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb.js'
import {createForms} from "../rpc/forms.twirp"
import {ImportRecaptchaEffect} from "../components/recaptcha"

const ValidateSchema = Yup.object().shape({
	description: Yup.string()
		.required('Required'),
});

export default () => {
	const [submitted, setSubmitted] = useState(false)
	const [recaptchaLoaded, setRecaptchaLoaded] = useState(false)
	const formsClient = createForms(`${process.env.GATSBY_BACKEND}`)
	useEffect(ImportRecaptchaEffect)
	let recaptchaInstance: Recaptcha | null
	return (
		<Page title="Safety Report">
			{(clients) => {
				if(submitted){
					return <p>Thank you for submitting your report. We will take action as soon as we are able.</p>
				}
				return (
					<Formik
						initialValues={{
							occurredOnDate: '',
							occurredOnTime: '',
							description: '',
							severity: 0,
							issuesBefore: false,
							resolution: '',
							resolutionOther: '',
							contactInformation: false,
							name: '',
							email: '',
							phoneNumber: '',
							recaptchaResponse: '',
						}}
						validationSchema={ValidateSchema}
						onSubmit={(values, {setSubmitting}) => {
							if (!recaptchaInstance) {
								console.error("recaptcha instance not loaded!?")
								return
							}
							if (values.recaptchaResponse === '') {
								recaptchaInstance.execute()
								return
							}
							clients.forms.safetyReport({
								occurredOn: {
									seconds: new Date(values.occurredOnDate + ' ' + values.occurredOnTime).valueOf()/1000
								},
								description: values.description,
								severity: values.severity,
								issuesBefore: values.issuesBefore,
								resolution: values.resolution,
								name: values.name,
								email: values.email,
								phoneNumber: values.phoneNumber,
								recaptchaResponse: values.recaptchaResponse
							}).then(() => {
								setSubmitted(true)
							}, e => {
								console.log(JSON.stringify(e))
							}).then(() => {
								setSubmitting(false)
							})
							
						}}
					>
						{({values, handleSubmit, submitForm, setFieldValue, isSubmitting}) => (
							<Form onSubmit={handleSubmit}>
								<Container>
									<Form.Row><FormField as="textarea" name="description" label="Describe the incident that occurred (required)"/></Form.Row>
									<Form.Row>
										<FormField type="date" name="occurredOnDate" label="Date the issue occurred" />
										<FormField type="time" name="occurredOnTime" label="Time the issue occurred" />
									</Form.Row>
									<p>Please Rank the severity of the issue</p>
									<Form.Row>
										<SlidingScale 
											granularity={10}
											smallEnd="Barely a problem" 
											bigEnd="I never want to see this person again" 
											label="severity"
											name="severity"
										/>
									</Form.Row>
									<p>What would you like to see done to resolve this issue?</p>
									<Form.Row>
										<RadioGroup 
											labels={[
												"No action, just raising awareness.",
												"I would like the organizers to keep an eye on the person causing the issue.",
												"I would like the organizers to speak to the person, without using details from the report.",
												"I would like the organizers to speak to the person, and they may use details from this report.",
												"I would like the organizers to set up a meeting with me and the person."
											]}
											name='resolution'
											otherName='resolutionOther'
										/>
									</Form.Row>
									<p>Have you add an issue with this person before?</p>
									<Form.Row>
										<RadioGroup
											labels={["Yes","No"]}
											name='issuesBefore'
										/>
									</Form.Row>
									<Form.Row><FormCheck label="I am okay with sharing my contact information" name="contactInformation"/></Form.Row>
									{values.contactInformation && (
										<Form.Row>
											<FormField label="Name" name="name" type="text"/>
											<FormField label="Email" name="email" type="email"/>
											<FormField label="Phone Number" name="phone" type="tel"/>
										</Form.Row>
									)}
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
									<Recaptcha
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
								</Container>
							</Form>
						)}
					</Formik>
				)
			}}
		</Page>
	)
}
