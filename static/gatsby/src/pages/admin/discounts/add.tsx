import React from 'react'
import Form from 'react-bootstrap/Form'
import Col from 'react-bootstrap/Col'
import AdminPage from '../../../components/AdminPage'
import Button from 'react-bootstrap/Button'
import ListGroup from 'react-bootstrap/ListGroup'
import FormField from '../../../components/FormField'
import WindowClose from '../../../components/WindowClose'
import {Formik} from 'formik'
import {createDiscount} from '../../../components/twirp'
import {dss} from '../../../rpc/discount.pb'


export default () => (
	<AdminPage title="Add Discount">
		{(gAuth) => {
			const discountClient = createDiscount(gAuth)

			return (
				<Formik
					initialValues={{
						code: '',
						singleDiscounts: []
					}}
					onSubmit={(values, {setSubmitting}) => {
						return discountClient.add({
							bundle: {
								code: values.code,
								discounts: values.singleDiscounts
							}
						}).then(req => {
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
							<Form.Row><Col>
								<FormField label="Code" name="code" type="text" />
							</Col></Form.Row>
							<Form.Row>
								<ListGroup style={{width: "100%"}}>
									{ values.singleDiscounts.map((sd, idx) => (
										<ListGroup.Item key={idx}>
											<Form.Row><Col className="text-right">
												<WindowClose onClick={() => setFieldValue('singleDiscounts', [...values.singleDiscounts.slice(0,idx), ...values.singleDiscounts.slice(idx+1)])} />
											</Col></Form.Row><Form.Row><Col>
												<FormField label="Name" name={`singleDiscounts[${idx}].name`} />
											</Col><Col>
												<FormField label="Applied To" name={`singleDiscounts[${idx}].appliedTo`} as="select">
													<option value={dss.PurchaseItem.FullWeekendPassPurchaseItem}>Full Weekend Pass</option>
													<option value={dss.PurchaseItem.DanceOnlyPassPurchaseItem}>Dance Only Pass</option>
													<option value={dss.PurchaseItem.MixAndMatchPurchaseItem}>Mix and Match</option>
													<option value={dss.PurchaseItem.SoloJazzPurchaseItem}>Solo Jazz</option>
													<option value={dss.PurchaseItem.TeamCompetitionPurchaseItem}>Team Competition</option>
													<option value={dss.PurchaseItem.TShirtPurchaseItem}>T-Shirt</option>
												</FormField>
											</Col></Form.Row>
										</ListGroup.Item>
									))}
								</ListGroup>
							</Form.Row>
							<Form.Row><Col>
								<Button onClick={() => {
									setFieldValue("singleDiscounts", [...values.singleDiscounts, {
										name: '',
										appliedTo: ''
									}])
								}}>Add New Single Discount</Button>
							</Col><Col>
								<Button type="submit" disabled={isSubmitting}>Submit Discount</Button>
							</Col></Form.Row>
						</Form>
					}
				</Formik>
			)
		}}
	</AdminPage>
)
