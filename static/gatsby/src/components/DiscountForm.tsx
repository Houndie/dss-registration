import React from 'react'
import {dss} from '../rpc/discount.pb'
import WindowClose from './WindowClose'
import FormField from './FormField'
import Button from 'react-bootstrap/Button'
import Col from 'react-bootstrap/Col'
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
		<Form.Row><Col>
			<FormField label="Code" name="code" type="text" />
		</Col></Form.Row>
		<Form.Row>
			<ListGroup style={{width: "100%"}}>
				{ values.discounts && values.discounts.map((sd, idx) => (
					<ListGroup.Item key={idx}>
						<Form.Row><Col className="text-right">
							<WindowClose onClick={() => setFieldValue('discounts', [...values.discounts.slice(0,idx), ...values.discounts.slice(idx+1)])} />
						</Col></Form.Row><Form.Row><Col>
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
						</Col></Form.Row>
					</ListGroup.Item>
				))}
			</ListGroup>
		</Form.Row>
		<Form.Row><Col>
			<Button onClick={() => {
				setFieldValue("discounts", [...values.discounts, {
					name: '',
					appliedTo: ''
				}])
			}}>Add New Single Discount</Button>
		</Col></Form.Row>
	</>
)
