import React, {useEffect, useState} from 'react'
import Page from '../components/Page.js'
import Form from "react-bootstrap/Form"
import Col from "react-bootstrap/Col"
import Container from "react-bootstrap/Container"
import Button from "react-bootstrap/Button"
import Spinner from "react-bootstrap/Spinner"
import FormField from '../components/FormField.js'
import FormCheck from '../components/FormCheck.js'
import SlidingScale from '../components/SlidingScale.js'
import RadioGroup from '../components/RadioGroup.js'
import {Formik} from 'formik'
import * as Yup from 'yup'
import Recaptcha from 'react-recaptcha'
import google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb.js'
import formsTwirp from "../rpc/forms_pb_twirp.js"
import {ImportRecaptchaEffect} from "../components/recaptcha.js"

const ValidateSchema = Yup.object().shape({
	description: Yup.string()
		.required('Required'),
});

export default () => {
	const [submitted, setSubmitted] = useState(false)
	const [recaptchaLoaded, setRecaptchaLoaded] = useState(false)
	const formsClient = formsTwirp.createFormsClient(`${process.env.GATSBY_BACKEND}`)
	useEffect(ImportRecaptchaEffect)
	let recaptchaInstance
	return (
		<Page title="Safety Report">
			{(submitted ? (
				<p>Thank you for submitting your report. We will take action as soon as we are able.</p>
			) : (

				<Formik
					initialValues={{
						occurredOnDate: '',
						occurredOnTime: '',
						description: '',
						severity: '',
						issuesBefore: '',
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
						if (values.recaptchaResponse === '') {
							recaptchaInstance.execute()
							return
						}

						const ts = new google_protobuf_timestamp_pb.Timestamp()
						ts.fromDate(new Date(values.occurredOnDate + ' ' + values.occurredOnTime))

						const req = new formsTwirp.SafetyReportReq()
						req.setOccurredOn(ts)
						req.setDescription(values.description)
						req.setSeverity(values.severity)
						req.setIssuesBefore(values.issuesBefore)
						req.setResolution(values.resolution)
						req.setName(values.Name)
						req.setEmail(values.Email)
						req.setPhoneNumber(values.PhoneNumber)
						req.setRecaptchaResponse(values.recaptchaResponse)

						formsClient.safetyReport(req).then(() => {
							setSubmitted(true)
							setSubmitting(false)
						}).catch((e) => {
							console.log(JSON.stringify(e))
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
			))}
		</Page>
	)
}
