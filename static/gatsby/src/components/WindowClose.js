import React from 'react'                   
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome' 
import { faWindowClose } from '@fortawesome/free-solid-svg-icons' 
                                      
const WindowClose = props => (<FontAwesomeIcon icon={faWindowClose} onClick={props.onClick} style={{cursor: "pointer", color: "red"}}/>)

export default WindowClose   
