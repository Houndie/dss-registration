import React from 'react'
import '../styles/style.scss'
import Container from 'react-bootstrap/Container'
import Footer from './Footer'
import Hero from './Hero'
import AuthGuard from './AuthGuard'
import {GoogleLoginResponse} from "react-google-login"

export type RequireAuth = {
	required: true
	callback: () => React.ReactNode
}

export type OptionalAuth = {
	required: false
	callback: () => React.ReactNode
}

interface GenericPageProps {
	title: string;
	menu: () => React.ReactNode
	requireAuth: RequireAuth | OptionalAuth
}

export default ({title, menu, requireAuth}: GenericPageProps) => (
	<>
		{menu()}
		<Hero image='/page_header.png' height='250px' title={title} />
		<Container>
			{ requireAuth.required ? (
				<AuthGuard>
					{requireAuth.callback}
				</AuthGuard>
			) : (() => { return requireAuth.callback()})()}
		</Container>
		<Footer />
	</>
)
