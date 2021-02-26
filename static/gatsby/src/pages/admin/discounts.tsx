import React, {useEffect, useState} from 'react'
import AdminPage from '../../components/AdminPage'
import {createDiscount} from "../../components/twirp"
import {dss} from "../../rpc/discount.pb"
import Row from 'react-bootstrap/Row'
import Card from 'react-bootstrap/Card'
import Accordion from 'react-bootstrap/Accordion'
import ListGroup from 'react-bootstrap/ListGroup'
import parseDollar from '../../components/parseDollar'

export default () => (
	<AdminPage title="Add Discount">
		{(gAuth) => {
			const [bundles, setBundles] = useState<dss.IDiscountBundle[]>([])
			const [discountClient, setDiscountClient] = useState<dss.Discount | null>(null)
			useEffect(() => {
				const dc = createDiscount(gAuth)

				dc.list({}).then(res => {
					setDiscountClient(dc)
					setBundles(res.bundles)
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
														console.error("no amount found")
														return ""
													}
													if(discount.amount.dollar) {
														return parseDollar(discount.amount.dollar)
													} else if (discount.amount.percent) {
														return "%" + discount.amount.percent
													} else if (discount.amount.squareNotFound) {
														return "discount not found in square"
													} else {
														console.error("unknown amount found")
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
