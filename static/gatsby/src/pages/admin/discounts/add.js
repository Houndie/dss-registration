import React from 'react'
import Form from 'react-bootstrap/Form'
import Col from 'react-bootstrap/Col'
import Page from '../../../components/Page'
import Button from 'react-bootstrap/Button'
import ListGroup from 'react-bootstrap/ListGroup'
import FormField from '../../../components/FormField'
import WindowClose from '../../../components/WindowClose'
import {Formik} from 'formik'
import discountTwirp from "../../../rpc/discount_pb_twirp.js"


export default () => (
	<Page title="Add Discount">
		{(gAuth) => {
			if(!gAuth) {
				return <p>You must be logged in to view this page!</p>
			}
			const discountClient = discountTwirp.createDiscountClient(`${process.env.GATSBY_BACKEND}`, {
				Authorization: `Bearer ${gAuth.accessToken}`
			})

			return (
				<Formik
					initialValues={{
						code: '',
						singleDiscounts: []
					}}
					onSubmit={(values, {setSubmitting}) => {
						const sds = values.singleDiscounts.map( singleDiscount => {
							const sd = new discountTwirp.SingleDiscount()
							sd.setName(singleDiscount.name)
							sd.setAppliedTo(singleDiscount.appliedTo)
							return sd
						})

						const b = new discountTwirp.DiscountBundle()
						b.setCode(values.code)
						b.setDiscountsList(sds)

						const req = new discountTwirp.DiscountAddReq();
						req.setBundle(b)
						
						return discountClient.add(req).then(req => {
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
													<option value={discountTwirp.PurchaseItem.FULLWEEKENDPASSPURCHASEITEM}>Full Weekend Pass</option>
													<option value={discountTwirp.PurchaseItem.DANCEONLYPASSPURCHASEITEM}>Dance Only Pass</option>
													<option value={discountTwirp.PurchaseItem.MIXANDMATCHPURCHASEITEM}>Mix and Match</option>
													<option value={discountTwirp.PurchaseItem.SOLOJAZZPURCHASEITEM}>Solo Jazz</option>
													<option value={discountTwirp.PurchaseItem.TEAMCOMPETITIONPURCHASEITEM}>Team Competition</option>
													<option value={discountTwirp.PurchaseItem.TSHIRTPURCHASEITEM}>T-Shirt</option>
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
	</Page>
)
