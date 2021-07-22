import React from 'react'
import {dss} from '../rpc/discount.pb'
import WindowClose from './WindowClose'
import FormField from './FormField'
import Button from 'react-bootstrap/Button'
import Col from 'react-bootstrap/Col'
import Row from 'react-bootstrap/Row'
import Form from 'react-bootstrap/Form'
import ListGroup from 'react-bootstrap/ListGroup'

type Bundle = {
	[P in keyof Required<dss.IDiscountBundle>]: Exclude<Required<dss.IDiscountBundle>[P], null>;
}

interface DiscountFormProps {
	values: Bundle
	setFieldValue: (field: string, value: any, shouldValidate?: boolean) => void
}

export default ({values, setFieldValue}: DiscountFormProps) => (
	<>
		<Row><Col>
			<FormField label="Code" name="code" type="text" />
		</Col></Row>
		{ values.discounts && values.discounts.length > 0 && (
			<Row className="mb-3"><Col>
				<ListGroup>
					{ values.discounts.map((sd, idx) => (
						<ListGroup.Item key={idx}>
							<Row><Col className="text-right">
								<WindowClose onClick={() => setFieldValue('discounts', [...values.discounts.slice(0,idx), ...values.discounts.slice(idx+1)])} />
							</Col></Row><Row><Col>
								<FormField label="Name" name={`discounts[${idx}].name`} />
							</Col><Col>
								<FormField label="Applied To" name={`discounts[${idx}].appliedTo`} as="select">
									<option value={dss.PurchaseItem.FullWeekendPassPurchaseItem}>Full Weekend Pass</option>
									<option value={dss.PurchaseItem.DanceOnlyPassPurchaseItem}>Dance Only Pass</option>
									<option value={dss.PurchaseItem.MixAndMatchPurchaseItem}>Mix and Match</option>
									<option value={dss.PurchaseItem.SoloJazzPurchaseItem}>Solo Jazz</option>
									<option value={dss.PurchaseItem.TeamCompetitionPurchaseItem}>Team Competition</option>
									<option value={dss.PurchaseItem.TShirtPurchaseItem}>T-Shirt</option>
								</FormField>
							</Col></Row>
						</ListGroup.Item>
					))}
				</ListGroup>
			</Col></Row>
		)}
		<Row><Col>
			<Button onClick={() => {
				setFieldValue("discounts", [...values.discounts, {
					name: '',
					appliedTo: ''
				}])
			}}>Add New Single Discount</Button>
		</Col></Row>
	</>
)
