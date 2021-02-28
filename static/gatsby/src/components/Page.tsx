import React, {useState} from 'react'
import Menu from './Menu'
import GenericPage, {OptionalAuth} from './GenericPage'
import {GoogleLoginResponse} from 'react-google-login'
import {Clients} from './auth'

interface PageProps {
	title: string
	children: (auth: Clients) => React.ReactNode
}

export default ({title, children}: PageProps) => (
	<GenericPage 
		title={title} 
		menu={(auth) => <Menu auth={auth}/>} 
		requireAuth={{ 
			required: false,
			callback: children 
		} as OptionalAuth}
	/>
)
