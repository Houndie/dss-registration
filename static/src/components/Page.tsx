import React from 'react'
import Menu from './Menu'
import GenericPage from './GenericPage'

interface PageProps {
	title: string
}

const Page: React.FC<PageProps> = ({ title, children }) => (
	<GenericPage title={title} menu={() => <Menu />}>
		{children}
	</GenericPage>
)

export default Page
