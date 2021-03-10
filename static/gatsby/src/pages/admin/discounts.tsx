import React, {useEffect, useState} from 'react'
import AdminPage from '../../components/AdminPage'
import {dss} from "../../rpc/discount.pb"
import Row from 'react-bootstrap/Row'
import Container from 'react-bootstrap/Container'
import Card from 'react-bootstrap/Card'
import Accordion from 'react-bootstrap/Accordion'
import ListGroup from 'react-bootstrap/ListGroup'
import Modal from 'react-bootstrap/Modal'
import parseDollar from '../../components/parseDollar'
import {Formik} from 'formik'
import Form from 'react-bootstrap/Form'
import DiscountForm from '../../components/DiscountForm'
import Col from 'react-bootstrap/Col'
import Button from 'react-bootstrap/Button'
import WithAlert, {ResponseKind, FormResponse} from '../../components/WithAlert'
import useTwirp from '../../components/useTwirp'

type ListItem = {
	isForm: boolean
} & dss.IDiscountBundle

export default () => (
	<AdminPage title="List Discounts">{() => (
		<WithAlert>{(setResponse) => {
			const [bundles, setBundles] = useState<ListItem[]>([])
			const [deleting, setDeleting] = useState<boolean>(false)
			const {discount} = useTwirp()
			useEffect(() => {
				discount().then(client => {
					return client.list({})
				}).then(res => {
					setBundles(res.bundles.map(bundle => { return {...bundle, isForm: false}}))
				}).catch( err => {
					setResponse({
						kind: ResponseKind.Bad,
						message: "Error fetching discounts: "+err
					})
				})
			}, [])
			return (
				<Accordion style={{width: "100%"}}>
					{bundles && bundles.map((bundle, idx) => (
						<Card key={idx} bg={((bundle.discounts && bundle.discounts.find((discount) => {
							return !discount.amount || discount.amount.squareNotFound
						})) ? 'warning' : 'light')} >
							<Accordion.Toggle as={Card.Header} eventKey={idx.toString()}>Code: {bundle.code}</Accordion.Toggle>
							<Accordion.Collapse eventKey={idx.toString()}><Card.Body>
								{ bundle.isForm ? (
									<Formik
										initialValues={{
											code: (bundle.code ? bundle.code : ''),
											discounts: (bundle.discounts ? bundle.discounts.map(({amount, ...tail}) => tail) : [])
										}}
										onSubmit={(values, {setSubmitting}) => {
											return discount().then(client => {
												return client.update({
													oldCode: bundle.code,
													bundle: values
												}).then(res => {
													setResponse({
														kind: ResponseKind.Good,
														message: "Discount successfully updated"
													})
													return client.get({
														code: values.code
													})
												}).then(res => {
													setBundles([...bundles.slice(0, idx), {
														...res.bundle,
														isForm: false
													}, ...bundles.slice(idx+1)])
												})
											}).catch(err => {
												setResponse({
													kind: ResponseKind.Bad,
													message: "Error updating/refetching discount: "+err
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
													<Button type="submit" disabled={isSubmitting}>Update Discount</Button>
												</Col><Col>
													<Button onClick={ () => {
														setBundles([...bundles.slice(0, idx), {
															...bundle,
															isForm: false
														}, ...bundles.slice(idx+1)])
													}}>Cancel</Button>
												</Col></Form.Row>
											</Form>
										}
									</Formik>
								):(
									<>
										<Row className="mb-3"><Col>
											<ListGroup style={{width: "100%"}}>
												{bundle.discounts && bundle.discounts.map((discount, discountIdx) => (
													<ListGroup.Item key={discountIdx}>
														Name: {discount.name}<br/>
														Applied To: {(() => {
															switch(discount.appliedTo) {
																case dss.PurchaseItem.FullWeekendPassPurchaseItem:
																	return "Full Weekend Pass"
																case dss.PurchaseItem.DanceOnlyPassPurchaseItem:
																	return "Dance Only Pass"
																case dss.PurchaseItem.MixAndMatchPurchaseItem:
																	return "Mix and Match"
																case dss.PurchaseItem.SoloJazzPurchaseItem:
																	return "Solo Jazz"
																case dss.PurchaseItem.TeamCompetitionPurchaseItem:
																	return "Team Competition"
																case dss.PurchaseItem.TShirtPurchaseItem:
																	return "T-Shirt"
																default:
																	return ""
															}
														})()}<br/>
														Amount: {(() => {
															if(!discount.amount) {
																setResponse({
																	kind: ResponseKind.Bad,
																	message: "No discount amount found"
																})
																return ""
															}
															if(discount.amount.dollar) {
																return parseDollar(discount.amount.dollar)
															} else if (discount.amount.percent) {
																return "%" + discount.amount.percent
															} else if (discount.amount.squareNotFound) {
																return "discount not found in square"
															} else {
																setResponse({
																	kind: ResponseKind.Bad,
																	message: "Unknown discount amount found"
																})
																return ""
															}
														})()}
													</ListGroup.Item>
												))}
											</ListGroup>
										</Col></Row><Row><Col>
											<Button onClick={() => {
												setBundles([...bundles.slice(0, idx), {
													...bundle,
													isForm: true
												}, ...bundles.slice(idx+1)])
											}}>Edit</Button>
										</Col><Col>
											<Button variant="danger" onClick={() => setDeleting(true)}>Delete</Button>
										</Col></Row>
										<Modal show={deleting} size="lg" onHide={() => setDeleting(false)} centered>
											<Formik
												initialValues={{}}
												onSubmit={() => {
													discount().then(client => {
														return client.delete({
															code: bundle.code
														})
													}).then(res => {
														setBundles([...bundles.slice(0, idx), ...bundles.slice(idx+1)])
														setResponse({
															kind: ResponseKind.Good,
															message: "Discount deleted successfully!"
														})
													}).catch(err => {
														setResponse({
															kind: ResponseKind.Bad,
															message: "Error deleting discount: "+err
														})
													}).finally(() => {
														setDeleting(false)
													})
												}}
											>
												{({handleSubmit}) => (
													<Form onSubmit={handleSubmit}>
														<Modal.Body>
															<p>Are you sure you want to delete this?</p>
														</Modal.Body>
														<Modal.Footer>
															<Button variant="secondary" onClick={() => setDeleting(false)}>Cancel</Button>
															<Button variant="danger" type="submit">Delete</Button>
														</Modal.Footer>
													</Form>
												)}
											</Formik>
										</Modal>
									</>
								)}
							</Card.Body></Accordion.Collapse>
						</Card>
					))}
				</Accordion>
			)
		}}</WithAlert>
	)}</AdminPage>
)
