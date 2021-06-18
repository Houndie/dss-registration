import React, {useState} from 'react'
import AdminMenu from './AdminMenu'
import GenericPage, {RequireAuth} from './GenericPage'

interface AdminPageProps {
	title: string
	children: () => React.ReactNode
}

export default ({title, children}: AdminPageProps) => (
	<GenericPage 
		title={title} 
		menu={() => <AdminMenu/>} 
		requireAuth={{ 
			required: true,
			callback: children 
		} as RequireAuth}
	/>
)
