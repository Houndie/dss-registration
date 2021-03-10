import React, {useState, useEffect} from 'react'
import {createPortal} from 'react-dom'
import Nav from "react-bootstrap/Nav"
import { useAuth0 } from '@auth0/auth0-react';

export default () => {
	const { isLoading, isAuthenticated, loginWithRedirect, logout } = useAuth0()
	if(isLoading) {
		return <Nav.Link href="#">...</Nav.Link>
	}
	if(!isAuthenticated) {
		return <Nav.Link href="#" onClick={() => loginWithRedirect()}>Login</Nav.Link>
	}
	return <Nav.Link href="#" onClick={() => logout({returnTo: window.location.origin})}>Logout</Nav.Link>
}
