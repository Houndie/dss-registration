import React from 'react'
import {useField, FieldHookConfig} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

interface OtherRadioProps {
	otherName: string
	touched: boolean
	error?: string
}

const OtherRadio = ({otherName, touched, error}: OtherRadioProps) => {
	const [field, meta] = useField({name: otherName})
	return (
		<Form.Row>
			<Col md="auto">
				<Form.Check inline type='radio' label="Other" isInvalid={touched && !!error} {...field} />
			</Col><Col>
				<Form.Control isInvalid={meta.touched && !!meta.error} {...field} />
			</Col>
		</Form.Row>
	)
}

type RadioGroupProps = {
	otherName?: string
	labels: string[]
} & FieldHookConfig<any>

export default ({ otherName, labels, ...props }: RadioGroupProps) => {
	const [field, meta] = useField(props);
	return (
		<Form.Group as={Col}>
			{labels.map((label) => (	
				<Form.Check type='radio' key={label} value={label} label={label} isInvalid={meta.touched && !!meta.error} onChange={field.onChange} name={props.name}/>
			))}
			{otherName && <OtherRadio otherName={otherName} touched={meta.touched} error={meta.error}/>}
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
