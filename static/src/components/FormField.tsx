import React from 'react'
import {useField, FieldHookConfig} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

type FormFieldProps = {
	label: string
	as?: 'input' | 'textarea' | 'select'
	children?: React.ReactNode
} & FieldHookConfig<any>

export default ({label, ...props}: FormFieldProps) => {
	const [field, meta] = useField(props);
	return (
		<Form.Group as={Col}>
			<Form.Label htmlFor={props.id || props.name}>{label}</Form.Label>
			<br />
			<Form.Control isInvalid={meta.touched && !!meta.error} {...field} as={props.as} children={props.children}/>
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
