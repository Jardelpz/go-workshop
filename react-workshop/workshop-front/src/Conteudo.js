import React from 'react'
import PropType from 'prop-types'

function Conteudo(props) {
    const {title, pagina} = props
    return (
        <div>
            <h1>{title}</h1>
            <h2>{pagina}</h2>
        </div>
    
    )
}

Conteudo.propType = {
    title: PropType.string.isRequired,
    pagina: PropType.number
}

Conteudo.defaultProps = {
    pagina: 0
}


export default Conteudo