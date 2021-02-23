import React, {useEffect, useState} from 'react'
import {AdminPage} from '../../components/Page'
import discountTwirp from "../../rpc/discount_pb_twirp.js"
import Row from 'react-bootstrap/Row'
import Card from 'react-bootstrap/Card'
import Accordion from 'react-bootstrap/Accordion'
import ListGroup from 'react-bootstrap/ListGroup'

export default () => (
	<AdminPage title="Add Discount">
		{(gAuth) => {
			const [bundles, setBundles] = useState(null)
			const [discountClient, setDiscountClient] = useState(null)
			useEffect(() => {
				const dc = discountTwirp.createDiscountClient(`${process.env.GATSBY_BACKEND}`, {
					Authorization: `Bearer ${gAuth.accessToken}`
				})
				dc.list(new discountTwirp.DiscountListReq()).then(res => {
					setDiscountClient(dc)
					setBundles(res.bundlesList)
				}).catch( err => {
					console.error(err)
				})
			}, [gAuth])
			return (
				<Row>
					<Accordion style={{width: "100%"}}>
						{bundles && bundles.map((bundle, idx) => (
							<Card key={idx}>
								<Accordion.Toggle as={Card.Header} eventKey={idx.toString()}>Code: {bundle.code}</Accordion.Toggle>
								<Accordion.Collapse eventKey={idx.toString()}><Card.Body>
									<ListGroup>
										{bundle.discountsList.map((discount, discountIdx) => (
											<ListGroup.Item key={discountIdx}>
												Name: {discount.name}<br/>
												Applied To: {(() => {
													switch(discount.appliedTo) {
														case discountTwirp.PurchaseItem.FULLWEEKENDPASSPURCHASEITEM:
															return "Full Weekend Pass"
														case discountTwirp.PurchaseItem.DANCEONLYPASSPURCHASEITEM:
															return "Dance Only Pass"
														case discountTwirp.PurchaseItem.MIXANDMATCHPURCHASEITEM:
															return "Mix and Match"
														case discountTwirp.PurchaseItem.SOLOJAZZPURCHASEITEM:
															return "Solo Jazz"
														case discountTwirp.PurchaseItem.TEAMCOMPETITIONPURCHASEITEM:
															return "Team Competition"
														case discountTwirp.PurchaseItem.TSHIRTPURCHASEITEM:
															return "T-Shirt"
														default:
															return ""
													}
												})()}
											</ListGroup.Item>
										))}
									</ListGroup>
								</Card.Body></Accordion.Collapse>
							</Card>
						))}
					</Accordion>
				</Row>
			)
		}}
	</AdminPage>
)
