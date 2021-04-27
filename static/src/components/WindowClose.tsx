import React from 'react'                   
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome' 
import { faWindowClose } from '@fortawesome/free-solid-svg-icons' 

interface WindowCloseProps {
	onClick: React.MouseEventHandler<SVGSVGElement>
}
                                      
export default ({onClick}: WindowCloseProps) => (<FontAwesomeIcon icon={faWindowClose} onClick={onClick} style={{cursor: "pointer", color: "red"}}/>)
