import React from 'react'
import Form from 'react-bootstrap/Form'
import Col from 'react-bootstrap/Col'
import AdminPage from '../../../components/AdminPage'
import Button from 'react-bootstrap/Button'
import {Formik} from 'formik'
import DiscountForm from '../../../components/DiscountForm'


export default () => (
	<AdminPage title="Add Discount">
		{(clients) => {
			return (
				<Formik
					initialValues={{
						code: '',
						discounts: []
					}}
					onSubmit={(values, {setSubmitting}) => {
						return clients.discount.add({
							bundle: values
						}).then(res => {
							console.log("success")
						}).catch(err => {
							console.error(err)
						}).then(() => {
							setSubmitting(false)
						})
					}}
				>
					{({values, isSubmitting, handleSubmit, setFieldValue}) => 
						<Form onSubmit={handleSubmit}>
							<DiscountForm values={values} setFieldValue={setFieldValue} />
							<Form.Row><Col>
								<Button type="submit" disabled={isSubmitting}>Submit Discount</Button>
							</Col></Form.Row>
						</Form>
					}
				</Formik>
			)
		}}
	</AdminPage>
)
