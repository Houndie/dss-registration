import React from 'react'
import { useAuth0 } from '@auth0/auth0-react';

interface AuthGuardProps {
	children: () => React.ReactNode
}

export default ({children}: AuthGuardProps) => {
	const { isLoading, isAuthenticated } = useAuth0()
	if(isLoading) {
		return <></>
	}
	if( isAuthenticated ){
		return <>{children()}</>
	}
	return <p>You must be logged in to view this page!</p>
}
