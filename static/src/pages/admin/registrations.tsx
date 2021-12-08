import React, {useState, useEffect} from 'react'
import AdminPage from '../../components/AdminPage'
import useTwirp from "../../components/useTwirp"
import { useAuth0 } from '@auth0/auth0-react';
import {dss} from "../../rpc/registration.pb"
import Col from "react-bootstrap/Col"
import Row from "react-bootstrap/Row"
import LoadingPage from "../../components/LoadingPage"
import PleaseVerifyEmail from "../../components/PleaseVerifyEmail"

const isPaid = (r: dss.IRegistrationInfo) => 
	((!r.fullWeekendPass || r.fullWeekendPass.squarePaid || r.fullWeekendPass.adminPaymentOverride) &&
		(!r.danceOnlyPass || r.danceOnlyPass.squarePaid || r.danceOnlyPass.adminPaymentOverride) &&
		(!r.soloJazz || r.soloJazz.squarePaid || r.soloJazz.adminPaymentOverride) &&
		(!r.mixAndMatch || r.mixAndMatch.squarePaid || r.mixAndMatch.adminPaymentOverride) &&
		(!r.teamCompetition || r.teamCompetition.squarePaid || r.teamCompetition.adminPaymentOverride) &&
		(!r.tshirt || r.tshirt.squarePaid || r.tshirt.adminPaymentOverride))

export default () => {
	const [myRegistrations, setMyRegistrations] = useState<dss.IRegistrationInfo[]|null>(null)
	const { registration } = useTwirp()
	const { isLoading, isAuthenticated, loginWithRedirect, user } = useAuth0()

	useEffect(() => {
		if(isLoading || !isAuthenticated || !user?.email_verified){
			setMyRegistrations(null)
			return
		}

		registration().then(client => {
			return client.list({})
		}).then(res => {
			setMyRegistrations(res.registrations)
		})
	}, [isLoading, isAuthenticated, user])

	return (
		<AdminPage title="My Registrations">
			{() => {
				if(isLoading) {
					return <LoadingPage/>
				}
				if( !isAuthenticated ){
					return <p>You must be logged in to view this page! <a href="#" onClick={() => loginWithRedirect()}>Login Now</a></p>
				}
				if(!user?.email_verified) {
					return <PleaseVerifyEmail/>
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
							<a href={"/admin/registration/"+r.id} key={r.id}>
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
		</AdminPage>
	)
}
