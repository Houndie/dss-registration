import React from 'react'
import {useField} from 'formik'
import Form from 'react-bootstrap/Form'
import Container from "react-bootstrap/Container"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"

export default ({ granularity, smallEnd, bigEnd, label, ...props }) => {
	const [field, meta] = useField({ ...props});
	let cols = [];
	for(let i = 0; i < granularity; i++) {
		cols.push(
			<Col key={i} className="text-center">
				<Form.Check inline aria-label={label+i} type='radio' {...field} {...props} value={i}/>
			</Col>
		)
	}
	return (
		<Form.Group as={Col}>
			<Container>
				<Row><Col>
					{smallEnd}
				</Col><Col className='text-right'>
					{bigEnd}
				</Col></Row>
				<Row>{cols}</Row>
			</Container>
			<Form.Control.Feedback type='invalid' tooltip>{meta.error}</Form.Control.Feedback>
		</Form.Group>
	);
};
