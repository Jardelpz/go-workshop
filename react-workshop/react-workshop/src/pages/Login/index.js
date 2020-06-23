import React from 'react'
import { Link } from 'react-router-dom'

function LoginPage(){
    return (
        <div>
            <h2>Login</h2>
            <Link to="/ofertas">Ir para ofertas</Link>
        </div>
    )
}


export default LoginPage