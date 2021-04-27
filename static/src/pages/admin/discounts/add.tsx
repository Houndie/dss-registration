import React, {useState} from 'react'
import Form from 'react-bootstrap/Form'
import Col from 'react-bootstrap/Col'
import AdminPage from '../../../components/AdminPage'
import Button from 'react-bootstrap/Button'
import Alert from 'react-bootstrap/Alert'
import {Formik} from 'formik'
import DiscountForm from '../../../components/DiscountForm'
import WithAlert, {ResponseKind, FormResponse} from '../../../components/WithAlert'
import useTwirp from '../../../components/useTwirp'

export default () => (
	<AdminPage title="Add Discount">
		{() => (
			<WithAlert>
				{(setResponse) => {
					const {discount} = useTwirp()
					return (
						<Formik
							initialValues={{
								code: '',
								discounts: []
							}}
							onSubmit={(values, {resetForm, setSubmitting}) => {
								return discount().then(client => {
									client.add({
										bundle: values
									})
								}).then(res => {
									setResponse({
										kind: ResponseKind.Good,
										message: "Discount added successfully!"
									})
									resetForm()
								}).catch(err => {
									setResponse({
										kind: ResponseKind.Bad,
										message: "Error adding discount: "+JSON.stringify(err)
									})
								}).finally(() => {
									setSubmitting(false)
								})
							}}
						>
							{({values, isSubmitting, handleSubmit, setFieldValue}) => 
								<Form onSubmit={handleSubmit}>
									<DiscountForm values={values} setFieldValue={setFieldValue} />
									<Form.Row className="mt-3"><Col>
										<Button type="submit" disabled={isSubmitting}>Submit Discount</Button>
									</Col></Form.Row>
								</Form>
							}
						</Formik>
					)
				}}
			</WithAlert>	
		)}
	</AdminPage>
)
