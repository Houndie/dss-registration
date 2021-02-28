import React from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Footer from './Footer'
import Hero from './Hero'
import AuthGuard from './AuthGuard'
import {GoogleLoginResponse} from "react-google-login"
import {useAuth, AuthResult, AuthedClients, Clients} from './auth'

export type RequireAuth = {
	required: true
	callback: (clients: AuthedClients) => React.ReactNode
}

export type OptionalAuth = {
	required: false
	callback: (clients: Clients) => React.ReactNode
}

interface GenericPageProps {
	title: string;
	menu: (auth: AuthResult) => React.ReactNode
	requireAuth: RequireAuth | OptionalAuth
}

export default ({title, menu, requireAuth}: GenericPageProps) => {
	const auth = useAuth(`${process.env.GATSBY_CLIENT_ID}`)
	return (
		<>
			{menu(auth)}
			<Hero image='/page_header.png' height='250px' title={title} />
			<Container>
				{ requireAuth.required ? (
					<AuthGuard clients={auth.clients}>
						{(clients) => {return requireAuth.callback(clients)}}
					</AuthGuard>
				) : (() => { return requireAuth.callback(auth.clients)})()}
			</Container>
			<Footer />
		</>
	)
}
