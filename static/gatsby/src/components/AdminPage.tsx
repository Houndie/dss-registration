import React, {useState} from 'react'
import AdminMenu from './AdminMenu'
import GenericPage, {RequireAuth} from './GenericPage'
import {GoogleLoginResponse} from 'react-google-login'

interface AdminPageProps {
	title: string
	children: (gAuth: GoogleLoginResponse) => React.ReactNode
}

export default ({title, children}: AdminPageProps) => (
	<GenericPage 
		title={title} 
		menu={(gAuth, setGAuth) => <AdminMenu gAuth={gAuth} setGAuth={setGAuth} />} 
		requireAuth={{ callback: children } as RequireAuth}
	/>
)
