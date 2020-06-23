import React, {useState, useEffect } from 'react'
import PropTypes from 'prop-types'

function Header({ text, value }) {
    const [active, setActive] = useState(false)

    const handleClickButton = () => {
        setActive(!active)
    }

    useEffect(() => {
        if (active){
            console.log("Botão ativado")
        }
    }, [active])

    return (
        <button onClick={() => handleClickButton()}>
            {active ? 'ativado': 'desativado'}
        </button>
        // <p>{text} {value}</p>
    )
}

Header.propTypes = {
    text: PropTypes.string.isRequired,
    value: PropTypes.string
}

Header.defaultProps = {
    value:"Sem valor inicial"
}

export default Header