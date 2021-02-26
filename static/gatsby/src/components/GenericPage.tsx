import React, {useState} from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Footer from './Footer'
import Hero from './Hero'
import AuthGuard from './AuthGuard'
import {GoogleLoginResponse} from "react-google-login"

export type RequireAuth = {
	required: true
	callback: (gAuth: GoogleLoginResponse) => React.ReactNode
}

export type OptionalAuth = {
	required: false
	callback: (gAuth: GoogleLoginResponse | null) => React.ReactNode
}

interface GenericPageProps {
	title: string;
	menu: (gAuth: GoogleLoginResponse | null, setGAuth: (arg0: GoogleLoginResponse | null) => void) => React.ReactNode
	requireAuth: RequireAuth | OptionalAuth
}

export default ({title, menu, requireAuth}: GenericPageProps) => {
	const [gAuth, setGAuth] = useState<GoogleLoginResponse | null>(null)
	return (
		<>
			{menu(gAuth, setGAuth)}
			<Hero image='/page_header.png' height='250px' title={title} />
			<Container>
				{ requireAuth.required ? (
					<AuthGuard gAuth={gAuth}>
						{(gAuth2) => {return requireAuth.callback(gAuth2)}}
					</AuthGuard>
				) : (() => { console.log("HERE"); console.log(requireAuth); return requireAuth.callback(gAuth)})()}
			</Container>
			<Footer />
		</>
	)
}
