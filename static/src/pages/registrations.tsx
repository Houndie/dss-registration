import React, {useState, useEffect} from 'react'
import Page from '../components/Page'
import useTwirp from "../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import {dss} from "../rpc/registration.pb"
import Col from "react-bootstrap/Col"
import Row from "react-bootstrap/Row"
import LoadingPage from "../components/LoadingPage"
import {isPaid} from "../components/RegistrationForm"

export default () => {
	const [myRegistrations, setMyRegistrations] = useState<dss.IRegistrationInfo[]|null>(null)
	const { registration } = useTwirp()
	const { isLoading, isAuthenticated, loginWithRedirect } = useAuth0()

	useEffect(() => {
		if(isLoading || !isAuthenticated){
			setMyRegistrations(null)
			return
		}

		registration().then(client => {
			return client.listByUser({})
		}).then(res => {
			setMyRegistrations(res.registrations)
		})
	}, [isLoading, isAuthenticated])

	return (
		<Page title="My Registrations">
			{() => {
				if(isLoading) {
					return <LoadingPage/>
				}
				if( !isAuthenticated ){
					return <p>You must be logged in to view this page! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
				}
				if( !myRegistrations ){
					return <LoadingPage/>
				}
				return (
					<>
						<b>
							<Row>
								<Col>Registered At</Col>
								<Col>First Name</Col>
								<Col>Last Name</Col>
								<Col>Email</Col>
								<Col>Paid</Col>
							</Row>
						</b>
						{ myRegistrations.map(r => (
							<a href={"/registration/"+r.id} key={r.id}>
								<Row>
									<Col>{(r.createdAt ? (new Date(r.createdAt)).toLocaleString() : null)}</Col>
									<Col>{r.firstName}</Col>
									<Col>{r.lastName}</Col>
									<Col>{r.email}</Col>
									<Col>{(isPaid(r) ? "yes" : "no")}</Col>
								</Row>
							</a>
						))}
					</>
				)
			}}
		</Page>
	)
}
