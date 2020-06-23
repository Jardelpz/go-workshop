import React, { useState, useEffect } from 'react'
import PropType from 'prop-types'
import './style.css'

function Index(props) {
    const {title, cor} = props
    const [active, setActive] = useState(false)
    
    const setButtonActive = () => {
        setActive(!active)
    }

    useEffect(()=> {
        if(active){
            document.getElementById("btn").style.background=`${cor}`;
        }else{
            document.getElementById("btn").style.background='#000'; 
        }  
    },[active])

    return (
        <div>
            <h1>
                {title}
            </h1>
            
            <button id="btn" onClick={()=>{ setButtonActive() }}>
                {active ? 'Ativado': 'Desativado'}
            </button>
        </div>
    )
}

Index.propType = {
    title: PropType.string.isRequired,
    cor: PropType.string
}

Index.defaultProps = {
    cor: '#FFF'
}


export default Index