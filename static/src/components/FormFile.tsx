import React from 'react'
import {useField, FieldHookConfig} from 'formik'
import Form from 'react-bootstrap/Form'
import Col from "react-bootstrap/Col"

type FormFileProps = {
	label: string
	myref?: React.MutableRefObject<HTMLInputElement|undefined>
	as?: 'input' | 'textarea' | 'select'
	children?: React.ReactNode
} & FieldHookConfig<any>

export default ({label, myref, ...props}: FormFileProps) => {
	const [field, meta, helper] = useField(props);
	return (
		<Col>
			<Form.Label htmlFor={props.id || props.name}>{label}</Form.Label>
			<br />
			<Form.Control 
				type="file" 
				isInvalid={meta.touched && !!meta.error} 
				name={field.name}
				as={props.as} 
				children={props.children} 
				ref={myref}
				onChange={(event: React.ChangeEvent<HTMLInputElement>) => { 
					if (event.currentTarget.files == null || event.currentTarget.files.length == 0) {
						helper.setValue(undefined)
						return
					}
					helper.setValue(event.currentTarget.files[0])
				}}/>
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Col>
	);
};
