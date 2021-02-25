import React, {useState} from 'react'
import Menu from './Menu'
import GenericPage, {OptionalAuth} from './GenericPage'
import {GoogleLoginResponse} from 'react-google-login'

interface PageProps {
	title: string
	children: (gAuth: GoogleLoginResponse | null) => React.ReactNode
}

export default ({title, children}: PageProps) => (
	<GenericPage 
		title={title} 
		menu={(gAuth, setGAuth) => <Menu gAuth={gAuth} setGAuth={setGAuth} />} 
		requireAuth={{ callback: children } as OptionalAuth}
	/>
)
