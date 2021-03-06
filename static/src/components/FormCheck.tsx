import React from 'react'
import {useField, FieldHookConfig} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

type FormCheckProps = {
	label: string
} & FieldHookConfig<any>

export default ({ label, ...props }: FormCheckProps) => {
	const [field, meta] = useField(props);
	return (
		<Form.Group as={Col}>
			<Form.Check label={label} isInvalid={meta.touched && !!meta.error} {...field} />
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
