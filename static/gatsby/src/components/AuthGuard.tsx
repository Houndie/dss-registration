import React from 'react'
import {GoogleLoginResponse} from "react-google-login"

interface AuthGuardProps {
	gAuth: GoogleLoginResponse | null
	children: (arg0: GoogleLoginResponse) => React.ReactNode
}

export default ({gAuth, children}: AuthGuardProps) => {
	if(!gAuth) {
		return <p>You must be logged in to view this page!</p>
	}
	return <>{children(gAuth)}</>
}
