import React from 'react'
import {useField} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

export default ({ label, ...props }) => {
	const [field, meta] = useField(props);
	return (
		<Form.Group as={Col}>
			<Form.Check label={label} isInvalid={meta.touched && !!meta.error} {...field} {...props} />
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
