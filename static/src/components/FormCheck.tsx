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
		<Col>
			<Form.Check label={label} isInvalid={meta.touched && !!meta.error} checked={field.value} {...field} disabled={props.disabled} />
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Col>
	);
};
