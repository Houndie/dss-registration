import React, {useState} from 'react'
import Alert from 'react-bootstrap/Alert'
import Col from 'react-bootstrap/Col'
import Row from 'react-bootstrap/Row'

export enum ResponseKind {
	Good = 0,
	Bad
}

export interface FormResponse {
	kind: ResponseKind
	message: string
}

interface WithChildrenProps {
	children: (arg0: (arg0: FormResponse|null) => void) => React.ReactNode
}

export default ({children}: WithChildrenProps) => {
	const [response, setResponse] = useState<FormResponse|null>(null)
	return (
		<>
			<Row><Col>
				{children(setResponse)}
			</Col></Row>
			{ response && (
				<Row className="mt-3"><Col>
					<Alert variant={response.kind == ResponseKind.Good ? 'success' : 'danger'}>
						{response.message}
					</Alert>
				</Col></Row>
			)}
		</>
	)
}
