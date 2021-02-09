import React from 'react'
import {useField} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

export default ({ title, otherName, labels, ...props }) => {
	const [field, meta] = useField(props);
	const [otherfield, othermeta] = (otherName ? useField({name: otherName}) : [undefined, undefined])
	return (
		<Form.Group as={Col}>
			{labels.map((label) => (	
				<Form.Check type='radio' key={label} value={label} label={label} isInvalid={meta.touched && !!meta.error} onChange={field.onChange} name={props.name}/>
			))}
			{otherName && 
				(
					<Form.Row>
						<Col md="auto">
							<Form.Check inline type='radio' label="Other" isInvalid={meta.touched && !!meta.error} {...field} {...props} />
						</Col><Col>
							<Form.Control isInvalid={othermeta.touched && !!othermeta.error} {...otherfield} {...props} />
						</Col>
					</Form.Row>
				)
			}
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
