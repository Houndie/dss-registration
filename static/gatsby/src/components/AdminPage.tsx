import React, {useState} from 'react'
import AdminMenu from './AdminMenu'
import GenericPage, {RequireAuth} from './GenericPage'
import {GoogleLoginResponse} from 'react-google-login'
import {AuthedClients} from './auth'

interface AdminPageProps {
	title: string
	children: (auth: AuthedClients) => React.ReactNode
}

export default ({title, children}: AdminPageProps) => (
	<GenericPage 
		title={title} 
		menu={(auth) => <AdminMenu auth={auth}/>} 
		requireAuth={{ 
			required: true,
			callback: children 
		} as RequireAuth}
	/>
)
