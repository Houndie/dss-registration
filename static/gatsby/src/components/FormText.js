import React from 'react'
import {useField} from 'formik'

const FormText = ({ label, ...props }) => {
	const [field, meta] = useField(props);
	return (
		<>
			<div className="form-group">
				<label htmlFor={props.id || props.name}>{label}</label>
				<br />
				<input className="form-control" {...field} {...props} />
				{meta.touched && meta.error ? (
					<>
						<br />
						<div className="error">{meta.error}</div>
					</>
				) : null}
			</div>
		</>
	);
};

export default FormText
