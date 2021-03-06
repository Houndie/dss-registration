import React, {useState} from 'react'
import Menu from './Menu'
import GenericPage, {OptionalAuth} from './GenericPage'

interface PageProps {
	title: string
	children: () => React.ReactNode
}

export default ({title, children}: PageProps) => (
	<GenericPage 
		title={title} 
		menu={() => <Menu/>} 
		requireAuth={{ 
			required: false,
			callback: children 
		} as OptionalAuth}
	/>
)
